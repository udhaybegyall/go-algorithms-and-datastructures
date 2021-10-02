package bitmanuplation

// returns the number of set bits in a binay representation of a number.
// or number of '1's in a binay representation of a number.
// Example: '1101' repreents 13 and it has setBits/'1's = 3.
func countSetBits(number int) int {

	count := 0
	temp_number := number

	for temp_number > 0 {

		count += temp_number & 1
		temp_number = temp_number >> 1
	}
	return count
}
