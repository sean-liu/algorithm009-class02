学习笔记
# 第四周


## 第九课 深度优先搜索和广度优先搜索

- 1. 深度优先搜索、广度优先搜索的实现和特性

	- 遍历搜索

		- 特点

			- 1. 每个节点都要访问一次

			- 2. 每个节点仅仅访问一次

			- 3. 访问顺序不限

		- 常见种类

			- 1. 深度优先搜索 (depth first search)

			- 2. 广度优先搜索 (breadth first search)

- 2. 实战题目解析：二叉树的层次遍历等问题

	- [102. 二叉树的层序遍历](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/)

		字节跳动、亚马逊、微软在半年内面试中考过，面试过程中出现频率排第三

		- 可能解法

			- DFS，记录level水平输出

			- BFS

	- [22. 括号生成](https://leetcode-cn.com/problems/generate-parentheses/)

		字节跳动、亚马逊、Facebook 在半年内面试中考过

		- 可能解法

			- DFS(其实就是递归)

	- [200. 岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)

		近半年内，亚马逊在面试中考查此题达到 350 次

		- 可能解法

			- 遍历+DFS找岛的边界

## 第十课 贪心算法

- 1. 贪心的实现、特性及实战题目解析

	- 贪心算法

		- 定义

			- 是一种在每一步选择中最优或者最好的解，从而希望导致结果是全局最好的或最优的算法。

		- 与动态规划、回溯的区别

			- 贪心：当下做局部最优判断，不会回退

			- 回溯：能够回退

			- 动态规划：最优判断+回退

		- 解决范围

			- 一些最优化问题

				- 求图中的最小生成树

				- 霍夫曼编码等

		- 不适用范围

			- 工程类和生活的问题

		- 特点

			- 效率高，耗资源少，接近最优解

		- 使用场景

			- 作为辅助算法，或解决一些不要求很精确的问题

		- 适用场景

			- 问题能拆成子问题，子问题的最优解能递推出最终问题的最优解。

				- 这种子问题的最优解称为最优子结构。

		- 关键点

			- 证明是否能用贪心

			- 以及一些不同的视角

	- 参考链接

		- [动态规划定义](https://zh.wikipedia.org/wiki/%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92)

	- [322. 零钱兑换](https://leetcode-cn.com/problems/coin-change/)

		- 关键点

			- 能用贪心的原因，是因为 20 10 5 1是倍数关系

			- 反例，如果硬币组合是 10 9 1求18就呵呵了

	- [455. 分发饼干](https://leetcode-cn.com/problems/assign-cookies/)

		亚马逊在半年内面试中考过

		- 可以使用原因

			- 如上

	- [122. 买卖股票的最佳时机 II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)

		亚马逊、字节跳动、微软在半年内面试中考过

		- 关键点

			- 模型简化成，只要计算高价和低价的累积差额，就是获利

	- [55. 跳跃游戏](https://leetcode-cn.com/problems/jump-game/)

		亚马逊、华为、Facebook 在半年内面试中考过

		- 关键点

			- 从后往前走贪心，从前往后不太方便，每一步都能有几种可能的情况

## 第十一课 二分查找

- 1. 二分查找的实现、特性及实战题目解析

	- 二分查找

		- 前提

			- 1. 目标函数单调性

			- 2. 存在上下界(bounded)

			- 3. 能通过索引访问

		- 代码模版

			left, right = 0, len(array) - 1  
			while left <= right:  
			   mid = (left + right) / 2  
			   if array[mid] == target:  
			      # find the target!!  
			      break or return result  
			   elif array[mid] < target:  
			      left = mid + 1  
			   else:  
				right = mid - 1

	- 参考链接

		- [Fast InvSqrt()](https://www.beyond3d.com/content/articles/8/)

	- [69. x 的平方根](https://leetcode-cn.com/problems/sqrtx/)

		字节跳动、微软、亚马逊在半年内面试中考过

		- 可能解法

			- 1. 二分查找

				- 1. 先证明可以二分查找

				- 2. 左界为1，右界为X

			- 2. 牛顿迭代法

		- 关键点

			- 考官观察

				- 你思维逻辑是否清晰

				- 代码能力是否强

				- Debug能力是否达标

			- 为什么能用二分查找的原因比较重要

	- [33. 搜索旋转排序数组](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)

		Facebook、字节跳动、亚马逊在半年内面试常考

		- 可能解法

			- 1. 暴力 还原O(logn)->升序->二分(O(logn))(写、总结)

			- 2. 二分查找，判断target是在有序那边还是无序那边

## 作业

- easy

	- [860. 柠檬水找零](https://leetcode-cn.com/problems/lemonade-change/)

		亚马逊在半年内面试中考过

		- [x] 20200612

		- [x] 20200613

		- [ ] 20200619

	- [455. 分发饼干](https://leetcode-cn.com/problems/assign-cookies/)

		亚马逊在半年内面试中考过

		- [x] 20200612

		- [x] 20200613

		- [ ] 20200619

	- [122. 买卖股票的最佳时机 II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)

		亚马逊、字节跳动、微软在半年内面试中考过

	- [874. 模拟行走机器人](https://leetcode-cn.com/problems/walking-robot-simulation/)

- medium

	- [127. 单词接龙](https://leetcode-cn.com/problems/word-ladder/)

		亚马逊在半年内面试常考

	- [200. 岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)

		近半年内，亚马逊在面试中考查此题达到 350 次

		- [x] 20200609

		- [x] 20200610

		- [ ] 20200616

	- 关键点：unionFindSet，在做岛屿数量的时候，挖了1小时的结果

	- [529. 扫雷游戏](https://leetcode-cn.com/problems/minesweeper/)

		亚马逊、Facebook 在半年内面试中考过

		- [x] 20200609

		- [x] 20200610

		- [ ] 20200616

	- [55. 跳跃游戏](https://leetcode-cn.com/problems/jump-game/)

		亚马逊、华为、Facebook 在半年内面试中考过

	- [33. 搜索旋转排序数组](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)

		Facebook、字节跳动、亚马逊在半年内面试常考

	- [74. 搜索二维矩阵](https://leetcode-cn.com/problems/search-a-2d-matrix/)

		亚马逊、微软、Facebook 在半年内面试中考过

	- [153. 寻找旋转排序数组中的最小值](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/)

		亚马逊、微软、字节跳动在半年内面试中考过

- hard

	- [126. 单词接龙 II](https://leetcode-cn.com/problems/word-ladder-ii/)

		微软、亚马逊、Facebook 在半年内面试中考过

	- [45. 跳跃游戏 II](https://leetcode-cn.com/problems/jump-game-ii/)

		亚马逊、华为、字节跳动在半年内面试中考过

- [ ] 使用二分查找，寻找一个半有序数组 [4, 5, 6, 7, 0, 1, 2] 中间无序的地方  
  说明：同学们可以将自己的思路、代码写在学习总结中

