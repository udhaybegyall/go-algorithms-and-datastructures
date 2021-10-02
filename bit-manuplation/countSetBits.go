package bitmanuplation

func countSetBits(number int) int {

	count := 0
	temp_number := number

	for temp_number > 0 {

		count += temp_number & 1
		temp_number = temp_number >> 1
	}
	return count
}
