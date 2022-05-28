package homework

func average(input [15]float32) (result float32) {
	for _, s := range input {
		result += s
	}
	return result / float32(len(input))
}
