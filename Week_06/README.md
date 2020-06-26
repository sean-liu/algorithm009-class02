学习笔记
# 第六周


## 第十二课 动态规划

- 1. 动态计划的实现及关注点

	- 关于递归、分治的感触

		- 感触

			- 1. 人肉递归低效、很累

			- 2. 找到最近最简单方法，将其分解成可重复解决的问题

			- 3. 数学归纳法思维(抵制人肉递归的诱惑)

		- 本质

			- 因为计算机的指令集的设计方式，所以归结到底就是寻找重复性

	- 动态规划

		- 定义

			- 1. 本质上就是一个递归或者分治问题

			- 2. 加一个最优子结构

		- 与递归、分治的区别

			- 共性

				- 找到重复子问题

			- 差异性

				- 最优子结构、中途可以淘汰次优解

		- 特点

			- 复杂度相对来说比较低，因为淘汰了次优解的缘故。

		- 还有个别名：动态递推

- 2. DP例题解析：Fibonacci数列、路径计数

	- 1. Fibonacci数列

		- 正常做法

			- 递归，此时代码可以往简洁方向优化，对执行效率来说，不会有影响，时间复杂度O(2^n)

		- 优化方式

			- 加缓存，复杂度降为O(n)

			- 自底向上，写for循环

	- [63. 不同路径 II](https://leetcode-cn.com/problems/unique-paths-ii/)

		谷歌、美团、微软在半年内面试中考过

		- 正常做法

			- 递归，由于只能往右和往下走，所以从start开始的下一步是由右边的那个点和下面的那个点组成。

		- Dp方式

			- 从终点开始倒推，最下面一排和最右边1排都是1，然后开始，往左上推，每一个点等于它的右边、下边的格子的数值的汇总

	- 关键点

		- 1. 最优子结构

			- opt[n] = best_of(opt[n-1], opt[n-2], ...)

		- 2. 存储中间状态

			- opt[I]

			- 这个是重点

		- 3. 递推公式(状态方程、DP方程)

			- Fib: opt[i] = opt[i-1] + opt[i-2]

			- 路径计数: opt[i,j] = opt[i+1][j] + opt[i][j+1](and check if a[i,j] 是否为空地)

			- 一般DP方程比较难找

- 3. DP例题解析：最长公共子序列

	- [1143. 最长公共子序列](https://leetcode-cn.com/problems/longest-common-subsequence/)

		字节跳动、谷歌、亚马逊在半年内面试中考过

		- 思考过程

			- 1. 思考S1=‘’， S2 = 任意，答案是0

			- 2. 思考s1=‘A’，s2=任意，答案是0，或者1(当A在s2中)

			- 3. 思考s1=‘。。。A’，s2=‘。。A’，答案是s1的子序列与s2的子序列的最长公共子序列+1

		- Dp公式

			- 1. 当s1[n-1] != s2[n-1]

				lcs[s1, s2] = max(lcs[s1-1,  s2], lcs[s1, s2-1])  
				  
				或  
				  
				lcs[s1, s2] = max(lcs[s1-1, s2], lcs[s1, s2-1], lcs[s1-1, s2-1])

			- 2. 当s1[n-1] == s2[n-1]

				lcs[s1, s2] = lcs[s1-1, s2-1] +1  
				  
				或  
				  
				Lcs[s1, s2] = max(lcs[s1-1, s2], lcs[s1, s2-1], lcs[s1-1, s2-1], lcs[s1-1, s2-1] + 1)

			- 3. 汇总

				If s1[n-1] == s2[n-1]  
					lcs[s1, s2] = lcs[s1-1, s2-1] + 1  
				If s1[n-1] != s2[n-1]  
					lcs[s1, s2] = max(lcs[s1-1, s2], lcs[s1, s2-1])

		- 关键点

			- 1. 会变成二位数组

			- 2. 下标要定义清楚

	- 动态规划小结

		- 1. 打破思维惯性，形成机器思维

		- 2. 理解复杂逻辑的关键

		- 3. 也是职业进阶的要领

	- 参考链接

		- [MIT 动态规划课程最短路径算法](https://www.bilibili.com/video/av53233912?from=search&seid=2847395688604491997)

- 4. 实战题目解析：三角形最小路径和

	- [70. 爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)

		阿里，腾讯，字节跳动，202005半年常考

		- 变种

			- 1. 可以走3种步伐，1、2、3

			- 2. 相邻2种步伐不能一样

	- [120. 三角形最小路径和](https://leetcode-cn.com/problems/triangle/)

		亚马逊、苹果、字节跳动在半年内面试考过

		- 可能解法

			- 1. 暴力，把所有可能的解法都递归跑一下，O(2^n)

			- 2. DP(bottom up)

				- A. 分治(子问题)

					- problem(I, j) = min(sub(I+1, j), sub(I+1, j+1)) + a[I,j]

				- B. 定义状态数组

					- f[i,j]

				- C. DP方程

					- f[i,j] = min(f[i+1, j], f[i+1, j+1]) + a[i, j]

		- [高票回答](https://leetcode.com/problems/triangle/discuss/38735/Python-easy-to-understand-solutions-(top-down-bottom-up)

- 5. 实战题目解析：最大子序列和

	- [53. 最大子序和](https://leetcode-cn.com/problems/maximum-subarray/)

		亚马逊、字节跳动在半年内面试常考

		- 可能解法

			- 1. 暴力，枚举所有开始和结束点，然后逐一求解，O(n^3)

			- 2. DP

				- A. 分治(子问题)

					- max_sum(i) = max(max_sum(i-1), 0) + a[I]

				- B. 状态数组定义

					- f[I]

				- C. DP方程

					- f[I] = Max(f[I-1], 0) + a[I]

	- [322. 零钱兑换](https://leetcode-cn.com/problems/coin-change/)

		亚马逊在半年内面试中常考

		- 可能解法

			- 1. 暴力，递归，画出所有状态

				- 状态树中，节点为0的，就是一个可能解，然后找level最低的，就是用硬币最少的

			- 2. BFS

			- 3. DP

				- A. 分治(子问题)

					- problem(i)= min(sub(i-k), k in(1, 2, 5)) +1

				- B. DP Array

					- f

				- C. DP 方程

					- f(n) = min(f(n-k), k in (1, 2, 5)) + 1

		- 看一下官方题解

	- 切记

		- 勿人肉递归，凭感觉，会吃亏

- 6. 实战题目解析：打架劫舍

	- [198. 打家劫舍](https://leetcode-cn.com/problems/house-robber/)

		字节跳动、谷歌、亚马逊在半年内面试中考过

		- DP演化

			- 1. 初版

				- 1. 状态定义

					- a[i] = 0...i 能偷到的最大值

					- 最终结果是a[i-1]

				- 有个问题

					- 不确定a[i] 是否被偷过，出不了dp方程

			- 2. 第二版

				- 1. 状态定义

					- a[i][0] = 0...i 不偷i时，能偷到的最大值

					- a[i][1] = 0...i 偷I时，能偷到的最大值

				- 2. DP方程

					- a[i][0] = max(a[i-1][0], a[i-1][1])

					- a[i][1] = max(a[i-1][0], 0) + nums[i]

				- 3. 最终答案

					- max(a[i][0], a[i][1])

			- 3. 第三版

				- 优化点

					- 尝试消掉状态数组中的额外维度

				- 更近一步优化点

					- 通过观察dp方程，发现其实只用当前最大，前一个最大，外加再以前一个最大就够了，可以优化空间复杂度

				- 1. 状态定义

					- a[i] = 0...i 必偷i的最大值

				- 2. DP方程

					- a[i] = max(a[i-1], a[i-2] + nums[i])

				- 3. 最终答案

					- max(a[i])

					- 所以在迭代的时候需要每一次用a[i]和res比较

	- 技巧

		- 当状态定义1维搞不定的时候，升维。

		- 一般2维dp是入门级的

		- 高级的需要3-4维dp

		- 状态空间定义看道行

		- 一般面试的时候，dp方程会很简单

		- 竞赛的时候，dp方程也很难

	- 参考链接

		- [一个方法团灭 6 道股票问题](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/solution/yi-ge-fang-fa-tuan-mie-6-dao-gu-piao-wen-ti-by-l-3/)

## 作业

- medium

	- [64. 最小路径和](https://leetcode-cn.com/problems/minimum-path-sum/)

	- [91. 解码方法](https://leetcode-cn.com/problems/decode-ways/)

	- [221. 最大正方形](https://leetcode-cn.com/problems/maximal-square/)

	- [621. 任务调度器](https://leetcode-cn.com/problems/task-scheduler/)

	- [647. 回文子串](https://leetcode-cn.com/problems/palindromic-substrings/)

- hard

	- [32. 最长有效括号](https://leetcode-cn.com/problems/longest-valid-parentheses/)

	- [72. 编辑距离](https://leetcode-cn.com/problems/edit-distance/)

		重点

	- [363. 矩形区域不超过 K 的最大数值和](https://leetcode-cn.com/problems/max-sum-of-rectangle-no-larger-than-k/)

	- [403. 青蛙过河](https://leetcode-cn.com/problems/frog-jump/)

	- [410. 分割数组的最大值](https://leetcode-cn.com/problems/split-array-largest-sum/)

	- [552. 学生出勤记录 II](https://leetcode-cn.com/problems/student-attendance-record-ii/)

	- [76. 最小覆盖子串](https://leetcode-cn.com/problems/minimum-window-substring/)

	- [312. 戳气球](https://leetcode-cn.com/problems/burst-balloons/)


