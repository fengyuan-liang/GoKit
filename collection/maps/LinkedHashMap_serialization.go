// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package maps

import (
	"bytes"
	"encoding/json"
)

func (m *LinkedHashMap[K, V]) MarshalJSON() ([]byte, error) {
	return m.ToJSON()
}

// ToJSON outputs the JSON representation of map.
func (m *LinkedHashMap[K, V]) ToJSON() ([]byte, error) {
	var b []byte
	buf := bytes.NewBuffer(b)

	buf.WriteRune('{')
	index := 0
	lastIndex := m.Size() - 1
	m.ForEach(func(k K, v V) {
		km, err := json.Marshal(k)
		if err != nil {
			return
		}
		buf.Write(km)

		buf.WriteRune(':')

		vm, err := json.Marshal(v)
		if err != nil {
			return
		}
		buf.Write(vm)

		if index != lastIndex {
			buf.WriteRune(',')
		}
		index++
	})

	buf.WriteRune('}')

	return buf.Bytes(), nil
}

// UnmarshalJSON @implements json.Unmarshaler
func (m *LinkedHashMap[K, V]) UnmarshalJSON(bytes []byte) error {
	return m.FromJSON(bytes)
}

func (m *LinkedHashMap[K, V]) FromJSON(data []byte) error {
	elements := make(map[K]V)
	err := json.Unmarshal(data, &elements)
	if err != nil {
		return err
	}
	// resize map
	m.Clear()
	for k, v := range elements {
		m.Put(k, v)
	}
	return nil
}
