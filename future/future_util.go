package future

import "sync"

// AllSuccess waits for all futures to complete and returns true if all are successful
func AllSuccess[T any](futures ...*Future[T]) (bool, error) {
	var wg sync.WaitGroup
	resultChan := make(chan bool, len(futures))
	errChan := make(chan error, len(futures))

	for _, f := range futures {
		wg.Add(1)
		go func(f *Future[T]) {
			defer wg.Done()
			_, err := f.Get()
			if err != nil {
				resultChan <- false
				errChan <- err
				return
			}
			resultChan <- f.Success
		}(f)
	}

	wg.Wait()
	close(resultChan)
	close(errChan)

	allSuccess := true
	var combinedError error
	for success := range resultChan {
		if !success {
			allSuccess = false
		}
	}

	for err := range errChan {
		if err != nil {
			combinedError = err
		}
	}

	return allSuccess, combinedError
}

// AnySuccess waits for any future to complete and returns true if any is successful
func AnySuccess[T any](futures ...*Future[T]) (bool, error) {
	var wg sync.WaitGroup
	resultChan := make(chan bool, len(futures))
	errChan := make(chan error, len(futures))

	for _, f := range futures {
		wg.Add(1)
		go func(f *Future[T]) {
			defer wg.Done()
			_, err := f.Get()
			if err != nil {
				resultChan <- false
				errChan <- err
				return
			}
			resultChan <- f.Success
		}(f)
	}

	go func() {
		wg.Wait()
		close(resultChan)
		close(errChan)
	}()

	for {
		select {
		case success, ok := <-resultChan:
			if !ok {
				return false, <-errChan
			}
			if success {
				return true, nil
			}
		case err := <-errChan:
			if err != nil {
				return false, err
			}
		}
	}
}
