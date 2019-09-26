由于是子串，说明串必须是连续的，因此，可以用滑动窗口，设窗口左端为s，右端为e
先枚举以s=0起始的所有子串，当子串有重复时，无论e怎么扩大，后面一定重复，因此就要移动s
但是，问题求最长，因此，直接从上一次的最大长度开始用新的s枚举

我这里二分答案，也可以

func lengthOfLongestSubstring(s string) int {
    var length, low, high, maxok int;
    var ss []rune = []rune(s)
    low = 1;
    high = 97;
    maxok = 0;
    for {
        if high < low {
            break;
        }
        var islengthok bool = false;
        length = (low + high) / 2;
        for i:=0;i<=len(s)-length;i++ {
            if length > len(s) {
                break;
            }
            var cur [140]int = [140]int{};
            var isloopok bool = true;
            for j:=i;j<i+length;j++ { 
                if cur[ss[j] - 0] == 1 {
                    isloopok = false;
                    break;
                } else {
                    cur[ss[j] - 0] = 1;
                }
            }
            if isloopok == true {
                maxok = length;
                islengthok = true;
                low = length + 1;
                break;
            }
        }
        if islengthok == false {
            high = length - 1;
        }
    }
    return maxok;
}