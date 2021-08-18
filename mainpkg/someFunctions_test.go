package mainpkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGcd(t *testing.T) {
	testTable := []struct {
		num1     int
		num2     int
		expected int
	}{
		{0, 12, 12},
		{3, 0, 3},
		{3, 12, 3},
		{25, 85, 5},
		{24, 32, 8},
	}

	for _, test := range testTable {
		assert.Equal(t, test.expected, gcd(test.num1, test.num2))
	}
}
