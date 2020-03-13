package data

import "os"

func main() {
	sum(10)
}

/**
使用这里的代码来解释 大O表示法

假设 每条语句的执行时间是 unit_time ， 那sum方法对数据规模是 n 来说总的执行时间是多少？
1.从代码的执行来看  sum := 0 和  i := 0 只执行一次 ， 所以执行时间各为 unit_time
2.从循环来看 sum += i   循环了 n次 所以执行时间为 n * unit_time

解：所以总的时间相加得
T(n) = unit_time + unit_time + n * (unit_time)
T(n) = 2unit_time + n * (unit_time)
T(n) = (2 + n) * unit_time

用 f(n) 来表示
f(n) = (2 + n) * unit_time

思考，当公式中的 n 无限大的时候 ， 常数项对于 公式的影响会越来越小，最后可以修改为
T(n) = O(f(n))

n ： 数据规模，通俗点说就是函数中的 n
f(n) ： 代码总的执行次数和 数据规模的关系
T(n) ： 代码的执行时间( 抽象的，这里不是实际的执行时间，是代码的执行时间和数据规模之间的关系)
O ： 表示代码的执行时间T(n)和 执行次数总和f(n) 成正比关系
 */
func sum(n int) {
	sum := 0
	i := 0
	for ; i < n; i++ {
		sum += i
	}

	os.Exit(sum)

}

/**
继续分析这段代码
假设每段代码的执行时间为 unit_time
1. 循环外的 每行代码的执行时间总和 unit_time + unit_time + unit_time
2. for + j =1 的执行时间 为 2n * unit_time
3. 循环内的循环 2n^2 * unit_time
4. 代码的总执行时间 T(n) = (2n^2 + 2n + 3) * unit_time
通过以上代码分析，我们不知道unit_time的具体执行时间是多少，但是可以分析出 执行时间T(n) 和 n的执行次数是成正比。

最后总结的T(n) = O(f(n))

n ： 数据规模，通俗点说就是函数中的 n
f(n) ： 代码总的执行次数和 数据规模的关系
T(n) ： 代码的执行时间( 抽象的，这里不是实际的执行时间，是代码的执行时间和数据规模之间的关系)
O ： 表示代码的执行时间T(n)和 执行次数总和f(n) 成正比关系

大 O 时间复杂度实际上并不具体表示代码的真正执行时间，而是表示 代码执行时间 随数据规模增长的变化趋势，所以也叫渐进时间复杂度
简称  时间复杂度
 */
func cal(n int) {
	sum := 0
	i := 1
	j := 1
	for ; i < n; i++ {
		j = 1
		for ; j <= n; j++ {
			sum = sum + i*j
		}
	}
	os.Exit(sum)
}

/**
时间复杂度分析方法 ：只关注循环执行次数最多的那段代码
大 O 这种复杂度表示方法只是表示一种趋势。 我们通常会忽略公式中的常量、低阶、系数，只需要记录一个最大阶的量级就可以，所以我们在分析一个算法、
一段代码的时间复杂度时，也只关注循环执行次数最多的那一段代码就可以了。这段代码的执行次数 n 的量级，就是整段代码的时间复杂度。


从下面的例子来看
下面的方法中有两处循环
第一段代码和 n 无关
第二段代码 O(n)
第三段代码 O(n^2)

加法法则，取两个和的最大的值
T1(n) = O(f(n))
T2(n) = O(g(n))
T(n) = T1(n) + T2(n) = max(O(f(n)),O(g(n))) = O(max(f(n),g(n)))

 */
func cal_one(n int) {
	/**
	unit_time + unit_time + 99unit_time
	最后 T(n) = 101 unit_time
	 */
	sum_1 := 0
	p := 1
	for ; p < 100; p++ {
		sum_1 = sum_1 + p
	}

	/**
	unit_time + unit_time + (n-1)unit_time
	T(n) = n * unit_time
	O(n)
	 */
	sum_2 := 0
	q := 1
	for ; q < n; q++ {
		sum_2 = sum_2 + q
	}

	/**
	unit_time + unit_time + unit_time + (n-1) * unit_time + (n-1)^2 * unit_time
	(n-1)^2 * unit_time
	O(n^2)
	 */
	sum_3 := 0
	i := 1
	j := 1
	for ; i <= n; i++ {
		j = 1
		for ; j <= n; j++ {
			sum_3 = sum_3 + i*j
		}
	}
}

/**
前面几行的的时间复杂度就是  T1(n) = O(n)

f函数不是一个简单的操作，那下面的时间复杂度就是
T(n) = T1(n) * T2(n) = O(n * n) = O(n^2)

 */
func cal_two(n int) {
	ret := 0
	i := 1
	for ; i < n; i++ {
		ret = ret + f(i)
	}
	os.Exit(ret)
}

func f(n int) int {
	sum := 0
	i := 1
	for ; i < n; i += 1 {
		sum = sum + i
	}
	return sum
}

/**
复杂度量级
常量阶  O(1)
指数阶  O(2 ^ n)
对数阶  O(log n)
阶乘阶  O(n!)
线性阶  O(n)
线性对数阶 O(n log n)
平方阶  O(n^2)
立方阶  O(n^3)
K次方阶 O(n^k)

对于以上的复杂度量级，可以粗略分为两类 ，多项式量级 和 非多项式量级 其中非多项式量级只有两个  O(2^n)  O(n!)
时间复杂度为非多项式量级的算法问题叫做 NP
当数据规模n越来越大时 ，非多量式量级的算法执行时间会急剧增加，求解问题会无限延长，所以，非多项时间复杂度的算法其实是非常低级的算法
主要来看一下多项式时间复杂度

1. O(1)
首先必须明确的概念，O(1) 只是常量级时间复杂度的一种表示方法，并不是只执行一行代码
比如下面这行代码有三行执行的代码  但是他的时间复杂度还是 O(1)

一般情况下，只要代码中不存在循环、递归，即使存在成千上万行代码，时间复杂度也是O(1)
 */

func O(n int) {
	i := 1
	j := 2
	n = i + j
	os.Exit(n)
}

/**
2. O(logn) O(nlogn)
对数阶时间复杂度，一般来说是最难分析的复杂度

一下的 循环 应该是一个等比数列
2^0  2^1  2^2  2^3  2^x = n
x = log 2^n  时间复杂度就是 O(log 2^n)
 */

func log_time(n int) {
	i := 1
	for i <= n {
		i = i*2
	}
}
