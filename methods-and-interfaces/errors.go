package main

import (
	"fmt"
	"math"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

type ErrNegativeSqrt float64

func (e MyError) Error() string {
	return fmt.Sprintf("at %v. %s", e.When, e.What)
}

func run() error {
	return &MyError{time.Now(), "an error throwed !"}
}

func errors() {
	fmt.Println("\n-- errors --")

	err := run()
	if err != nil {
		fmt.Println(err)
	}
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number:", float64(e))
}

// error exercise
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := x / 2
	const delta = 0.0000001
	i := 0
	for {
		prevZ := z
		z -= (z*z - x) / (2 * z)
		i++
		if math.Abs(prevZ-z) < delta {
			break
		}
	}

	return z, nil
}
