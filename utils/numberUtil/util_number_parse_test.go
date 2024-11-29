// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package numberUtil

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestParseString(t *testing.T) {
	tests := []struct {
		input    any
		expected string
	}{
		{123, "123"},
		{int32(123), "123"},
		{int64(1234567890123456789), "1234567890123456789"},
		{uint(123), "123"},
		{float32(123.456), "123.456"},
		{123.456, "123.456"},
		{"test string", "test string"},
		{nil, ""},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("input=%v", tt.input), func(t *testing.T) {
			// Redirect log output to a buffer
			var buf bytes.Buffer
			log.SetOutput(&buf)

			result := ParseString(tt.input)

			if result != tt.expected {
				t.Errorf("ParseString(%v) = %v, want %v", tt.input, result, tt.expected)
			}

			if tt.input == nil {
				logOutput := buf.String()
				expectedLog := ""
				if !strings.Contains(logOutput, expectedLog) {
					t.Errorf("Expected log output to contain %q, but got %q", expectedLog, logOutput)
				}
			}
		})
	}
}

func TestParseNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
		wantErr  bool
	}{
		{"123", 123, false},
		{"123.456", 123.456, false},
		{"abc", 0, true},
	}

	for _, tt := range tests {
		result, err := ParseNumber(tt.input)
		if (err != nil) != tt.wantErr {
			t.Errorf("ParseNumber(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			continue
		}
		if result != tt.expected {
			t.Errorf("ParseNumber(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestParseInt64(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
		wantErr  bool
	}{
		{"123", 123, false},
		{"-123", -123, false},
		{"abc", 0, true},
	}

	for _, tt := range tests {
		result, err := ParseInt64(tt.input)
		if (err != nil) != tt.wantErr {
			t.Errorf("ParseInt64(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			continue
		}
		if result != tt.expected {
			t.Errorf("ParseInt64(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestParseInt32(t *testing.T) {
	tests := []struct {
		input    string
		expected int32
		wantErr  bool
	}{
		{"123", 123, false},
		{"-123", -123, false},
		{"abc", 0, true},
	}

	for _, tt := range tests {
		result, err := ParseInt32(tt.input)
		if (err != nil) != tt.wantErr {
			t.Errorf("ParseInt32(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			continue
		}
		if result != tt.expected {
			t.Errorf("ParseInt32(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestParseUint(t *testing.T) {
	tests := []struct {
		input    string
		expected uint
		wantErr  bool
	}{
		{"123", 123, false},
		{"-123", 0, true},
		{"abc", 0, true},
	}

	for _, tt := range tests {
		result, err := ParseUint(tt.input)
		if (err != nil) != tt.wantErr {
			t.Errorf("ParseUint(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			continue
		}
		if result != tt.expected {
			t.Errorf("ParseUint(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestParseFloat64(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
		wantErr  bool
	}{
		{"123.456", 123.456, false},
		{"123", 123, false},
		{"abc", 0, true},
	}

	for _, tt := range tests {
		result, err := ParseFloat64(tt.input)
		if (err != nil) != tt.wantErr {
			t.Errorf("ParseFloat64(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			continue
		}
		if result != tt.expected {
			t.Errorf("ParseFloat64(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}
