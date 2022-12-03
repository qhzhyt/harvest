package utils

func Map[I interface{}, O interface{}](inputs []I, mapFunc func(I) O) []O {
	outputs := make([]O, len(inputs))
	for i, input := range inputs {
		outputs[i] = mapFunc(input)
	}
	return outputs
}
