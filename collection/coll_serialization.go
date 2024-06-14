// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package collection

// JSONSerializer provides JSON serialization
type JSONSerializer interface {
	// ToJSON outputs the JSON representation of collection's elements.
	ToJSON() ([]byte, error)
	// MarshalJSON @implements json.Marshaler
	MarshalJSON() ([]byte, error)
}

// JSONDeserializer provides JSON deserialization
type JSONDeserializer interface {
	// FromJSON populates collection's elements from the input JSON representation.
	FromJSON([]byte) error
	// UnmarshalJSON @implements json.Unmarshaler
	UnmarshalJSON([]byte) error
}
