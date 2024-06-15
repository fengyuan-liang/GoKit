// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package maps

import (
	"errors"
	"github.com/fengyuan-liang/GoKit/collection"
)

// ToJSON outputs the JSON representation of the map.
func (s *SynchronizedMap[K, V]) ToJSON() ([]byte, error) {
	if serializer, ok := s.m.(collection.JSONSerializer); ok {
		return serializer.MarshalJSON()
	}
	return nil, errors.New(" unknown map type")
}

// FromJSON populates the map from the input JSON representation.
func (s *SynchronizedMap[K, V]) FromJSON(data []byte) (err error) {
	if deserializer, ok := s.m.(collection.JSONDeserializer); ok {
		err = deserializer.UnmarshalJSON(data)
	}
	return
}

// UnmarshalJSON @implements json.Unmarshaler
func (s *SynchronizedMap[K, V]) UnmarshalJSON(bytes []byte) error {
	return s.FromJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (s *SynchronizedMap[K, V]) MarshalJSON() ([]byte, error) {
	return s.ToJSON()
}
