学习笔记
# 第一周


## 第三课

### 1. 数组、链表、跳表的基本实现和特性

- Array

	- 时间复杂度

		- Access, O(1)

		- Search, O(n)

		- Insertion, O(n)

		- Deletion, O(n)

	- 特点

		- 随机访问时间极快，插入和删除偏慢

	- 老师剖析了java的arraylist

		- 看了add

		- 还看了ensureCapacity

- Linked List

	- 定义

		- 链表由0到n个节点构成，每个节点包含通向下一个节点的链接，以及其本身的值构成

	- 由来

		- 为了弥补数组的缺点，添加和删除元素比较慢，产生了链表

	- 时间复杂度

		- Access, O(n)

		- Search, O(n)

		- Insertion, O(1)

		- Deletion, O(1)

	- 扩展

		- 当尾巴节点指向头节点，该结构称为循环链表

		- 当节点中包含了前链接和后链接的时候，该结构称为双链表

	- 细节

		- 介绍了如何插入和删除元素的细节步骤，以及特点

- Skip List

	- 定义

		- 跳表是由普通链表，以及多层级的索引组成，并且只适用于有序元素的情况。

	- 由来

		- 为了弥补链表在有序数据查询慢的缺点，出现了跳表

	- 时间复杂度

		- Access, O(logn)

		- Search, O(logn)

		- Insertion, O(logn)

		- Deletion, O(logn)

	- 特点

		- 跳表对标的是平衡树和二分查找，是一种 插入/删除/搜索 都是O(logn)的数据结构。1989年出现。比其他结构晚了近30年

		- 优势是原理简单、容易实现、方便扩展、效率更高。

			- 在一些热门项目里替代了平衡树，如Redis，LevelDB等。

	- 细节

		- 跳表的查找加速如何实现

			- 增加其他维度的信息，在此上下文中，即增加多级索引。

				- 每一级索引的个数由在下一层元素个数/2得到

		- 查询的时间复杂度分析

			- 第k级索引节点的个数是n/(2^k)

			- 假设索引有h级，最高级的索引有2个节点，则n/(2^h)=2，从而h=log2(n)-1

		- 现实中跳表的形态

			- 不是如上所示的规整的形状，由于有动态的插入和删除动作

		- 空间复杂度分析

			- 由于所有索引的大小是由每层索引个数汇总，所以总节点数=Sum(n/(2^k)(k=1,2,3...))数量级是在n级的，所以空间复杂度为O(n)

- 应用

	- LRU Cache - Linked List

	- Redis - Skip List

- 小结

	- 跳表：升维思想+空间换时间

### 2. 实战题目解析：移动零

- 练习步骤

	- 1. 5-10分钟：读题和思考

	- 2. 有思路：自己开始做和写代码；不然，就看题解

	- 3. 默写、背诵到熟练

	- 4. 然后开始自己写(闭卷)

- [283. 移动零](https://leetcode-cn.com/problems/move-zeroes/)

	- 可能解法

		- 1. loop，count zeroes

		- 2. 开新数组，loop，非零从左开始放，零从右开始放

		- 3. 多开一个index j，用于标示非0位置，每次用完做 ++操作

	- 需要养成的好的点

		- 测试一些边界情况

	- 关键点

		- [x] 方法3中的index 方法要熟练，需要练习一下

### 3. 实战项目解析：盛水最多的容器、爬楼梯

- [11. 盛最多水的容器](https://leetcode-cn.com/problems/container-with-most-water/)

	- 可能解法

		- 1. 枚举left bar，和right bar，(left index-right index)*min(height), O(n^2)

		- 2. 左右边界 i, j 向中间收敛

	- 关键点

		- [x] 双循环遍历，自己写一下 

		- 左右夹逼是一个普遍使用的套路

- [70. 爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)

	- 可能解法

		- 1. 写递归公式

	- 关键点

		- 做题没思路的一般处理过程

			- 1. 尝试暴力破解

			- 2. 设想基本情况

			- 3. 尝试找出最近重复子问题

				- 之所以上面的方式可以生效是因为，计算机本质上是if else for loop 还有recursion

### 4. 实战项目解析：3数之和、环形链表

- [1. 两数之和](https://leetcode-cn.com/problems/two-sum/)

	- 可能解法

		- 1. 暴力法，2个循环，时间复杂度O(n^2)

		- 2. 用hashmap存先存一份array数据，然后在过一遍，得出结果，时间O(n)，空间O(n)

		- 3. 用hashmap一边循环，针对每个数字计算出需要的数字作为key，然后等待后记数字来查是否满足，时间O(n)，空间O(n)

- [15. 三数之和](https://leetcode-cn.com/problems/3sum/)

	- 可能解法

		- 1. 暴力求解，O(n^3)

		- 2. 用hash：两重暴力+hash

		- 3. 夹逼：排序后，3个指针，kij，k循环一次，在每个k元素时，ij从左右往中间夹逼，时间O(n^2)

	- 关键点

		- 夹逼一般只能在有序情况下进行

## 第四课

### 1. 栈和队列的基本实现和特性

- Stack

	- 特点

		- 先进后出FILO

	- 时间复杂度

		- Access, O(n)

		- Search, O(n)

		- Insertion, O(1)

		- Deletion, O(1)

	- java实现的文档

		- https://docs.oracle.com/javase/10/docs/api/java/util/Stack.html

- Queue

	- 特点

		- 先进先出FIFO

	- 时间复杂度

		- Access, O(n)

		- Search, O(n)

		- Insertion, O(1)

		- Deletion, O(1)

	- java实现的文档

		- https://docs.oracle.com/javase/10/docs/api/java/util/Queue.html

- Deque(Double-ended Queue)

	- 定义

		- 其实就是一个2端都可以出和入元素的队列

	- 特点

		- 2端可以做出和入元素的操作，所以既可以是Queue也可以是Stack

	- java实现的文档

		- https://docs.oracle.com/javase/10/docs/api/java/util/Deque.html

- 工程实现

	- 主流语言基本都有各自实现

	- 如何搜索相关源代码，google，关键字即可

- Priority Queue

	- 时间复杂度

		- Insertion, O(1)

		- Deletion, O(logn)

	- 具体底层实现

		- heap(堆)

		- bst(二叉搜索树)

		- treap

	- java实现的文档

		- https://docs.oracle.com/javase/10/docs/api/java/util/PriorityQueue.html

### 2. 实战题目解析

- 预习题目

	- [20. 有效的括号](https://leetcode-cn.com/problems/valid-parentheses/description/)

		- 可能解法

			- 1. 暴力：不断地replace匹配的括号，O(n^2)

				- 情况1. ()[]{}

				- 情况2. ((()))

			- 2. 用stack，O(n)

	- [155. 最小栈](https://leetcode-cn.com/problems/min-stack/)

		- 可能解法

			- 1. 2个栈来实现

- 实战题目

	- [84. 柱状图中最大的矩形](https://leetcode-cn.com/problems/largest-rectangle-in-histogram/)

		- 可能解法

			- 1. 暴力1，O(n^3)

				for i -> 0, n-2  
					for j -> i+1, n-1  
						(i, j)-> 最小高度，area  
						update(max)

			- 2. 暴力2，O(n^2)

				For i -> 0, n-1  
					找左右边界  
					area = (right - left)*height[i]  
					update max  
				

			- 3. 用stack，需要维护一个有序的stack来记录，下标和高度，O(n)

	- [239. 滑动窗口最大值](https://leetcode-cn.com/problems/sliding-window-maximum/)

		- 可能解法

			- 1. 暴力，O(n*k)

			- 2. deque(sliding window)

- 关键点

	- array类

		- 双指针方式有时候有奇效可以尝试着用一下

	- 如果有最近相关性就往stack上走

		- 最近相关性就是诸如此类的结构：([])

	- 如果遇到只能用队列来实现栈，或者只能用栈来实现队列

		- 2个队列和2个栈有奇效

## 作业

### 题目集

- easy

	- [x] [26. 删除排序数组中的重复项](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/)

	- [189. 旋转数组](https://leetcode-cn.com/problems/rotate-array/)

		- [x] 20200518

		- [x] 20200519

		- [ ] 20200525

	- [88. 合并两个有序数组](https://leetcode-cn.com/problems/merge-sorted-array/)

		- [x] 20200521

		- [x] 20200522

		- [ ] 20200528

	- [ ] [21. 合并两个有序链表](https://leetcode-cn.com/problems/merge-two-sorted-lists/)

	- [ ] [1. 两数之和](https://leetcode-cn.com/problems/two-sum/)

	- [x] [283. 移动零](https://leetcode-cn.com/problems/move-zeroes/)

	- [x] [66. 加一](https://leetcode-cn.com/problems/plus-one/)

	- [x] 用 add first 或 add last 这套新的 API 改写 Deque 的代码

	- [ ] 分析 Queue 和 Priority Queue 的源码

- medium

	- [ ] [641. 设计循环双端队列](https://leetcode-cn.com/problems/design-circular-deque/)

- hard

	- [ ] [42. 接雨水](https://leetcode-cn.com/problems/trapping-rain-water/)

### 要求选二即可
