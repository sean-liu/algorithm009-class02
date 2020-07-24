学习总结
感受

整个算法训练营给了我一条很清晰的路径，来提高我在算法这一块的功底。
我第一次意识到了，过遍数，不光对于背单词有奇效，同时也对于掌握算法来说也有奇效的。
直到现在，我对于老师在讲述实战例题时候的表情以及讲述的内容依旧记忆犹新，印象深刻。
课程结束以后，刷题不会停，做时间的朋友，争取有一天可以刷完leetcode上所有的题目。

以下为具体的课程的相对应的笔记以及总结

# 预习周


## 第一课

- 精通一个领域

	- Chunk it up

		- 庖丁解牛

		- 脉络连接

	- Deliberate Practicing

		- 刻意练习-过遍数

		- 练习缺陷, 弱点地方

		- 不舒服, 不爽, 枯燥

		- 生活中例子: 乒乓球, 台球, 游戏等

	- Feedback

		- 即时反馈

		- 主动型反馈(自己去找)

			- 高手代码(GitHub, LeetCode等)

			- 第一视角直播

		- 被动式反馈(高手给你指点)

			- code review

			- 教练看你打, 给你反馈

- 刷题技巧

	- 切题四件套

		- Clarification

		- Possible solutions

			- compare(time/space)

			- optimal(加强)

		- Coding(多写)

		- Test cases

	- 五步刷题法(五毒神掌)

		- 第一遍

			- 5分钟: 读题 + 思考

			- 直接看解法: 注意! 多解法, 比较解法优劣

			- 背诵, 默写好的解法

		- 第二遍

			- 马上自己写->leetcode提交

			- 多种解法比较, 体会->优化!

		- 第三遍

			- 过了1天后, 再重复做题

			- 不同解法的熟练度->专项练习

		- 第四遍

			- 过了一周: 反复回来练习相同题目

		- 第五遍

			- 面试前一周恢复性训练

	- 做算法题的最大误区: 只做一遍

## 第二课

- 第一节 训练环境设置, 编码技巧和Code Style

- 第二节 时间复杂度, 空间复杂度


# 2nd week


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

		亚马逊、Facebook、谷歌在半年内面试中考过

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

			- 当把链表从1维扩展成2维的时候(具体来说，就是当next指针不只是指向了1个后记结点的时候)，就变成了树。

		- 定义

			- 树就是有向无环的图，是由一个一个结点(包含子结点们，以及自身值的结构)组成的。

	- 二叉树 (Binary Tree)

		- 定义

			- 每个结点上只包含左右2个结点的树

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

		- linkedlist是tree的特例(后记结点只有一个)，tree是graph的特例(graph可以包含环)

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

			- 结点的值大于等于子结点的值

		- 实现细节

			- 一般是基于数组

			- 父子结点寻址方式

				- 假设根结点索引是0

				- 左子结点：2i+1

				- 右子结点：2i+2

				- 父结点：floor((i-1)/2)

			- 插入操作

				- 1. 将新元素放入数组最后一个位置

				- 2. 从新元素开始，按照逐级地和父元素比较，以及挪动元素，直到合适的位置。(HeapifyUp过程)

			- 删除操作

				- 1. 将栈顶元素删除，并把栈尾元素放到栈顶

				- 2. 从根结点开始，在左右子结点中，把最大的那个移到父结点位置，一路到底。(HeapifyDown过程)

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

	- 参考链接

		- [连通图个数](https://leetcode-cn.com/problems/number-of-islands/)

		- [拓扑排序（Topological Sorting）](https://zhuanlan.zhihu.com/p/34871092)

		- [最短路径（Shortest Path）Dijkstra](https://www.bilibili.com/video/av25829980?from=search&seid=13391343514095937158)

		- [最小生成树（Minimum Spanning Tree）](https://www.bilibili.com/video/av84820276?from=search&seid=17476598104352152051)

## 作业

- easy

	- [x] [1. 两数之和](https://leetcode-cn.com/problems/two-sum/)

	- [242. 有效的字母异位词](https://leetcode-cn.com/problems/valid-anagram/)

		亚马逊、Facebook、谷歌在半年内面试中考过

		- [x] 20200525

		- [x] 20200526

		- [x] 20200601

	- [589. N叉树的前序遍历](https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/)

		亚马逊在半年内面试中考过

		- [x] 20200526

		- [x] 20200527

		- [x] 20200602

- medium

	- [49. 字母异位词分组](https://leetcode-cn.com/problems/group-anagrams/)

		亚马逊在半年内面试中常考

		- [x] 20200525

		- [x] 20200526

		- [x] 20200601

	- [94. 二叉树的中序遍历](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)

		亚马逊、微软、字节跳动在半年内面试中考过

		- [x] 20200529

		- [x] 20200530

		- [x] 20200605

	- [144. 二叉树的前序遍历](https://leetcode-cn.com/problems/binary-tree-preorder-traversal/)

		谷歌、微软、字节跳动在半年内面试中考过

		- [x] 20200530

		- [x] 20200531

		- [x] 20200606

	- [429. N叉树的层序遍历](https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/)

		- [x] 20200531

		- [x] 20200601

		- [x] 20200607

	- [264. 丑数 II](https://leetcode-cn.com/problems/ugly-number-ii/)

		- [x] 20200531

		- [x] 20200601

		- [x] 20200607

	- [347. 前 K 个高频元素](https://leetcode-cn.com/problems/top-k-frequent-elements/)

		- [x] 20200531

		- [x] 20200601

		- [x] 20200607

- [ ] 写一个关于HashMap的小总结

	对于不熟悉Java语言的同学，此项作业可选做。

- [ ] 自学HeapSort

	自学 https://www.geeksforgeeks.org/heap-sort/

- 依旧选二即可


# 1st week


## 第三课 数组、链表、跳表

- 1. 数组、链表、跳表的基本实现和特性

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

			- 链表由0到n个结点构成，每个结点包含通向下一个结点的链接，以及其本身的值构成

		- 由来

			- 为了弥补数组的缺点，添加和删除元素比较慢，产生了链表

		- 时间复杂度

			- Access, O(n)

			- Search, O(n)

			- Insertion, O(1)

			- Deletion, O(1)

		- 扩展

			- 当尾巴结点指向头结点，该结构称为循环链表

			- 当结点中包含了前链接和后链接的时候，该结构称为双链表

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

				- 第k级索引结点的个数是n/(2^k)

				- 假设索引有h级，最高级的索引有2个结点，则n/(2^h)=2，从而h=log2(n)-1

			- 现实中跳表的形态

				- 不是如上所示的规整的形状，由于有动态的插入和删除动作

			- 空间复杂度分析

				- 由于所有索引的大小是由每层索引个数汇总，所以总结点数=Sum(n/(2^k)(k=1,2,3...))数量级是在n级的，所以空间复杂度为O(n)

	- 应用

		- LRU Cache - Linked List

		- Redis - Skip List

	- 小结

		- 跳表：升维思想+空间换时间

- 2. 实战题目解析：移动零

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

			- 方法3中的index 方法要熟练，需要练习一下

- 3. 实战项目解析：盛水最多的容器、爬楼梯

	- [11. 盛最多水的容器](https://leetcode-cn.com/problems/container-with-most-water/)

		- 可能解法

			- 1. 枚举left bar，和right bar，(left index-right index)*min(height), O(n^2)

			- 2. 左右边界 i, j 向中间收敛

		- 关键点

			- 双循环遍历，自己写一下 

			- 左右夹逼是一个普遍使用的套路

	- [70. 爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)

		阿里巴巴、腾讯、字节跳动在半年内面试常考

		- 可能解法

			- 1. 写递归公式

		- 关键点

			- 做题没思路的一般处理过程

				- 1. 尝试暴力破解

				- 2. 设想基本情况

				- 3. 尝试找出最近重复子问题

					- 之所以上面的方式可以生效是因为，计算机本质上是if else for loop 还有recursion

- 4. 实战项目解析：3数之和、环形链表

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

## 第四课 栈、队列、优先队列、双端队列

- 1. 栈和队列的基本实现和特性

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

- 2. 实战题目解析

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

- 题目集

	- easy

		- [26. 删除排序数组中的重复项](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/)

			- [x] 20200521

			- [x] 20200524

		- [189. 旋转数组](https://leetcode-cn.com/problems/rotate-array/)

			- [x] 20200518

			- [x] 20200519

			- [x] 20200525

		- [88. 合并两个有序数组](https://leetcode-cn.com/problems/merge-sorted-array/)

			- [x] 20200521

			- [x] 20200522

			- [x] 20200528

		- [21. 合并两个有序链表](https://leetcode-cn.com/problems/merge-two-sorted-lists/)

			- [x] 20200523

			- [x] 20200524

			- [x] 20200530

		- [1. 两数之和](https://leetcode-cn.com/problems/two-sum/)

			- [x] 20200523

			- [x] 20200524

			- [x] 20200530

		- [283. 移动零](https://leetcode-cn.com/problems/move-zeroes/)

			华为，字节跳动，202005半年内常考

			- [x] 20200520

			- [x] 20200521

			- [x] 20200527

		- [66. 加一](https://leetcode-cn.com/problems/plus-one/)

			- [x] 20200518

			- [x] 20200524

	- medium

		- [641. 设计循环双端队列](https://leetcode-cn.com/problems/design-circular-deque/)

			- [x] 20200524

			- [x] 20200525

			- [x] 20200531

	- hard

		- [42. 接雨水](https://leetcode-cn.com/problems/trapping-rain-water/)

			- [x] 20200524

			- [x] 20200525

			- [x] 20200531

	- [ ] 分析 Queue 和 Priority Queue 的源码

	- [x] 用 add first 或 add last 这套新的 API 改写 Deque 的代码

- 要求选二即可


# 3rd week


## 第七课 泛型递归、树的递归

- 1. 递归的实现、特性以及思维要点

	- 回顾

		- 树相关的题目为何都是递归？

			- 1. 由于树结点的定义

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

- easy

	- [169. 多数元素](https://leetcode-cn.com/problems/majority-element/description/)

		亚马逊、字节跳动、Facebook 在半年内面试中考过。
		简单但高频

		- [x] 20200606

		- [x] 20200607

		- [x] 20200613

- medium

	- [236. 二叉树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)

		Facebook 在半年内面试常考

		- [x] 20200604

		- [x] 20200605

		- [x] 20200611

	- [105. 从前序与中序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)

		字节跳动、亚马逊、微软在半年内面试中考过

		- [x] 20200604

		- [x] 20200605

		- [x] 20200611

	- [77. 组合](https://leetcode-cn.com/problems/combinations/)

		微软、亚马逊、谷歌在半年内面试中考过

		- [x] 20200604

		- [x] 20200605

		- [x] 20200611

	- [46. 全排列](https://leetcode-cn.com/problems/permutations/)

		字节跳动在半年内面试常考

		- [x] 20200604

		- [x] 20200605

		- [x] 20200611

	- [47. 全排列 II](https://leetcode-cn.com/problems/permutations-ii/)

		亚马逊、字节跳动、Facebook 在半年内面试中考过

		- [x] 20200604

		- [x] 20200605

		- [x] 20200611


# 4th week


## 第9课 深度优先搜索和广度优先搜索

- 1. 深度优先搜索、广度优先搜索的实现和特性

	- 遍历搜索

		- 特点

			- 1. 每个结点都要访问一次

			- 2. 每个结点仅仅访问一次

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

## 第10课 贪心算法

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

## 第11课 二分查找

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

		- [x] 20200619

	- [455. 分发饼干](https://leetcode-cn.com/problems/assign-cookies/)

		亚马逊在半年内面试中考过

		- [x] 20200612

		- [x] 20200613

		- [x] 20200619

	- [122. 买卖股票的最佳时机 II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)

		亚马逊、字节跳动、微软在半年内面试中考过

		- [x] 20200621

		- [x] 20200622

		- [x] 20200628

	- [874. 模拟行走机器人](https://leetcode-cn.com/problems/walking-robot-simulation/)

		- [x] 20200621

		- [x] 20200622

		- [x] 20200628

- medium

	- [127. 单词接龙](https://leetcode-cn.com/problems/word-ladder/)

		亚马逊在半年内面试常考

	- [200. 岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)

		近半年内，亚马逊在面试中考查此题达到 350 次

		- [x] 20200609

		- [x] 20200610

		- [x] 20200616

	- 关键点：unionFindSet，在做岛屿数量的时候，挖了1小时的结果

	- [529. 扫雷游戏](https://leetcode-cn.com/problems/minesweeper/)

		亚马逊、Facebook 在半年内面试中考过

		- [x] 20200609

		- [x] 20200610

		- [x] 20200616

	- [55. 跳跃游戏](https://leetcode-cn.com/problems/jump-game/)

		亚马逊、华为、Facebook 在半年内面试中考过

		- [x] 20200617

		- [x] 20200618

		- [x] 20200624

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

		- [x] 20200616

		- [x] 20200617

		- [x] 20200623

- [ ] 使用二分查找，寻找一个半有序数组 [4, 5, 6, 7, 0, 1, 2] 中间无序的地方  
  说明：同学们可以将自己的思路、代码写在学习总结中


# 6th week


## 第12课 动态规划

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

				- 状态树中，结点为0的，就是一个可能解，然后找level最低的，就是用硬币最少的

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

		- [x] 20200627

		- [x] 20200628

		- [x] 20200704

	- [91. 解码方法](https://leetcode-cn.com/problems/decode-ways/)

		- [x] 20200627

		- [x] 20200628

		- [x] 20200704

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


# 7th week


## 第13课 字典树和并查集

- 1. Trie树的基本实现和特性

	- 复习

		- 树

			- 跟结点是level0

		- 二叉搜索树

			- 左子树小于根结点，右子树大于根结点

			- 中序遍历可以顺序输出

	- 字典树(Trie)

		- 数据结构

			- 结点不存完整字符串，只存单字母，和标示字符串结束的终止符，另外加上26个可能的指针组成的数组

		- 基本性质

			- 结点本身不存完整字符串

			- 根结点到某一叶子结点上，经过的结点，即为完整字符串。

			- 每个结点的子结点的路径代表的字符串不相同

		- 核心思想

			- 空间换时间

			- 利用字符串公共前缀来减少时间消耗

	- [208. 实现 Trie (前缀树)](https://leetcode-cn.com/problems/implement-trie-prefix-tree/)

		- 必要方法

			- insert

			- search

			- starwith

- 2. Trie树实战题目解析：单词搜索2

	- [212. 单词搜索 II](https://leetcode-cn.com/problems/word-search-ii/)

		亚马逊、微软、苹果在半年内面试中考过

		- 可能解法

			- 1. words遍历，对于每个words去board上一个一个查过去，O(N*n*m*4^k), N是words个数，n、m是board的长和宽，k为单词平均长度

			- 2. trie

				- 遍历words，建立trie树

				- boards，dfs

				- O(n*m*4*3^(k-1))，n*m是board总数，4是第一次可以尝试4个方向，3是后面的只能尝试其他3个没试过的方向。

		- 关键点

			- 是4连通(上下左右)，不是8连通(斜方也算)

			- 4连通遍历，使用各个方向的单位数组做操作

- 3. 并查集的基本实现、特性和实战题目解析

	- 并查集(Disjoint set)

		- 适用场景

			- 组团、配对问题

			- group or not

		- 基本操作

			- makeSet(s)

				- 建立一个新的并查集，其中包含s个单元素集合

			- unionSet(x, y)

				- 把x与y所在的集合合并，要求x和y所在集合不相交，相交则不操作。

			- find(x)

				- 找到x所在的集合代表，用于判定2个x是否属于同一个集合

		- 特殊操作

			- 在调用find的时候，到集合代表的路径会被缩短。

	- 参考链接

		- [并查集代码模板](https://shimo.im/docs/VtcxL0kyp04OBHak/read)

	- [547. 朋友圈](https://leetcode-cn.com/problems/friend-circles/)

		亚马逊、Facebook、字节跳动在半年内面试中考过

		- 可能解法

			- 1. DFS、BFS 与岛屿问题类似

			- 2. 并查集

## 第14课 高级搜索

- 1. 剪枝的实现和特性(Pruning)

	- 初级搜索

		- 朴素搜索

		- 优化方式

			- 不重复

			- 剪枝

		- 搜索方向

			- DFS

			- BFS

		- 基于方向的优化

			- 双向BFS搜索

			- A*启发式搜索

	- 剪枝

		- 用一些条件剔除了和目标一定无关的选项

	- 回溯

		- 和递归差不多，与递归的区别是，如果当前分支不行的话，会恢复相关的状态

	- 参考链接

		- [DFS 代码模板](https://shimo.im/docs/UdY2UUKtliYXmk8t/read)

		- [BFS 代码模板](https://shimo.im/docs/ZBghMEZWix0Lc2jQ/read)

		- [AlphaZero Explained](https://nikcheerla.github.io/deeplearningschool/2018/01/01/AlphaZero-Explained/)

		- [棋类复杂度](https://u.geekbang.org/lesson/14?article=251083)

- 2. 剪枝实战题目解析：数独

	- [70. 爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)

		阿里，腾讯，字节跳动，202005半年常考

		- 用递归再做一下

	- [22. 括号生成](https://leetcode-cn.com/problems/generate-parentheses/)

		字节跳动、亚马逊、Facebook 在半年内面试中考过

		- 用剪枝再练一下

		- 可以用dp练一下

	- [51. N皇后](https://leetcode-cn.com/problems/n-queens/)

		亚马逊、苹果、字节跳动在半年内面试中考过

		- 关键点

			- 记录重复的方式：撇、横、捺

				- 公式，对于(p, q)

					- 横竖

						- p != x && q != y

					- 撇

						- p + q != x + y

					- 捺

						- p - q != x - y

		- 比较好的解法

			- https://leetcode.com/problems/n-queens/discuss/19808/Accepted-4ms-c++-solution-use-backtracking-and-bitmask-easy-understand

			- https://leetcode.com/problems/n-queens/discuss/19810/Fast-short-and-easy-to-understand-python-solution-11-lines-76ms

	- [36. 有效的数独](https://leetcode-cn.com/problems/valid-sudoku/)

		亚马逊、苹果、微软在半年内面试中考过

		- 关键点

			- 九宫格的表示

				- box_index = (row/3)*3+column/3

	- [37. 解数独](https://leetcode-cn.com/problems/sudoku-solver/)

		亚马逊、华为、微软在半年内面试中考过

		- 可能解法

			- 1. 回溯

				- 正常回溯，ij遍历棋盘格

			- 2. 回溯

				- 特点，它把棋盘上9*9看作是81层level的回溯，到最底层则成功

			- 3. 回溯

				- 特点

					- 接近人类的处理方式

					- 这个算法在空元素选择上，可以用A*算法优化

				- 步骤

					- 1. 初始化所有可用的数字加入行、列、块的dict

					- 2. 找出所有的空格子

					- 3. 用backtrack，循环所有的空格子。

					- 4. 退出条件，空格子数组为空

- 3. 双向BFS的实现、特性和题解

	- 双向bfs

		- 同时从开始点和终点做bfs搜索，当有结点在2边都存在时，即找到

	- [127. 单词接龙](https://leetcode-cn.com/problems/word-ladder/)

		亚马逊、Facebook、谷歌在半年内面试中考过

		- 可能解法

			- 1. BFS

			- 2. DFS

				- 要及时剪枝

			- 3. Two-ended bfs

				- 2个set开始处理，每一次从少的那个set开始处理

				- 和bfs区别是bfs用queue，这边是set

- 4. 启发式搜索的实现、特性和题解

	- A*搜索

		- 定义

			- 就是普通的BFS或者DFS，在queue里出队列的时候，做一定的选择

		- 启发式函数(h(n))

			- 用来评价结点的分数，分数越高，越容易完成目标

	- [1091. 二进制矩阵中的最短路径](https://leetcode-cn.com/problems/shortest-path-in-binary-matrix/)

		亚马逊、字节跳动、Facebook 在半年内面试中考过

		- 可能解法

			- 1. dp

			- 2. BFS

			- 3. A*

				- 就是每一次bfs以后，用h(n)计算哪个结点比较有前途

	- [773. 滑动谜题](https://leetcode-cn.com/problems/sliding-puzzle/)

		微软、谷歌、Facebook 在半年内面试中考过

		- 可能解法

			- 1. DFS

			- 2. BFS

				- 比dfs好一点

			- 3. A*

				- 1. 定义一个方向向量

					- 当棋盘是2*3的时候，需要创建一个map是0-5的字典，值是能移动的index

	- 关键点

		- 参考链接都是比较好的解法，需要自己看一下

	- 参考链接

		- [A* 代码模板](https://shimo.im/docs/8CzMlrcvbWwFXA8r/read)

		- [相似度测量方法](https://dataaspirant.com/2015/04/11/five-most-popular-similarity-measures-implementation-in-python/)

		- [二进制矩阵中的最短路径的 A* 解法](https://leetcode.com/problems/shortest-path-in-binary-matrix/discuss/313347/A*-search-in-Python)

		- [8 puzzles 解法比较](https://zxi.mytechroad.com/blog/searching/8-puzzles-bidirectional-astar-vs-bidirectional-bfs/)

## 第15课 AVL和红黑树的实现和特性

- 平衡二叉树

	- 存在目的

		- 为了使搜索速度逼近理论值，需要时时平衡左右结点的个数

	- 实现

		- 2-3 tree

		- AA tree

		- AVL tree

		- B tree

		- Red Black tree

		- Scapegoat tree

		- splay tree

		- treap

		- 以上都需要了解一下

	- AVL tree

		- 发明者

			- G.M.Adelson-Velsky & Evgenii Landis

		- Balanced Factor

			- 定义

				- 左子树的高度-右子树的高度(有时候相反)

			- 取值范围

				- -1，0，1

		- 通过旋转操作来进行平衡(4种)

			- 1. 左旋

				- 子树形态

					- 右右子树，ABC，A是根，C是叶子

				- 转法

					- 把B做根，A变左叶子，C变右叶

			- 2. 右旋

				- 子树形态

					- 左左子树，CBA，C是叶子，A是根

				- 转法

					- 把B做根，A变右叶子，C变左叶子

			- 3. 左右旋

				- 子树形态

					- 左右子树，BCA，A是根，C是叶子，B<C<A

				- 转法

					- 1. B变叶子结点，C变中间结点

					- 2. C变根，B为左叶子，A为右叶子

			- 4. 右左旋

				- 子树形态

					- 右左子树，ACB，A是根，C是叶子，A<C<B

				- 转法

					- 1. B变叶子，C变中间

					- 2. C变根，A为左叶子，B为右叶子

			- 当左右子树下有其他子树时，那些子树应该滑到新的结点的子结点的空余位置上

		- [Self-balancing binary search tree](https://en.wikipedia.org/wiki/Self-balancing_binary_search_tree)

		- 不足

			- 结点需要存储额外信息

			- 调整次数频繁

- 近似平衡二叉树

	- 特点

		- 相对于AVL的话，有一些tradeoff

	- Red Black Tree

		- 特点

			- 确保左右子树的高度差小于2倍

		- 限定条件

			- 每个结点要么时红色，要么是黑色

			- 根结点是黑色

			- 每个叶子结点(NIL结点，空结点)是黑色的

			- 不能有相邻接的2个红色结点

			- 从任一结点到其每个叶子的所有路径都包含相同数目的黑色结点

		- 关键性质

			- 从根到叶子的最长的可能路径不多于最短的可能路径的2倍长

- Red Black Tree与AVL Tree对比

	- lookup speed: AVL > RB

		- AVL trees provide faster lookups than Red Black Trees because they are more strictly balanced.

	- insertion and deletion speed: RB > AVL

		- Red Black Trees provide faster insertion and removal operations than AVL trees as fewer rotations are done due to relatively relaxed balancing.

	- extra space: RB < AVL

		- AVL trees store balance factors or heights with each node, thus requires storage for an integer per node whereas Red Black Tree requires only 1 bit of information per node.

	- Usage: RB in codes, AVL in db

		- Red Black Trees are used in most of the language librarieslike map, multimap, multisetin C++ whereas AVL trees are used in databases where faster retrievals are required. 

## 作业

- easy

	- [x] [70. 爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)

		阿里，腾讯，字节跳动，202005半年常考

- medium

	- [208. 实现 Trie (前缀树)](https://leetcode-cn.com/problems/implement-trie-prefix-tree/)

		亚马逊、微软、谷歌在半年内面试中考过

	- [547. 朋友圈](https://leetcode-cn.com/problems/friend-circles/)

		亚马逊、Facebook、字节跳动在半年内面试中考过

	- [200. 岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)

		近半年内，亚马逊在面试中考查此题达到 361 次

		- [ ] 用union found set实现一下

	- [130. 被围绕的区域](https://leetcode-cn.com/problems/surrounded-regions/)

		亚马逊、eBay、谷歌在半年内面试中考过

	- [36. 有效的数独](https://leetcode-cn.com/problems/valid-sudoku/)

		亚马逊、苹果、微软在半年内面试中考过

	- [22. 括号生成](https://leetcode-cn.com/problems/generate-parentheses/)

		亚马逊、Facebook、字节跳动在半年内面试中考过

	- [127. 单词接龙](https://leetcode-cn.com/problems/word-ladder/)

		亚马逊、Facebook、谷歌在半年内面试中考过

		- 用双向bfs

		- [x] 20200723

		- [ ] 20200724

		- [ ] 20200730

	- [433. 最小基因变化](https://leetcode-cn.com/problems/minimum-genetic-mutation/)

		谷歌、Twitter、腾讯在半年内面试中考过

		- 双向bfs

- hard

	- [212. 单词搜索 II](https://leetcode-cn.com/problems/word-search-ii/)

		亚马逊、微软、苹果在半年内面试中考过

	- [51. N皇后](https://leetcode-cn.com/problems/n-queens/)

		亚马逊、苹果、字节跳动在半年内面试中考过

	- [37. 解数独](https://leetcode-cn.com/problems/sudoku-solver/)

		亚马逊、华为、微软在半年内面试中考过

- 总结双向 BFS 代码模版，请同学们提交在学习总结中


# 覃老师建议的leetcode题目


## array

- easy

	- [70. 爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)

		阿里，腾讯，字节跳动，202005半年常考

		- [x] 20200521

		- [x] 20200522

		- [x] 20200528

- medium

	- [11. 盛最多水的容器](https://leetcode-cn.com/problems/container-with-most-water/)

		腾讯，百度，字节跳动在近半年内常考(202005)

		- [x] 20200521

		- [x] 20200522

		- [x] 20200528

	- [15. 三数之和](https://leetcode-cn.com/problems/3sum/)

		(高频老题)国内，国际大厂历年面试高频老题

		- [x] 20200616

		- [x] 20200617

		- [x] 20200623

## linkedlist

解法固定，熟能生巧

- easy

	- [141. 环形链表](https://leetcode-cn.com/problems/linked-list-cycle/)

		阿里，字节跳动，腾讯202005半年内常考

	- [206. 反转链表](https://leetcode-cn.com/problems/reverse-linked-list/)

		字节跳动，亚马逊202005半年内常考

		- [x] 20200621

		- [x] 20200622

		- [x] 20200628

- medium

	- [24. 两两交换链表中的结点](https://leetcode-cn.com/problems/swap-nodes-in-pairs/)

		阿里，字节跳动202005半年内常考

	- [142. 环形链表 II](https://leetcode-cn.com/problems/linked-list-cycle-ii/)

- hard

	- [25. K 个一组翻转链表](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/)

		字节跳动，猿辅导202005半年内常考

## stack

- easy

	- [20. 有效的括号](https://leetcode-cn.com/problems/valid-parentheses/)

		亚马逊、JPMorgan 在半年内面试常考

	- [155. 最小栈](https://leetcode-cn.com/problems/min-stack/)

		亚马逊在半年内面试常考

- hard

	- [84. 柱状图中最大的矩形](https://leetcode-cn.com/problems/largest-rectangle-in-histogram/)

		亚马逊、微软、字节跳动在半年内面试中考过

	- [239. 滑动窗口最大值](https://leetcode-cn.com/problems/sliding-window-maximum/)

		亚马逊在半年内面试常考

## Tree

- easy

	- [590. N叉树的后序遍历](https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/)

		亚马逊在半年内面试中考过

- medium

	- [102. 二叉树的层序遍历](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/)

		字节跳动、亚马逊、微软在半年内面试中考过，面试过程中出现频率排第三

## Heap

- easy

	- [面试题40. 最小的k个数](https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/)

## recursion

- easy

	- [226. 翻转二叉树](https://leetcode-cn.com/problems/invert-binary-tree/)

		谷歌、字节跳动、Facebook 在半年内面试中考过

	- [104. 二叉树的最大深度](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/)

		亚马逊、微软、字节跳动在半年内面试中考过

	- [111. 二叉树的最小深度](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/)

		Facebook、字节跳动、谷歌在半年内面试中考过

- medium

	- [22. 括号生成](https://leetcode-cn.com/problems/generate-parentheses/)

		字节跳动、亚马逊、Facebook 在半年内面试中考过

	- [98. 验证二叉搜索树](https://leetcode-cn.com/problems/validate-binary-search-tree/)

		谷歌、字节跳动、Facebook 在半年内面试中考过

- hard

	- [297. 二叉树的序列化与反序列化](https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/)

		Facebook、亚马逊在半年内面试常考

## divide&conquer

- medium

	- [50. Pow(x, n)](https://leetcode-cn.com/problems/powx-n/)

		Facebook 在半年内面试常考

## backtracking

- medium

	- [78. 子集](https://leetcode-cn.com/problems/subsets/)

		Facebook、字节跳动、亚马逊在半年内面试中考过

	- [17. 电话号码的字母组合](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/)

		亚马逊在半年内面试常考

- hard

	- [51. N皇后](https://leetcode-cn.com/problems/n-queens/)

		字节跳动、苹果、谷歌在半年内面试中考过

## BFS/DFS

- medium

	- [433. 最小基因变化](https://leetcode-cn.com/problems/minimum-genetic-mutation/)

	- [515. 在每个树行中找最大值](https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/)

		微软、亚马逊、Facebook 在半年内面试中考过

## greedy

- medium

	- [322. 零钱兑换](https://leetcode-cn.com/problems/coin-change/)

		亚马逊在半年内面试中常考

## binarysearching

- easy

	- [69. x 的平方根](https://leetcode-cn.com/problems/sqrtx/)

		字节跳动、微软、亚马逊在半年内面试中考过

	- [367. 有效的完全平方数](https://leetcode-cn.com/problems/valid-perfect-square/)

		亚马逊在半年内面试中考过

## dynamic programming

- normal

	- easy

		- [53. 最大子序和](https://leetcode-cn.com/problems/maximum-subarray/)

			亚马逊、字节跳动在半年内面试常考

		- [198. 打家劫舍](https://leetcode-cn.com/problems/house-robber/)

			字节跳动、谷歌、亚马逊在半年内面试中考过

		- [121. 买卖股票的最佳时机](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/)

			亚马逊、字节跳动、Facebook 在半年内面试中常考

		- [122. 买卖股票的最佳时机 II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)

			亚马逊、字节跳动、微软在半年内面试中考过

	- medium

		- [1143. 最长公共子序列](https://leetcode-cn.com/problems/longest-common-subsequence/)

			字节跳动、谷歌、亚马逊在半年内面试中考过

		- [120. 三角形最小路径和](https://leetcode-cn.com/problems/triangle/)

			亚马逊、苹果、字节跳动在半年内面试考过

		- [152. 乘积最大子数组](https://leetcode-cn.com/problems/maximum-product-subarray/)

			亚马逊、字节跳动、谷歌在半年内面试中考过

		- [213. 打家劫舍 II](https://leetcode-cn.com/problems/house-robber-ii/)

			字节跳动在半年内面试中考过

		- [309. 最佳买卖股票时机含冷冻期](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/)

			谷歌、亚马逊在半年内面试中考过

		- [714. 买卖股票的最佳时机含手续费](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/)

	- hard

		- [123. 买卖股票的最佳时机 III](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/)

			字节跳动在半年内面试中考过

		- [188. 买卖股票的最佳时机 IV](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/)

			谷歌、亚马逊、字节跳动在半年内面试中考过

- advance

	- medium

		- [62. 不同路径](https://leetcode-cn.com/problems/unique-paths/)

			Facebook、亚马逊、微软在半年内面试中考过

		- [63. 不同路径 II](https://leetcode-cn.com/problems/unique-paths-ii/)

			谷歌、美团、微软在半年内面试中考过

		- [322. 零钱兑换](https://leetcode-cn.com/problems/coin-change/)

			亚马逊在半年内面试中常考

		- [279. 完全平方数](https://leetcode-cn.com/problems/perfect-squares/)

			亚马逊、谷歌在半年内面试中考过

		- [55. 跳跃游戏](https://leetcode-cn.com/problems/jump-game/)

			亚马逊、华为、Facebook 在半年内面试中考过

		- [518. 零钱兑换 II](https://leetcode-cn.com/problems/coin-change-2/)

			亚马逊、字节跳动在半年内面试中考过

	- hard

		- [72. 编辑距离](https://leetcode-cn.com/problems/edit-distance/)

			重点

		- [45. 跳跃游戏 II](https://leetcode-cn.com/problems/jump-game-ii/)

			亚马逊、华为、字节跳动在半年内面试中考过

		- [980. 不同路径 III](https://leetcode-cn.com/problems/unique-paths-iii/)

			谷歌在半年内面试中考过

## trie

- medium

	- [208. 实现 Trie (前缀树)](https://leetcode-cn.com/problems/implement-trie-prefix-tree/)

		亚马逊、微软、谷歌在半年内面试中考过

- hard

	- [212. 单词搜索 II](https://leetcode-cn.com/problems/word-search-ii/)

		亚马逊、微软、苹果在半年内面试中考过

## disjoint set

- medium

	- [547. 朋友圈](https://leetcode-cn.com/problems/friend-circles/)

		亚马逊、Facebook、字节跳动在半年内面试中考过

	- [200. 岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)

		近半年内，亚马逊在面试中考查此题达到 361 次

		- [ ] 用union found set实现一下

	- [130. 被围绕的区域](https://leetcode-cn.com/problems/surrounded-regions/)

		亚马逊、eBay、谷歌在半年内面试中考过

## pruning

- easy

	- [70. 爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)

		阿里，腾讯，字节跳动，202005半年常考

- medium

	- [22. 括号生成](https://leetcode-cn.com/problems/generate-parentheses/)

		字节跳动、亚马逊、Facebook 在半年内面试中考过

## heuristic search(A*)

- medium

	- [1091. 二进制矩阵中的最短路径](https://leetcode-cn.com/problems/shortest-path-in-binary-matrix/)

		亚马逊、字节跳动、Facebook 在半年内面试中考过

- hard

	- [773. 滑动谜题](https://leetcode-cn.com/problems/sliding-puzzle/)

		微软、谷歌、Facebook 在半年内面试中考过

## bit operation

- easy

	- [191. 位1的个数](https://leetcode-cn.com/problems/number-of-1-bits/)

		Facebook、苹果在半年内面试中考过

	- [231. 2的幂](https://leetcode-cn.com/problems/power-of-two/)

		谷歌、亚马逊、苹果在半年内面试中考过

	- [190. 颠倒二进制位](https://leetcode-cn.com/problems/reverse-bits/)

		苹果在半年内面试中考过

- medium

	- [338. 比特位计数](https://leetcode-cn.com/problems/counting-bits/)

		字节跳动、Facebook、MathWorks 在半年内面试中考过

- hard

	- [51. N皇后](https://leetcode-cn.com/problems/n-queens/)

		字节跳动、苹果、谷歌在半年内面试中考过

	- [52. N皇后 II](https://leetcode-cn.com/problems/n-queens-ii/)

		亚马逊在半年内面试中考过

## LRU Cache

- medium

	- [146. LRU缓存机制](https://leetcode-cn.com/problems/lru-cache/)

		亚马逊、字节跳动、Facebook、微软在半年内面试中常考

## sorting

- easy

	- [1122. 数组的相对排序](https://leetcode-cn.com/problems/relative-sort-array/)

		谷歌在半年内面试中考过

	- [242. 有效的字母异位词](https://leetcode-cn.com/problems/valid-anagram/)

		Facebook、亚马逊、谷歌在半年内面试中考过

- medium

	- [1244. Leetcode排行榜](https://leetcode-cn.com/problems/design-a-leaderboard/)

		Bloomberg 在半年内面试中考过

	- [56. 合并区间](https://leetcode-cn.com/problems/merge-intervals/)

		Facebook、字节跳动、亚马逊在半年内面试中常考

- hard

	- [493. 翻转对](https://leetcode-cn.com/problems/reverse-pairs/)

		字节跳动在半年内面试中考过

## advancing dp

- hard

	- [72. 编辑距离](https://leetcode-cn.com/problems/edit-distance/)

		字节跳动、亚马逊、谷歌在半年内面试中考过

## string related

- easy

	- [709. 转换成小写字母](https://leetcode-cn.com/problems/to-lower-case/)

		谷歌在半年内面试中考过

	- [58. 最后一个单词的长度](https://leetcode-cn.com/problems/length-of-last-word/)

		苹果、谷歌、字节跳动在半年内面试中考过

	- [771. 宝石与石头](https://leetcode-cn.com/problems/jewels-and-stones/)

		亚马逊在半年内面试中考过

	- [14. 最长公共前缀](https://leetcode-cn.com/problems/longest-common-prefix/)

		亚马逊、谷歌、Facebook 在半年内面试中考过

	- [344. 反转字符串](https://leetcode-cn.com/problems/reverse-string/)

		亚马逊、谷歌、苹果在半年内面试中考过

	- [541. 反转字符串 II](https://leetcode-cn.com/problems/reverse-string-ii/)

		亚马逊在半年内面试中考过

	- [557. 反转字符串中的单词 III](https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/)

		微软、字节跳动、华为在半年内面试中考过

	- [917. 仅仅反转字母](https://leetcode-cn.com/problems/reverse-only-letters/)

		字节跳动在半年内面试中考过

	- [125. 验证回文串](https://leetcode-cn.com/problems/valid-palindrome/)

		Facebook 在半年内面试中常考

- medium

	- [151. 翻转字符串里的单词](https://leetcode-cn.com/problems/reverse-words-in-a-string/)

		微软、字节跳动、苹果在半年内面试中考过


# 8th week


## 第16课 位运算

- 1. 位运算基础及实战要点

	- 具体操作符

		- 左移

			- <<

				- 0011->0110

			- 左边如果超了就没了

		- 右移

			- >>

				- 0110->0011

		- 按位或

			- |

				- 0011, 1011->1011

		- 按位与

			- &

				- 0011, 1011->0011

		- 按位取反

			- ~

				- 0011->1100

		- 按位异或

			- ^

				- 0011, 1011->1000

			- 相同为0，不同为1。也可以用“不进位加法”来理解。

			- 特点

				- x^0 = x

				- x^x = 0

				- x^1s = ~x

					- 1s = ~0

				- x^(~x) = 1s

				- c = a^b -> a^c = b, b^c = a

					- 交换律

				- a^b^c = a^(b^c) = (a^b)^c

					- 结合律

	- 指定位置的位运算

		- 针对x最右边的n位(含)清零

			- x&(~0<<n)

				- ~0<<n，形如111110000，与的话，那些0位全清了

		- 针对x最左边的n位(含)清零

			- x&((1<<n)-1)

		- 获取x的第n位

			- 值

				- (x>>n)&1

			- 幂值

				- x&(1<<n)

		- 针对x的第n位

			- 设为1

				- x|(1<<n)

			- 设为0

				- x&(~(1<<n))

	- 实战要点

		- 判断奇偶

			- x%2==1 -> (x&1)==1

			- x%2==0 -> (x&1)==0

			- 建议多用位运算，但一般来说，高级编译器都会优化

		- x/2 -> x>>1

			- 例：mid=(left+right)/2 -> mid = (left+right)>>1

		- x=x&(x-1)

			- 清理最低位置的1，注，是最低位置，10 -> 8自己感受

		- x&-x -> 获取最低位的1

		- x&~x -> 0

	- 参考链接

		- [如何从十进制转换为二进制](https://zh.wikihow.com/%E4%BB%8E%E5%8D%81%E8%BF%9B%E5%88%B6%E8%BD%AC%E6%8D%A2%E4%B8%BA%E4%BA%8C%E8%BF%9B%E5%88%B6)

- 2. 位运算实战题目解析

	- [191. 位1的个数](https://leetcode-cn.com/problems/number-of-1-bits/)

		Facebook、苹果在半年内面试中考过

		- 可能解法

			- 暴力法，有几位循环几次

				- 1. 一位一位看

				- 2. 先%2，再/2

				- 3. 先>>1，再&1

			- 4. x=x&(x-1)打掉最低位，有几位循环几次 

	- [231. 2的幂](https://leetcode-cn.com/problems/power-of-two/)

		谷歌、亚马逊、苹果在半年内面试中考过

		- 可能解法

			- 直接用x=x&(x-1)打掉最低位，需要等于0，不是零就不满足

	- [190. 颠倒二进制位](https://leetcode-cn.com/problems/reverse-bits/)

		苹果在半年内面试中考过

		- 可能解法

			- 1. int -> string -> reverse，(这种肯定挂)

			- 2. 直接用位运算翻转int的各个位数

	- [52. N皇后 II](https://leetcode-cn.com/problems/n-queens-ii/)

		亚马逊在半年内面试中考过

		- 关键点

			- 用位运算来检测是否被进攻

			- 递归的时候不需要状态恢复，因为int是copy

	- [338. 比特位计数](https://leetcode-cn.com/problems/counting-bits/)

		字节跳动、Facebook、MathWorks 在半年内面试中考过

		- 关键点

			- DP+位运算

## 第17课 布隆过滤器、LRU Cache

- 1. 布隆过滤器的实现及应用

	- 1. Bloom Filter

		- 定义

			- 由一个很长的二进制向量和一系列随机映射函数组成。

		- 作用

			- 用于检索一个元素是否在一个集合中。

		- 特点

			- 优点

				- 空间效率和查询时间都远远超一般的算法

			- 缺点

				- 有一定误识别率和删除困难

			- 待判定元素判定不在Bloom里面正确率100%，在Bloom里面的话，正确率不是100%

		- 案例

			- 比特币网络

			- 分布式系统(Map-Reduce)-hadoop、search engine

			- Redis缓存

			- 垃圾邮件、评论等的过滤

		- 实现

			- Python

				- [布隆过滤器 Python 代码示例](https://shimo.im/docs/UITYMj1eK88JCJTH/read)

				- [布隆过滤器 Python 实现示例](https://www.geeksforgeeks.org/bloom-filters-introduction-and-python-implementation/)

				- [高性能布隆过滤器 Python 实现示例](https://github.com/jhgg/pybloof)

			- Java

				- [布隆过滤器 Java 实现示例 1](https://github.com/lovasoa/bloomfilter/blob/master/src/main/java/BloomFilter.java)

				- [布隆过滤器 Java 实现示例 2](https://github.com/Baqend/Orestes-Bloomfilter)

		- 参考链接

			- [布隆过滤器的原理和实现](https://www.cnblogs.com/cpselvis/p/6265825.html)

			- [使用布隆过滤器解决缓存击穿、垃圾邮件识别、集合判重](https://blog.csdn.net/tianyaleixiaowu/article/details/74721877)

- 2. LRU Cache的实现、应用和题解

	- Cache

		- 1. 记忆

		- 2. 钱包 - 储物柜

		- 3. 代码模块

	- LRU Cache

		- 关键点

			- 大小

			- 替换策略

				- LRU

					- least recently used

				- LFU

					- least frequently used

		- 实现方式

			- HashTable+DoubleLinkedList

		- 效率

			- O(1)查询

			- O(1)修改、更新

	- [146. LRU缓存机制](https://leetcode-cn.com/problems/lru-cache/)

		亚马逊、字节跳动、Facebook、微软在半年内面试中常考

		- 没什么花头，考代码功底，一定要做一下

	- 参考链接

		- [Understanding the Meltdown exploit – in my own simple words](https://www.sqlpassion.at/archive/2018/01/06/understanding-the-meltdown-exploit-in-my-own-simple-words/)

			- 介绍的是CPU的Socket

		- [替换算法总览](https://en.wikipedia.org/wiki/Cache_replacement_policies)

		- [LRU Cache Python 代码示例](https://shimo.im/docs/CoyPAyXooGcDuLQo/read)

## 第18课 排序算法

- 1. 初级排序和高级排序的实现和特性

	- 参考链接

		- [十大经典排序算法](https://www.cnblogs.com/onepixel/p/7674659.html)

		- [9种经典排序算法可视化动画](https://www.bilibili.com/video/av25136272)

		- [6分钟看完15种排序算法动画展示](https://www.bilibili.com/video/av63851336)

		- [快速排序代码示例](https://shimo.im/docs/TX9bDbSC7C0CR5XO/read)

		- [归并排序代码示例](https://shimo.im/docs/sDXxjjiKf3gLVVAU/read)

		- [堆排序代码示例](https://shimo.im/docs/M2xfacKvwzAykhz6/read)

	- 排序算法

		- 比较类排序

			- 定义

				- 通过比较决定元素间的相对次序，时间复杂度，最低到O(nlogn)，也称非线性时间比较类排序

			- Exchange Sort

				- Bubble Sort

					- 时间

						- 平均|最坏

							- O(n^2)

						- 最好

							- O(n)

					- 空间

						- O(1)

					- 稳定性

						- 稳定

					- 过程

						- 嵌套循环，每次查看相邻的元素，如果逆序，则交换

				- Quick Sort

					- 时间

						- 平均|最好

							- O(nlogn)

						- 最坏

							- O(n^2)

					- 空间

						- O(nlogn)

					- 稳定性

						- 不稳定

					- 过程

						- 取标杆pivot，将小元素放pivot左边，大元素放右边，然后依次对左边和右边的子数组继续快排，递归

			- Insertion Sort

				- Insertion Sort

					- 时间

						- 平均|最坏

							- O(n^2)

						- 最好

							- O(n)

					- 空间

						- O(1)

					- 稳定性

						- 稳定

					- 过程

						- 从前到后逐步构建有序序列；对于未排序元素，在已排序序列中从后向前扫描，找到相应位置并插入

				- 希尔排序

					- 时间

						- 平均

							- O(n^1.3)

						- 最坏

							- O(n^2)

						- 最好

							- O(n)

					- 空间

						- O(1)

					- 稳定性

						- 不稳定

			- Selection Sort

				- Selection Sort

					- 时间

						- 平均

							- O(n^2)

						- 最坏

							- O(n^2)

						- 最好

							- O(n^2)

					- 空间

						- O(1)

					- 稳定性

						- 不稳定

					- 过程

						- 每次找最小值，然后放到待排数组的起始位置

				- Heap Sort

					- 时间

						- O(nlogn)

					- 空间

						- O(1)

					- 稳定性

						- 不稳定

					- 过程

						- 1. 数组元素依次建立小顶堆

						- 2. 依次取堆顶元素，并删除

			- Merge Sort

				- 二路并归排序

				- 多路并归排序

				- 时间

					- O(nlogn)

				- 空间

					- O(n)

				- 稳定性

					- 稳定

				- 过程

					- 1. 把长度n的数组分成2个n/2的子数组

					- 2. 对这2个子序列采用merge sort

					- 3. 把这2个排序好的子序列合并成最终的排序序列

			- Merge Sort 与 Quick Sort相比

				- 有相似性，但过程相反

				- Merge Sort：先排序左右子数组，然后合并2个有序子序列。

				- Quick Sort：先划出左右子数组，左小于右，然后在对左右子数组进行排序

		- 非比较类排序

			- 定义

				- 不通过比较来决定元素间的相对次序，时间复杂度，最低到O(n)，也称线性时间非比较类排序

			- Counting Sort

				- 时间

					- O(n+k)

				- 空间

					- O(n+k)

				- 稳定性

					- 稳定

				- 过程

					- 1. 数据源必须是有确定范围的整数。

					- 2. 将输入的数据值转化为键存储在额外开辟的数组空间中；

					- 3. 然后依次把计数大于1的填充回原数组

			- Bucket Sort

				- 时间

					- 平均

						- O(n+k)

					- 最坏

						- O(n^2)

					- 最好

						- O(n)

				- 空间

					- O(n+k)

				- 稳定性

					- 稳定

				- 过程

					- 1. 假设输入数据服从均匀分布

					- 2. 将数据分到有限数量的桶里

					- 3. 每个桶再分别排序(使用其他排序方法，或者递归方式继续调用桶排序)

			- Radix Sort(基数)

				- 时间

					- O(n*k)

				- 空间

					- O(n+k)

				- 稳定性

					- 稳定

				- 过程

					- 按照低位先排序，然后收集

					- 再按照高位排序，再收集

					- 直到最高位

					- 有时候有些属性是有优先级顺序的，先按低优先级排序，再按高优先级排序

- 2. 特殊排序及实战题目详解

	- [242. 有效的字母异位词](https://leetcode-cn.com/problems/valid-anagram/)

		Facebook、亚马逊、谷歌在半年内面试中考过

		- 可以自己写排序算法，而不是用系统库

	- [56. 合并区间](https://leetcode-cn.com/problems/merge-intervals/)

		Facebook、字节跳动、亚马逊在半年内面试中常考

		- 推荐做法

			- 排序 然后干

	- [493. 翻转对](https://leetcode-cn.com/problems/reverse-pairs/)

		字节跳动在半年内面试中考过

		- 关键点

			- 面试中会出现很多和逆序对相关的题目

		- 可能解法

			- 1. 暴力法，2次嵌套循环

			- 2. merge-sort

				- 主要是在排序中，加入了逆序对统计的动作

			- 3. 树状数组

## 作业

- easy

	- [191. 位1的个数](https://leetcode-cn.com/problems/number-of-1-bits/)

		Facebook、苹果在半年内面试中考过

		- [x] 20200711

		- [x] 20200712

		- [x] 20200718

	- [231. 2的幂](https://leetcode-cn.com/problems/power-of-two/)

		谷歌、亚马逊、苹果在半年内面试中考过

		- [x] 20200711

		- [x] 20200712

		- [x] 20200718

	- [190. 颠倒二进制位](https://leetcode-cn.com/problems/reverse-bits/)

		苹果在半年内面试中考过

	- [1122. 数组的相对排序](https://leetcode-cn.com/problems/relative-sort-array/)

		谷歌在半年内面试中考过

	- [242. 有效的字母异位词](https://leetcode-cn.com/problems/valid-anagram/)

		Facebook、亚马逊、谷歌在半年内面试中考过

- medium

	- [146. LRU缓存机制](https://leetcode-cn.com/problems/lru-cache/)

		亚马逊、字节跳动、Facebook、微软在半年内面试中常考

	- [1244. Leetcode排行榜](https://leetcode-cn.com/problems/design-a-leaderboard/)

		Bloomberg 在半年内面试中考过

	- [56. 合并区间](https://leetcode-cn.com/problems/merge-intervals/)

		Facebook、字节跳动、亚马逊在半年内面试中常考

- hard

	- [51. N皇后](https://leetcode-cn.com/problems/n-queens/)

		字节跳动、苹果、谷歌在半年内面试中考过

	- [52. N皇后 II](https://leetcode-cn.com/problems/n-queens-ii/)

		亚马逊在半年内面试中考过

	- [493. 翻转对](https://leetcode-cn.com/problems/reverse-pairs/)

		字节跳动在半年内面试中考过

- 用自己熟悉的编程语言，手写各种初级排序代码，提交到学习总结中。


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

	- 参考链接

		- [不可变字符串](https://lemire.me/blog/2017/07/07/are-your-strings-immutable/)

		- [Atoi 代码示例](https://shimo.im/docs/5kykuLmt7a4DdjSP/read)

	- [5. 最长回文子串](https://leetcode-cn.com/problems/longest-palindromic-substring/)

		亚马逊、字节跳动、华为在半年内面试中常考

		- 自己写

	- 字符串基础问题

		- [709. 转换成小写字母](https://leetcode-cn.com/problems/to-lower-case/)

			谷歌在半年内面试中考过

		- [58. 最后一个单词的长度](https://leetcode-cn.com/problems/length-of-last-word/)

			苹果、谷歌、字节跳动在半年内面试中考过

		- [771. 宝石与石头](https://leetcode-cn.com/problems/jewels-and-stones/)

			亚马逊在半年内面试中考过

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

	- 字符串操作问题

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

		- [541. 反转字符串 II](https://leetcode-cn.com/problems/reverse-string-ii/)

			亚马逊在半年内面试中考过

		- [151. 翻转字符串里的单词](https://leetcode-cn.com/problems/reverse-words-in-a-string/)

			微软、字节跳动、苹果在半年内面试中考过

			- 可能解法

				- 1. split，reverse，join

				- 2. reverse 所有的char，然后再翻转每个单词

		- [557. 反转字符串中的单词 III](https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/)

			微软、字节跳动、华为在半年内面试中考过

		- [917. 仅仅反转字母](https://leetcode-cn.com/problems/reverse-only-letters/)

			字节跳动在半年内面试中考过

	- 异位词问题

		- [242. 有效的字母异位词](https://leetcode-cn.com/problems/valid-anagram/)

			亚马逊、Facebook、谷歌在半年内面试中考过

			- 可能解法

				- 1. 暴力，sort，sorted_str看是否相等

				- 2. hash map-> 统计词频

		- [49. 字母异位词分组](https://leetcode-cn.com/problems/group-anagrams/)

			亚马逊在半年内面试中常考

		- [438. 找到字符串中所有字母异位词](https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/)

			Facebook 在半年内面试中常考

			- 关键点

				- 有个大小为目标字符串的窗口，然后平移

	- 回文串问题

		- [125. 验证回文串](https://leetcode-cn.com/problems/valid-palindrome/)

			Facebook 在半年内面试中常考

		- [680. 验证回文字符串 Ⅱ](https://leetcode-cn.com/problems/valid-palindrome-ii/)

			Facebook 在半年内面试中常考

		- [5. 最长回文子串](https://leetcode-cn.com/problems/longest-palindromic-substring/)

			亚马逊、字节跳动、华为在半年内面试中常考

- 2. 高级字符串算法

	- 最长子串、子序列问题

		- [1143. 最长公共子序列](https://leetcode-cn.com/problems/longest-common-subsequence/)

			字节跳动、谷歌、亚马逊在半年内面试中考过

		- [72. 编辑距离](https://leetcode-cn.com/problems/edit-distance/)

			字节跳动、亚马逊、谷歌在半年内面试中考过

		- [5. 最长回文子串](https://leetcode-cn.com/problems/longest-palindromic-substring/)

			亚马逊、字节跳动、华为在半年内面试中常考

	- 字符串+DP问题

		- [10. 正则表达式匹配](https://leetcode-cn.com/problems/regular-expression-matching/)

			Facebook、微软、字节跳动在半年内面试中考过

			- [题解](https://leetcode-cn.com/problems/regular-expression-matching/solution/ji-yu-guan-fang-ti-jie-gen-xiang-xi-de-jiang-jie-b/)

		- [44. 通配符匹配](https://leetcode-cn.com/problems/wildcard-matching/)

			Facebook、微软、字节跳动在半年内面试中考过

		- [115. 不同的子序列](https://leetcode-cn.com/problems/distinct-subsequences/)

			MathWorks 在半年内面试中考过

## 作业

- easy

	- [387. 字符串中的第一个唯一字符](https://leetcode-cn.com/problems/first-unique-character-in-a-string/)

		亚马逊、微软、Facebook 在半年内面试中考过

		- [x] 20200719

		- [x] 20200720

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

