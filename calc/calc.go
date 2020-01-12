package calc

import "errors"

//Add returns sum of integers
func Add(numbers ...int) (error, int) {
	sum := 0

	if len(numbers) < 2 {
		return errors.New("provide more than 1 number"), sum
	} else {
		for _, x := range numbers {
			sum += x
		}
	}
	return nil, sum
}
