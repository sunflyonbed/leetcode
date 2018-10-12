package leetcode

func isMatch(s string, p string) (bool, [][]bool) {
	byteS := []byte(s)
	byteP := []byte(p)
	sizeS := len(byteS)
	sizeP := len(byteP)
	if sizeS == 0 && sizeP == 0 {
		return true, nil
	}
	//dp[i][j] represents whether the substring of s with length i from right can be matched with the substring of p with length j from right.
	dp := make([][]bool, sizeS+1)
	dp[0] = make([]bool, sizeP+1)
	dp[0][0] = true
	for i := 0; i < sizeP; i++ {
		if byteP[i] == '*' && dp[0][i-1] {
			dp[0][i+1] = true
		}
	}
	for i := 0; i < sizeS; i++ {
		if dp[i+1] == nil {
			dp[i+1] = make([]bool, sizeP+1)
		}
		for j := 0; j < sizeP; j++ {
			if byteP[j] == '.' {
				dp[i+1][j+1] = dp[i][j]
			}
			if byteP[j] == byteS[i] {
				dp[i+1][j+1] = dp[i][j]
			}
			if byteP[j] == '*' {
				if byteP[j-1] != byteS[i] && byteP[j-1] != '.' {
					dp[i+1][j+1] = dp[i+1][j-1]
				} else {
					dp[i+1][j+1] = dp[i+1][j] || dp[i][j+1] || dp[i+1][j-1]
				}
			}
		}
	}
	return dp[sizeS][sizeP], dp
}
