package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//0 < frame.length <= 200
//0 < frame[0].length <= 200

func jewelleryValue(frame [][]int) int {
	rowNum := len(frame)    //2
	colNum := len(frame[0]) //3
	dp := make([][]int, rowNum)
	for i := 0; i < rowNum; i++ {
		dp[i] = make([]int, colNum)
	}
	dp[0][0] = frame[0][0]
	for i := 1; i < colNum; i++ {
		dp[0][i] = frame[0][i] + dp[0][i-1]
	}
	for j := 1; j < rowNum; j++ {
		dp[j][0] = frame[j][0] + dp[j-1][0]
	}
	for i := 1; i < rowNum; i++ {
		for j := 1; j < colNum; j++ {
			if dp[i-1][j] > dp[i][j-1] {
				dp[i][j] = dp[i-1][j] + frame[i][j]
			} else {
				dp[i][j] = dp[i][j-1] + frame[i][j]
			}
		}
	}
	return dp[rowNum-1][colNum-1]
}

// 提示：
//
// 0 <= s.length <= 5 * 104
// s 由英文字母、数字、符号和空格组成
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	left, right := 0, 0
	maxLength := 0
	result := make(map[string]struct{})
	for left < len(s) && right < len(s) {
		_, ok := result[string(s[right])]
		if !ok {
			result[string(s[right])] = struct{}{}
			right++
			maxLength = max(maxLength, right-left)
		} else {
			delete(result, string(s[left]))
			left++
		}
	}
	return maxLength
}

// 丑数 就是只包含质因数 2、3 和 5 的正整数。
//
// 给你一个整数 n ，请你判断 n 是否为 丑数 。如果是，返回 true ；否则，返回 false 。
//
// 示例 1：
//
// 输入：n = 6
// 输出：true
// 解释：6 = 2 × 3
// 示例 2：
//
// 输入：n = 1
// 输出：true
// 解释：1 没有质因数，因此它的全部质因数是 {2, 3, 5} 的空集。习惯上将其视作第一个丑数。
// 示例 3：
//
// 输入：n = 14
// 输出：false
// 解释：14 不是丑数，因为它包含了另外一个质因数 7 。
//
// 提示：
//
// -231 <= n <= 231 - 1
func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}
	remain := n
	for remain > 1 {
		if remain%2 != 0 && remain%3 != 0 && remain%5 != 0 {
			return false
		}
		if remain%2 == 0 {
			remain = remain / 2
		}
		if remain%3 == 0 {
			remain = remain / 3
		}
		if remain%5 == 0 {
			remain = remain / 5
		}
	}
	return true
}

func isUglyV2(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}
	if n%2 != 0 && n%3 != 0 && n%5 != 0 {
		return false
	}

	return true
}

func Do(ctx context.Context, wg *sync.WaitGroup) {
	ctx, cancle := context.WithTimeout(ctx, time.Second*2)
	defer func() {
		cancle()
		wg.Done()
	}()

	done := make(chan struct{}, 1) //执行成功的channel
	go func(ctx context.Context) {
		fmt.Println("go goroutine")
		time.Sleep(time.Second * 3)
		done <- struct{}{} //发送完成的信号
	}(ctx)

	select {
	case <-ctx.Done(): //超时
		fmt.Printf("timeout,err:%v\n", ctx.Err())
	case <-time.After(3 * time.Second): //超时第二种方法
		fmt.Printf("after 1 sec.")
	case <-done: //程序正常结束
		fmt.Println("done")
	}

}

// 给定一个字符串 s ，找到 它的第一个不重复的字符，并返回它的索引 。如果不存在，则返回 -1 。
//
// 示例 1：
//
// 输入: s = "leetcode"
// 输出: 0
// 示例 2:
//
// 输入: s = "loveleetcode"
// 输出: 2
// 示例 3:
//
// 输入: s = "aabb"
// 输出: -1
//
// 提示:
//
// 1 <= s.length <= 105
// s 只包含小写字母

func firstUniqChar(s string) int {
	chCnt := [26]int{}
	for i := 0; i < len(s); i++ {
		chCnt[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if chCnt[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}

//在股票交易中，如果前一天的股价高于后一天的股价，则可以认为存在一个「交易逆序对」。请设计一个程序，输入一段时间内的股票交易记录 record，返回其中存在的「交易逆序对」总数。
//
//
//
//示例 1:
//
//输入：record = [9, 7, 5, 4, 6]
//输出：8
//解释：交易中的逆序对为 (9, 7), (9, 5), (9, 4), (9, 6), (7, 5), (7, 4), (7, 6), (5, 4)。
//
//
//限制：
//
//0 <= record.length <= 50000

func reversePairs(record []int) int {
	if len(record) < 2 {
		return 0
	}
	dp := make([]int, len(record))
	dp[0] = 0
	tmp := 0
	for i := 1; i < len(record); i++ {
		tmp = 0
		for j := 0; j < i; j++ {
			if record[j] > record[i] {
				tmp++
			}
		}
		dp[i] = dp[i-1] + tmp
	}
	return dp[len(dp)-1]
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	newHeadA, newHeadB := headA, headB
	lengthAll := 0
	for newHeadA != nil {
		lengthAll++
		newHeadA = newHeadA.Next
	}

	for newHeadB != nil {
		lengthAll++
		newHeadB = newHeadB.Next
	}
	newHeadA, newHeadB = headA, headB
	fmt.Println(lengthAll, newHeadA, newHeadB)
	length := 0
	for length < lengthAll && headA != nil && headB != nil {
		fmt.Println("A=", headA.Val, " B=", headB.Val, " length=", length)
		if headB == headA {
			return headA
		}

		if headA.Next == nil {
			headA.Next = newHeadB
		}
		if headB.Next == nil {
			headB.Next = newHeadA
		}
		headA = headA.Next
		headB = headB.Next
		length++
	}
	return nil
}

func getIntersectionNodeV2(headA, headB *ListNode) *ListNode {
	newHeadA, newHeadB := headA, headB
	lengthAll := 0
	for newHeadA != nil {
		lengthAll++
		newHeadA = newHeadA.Next
	}

	for newHeadB != nil {
		lengthAll++
		newHeadB = newHeadB.Next
	}
	newHeadA, newHeadB = headA, headB
	fmt.Println(lengthAll, newHeadA, newHeadB)
	length := 0
	for length < lengthAll {
		fmt.Println("A=", headA.Val, " B=", headB.Val, " length=", length)
		if headB == headA {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
		length++

		if headA == nil {
			headA = newHeadB
		}
		if headB == nil {
			headB = newHeadA
		}
	}
	return nil
}

// 0 <= scores.length <= 105
// -109 <= scores[i] <= 109
// scores 是一个非递减数组
// -109 <= target <= 109
func countTarget(scores []int, target int) int {
	l := BinarySearch(scores, target, func(intA, intB int) bool {
		return intA > intB
	})
	r := BinarySearch(scores, target, func(intA, intB int) bool {
		return intA >= intB
	})
	fmt.Println(l, r)
	return r - l
}

func BinarySearch(scores []int, target int, compare func(intA, intB int) bool) int {
	l, r := 0, len(scores)
	for l < r {
		mid := l + (r-l)/2
		if compare(target, scores[mid]) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func testWithGoroutineTimeOut() {
	var wg sync.WaitGroup
	done := make(chan struct{})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(time.Second * time.Duration(1))
			defer wg.Done()
		}()
	}
	// wg.Wait()此时也要go出去,防止在wg.Wait()出堵住
	go func() {
		wg.Wait()
		close(done)
	}()
	select {
	// 正常结束完成
	case <-done:
		fmt.Println("done")
	// 超时
	case <-time.After(500 * time.Millisecond):
		fmt.Println("timeout")
	}
}

// n == nums.length
// 1 <= n <= 104
// 0 <= nums[i] <= n
func missingNumber(nums []int) int {
	sum := 0
	for i := 1; i <= len(nums); i++ {
		sum += i
	}
	for _, num := range nums {
		sum -= num
	}
	return sum
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 提示：
//
// 1 ≤ cnt ≤ 二叉搜索树元素个数
func findTargetNode(root *TreeNode, cnt int) int {
	ReverseMidOrder(root, &cnt)
	return cnt
}

func ReverseMidOrder(root *TreeNode, cnt *int) int {
	if root.Right != nil {
		ReverseMidOrder(root.Left, cnt)
	}
	if *cnt == 0 {
		return root.Val
	}
	if root.Left != nil {
		*cnt--
		ReverseMidOrder(root.Right, cnt)
	}
	return -1
}

func midOrder(root *TreeNode, orders *[]int) {
	if root.Left != nil {
		midOrder(root.Left, orders)
	}
	*orders = append(*orders, root.Val)
	if root.Right != nil {
		midOrder(root.Right, orders)
	}
}

// 1 ≤ cnt ≤ 二叉搜索树元素个数
func kthLargest(root *TreeNode, k int) int {
	count := 0
	var f func(root *TreeNode) int
	f = func(root *TreeNode) int {
		if root.Right != nil {
			f(root.Right)
		}
		k--
		if k == 0 {
			count = root.Val
		}
		if root.Left != nil {
			f(root.Left)
		}
		return count
	}
	return f(root)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func maxDepthV2(root *TreeNode) int {
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftDeep := dfs(root.Left)
		rightDeep := dfs(root.Right)
		return max(leftDeep, rightDeep) + 1
	}
	return dfs(root)
}

func twoSumV2(price []int, target int) []int {
	if price[0] > target {
		return []int{}
	}
	l, r := 0, len(price)-1
	f := func(price []int, target int) bool {

		l, r := 0, len(price)-1
		for l <= r {
			mid := l + (r-l)/2
			if target == price[mid] {
				return true
			} else if target > price[mid] {
				l = mid + 1
			} else {
				r = mid
			}
		}
		return false
	}

	for price[r] > target && l <= r {
		r--
	}
	for l < r {
		fmt.Println(l, r)
		fmt.Println(price[l:r+1], target-price[l])
		if f(price[l:r+1], target-price[l]) == true && l <= r {
			return []int{price[l], target - price[l]}
		}
		l++
	}
	return []int{}
}

func twoSumV3(price []int, target int) []int {
	m := make(map[int]struct{}, len(price))
	for i := 0; i < len(price); i++ {
		m[price[i]] = struct{}{}
	}
	for i := 0; i < len(price); i++ {
		_, ok := m[target-price[i]]
		if ok {
			return []int{price[i], target - price[i]}
		}
	}
	return []int{}
}

// 例 1：
//
// 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
// 输出：[3,3,5,5,6,7]
// 解释：
// 滑动窗口的位置                最大值
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
// 示例 2：
//
// 输入：nums = [1], k = 1
// 输出：[1]
//
// 提示：
//
// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104
// 1 <= k <= nums.length
func maxSlidingWindow(nums []int, k int) []int {
	dp := make([]int, len(nums)-k+1)
	dp[0] = maxRange(nums[:k])
	// dp[i] = max(nums[i],nums[i],...,nums[i+k-1])
	for i := 1; i < len(nums)-k+1; i++ {
		if nums[i+k-1] > dp[i-1] {
			dp[i] = nums[i+k-1]
		} else if nums[i-1] < dp[i-1] {
			dp[i] = dp[i-1]
		} else {
			dp[i] = maxRange(nums[i : i+k])
		}
	}
	return dp
}

func maxRange(nums []int) int {
	maxT := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxT {
			maxT = nums[i]
		}
	}
	return maxT
}

func maxSlidingWindowV2(nums []int, k int) []int {
	// 特殊情况：如果原始数组为空或者滑动窗口大小为空，则返回空数组
	if len(nums) == 0 || k == 0 {
		return []int{}
	}
	// 队列中储存nums数组中元素的下标，队列中所有下标都要在窗口范围内
	// 队头是每次遍历的最大值的下标，队列中储存的下标对应nums中的元素大小是递减的，方便我们获取每次窗口滑动一位后的最大值
	queue := make([]int, 0)
	// 结果集
	res := make([]int, 0)
	// 滑动窗口右移时
	// 如果还未形成窗口：刚开始遍历时，滑动窗口的左右边界left和right都为0，有一个形成窗口的过程，也就是使整个窗口都在nums数组中；left不动，right向右移，此过程没有最大值
	for i := 0; i < k; i++ {
		// 将right新增的元素nums[right]和队尾的元素nums[queue[len(queue)-1]]的大小进行比较，如果队尾元素大小比新数小，则把队尾元素去掉；直到队尾元素大小比新数大或者队列为空的时候才停止
		for len(queue) > 0 && nums[i] > nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		// 此时将这个新数的索引right加入队列；这样可以保证队列里的元素的大小是递减的，还可以保证队列内元素的索引是升序的（使得在后续每次移除队首的时候最多只需要移除一个元素）
		queue = append(queue, i)
	}
	fmt.Println(queue, res)
	// 将刚形成窗口时的最大值加入结果集中
	res = append(res, nums[queue[0]])
	fmt.Println(queue, res)
	// 如果已经形成窗口：当窗口形成时，也就是整个窗口都已经在nums数组中；left和right一起向右移动，此过程每移动一次就产生一个最大值
	for i := k; i < len(nums); i++ {
		// 看队首的元素queue[0]储存的索引，是否还在滑动窗口范围内，也就是判断是否新的left>queue[0]，如果不在了则将队列最前端移除
		fmt.Print(" 	i-->", i, "  ")
		fmt.Print(" 	a-->", queue, res)
		if i-k+1 > queue[0] {
			queue = queue[1:]
		}
		fmt.Print(" 	b-->", queue, res)
		// 将right新增的元素nums[right]和队尾的元素nums[queue[len(queue)-1]]的大小进行比较，如果队尾元素大小比新数小，则把队尾元素去掉；直到队尾元素大小比新数大或者队列为空的时候才停止
		for len(queue) > 0 && nums[i] > nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
			fmt.Print(" 	c-->", queue, res)
		}
		// 此时将这个新数的索引right加入队列
		queue = append(queue, i)
		fmt.Print(" 	d-->", queue, res)
		// 滑动窗口移动后，将队列首部保存的新的最大值，加入到结果集中
		res = append(res, nums[queue[0]])
		fmt.Println(" 	e-->", queue, res)
	}

	// 返回结果集
	return res
}

func maxSlidingWindowV3(nums []int, k int) []int {
	ans := make([]int, 0, len(nums)-k+1) // 预先分配好空间
	q := []int{}
	for i, x := range nums {
		// 1. 入
		fmt.Printf(" 	i=%d ,x=%d", i, x)
		for len(q) > 0 && nums[q[len(q)-1]] <= x {
			q = q[:len(q)-1] // 维护 q 的单调性
			fmt.Print(" 	a-->", q, ans)
		}
		q = append(q, i) // 入队
		fmt.Print(" 	b-->", q, ans)
		// 2. 出
		if i-q[0] >= k { // 队首已经离开窗口了
			q = q[1:] // Go 的切片是 O(1) 的
		}
		fmt.Print(" 	b-->", q, ans)
		// 3. 记录答案
		if i >= k-1 {
			// 由于队首到队尾单调递减，所以窗口最大值就是队首
			ans = append(ans, nums[q[0]])
			fmt.Print(" 	e-->", q, ans)
		}
		fmt.Println(" 	end line")
	}
	return ans
}

// 社团共有 num 位成员参与破冰游戏，编号为 0 ~ num-1。成员们按照编号顺序围绕圆桌而坐。社长抽取一个数字 target，从 0 号成员起开始计数，排在第 target 位的成员离开圆桌，且成员离开后从下一个成员开始计数。请返回游戏结束时最后一位成员的编号。
//
// 示例 1：
//
// 输入：num = 7, target = 4
// 输出：1
// 示例 2：
//
// 输入：num = 12, target = 5
// 输出：0
//
// 提示：
//
// 1 <= num <= 10^5
// 1 <= target <= 10^6

// ===== >   业务逻辑层  < ========
func main() {
	root := &TreeNode{Val: 7}
	root.Left = &TreeNode{Val: 3, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 5}}
	root.Right = &TreeNode{Val: 9}

}
