package p2


func Product_all_expect_at_index(vals []int) []int{

	var result = make([]int,len(vals))
	result[0] = 1

	for i := 1; i < len(vals) ; i++{
		result[i] = result[i-1] * vals[i-1]
	}

	right_product := 1
	for i:= len(vals) - 1; i >= 0; i--{
		result[i] *=  right_product
		right_product *= vals[i]
	}

	return result
}