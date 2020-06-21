package solutions

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

// // loop for once, pick each time's peak and valley and get sum of profit.
// func maxProfit(prices []int) int {

// 	maxProfit := 0
// 	for i, peak, valley := 0, 0, 0; i < len(prices)-1; {
// 		for ; i < len(prices)-1 && prices[i] >= prices[i+1]; i++ {
// 		}
// 		valley = prices[i]
// 		for ; i < len(prices)-1 && prices[i] <= prices[i+1]; i++ {
// 		}
// 		peak = prices[i]
// 		maxProfit += peak - valley
// 	}

// 	return maxProfit
// }

// loop for once, based on previous solution, focus on compare current and previous value to decide whether to sum or not.
func maxProfit(prices []int) int {

	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}

	return maxProfit
}
