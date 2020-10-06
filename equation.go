package equation

import (
	"fmt"
	"math/rand"
	"time"
)

type Equation struct {
	a, b, c int
	b24ac   int
	r1, r2  Root
}

type Root struct {
	plus     bool
	up, down int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (e *Equation) GetEquationString() string {
	var result string
	if e.a == 1 {
		result += fmt.Sprintf("x²")
	} else {
		result += fmt.Sprintf("%dx²", e.a)
	}
	if e.b > 0 {
		result += fmt.Sprintf("+%dx", e.b)
	} else {
		result += fmt.Sprintf("%dx", e.b)
	}
	if e.c > 0 {
		result += fmt.Sprintf("+%d", e.c)
	} else {
		result += fmt.Sprintf("%d", e.c)
	}
	result += fmt.Sprintf("=0\n")
	return result
}

func GenerateRandomEquation(Range int) *Equation {
	e := new(Equation)
	if Range < 0 {
		return e
	}
	squr := 2 * (rand.Intn(Range) + 1)
	e.b24ac = squr * squr

	e.b = 2 * (rand.Intn(Range) + 1)
	ac := -((e.b24ac - (e.b * e.b)) / 4)
	acSlice := shuffle(getPrimeNumberSlice(ac))
	aElementCount := rand.Intn(len(acSlice) + 1)
	e.a = 1
	for i := 0; i < aElementCount; i++ {
		e.a = e.a * acSlice[i]
	}
	e.c = ac / e.a

	e.r1.up = ((-e.b) + squr)
	e.r1.down = (2 * e.a)
	if (e.r1.up >= 0) == (e.r1.down >= 0) {
		e.r1.plus = true
	}
	e.r2.up = ((-e.b) - squr)
	e.r2.down = (2 * e.a)
	if (e.r2.up >= 0) == (e.r2.down >= 0) {
		e.r2.plus = true
	}
	return e
}

func getPrimeNumberSlice(input int) []int {
	if input < 0 {
		input = -input
	}
	result := []int{}
	max := input
	for i := 2; i < max; i++ {
		if input%i == 0 {
			input = input / i
			result = append(result, i)
			i--
		}
	}
	return result
}

func shuffle(vals []int) []int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]int, len(vals))
	n := len(vals)
	for i := 0; i < n; i++ {
		randIndex := r.Intn(len(vals))
		ret[i] = vals[randIndex]
		vals = append(vals[:randIndex], vals[randIndex+1:]...)
	}
	return ret
}
