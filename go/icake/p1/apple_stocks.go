package p1



func Best_price(prices []int) int {

	if len(prices) < 2 {
		return 0
	}

	min_price := prices[0]
	max_profit := prices[1] - prices[0]

	for _,val := range prices{

		if val - min_price > max_profit{
			max_profit = val - min_price
		}

		if val < min_price {
			min_price = val
		}
	}

	return max_profit
}