package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	for _, value := range input {
		result += value
	}
	result = result / float32(len(input))
	return result
}
