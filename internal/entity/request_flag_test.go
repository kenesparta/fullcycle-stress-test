package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestFlag_Validate(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected bool
	}{
		{
			name:     "http://example.com",
			url:      "http://example.com",
			expected: true,
		},
		{
			name:     "https://example.com",
			url:      "https://example.com",
			expected: true,
		},
		{
			name:     "ftp://example.com",
			url:      "ftp://example.com",
			expected: false,
		},
		{
			name:     "://example.com",
			url:      "://example.com",
			expected: false,
		},
		{
			name:     "http://",
			url:      "http://",
			expected: false,
		},
		{
			name:     "http://a",
			url:      "http://a",
			expected: true,
		},
		{
			name:     "http://example.com?id=1234&token=abcde",
			url:      "http://example.com?id=1234&token=abcde",
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, isValidURL(tc.url))
		})
	}
}
