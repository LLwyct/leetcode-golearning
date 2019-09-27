// 定义 p(i,j)为串s中的i到j是否为回文串
// 则有 p(i,j) == is(p(i+1, j-1) && s[i] == s[j])
// dp即可，注意要预处理长度为1和2的回文串，因为它们不受状态公式的影响

var dp [1007][1007]bool;
func longestPalindrome(s string) string {
    dp = [1007][1007]bool{};
    var l int = len(s);
	var mini, maxj int = 0, 0;
	// 初始化长度为1的串
    for i:=0;i<l;i++ {
        dp[i][i] = true;
	}
	// 初始化长度为2的串
    for i:=0;i<l-1;i++ {
        if s[i] == s[i + 1] {
            dp[i][i+1] = true;
            mini, maxj = i, i+1;
        }
	}
	// dp
    if l>=3 {
        for i:=l-2;i>=0;i-- {
            for j:=i+2;j<l;j++ {
                if dp[i+1][j-1] == true && s[i] == s[j] {
                    dp[i][j] = true;
                    if j - i > maxj - mini {
                        mini, maxj = i, j;
                    }
                }
            }
        }
    } else {
        if l == 2 && s[0] == s[1] {
            return s;
        } else if l == 2 {
            return string(s[0]);
        } else if l == 1 {
            return s;
        } else {
            return "";
        }
    }
    return s[mini:maxj+1]
}