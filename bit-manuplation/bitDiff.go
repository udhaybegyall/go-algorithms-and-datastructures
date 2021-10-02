package bitmanuplation

// return the number of bits that needs to be changed
// so that number one becomes number two
func bitDiff(number_one int, number_two int) int {
	return countSetBits(number_one ^ number_two)
}
