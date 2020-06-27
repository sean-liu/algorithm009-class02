package solutions

// https://leetcode-cn.com/problems/decode-ways/

// dp approach, status array dp[i] represents number of decoding way of s[i:]
func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}

	dp := make([]int, 3)
	dp[len(s)%3] = 1 // used for sum in it's combined case
	if s[len(s)-1] != '0' {
		dp[(len(s)-1)%3] = 1
	}

	for i := len(s) - 2; i >= 0; i-- {
		firstDigit, secondDigit := s[i]-'0', s[i+1]-'0'
		if firstDigit == 0 {
			dp[i%3] = 0
		} else {
			combinedIndex := firstDigit*10 + secondDigit
			dp[i%3] = dp[(i+1)%3]
			if combinedIndex <= 26 {
				dp[i%3] += dp[(i+2)%3]
			}
		}
	}

	return dp[0]
}
