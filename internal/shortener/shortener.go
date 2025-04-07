package shortener

import (
	"fmt"
	"math"
)

func ShortenNumber(n int) string {
	num := float64(n)

	if math.Abs(num) >= 1000000000 {
		num /= 1000000000
		return fmt.Sprintf("%.1f bil", num)
	} else if math.Abs(num) >= 1000000 {
		num /= 1000000
		return fmt.Sprintf("%.1f mil", num)
	} else if math.Abs(num) >= 1000 {
		num /= 1000
		return fmt.Sprintf("%.1f k", num)
	}

	return fmt.Sprint(num)
}
