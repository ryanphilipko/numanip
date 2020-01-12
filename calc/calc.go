package calc

//Add returns sum of 2 integers
func Add(numbers ...int) int {
	sum := 0
	for _, x := range numbers {
		sum += x
	}
	return sum
}
