package solutions

// https://leetcode-cn.com/problems/decode-ways/

// dp approach, status array dp[i] represents number of decoding way of s[i:]
func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[len(s)] = 1
	if s[len(s)-1] != '0' {
		dp[len(s)-1] = 1
	}

	for i := len(s) - 2; i >= 0; i-- {
		first := s[i] - '0'
		second := s[i+1] - '0'
		if first != 0 {
			dp[i] = dp[i+1]
			if combinedIndex := first*10 + second; combinedIndex <= 26 {
				dp[i] += dp[i+2]
			}
		}
	}

	return dp[0]
}
