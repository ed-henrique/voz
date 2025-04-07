package shortener

import (
	"fmt"
	"testing"
)

func TestShortenNumber(t *testing.T) {
	type test struct {
		n        int
		expected string
	}

	tt := []test{
		{100, "100"},
		{1000, "1.0 k"},
		{1352, "1.4 k"},
		{723849, "723.8 k"},
		{1238097459, "1.2 bil"},
	}

	for _, it := range tt {
		t.Run(fmt.Sprintf("check for %d", it.n), func(t *testing.T) {
			got := ShortenNumber(it.n)

			if got != it.expected {
				t.Errorf("got %s expected %s", got, it.expected)
			}
		})
	}
}
