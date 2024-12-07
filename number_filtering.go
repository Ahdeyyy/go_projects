package exercises

import (
	"math"
)

type Cond = func(int) bool

func Filter_nums(input []int, cond Cond) []int {
	output := make([]int, len(input))
	i := 0
	for _, v := range input {
		if cond(v) {
			output[i] = v
			i++
		}
	}

	return output

}

func IsEven(num int) bool {
	return num%2 == 0
}

func IsOdd(num int) bool {
	return !IsEven(num)
}

func IsPrime(num int) bool {
	n := float64((fact(num-1) + 1) / num)
	return int(math.Mod(n*1000.0, 1000.0)) == 0
}

func IsOddPrime(num int) bool {
	return IsPrime(num) && IsOdd(num)
}

func IsMultipleOf(num int) func(int) bool {
	return func(i int) bool {
		return i%num == 0
	}
}
func FilterByConds(input []int, cond func(int, ...Cond) bool, functions ...func(int) bool) []int {
	output := make([]int, len(input))

	i := 0
	for _, v := range input {
		if cond(v, functions...) {
			output[i] = v
			i++
		}
	}

	return output
}

func All(num int, fns ...Cond) bool {
	for _, fn := range fns {
		if fn(num) == false {
			return false
		}
	}
	return true
}

func Any(num int, fns ...Cond) bool {
	for _, fn := range fns {
		if fn(num) == true {
			return true
		}
	}
	return false
}

func IsGreater(num int) func(int) bool {
	return func(i int) bool {
		return i > i
	}
}

func IsOddGreaterAndMultipleOf(g int, m int) func(int) bool {

	return func(i int) bool {
		return IsOdd(i) && IsGreater(g)(i) && IsMultipleOf(m)(i)
	}
}

func IsEvenMultipleOf(mul int) func(int) bool {
	return func(i int) bool {
		return IsEven(i) && IsMultipleOf(mul)(i)
	}
}

func fact(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	return num * fact(num-1)
}
