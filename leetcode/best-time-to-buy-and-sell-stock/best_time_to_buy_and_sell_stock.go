package best_time_to_buy_and_sell_stock

func maxProfitWithBruceForce(prices []int) int {
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			diff := prices[j] - prices[i]
			if diff > maxProfit {
				maxProfit = diff
			}
		}
	}
	return maxProfit
}

func maxProfitWithDP(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	maxProfit, minPrice := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		diff := prices[i] - minPrice
		if diff > maxProfit {
			maxProfit = diff
		}
		if minPrice > prices[i] {
			minPrice = prices[i]
		}
	}
	return maxProfit
}
