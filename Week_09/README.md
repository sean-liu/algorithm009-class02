学习笔记
# 9th week


## 第19课 高级动态规划

- 1. 动态规划、状态转移方程串讲

	- 递归、分治、回溯、动态规划复习

		- 分治

			- 感触

				- 1. 人肉递归不合理

				- 2. 找到最近最简方法，将其拆解成可重复解决的问题

				- 3. 数学归纳法思维

			- 本质

				- 寻找重复性->计算机指令集

		- 动态规划

			- 1. simplifying a complicated problem by breaking it down into simpler sub-problems in a recursive manner

			- 2. Divide & conquer + optimal substructure

			- 3. 顺推形式：动态递推，一般可以从下往上推

			- 复杂点

				- 怎么定义状态

				- 状态方程怎么得出

					- 1. 可以是相加

					- 2. min或者max，一些值

					- 3. 或者累加累减

					- 4. 小循环之前k个值

		- 关键点

			- 动态规划 和 递归或者分治 没有根本上区别(关键看有无最优的子结构)

			- 拥有共性

				- 找重复子问题

			- 差异性

				- 最优子结构、中途可以淘汰次优解

	- 常见的DP题目和状态方程

		- [70. 爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)

			阿里，腾讯，字节跳动，202005半年常考

			- 和硬币置换一般差不多

			- 都是fabnacci的变种

			- Dp公式

				- f(n) = f(n-1) + f(n-2), f(1) = 1, f(0) = 0

			- 空间优化

				- 可以到2个变量，参见70题最优解

		- [62. 不同路径](https://leetcode-cn.com/problems/unique-paths/)

			Facebook、亚马逊、微软在半年内面试中考过

			- Dp公式

				- f(x, y) = f(x-1, y) + f(x, y-1)

			- 空间优化

				- 可以到1维dp方程

		- [198. 打家劫舍](https://leetcode-cn.com/problems/house-robber/)

			字节跳动、谷歌、亚马逊在半年内面试中考过

			- 限制条件

				- 相邻的不能打

			- 方案一

				- dp[i]状态定义

					- max $ of robbing A[0 -> i] && nums[i]

				- dp公式

					- dp[i] = max(dp[i-2] + nums[i], dp[i-1])

			- 方案二

				- dp[i][0]状态定义

					- max $ of robbing A[0 -> i] && !nums[i]

				- dp[i][1]状态定义

					- max $ of robbing A[0 -> i] && nums[i]

				- 上面的状态定义方式，在高级dp中会经常会用到

				- dp公式

					- dp[i][0] = max(dp[i-1][0], dp[i-1][1])

					- dp[i][1] = dp[i-1][0] + nums[i]

		- [64. 最小路径和](https://leetcode-cn.com/problems/minimum-path-sum/)

			- dp[i][j]状态定义

				- minPath(A[1->i][1->j])

			- dp公式

				- dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + A[i][j]

		- 买卖股票

			- 参考链接(必看)

				- [一个方法团灭 6 道股票问题](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/solution/yi-ge-fang-fa-tuan-mie-6-dao-gu-piao-wen-ti-by-l-3/)

			- dp[i][k][0 or 1] (0<=i<=n-1, 1<=k<=K)

				- i为天数

				- k为最多交易次数

				- [0, 1]为是否持有股票

				- 总状态数

					- n*K*2种状态

			- dp模版

				- for 0 <= i < n

				- for 1 <= k <= K

				- for s in {0, 1}

				- dp[i][k][s] = max(buy, sell, rest)

			- dp公式

				- dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])

					- max(选择rest, 选择sell)

					- 解释：今天我没有持有股票

						- 我昨天就没有持有，然后今天选择rest，所以我今天还是没有持有

						- 我昨天持有股票，但是我今天sell了，所以我今天没有持有股票了

				- dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])

					- max(选择rest, 选择buy)

					- 解释：今天我持有股票

						- 我昨天就持有股票，然后今天选择rest，所以我今天还持有股票

						- 我昨天本来没有持有，但我今天选择buy，所以我今天就持有股票了

			- 初始状态

				- dp[-1][k][0] = dp[i][0][0] = 0

				- dp[-1][k][1] = dp[i][0][1] = -infinity

- 2. 高级动态规划题目详解

	- 复杂度来源

		- 1. 状态拥有更多维度(2维、3维、或者更多、甚至需要压缩)

		- 2. 状态方程更加复杂

	- 本质

		- 逻辑思维

	- 爬楼梯问题改进

		- 变种

			- 可以走，1步2步3步

				- Dp 方程 多一项，变为：a[I] = a[i-1]+a[i-2]+a[i-3]

			- 可以走的步的种类x1，x2，x3

				- 双循环，第二层循环这个步伐数组，用累加做，类似a[i] += a[i-x[j]]

			- 前后不能走相同的步伐

				- 需要二维a[i][k]，其中k代表a[i]的最后一步是k步

		- 基本代码

			if n <= 1 {  
				return n  
			}  
			  
			a := make([]int, n)  
			  
			a[0], a[1] = 1, 2  
			  
			for i:=2; i < n; I++ {  
				a[i] = a[i-1] + a[i-2]  
			}  
			  
			return a[n-1]

	- [72. 编辑距离](https://leetcode-cn.com/problems/edit-distance/)

		字节跳动、亚马逊、谷歌在半年内面试中考过

		- 关键点

		- 可能解法

			- 1. 双端BFS

				- 每一层扩展变化一次能到的单词，需要相关联的变化，不是任意变化

			- 2. dp

				- dp[i][j]为 word1[0:i] 与 word2[0:j]的编辑距离

				- 这样几种情况

					- 1. If w1[i] == w2[j]

						- dp[i][j] = dp[i-1][j-1]

					- 2. If w1[i] != w2[j]

						- dp[i][j] = min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1

					- 3. If w1 == nil

						- edit_dist = len(w2)

					- 4. If w2 == nil

						- edit_dist = len(w1)

				- 官方题解的话，图不错

## 第20课 字符串算法

- 1. 字符串基础知识和引申题目

	- 可变与不可变

		- 有的语言是可变的，需要留意一下

	- 比较的时候，有些语言是比较reference，需要注意

	- [387. 字符串中的第一个唯一字符](https://leetcode-cn.com/problems/first-unique-character-in-a-string/)

		亚马逊、微软、Facebook 在半年内面试中考过

		- 可能解法

			- 1. 暴力

				- 循环i所有char，然后循环j所有i以后的字符串

				- O(n^2)

			- 2. hashmap

				- Time O(n)

				- 过2遍，第一遍统计出现次数，第二遍找到第一个char，次数是1的

	- [8. 字符串转换整数 (atoi)](https://leetcode-cn.com/problems/string-to-integer-atoi/)

		亚马逊、微软、Facebook 在半年内面试中考过

		- 关键点

			- 想清楚每一步要干什么

	- 参考链接

		- [不可变字符串](https://lemire.me/blog/2017/07/07/are-your-strings-immutable/)

		- [Atoi 代码示例](https://shimo.im/docs/5kykuLmt7a4DdjSP/read)

	- [14. 最长公共前缀](https://leetcode-cn.com/problems/longest-common-prefix/)

		亚马逊、谷歌、Facebook 在半年内面试中考过

		- 可能算法

			- 1.  暴力

				- 从最短的单词开始，一位一位的比对前缀

				- O(n^2*m)，m是单词平均长度

			- 2. 把单词平列，然后一位一位看，如果有某个单词的前缀不同，那就确定了，前面是一样的

				- O(n^2)

			- 3. trie树

	- [344. 反转字符串](https://leetcode-cn.com/problems/reverse-string/)

		亚马逊、谷歌、苹果在半年内面试中考过

		- 关键点

			- 写熟双指针前后循环方式

			- 双层嵌套循环，写熟

	- [151. 翻转字符串里的单词](https://leetcode-cn.com/problems/reverse-words-in-a-string/)

		微软、字节跳动、苹果在半年内面试中考过

		- 可能解法

			- 1. split，reverse，join

			- 2. reverse 所有的char，然后再翻转每个单词

	- [438. 找到字符串中所有字母异位词](https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/)

		Facebook 在半年内面试中常考

		- 关键点

			- 有个大小为目标字符串的窗口，然后平移

	- [5. 最长回文子串](https://leetcode-cn.com/problems/longest-palindromic-substring/)

		亚马逊、字节跳动、华为在半年内面试中常考

		- 自己写

- 2. 高级字符串算法

## 作业

- easy

	- [387. 字符串中的第一个唯一字符](https://leetcode-cn.com/problems/first-unique-character-in-a-string/)

		亚马逊、微软、Facebook 在半年内面试中考过

		- [x] 20200719

		- [ ] 20200720

		- [ ] 20200726

	- [541. 反转字符串 II](https://leetcode-cn.com/problems/reverse-string-ii/)

		亚马逊在半年内面试中考过

	- [557. 反转字符串中的单词 III](https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/)

		微软、字节跳动、华为在半年内面试中考过

	- [917. 仅仅反转字母](https://leetcode-cn.com/problems/reverse-only-letters/)

		字节跳动在半年内面试中考过

	- [205. 同构字符串](https://leetcode-cn.com/problems/isomorphic-strings/)

		谷歌、亚马逊、微软在半年内面试中考过

	- [680. 验证回文字符串 Ⅱ](https://leetcode-cn.com/problems/valid-palindrome-ii/)

		Facebook 在半年内面试中常考

	- [746. 使用最小花费爬楼梯](https://leetcode-cn.com/problems/min-cost-climbing-stairs/)

		亚马逊在半年内面试中考过

- medium

	- 在学习总结中，写出[不同路径 2](https://leetcode-cn.com/problems/unique-paths-ii/) 这道题目的状态转移方程。

	- [151. 翻转字符串里的单词](https://leetcode-cn.com/problems/reverse-words-in-a-string/)

		微软、字节跳动、苹果在半年内面试中考过

	- [300. 最长上升子序列](https://leetcode-cn.com/problems/longest-increasing-subsequence/)

		字节跳动、亚马逊、微软在半年内面试中考过

	- [x] [91. 解码方法](https://leetcode-cn.com/problems/decode-ways/)

		字节跳动、亚马逊、Facebook 在半年内面试中考过

	- [8. 字符串转换整数 (atoi)](https://leetcode-cn.com/problems/string-to-integer-atoi/)

		亚马逊、微软、Facebook 在半年内面试中考过

	- [438. 找到字符串中所有字母异位词](https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/)

		Facebook 在半年内面试中常考

	- [5. 最长回文子串](https://leetcode-cn.com/problems/longest-palindromic-substring/)

		亚马逊、字节跳动、华为在半年内面试中常考

- hard

	- [32. 最长有效括号](https://leetcode-cn.com/problems/longest-valid-parentheses/)

		亚马逊、字节跳动、华为在半年内面试中考过

	- [818. 赛车](https://leetcode-cn.com/problems/race-car/)

		谷歌在半年内面试中考过

	- [44. 通配符匹配](https://leetcode-cn.com/problems/wildcard-matching/)

		Facebook、微软、字节跳动在半年内面试中考过

	- [115. 不同的子序列](https://leetcode-cn.com/problems/distinct-subsequences/)

		MathWorks 在半年内面试中考过

	- [85. 最大矩形](https://leetcode-cn.com/problems/maximal-rectangle/)

