package best_time_to_buy_and_sell_stock

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

//func maxProfit(prices []int) int {
//	maxProfit := 0
//	min := math.MaxInt64
//	maxIndex := 0
//
//	for i := 0; i < len(prices)-1; i++ {
//		if prices[i] < min {
//			min = prices[i]
//		} else {
//			continue
//		}
//		if i >= maxIndex {
//			maxIndex = i + 1
//			for j := i + 1; j < len(prices); j++ {
//				if prices[j] >= prices[maxIndex] {
//					maxIndex = j
//				}
//			}
//		}
//		if tmp := prices[maxIndex] - min; tmp > maxProfit {
//			maxProfit = tmp
//		}
//	}
//
//	return maxProfit
//}

func maxProfit(prices []int) int {
	maxProfit := 0
	min := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			maxProfit = max(maxProfit, prices[i]-min)
		}
	}

	return maxProfit
}
