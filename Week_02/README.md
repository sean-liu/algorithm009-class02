学习笔记
# 第二周


## 第五课 哈希表、映射、集合

- 1. 哈希表、映射、集合的实现与特性

	- 哈希表(Hash table)

		- 定义

			- 也叫散列表，是根据关键码值(Key value)而直接进行访问的数据结构。 它通过把关键码值映射到表中一个位置来访问记录，以加快查找的 速度。

			- 其中，映射函数叫作散列函数(Hash Function)，存放记录的数组 叫作哈希表(或散列表)。

		- 扩展

			- Hash Function

				- 定义

					- 把key转化成一个特定的数值的函数称为hash function

				- 举例

					- 在java中，“lies”->429

			- Hash collision

				- 定义

					- Key通过hash function转化成数值的时候，有时候会出现不同的key的结果是相同的，这种情况就是 hash collision

		- 结构

			- 由一个数组来存储被hash的key的值，如果值相同，会再接上链表

		- 时间复杂度

			除非在碰撞很严重的情况下，全都会退化成O(n)

			- Search, O(1)

			- Insertion, O(1)

			- Deletion, O(1)

		- 工程实践

			- 电话簿

			- 用户信息表

			- 缓存(LRU Cache)

			- Redis

		- 具体实现

			- Java

				- Set

					- https://docs.oracle.com/en/java/javase/12/docs/api/ java.base/java/util/Set.html

					- TreeSet, HashSet

				- Map

					- https://docs.oracle.com/en/java/javase/12/docs/api/ java.base/java/util/Map.html

					- HashMap, Hashtable, ConcurrentHashMap

				- 全是接口，有很多具体实现，可以都看看每个实现的应用场景

- 2. 实战题目解析：有效的字母异位词等问题

	- [242. 有效的字母异位词](https://leetcode-cn.com/problems/valid-anagram/)

		- 可能解法

			- 1. 暴力，sort，sorted_str看是否相等

			- 2. hash map-> 统计词频

	- [1. 两数之和](https://leetcode-cn.com/problems/two-sum/)

		- 可能解法

			- 1. 用hashmap把所有元素存进去，然后第二次循环去检查是否有match

			- 2. 一次hashmap法，循环的时候看hashmap中是否有命中，没有命中的话，算出target-当前元素的值作为key，塞到hashmap中

## 第六课 树、二叉树、二叉搜索树

- 1. 树、二叉树、二叉搜索树的实现和特性

	- 树 (Tree)

		- 由来

			- 当把链表从1维扩展成2维的时候(具体来说，就是当next指针不只是指向了1个后记节点的时候)，就变成了树。

		- 定义

			- 树就是有向无环的图，是由一个一个节点(包含子节点们，以及自身值的结构)组成的。

	- 二叉树 (Binary Tree)

		- 定义

			- 每个节点上只包含左右2个节点的树

		- 遍历

			- 1. 前序 (Pre-order): 根-左-右

			- 2. 中序 (In-order): 左-根-右

			- 3. 后序 (Post-order): 左-右-根

	- 二叉搜索树 (Binary Search Tree)

		- 定义

			- 二叉搜索树，也称二叉搜索树、有序二叉树(Ordered Binary Tree)、 排序二叉树(Sorted Binary Tree)，是指一棵空树或者具有下列性质的 二叉树:  
			  1. 左子树上所有结点的值均小于它的根结点的值;  
			  2. 右子树上所有结点的值均大于它的根结点的值;  
			  3. 以此类推:左、右子树也分别为二叉查找树。 (这就是 重复性!)

		- 特点

			- 中序遍历 是 升序排列

		- 时间复杂度

			- Search, O(logn)

			- Insertion, O(logn)

			- Deletion, O(logn)

	- 图 (Graph)

	- 特殊关系

		- linkedlist是tree的特例(后记节点只有一个)，tree是graph的特例(graph可以包含环)

	- 为什么树的面试题解法一般都是递归?

		- 由于它的结构导致

- 2. 实战题目解析：二叉树的中序遍历

	- [94. 二叉树的中序遍历](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)

		亚马逊、微软、字节跳动在半年内面试中考过

		- 可能解法

			- 1. 递归

			- 2. 堆栈

			- 3. 莫里斯

## 第六课2 堆和二叉堆、图

- 1. 堆和二叉堆的实现和特性

	- 堆(Heap)

		- 定义

			- 能迅速找到最大值或者最小值的数据结构。注：同时只能找一个方向。

		- 常见操作

			- findMax, O(1)

			- deleteMax, O(logn)

			- Insert, O(logn)

		- [各种实现堆的比较](https://www.geeksforgeeks.org/heap-data-structure/)

	- 二叉堆(Binary Heap)

		- 定义

			- 是堆的一种实现，可以由数组实现

		- 性质

			- 底层是完全二叉树

			- 节点的值大于等于子节点的值

		- 实现细节

			- 一般是基于数组

			- 父子节点寻址方式

				- 假设根节点索引是0

				- 左子节点：2i+1

				- 右子节点：2i+2

				- 父节点：floor((i-1)/2)

			- 插入操作

				- 1. 将新元素放入数组最后一个位置

				- 2. 从新元素开始，按照逐级地和父元素比较，以及挪动元素，直到合适的位置。(HeapifyUp过程)

			- 删除操作

				- 1. 将栈顶元素删除，并把栈尾元素放到栈顶

				- 2. 从根节点开始，在左右子节点中，把最大的那个移到父节点位置，一路到底。(HeapifyDown过程)

- 2. 实战题目解析：最小的k个数、滑动窗口最大值等问题

	- [面试题40. 最小的k个数](https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/)

		- 可能解法

			- 1. 排序, O(nlogn)

			- 2. heap, O(nlogk)

			- 3. quicksort

	- [239. 滑动窗口最大值](https://leetcode-cn.com/problems/sliding-window-maximum/)

		亚马逊在半年内面试常考

		- 可能解法

			- 1. 用heap

	- [347. 前 K 个高频元素](https://leetcode-cn.com/problems/top-k-frequent-elements/)

		- 可能解法

			- 1. 用heap，和hashmap，先进map出count，然后heap用count做权，取前4个

- 3. 图的实现和特性

	- 图(Graph)

		- 定义

			- 包含用来描述一个图所需要的点和线的集合，即是图

		- 属性

			- Graph(V, E)

			- V: 点(Vertex)  

				- 1. 度

					- 分入度：从别的点连过来的数量，和出度：从自身连出去的数量

				- 2. 与其他的Vertex是否通

			- E: 边(Edge)

				- 1. 有向与无向 (单行线)

				- 2. 权重 (边长)

		- 储存边的方式

			- 邻接矩阵

				- 当点a与点b可以连同时，在a行b列，以及b行a列，填上1的方式来记录

			- 邻接表

				- 链表方式，存储，从当前点能出发的边，以及点

		- 分类

			- 无向无权图

			- 有向无权图

			- 无向有权图

			- 有向有权图

		- 相关算法

			- DFS

				- 深度优先遍历

			- BFS

				- 广度优先遍历

			- 注意带上visited集合，用于标记，是否已经访问过某个点

		- 高级算法

			- [连通图个数](https://leetcode-cn.com/problems/number-of-islands/)

			- [拓扑排序](https://zhuanlan.zhihu.com/p/34871092)

			- [Dijkstra最短路径算法](https://www.bilibili.com/video/av25829980?from=search&seid=13391343514095937158)

			- [最小生成树](https://www.bilibili.com/video/av84820276?from=search&seid=17476598104352152051)

## 作业

- easy

	- [x] [1. 两数之和](https://leetcode-cn.com/problems/two-sum/)

	- [242. 有效的字母异位词](https://leetcode-cn.com/problems/valid-anagram/)

		亚马逊、Facebook、谷歌在半年内面试中考过

		- [x] 20200525

		- [x] 20200526

		- [ ] 20200601

	- [589. N叉树的前序遍历](https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/)

		亚马逊在半年内面试中考过

		- [x] 20200526

		- [x] 20200527

		- [ ] 20200602

- medium

	- [49. 字母异位词分组](https://leetcode-cn.com/problems/group-anagrams/)

		亚马逊在半年内面试中常考

		- [x] 20200525

		- [x] 20200526

		- [ ] 20200601

	- [94. 二叉树的中序遍历](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)

		亚马逊、微软、字节跳动在半年内面试中考过

		- [x] 20200529

		- [ ] 20200530

		- [ ] 20200605

	- [ ] [144. 二叉树的前序遍历](https://leetcode-cn.com/problems/binary-tree-preorder-traversal/)

		谷歌、微软、字节跳动在半年内面试中考过

	- [ ] [429. N叉树的层序遍历](https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/)

	- [ ] [264. 丑数 II](https://leetcode-cn.com/problems/ugly-number-ii/)

	- [ ] [347. 前 K 个高频元素](https://leetcode-cn.com/problems/top-k-frequent-elements/)

- [ ] 写一个关于HashMap的小总结

	对于不熟悉Java语言的同学，此项作业可选做。

- [ ] 自学HeapSort

	自学 https://www.geeksforgeeks.org/heap-sort/

- 依旧选二即可

