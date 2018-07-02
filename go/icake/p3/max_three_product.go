package p3

func Get_max_three_product(values []int) int {

	if len(values) < 2{
		return 0
	}

	max_value := max(values[0], values[1])
	min_value := min(values[0], values[1])

	max_product_two := values[0] * values[1]
	min_product_two := values[0] * values[1]

	result := max_product_two * values[2]

	for i := 2; i < len(values); i++{
		val := values[i]
		result = max(result, max(max_product_two*val, min_product_two*val))

		max_product_two = max(max_product_two, max(val*max_value, val*min_value))
		min_product_two = min(min_product_two, min(val*max_value, val*min_value))

		max_value = max(max_value, val)
		min_value = min(min_value, val)

		}

		return result

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


func max(a, b int) int{
	if a > b{
		return a
	}

	return b
}