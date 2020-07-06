学习笔记

针对 单词搜索2的trie树方式的时间复杂度分析如下
O(n*m*4*3^(k-1))，n*m是board总数，4是第一次可以尝试4个方向，3是后面的只能尝试其他3个没试过的方向。

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

