学习笔记
# 第三周


## 第七课 泛型递归、树的递归

- 1. 递归的实现、特性以及思维要点

	- 回顾

		- 树相关的题目为何都是递归？

			- 1. 由于树节点的定义

			- 2. 重复性(自相似性)

	- 递归

		- 定义

			- 通过调用自身的方式来进行循环的方式

		- 结构

			- 1. 退出条件

			- 2. 当层逻辑

			- 3. 下层调用

			- 4. 当层状态清理

		- 关键点

			- 1. 不要在脑中进行模拟递归

			- 2. 找最近最简单方法，总结可重复结构

			- 3. 数学归纳法思维

- 2. 实战题目解析：爬楼梯、括号生成等问题

	- [70. 爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)

		阿里巴巴、腾讯、字节跳动在半年内面试常考

		- 关键点

			- N直接出有点难，先从1，2，3开始

	- [22. 括号生成](https://leetcode-cn.com/problems/generate-parentheses/)

		字节跳动在半年内面试中考过

		- 关键点

			- 1. 递归函数体的参数，可能因为具体情况会调整

			- 2. 有时候当层逻辑可以合并到下层调用

			- 3. 终结条件有时候可以挪到下层调用上

		- 技巧

			- 通过对终结条件前移到下层调用之前的这种手段，叫做剪枝

	- [98. 验证二叉搜索树](https://leetcode-cn.com/problems/validate-binary-search-tree/)

		谷歌、字节跳动、Facebook 在半年内面试中考过

		- 可能解法

			- 1. 写个递归去验证

			- 2. 利用bst中序遍历是递增的，来验证

	- 参考链接

		- [如何优雅地计算斐波那契数列](https://time.geekbang.org/dailylesson/detail/100028406)

## 第八课 分治、回溯

- 1. 分治、回溯的实现和特性

	- 分治 (Divide & Conquer)

		- 定义

			- 实质就是，分解问题，找重复子问题，然后汇总子问题的解。

			- 分治是一种思想，递归是一种具体的编程方法。

		- 结构

			- 1. 退出条件

			- 2. 拆分子问题(当层逻辑)

			- 3. 子问题求解(下层调用)

			- 4. 子问题汇总

			- 5. 当层状态清理

		- 关键点

			- 处理当层时，不要越级，或者越太多级。

	- 回溯 (Backtracking)

		- 定义

			- 本质上来说，就是利用递归的方法，在所有解题空间内，找到正确的答案。

			- 有时候会有，找不到正确答案，同时耗费掉指数时间复杂度的时间的最坏情况发生。

- 2. 实战题目解析：Pow(x,n)、子集

	- [50. Pow(x, n)](https://leetcode-cn.com/problems/powx-n/)

		Facebook 在半年内面试常考

		- 可能解法

			- 1. 暴力法，有几个n就循环几次*x动作，时间O(n)

			- 2. 分治，pow(x,n)->pow(x,n/2)^2，根据n的奇偶性，追加*x，时间O(logn)

		- 关键点

			- 确定参数边界

	- [78. 子集](https://leetcode-cn.com/problems/subsets/)

		Facebook、字节跳动、亚马逊在半年内面试中考过

		- 可能解法

			- 1. 回溯，把每个可能出现的符号按照要么选要么不选的方式进行组合，O(2^n)

			- 2. 迭代，最初从空集开始，for循环把所有元素都加进去，然后把当次循环的所有结果加回结果集合，O(2^n)

	- 参考链接

		- [牛顿迭代法原理](http://www.matrix67.com/blog/archives/361)

		- [牛顿迭代法代码](http://www.voidcn.com/article/p-eudisdmk-zm.html)

- 3. 实战题目解析：电话号码的字母组合、n皇后

	- [169. 多数元素](https://leetcode-cn.com/problems/majority-element/description/)

		亚马逊、字节跳动、Facebook 在半年内面试中考过。
		简单但高频

		- 自己练有很多解法，都可以看看

	- [17. 电话号码的字母组合](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/)

		亚马逊在半年内面试常考

		- 关键点

			- 可以参照子集生成，把每个数字对应成一串字符集，然后全组合

	- [51. N皇后](https://leetcode-cn.com/problems/n-queens/)

		字节跳动、苹果、谷歌在半年内面试中考过

		- 关键点

			- 斜边的标示方法

## 作业

- medium

	- [236. 二叉树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)

		Facebook 在半年内面试常考

		- [x] 20200604

		- [x] 20200605

		- [ ] 20200611

	- [105. 从前序与中序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)

		字节跳动、亚马逊、微软在半年内面试中考过

		- [x] 20200604

		- [x] 20200605

		- [ ] 20200611

	- [77. 组合](https://leetcode-cn.com/problems/combinations/)

		微软、亚马逊、谷歌在半年内面试中考过

		- [x] 20200604

		- [x] 20200605

		- [ ] 20200611

	- [46. 全排列](https://leetcode-cn.com/problems/permutations/)

		字节跳动在半年内面试常考

		- [x] 20200604

		- [x] 20200605

		- [ ] 20200611

	- [47. 全排列 II](https://leetcode-cn.com/problems/permutations-ii/)

		亚马逊、字节跳动、Facebook 在半年内面试中考过

		- [x] 20200604

		- [x] 20200605

		- [ ] 20200611
