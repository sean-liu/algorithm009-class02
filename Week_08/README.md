学习笔记
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

		- [ ] 20200718

	- [231. 2的幂](https://leetcode-cn.com/problems/power-of-two/)

		谷歌、亚马逊、苹果在半年内面试中考过

		- [x] 20200711

		- [x] 20200712

		- [ ] 20200718

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

