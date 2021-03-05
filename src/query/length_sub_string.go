package query

// 示例 1:
//
// 输入: "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 示例 2:
//
// 输入: "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:
//
// 输入: "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

// 思路： 滑动窗口方式，
// 窗口内的值的和保存在一个变量中；
// 通过不断的往右滑动来算出当前窗口的值，并与保存的最大值作比较；
// 当窗口滑动到最右边时终止滑动；

func LengthOfSubString(s string) int {
	
	// 用来存储字符串的位置，每次进行滑动记录字符串在窗口的位置
	m := make(map[rune]int)
	start, max := -1, 0
	
	for k, v := range s {
		
		// 如果之前字符串出现过了，并且这次出现比之前的位置靠后，把窗口移动从新开始窗口
		if last, ok := m[v]; ok && last > start {
			start = last
		}
		
		m[v] = k
		
		// 窗口结束 - 窗口界截止点大于当前最大 重新赋值
		if k-start > max { // 保存最大值
			max = k - start
		}
		
	}
	
	return max
}