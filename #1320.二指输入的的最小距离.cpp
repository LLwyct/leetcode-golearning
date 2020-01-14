dp,[k][i][j]分别代表，处理完第k个字符后，左手在i，右手在j的最小路径

状态转移方程[k][i][j]可能是由上一次左手在word[i-1]，转移而来，也可能是由上一次右手在word[i-1]转移而来
dp[k][i][j] = min(dp[k-1][i][pre] + distance(pre, j), dp[k-1][pre][j] + distance(pre, i));

class Solution {
public:
    int minimumDistance(string word) {
        const int INF = 1e7;
        const int wordLength = word.size();
        int dp[303][26][26] = {0};
        for (int i=0;i<26;++i) {
            for (int j=0;j<26;++j) {
                dp[0][i][j] = 0;
            }
        }
        for (int k=1;k<word.size();++k) {
            int pre = word[k-1] - 'A';
            for (int i=0;i<26;++i) {
                for (int j=0;j<26;++j) {
                    dp[k][i][j] = min(dp[k-1][i][pre] + distance(pre, j), dp[k-1][pre][j] + distance(pre, i));
                }
            }
        }
        int ans = INF;
        int endWord = word[wordLength-1] - 'A';
        for (int i=0;i<26;++i) {
            if (dp[wordLength - 1][endWord][i] <= ans) {
                ans = dp[wordLength - 1][endWord][i];
            }
            if (dp[wordLength - 1][i][endWord] <= ans) {
                ans = dp[wordLength - 1][i][endWord];
            }
        }
        return ans;
    }
    int distance (int a, int b) {
        int a_row = a / 6;
        int a_col = a % 6;
        int b_row = b / 6;
        int b_col = b % 6;
        return abs(a_col - b_col) + abs(a_row - b_row);
    }
};
