// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sets

import "encoding/json"

// ToJSON outputs the JSON representation of the set.
func (set *HashSet[T]) ToJSON() ([]byte, error) {
	return json.Marshal(set.Values())
}

// FromJSON populates the set from the input JSON representation.
func (set *HashSet[T]) FromJSON(data []byte) error {
	var elements []T
	err := json.Unmarshal(data, &elements)
	if err == nil {
		set.Clear()
		set.Add(elements...)
	}
	return err
}

// UnmarshalJSON @implements json.Unmarshaler
func (set *HashSet[T]) UnmarshalJSON(bytes []byte) error {
	return set.FromJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (set *HashSet[T]) MarshalJSON() ([]byte, error) {
	return set.ToJSON()
}
