// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package stream

func Of[In any, Out any](elements []In) IStream[In, Out] {
	if len(elements) == 0 {
		return nil
	}
	stream := new(Stream[In, Out])
	stream.Data = elements
	return stream
}
