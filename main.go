package main

import (
	"fmt"
	"math"
)

func isNegativeNumber(x int) bool {
	if x >= 0 {
		return false
	} else {
		return true
	}
}

func balikanAngka(x int) int {
	x = int(math.Abs(float64(x)))

	var angkaTerbalik int
	for x > 0 {
		sisa := x % 10
		angkaTerbalik = angkaTerbalik*10 + sisa
		x /= 10
	}

	if isNegativeNumber(x) {
		angkaTerbalik = angkaTerbalik * -1
	}

	return angkaTerbalik
}

func isPalindrom(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	xBalikAngka := balikanAngka(x)

	return xBalikAngka == x
}

func bilanganPrima(n int) (prima []int) {
	b := make([]bool, n)
	for i := 2; i < n; i++ {
		if b[i] {
			continue
		}

		prima = append(prima, i)

		for x := i * i; x < n; x += i {
			b[x] = true
		}
	}
	return
}

func main() {
	var pal []int
	prime := bilanganPrima(200)
	for _, v := range prime {
		if isPalindrom(v) {
			pal = append(pal, v)
		}
	}

	fmt.Print(pal)
}
