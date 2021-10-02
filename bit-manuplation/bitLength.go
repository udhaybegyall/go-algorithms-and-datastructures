package bitmanuplation

func bitLength(number int32) int32 {

	var bitCount uint32
	bitCount = 0

	for 1<<bitCount <= number {
		bitCount += 1
	}
	return int32(bitCount)
}
