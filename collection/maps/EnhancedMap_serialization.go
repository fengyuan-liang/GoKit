// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package maps

import "encoding/json"

// ToJSON outputs the JSON representation of the map.
func (m *EnhancedMap[K, V]) ToJSON() ([]byte, error) {
	copyMap := make(map[K]V, m.Size())
	for k, v := range m.m {
		copyMap[k] = v
	}
	return json.Marshal(&copyMap)
}

// FromJSON populates the map from the input JSON representation.
func (m *EnhancedMap[K, V]) FromJSON(data []byte) error {
	rawMap := make(map[K]V)
	err := json.Unmarshal(data, &rawMap)
	if err == nil {
		m.Clear()
		for k, v := range rawMap {
			m.Put(k, v)
		}
	}
	return err
}

// UnmarshalJSON @implements json.Unmarshaler
func (m *EnhancedMap[K, V]) UnmarshalJSON(bytes []byte) error {
	return m.FromJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (m *EnhancedMap[K, V]) MarshalJSON() ([]byte, error) {
	return m.ToJSON()
}
