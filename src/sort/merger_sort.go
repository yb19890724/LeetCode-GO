package sort

// 归并排序

// https://segmentfault.com/a/1190000022288880

// 稳定排序算法

// 时间复杂度：最好 最坏 平均 都是 O(n*logn)

// 空间复杂度: O(n)

// 流程：
// 1. 不断的将当前序列平均分割成2个子例子
// 一直分到不能再分割 序列中只剩下一个元素
// 2. 在不断的将2个子序列合并成一个有序序列

// 例：
//       0 1 2 3 4 5
// []int{2 4 8 8 8 12}

// [begin,end) 左闭右开区间  范围进行合并
// [2 4 8)  [8 8 12)


// 自顶向下归并排序，排序范围在 [begin,end) 的数组
func MergeSort(data []int, begin int, end int) []int{
	// 元素数量大于1时才进入递归
	if end-begin > 1 {
		
		// 将数组一分为二，分为 array[begin,mid) 和 array[mid,high)
		mid := (begin + end) >>1
		
		// 先将左边排序好
		MergeSort(data, begin, mid)
		
		// 再将右边排序好
		MergeSort(data, mid, end)
		
		// 两个有序数组进行合并
		merge(data, begin, mid, end)
	}
	return data
}





// merge @todo  中的排序方式
// @todo ai 是覆盖到哪了
// @todo [] 代表已经被覆盖使用了
// @todo <> 比较的两个值
// @todo {} 被覆盖的值
// @todo | | ai 和 ri的位置

// 情况1：右边先结束【ri先到8了】 (此时ri已经越届了,那么直接把 备份的最后一个拿过来就可以了)

// [3,6,11,18,8,10,12,14]

// 步骤分析： （每一轮是上一轮的结果)
// 1.比较覆盖：（li 和 ri 位置的进行比较） 3 < 8  (所以要把3覆盖到 ai的位置)
// li++ && ai++ (ri不变！！！)
// (li == 0)  <3>,6,11,18(备份)  (ai==0) [{3},6,11,18),[<8>|,10,12,14)  (ri == 4)

// --------------------------------------
// 2.比较覆盖：（li 和 ri 位置的进行比较） 6 < 8  (所以要把6覆盖到 ai的位置)
// li++ && ai++ (ri不变！！！)
// (li == 1)  [3],<6>,11,18(备份)  (ai==1) [{3,6},11,18),[<8>|,10,12,14)  (ri == 4)

// --------------------------------------
// 3.比较覆盖：（li 和 ri 位置的进行比较） 11 > 8  (所以要把8覆盖到 ai的位置)
// ai++ & ri++ (li 不便)
// (li == 2)  [3],[6],<11>,18(备份)  (ai==2) [{3,6,8},18),[<8>|,10,12,14)  (ri == 4)

// --------------------------------------
// 4.比较覆盖   li 和 ri 位置的进行比较） 11 > 10  (所以要把10覆盖到 ai的位置)
// ai++ & ri++ (li 不便)
// (li == 2)  [3],[6],<11>,18(备份)  (ai==3) [{3,6,8,10}),[8,<10>|,12,14)  (ri == 5)

// --------------------------------------
// 5.比较覆盖   li 和 ri 位置的进行比较） 11 < 12  (所以要把11覆盖到 ai的位置)
// li++ && ai++ (ri不变！！！)
// (li == 2)  [3],[6],<11>,18(备份)  (ai==4) [{3,6,8,10),[8}|,10,<12>|,14)  (ri == 6)

// 6.比较覆盖   li 和 ri 位置的进行比较） 18 < 12  (所以要把12覆盖到 ai的位置)
// ai++ & ri++ (li 不便)
// (li == 3)  [3],[6],[11],<18>(备份)  (ai==5) [3,6,8,10),[11},10|,<12>|,14)  (ri == 6)

// 7.比较覆盖   li 和 ri 位置的进行比较） 18 < 14  (所以要把14覆盖到 ai的位置)
// ai++ & ri++ (li 不便)
// (li == 3)  [3],[6],[11],<18>(备份)  (ai==6) [3,6,8,10),[11,12},12|,<14>|)  (ri == 7)

// 8.比较覆盖   (此时ri已经越届了,那么直接把 备份的最后一个拿过来就可以了)
// ai++ & ri++ (li 不便)
// (li == 3)  [3],[6],[11],18(备份)  (ai==7) [3,6,8,10),[11,12,14,14}|)| (ri == 8)

// 最后结果
// (li == 3)  [3],[6],[11],[18](备份)  (ai==8) [3,6,8,10),[11,12,14,18})|| (ri == 8)



// 情况2：左边先结束【ai先到4了】 整个归并就结束了(剩下的肯定是大的)

// [3,6,9,11,8,10,12,14]
// 最后： li ==4 3,6,9,11 (备份)  (ai==6) [{3,6,8,9),[10,11},12,14}) (ri == 6)


// 将[begin,mid) 和 [min,end) 范围的序列合并成一个有序序列
// 重点：merge的两个序列是在同一个数组里面
// 为了完成merge操作，将[begin,mid)备份出来一份，不用备份所有
// 重复将最小的放在最左面，
func merge(data []int, begin int, mid int, end int) []int {
	
	// 申请额外的空间来合并两个有序数组，这两个数组是 array[begin,mid),array[mid,end)
	leftSize := mid - begin         // 左边数组的长度
	rightSize := end - mid          // 右边数组的长度
	newSize := leftSize + rightSize // 辅助数组的长度
	result := make([]int, 0, newSize)
	
	l, r := 0, 0
	for l < leftSize && r < rightSize {
		lValue := data[begin+l] // 左边数组的元素
		rValue := data[mid+r]   // 右边数组的元素
		// 小的元素先放进辅助数组里
		if lValue < rValue {
			result = append(result, lValue)
			l++
		} else {
			result = append(result, rValue)
			r++
		}
	}
	
	// 将剩下的元素追加到辅助数组后面
	
	// 左面先结束或者右面先结束
	result = append(result, data[begin+l:mid]...)
	result = append(result, data[mid+r:end]...)
	
	// 将辅助数组的元素复制回原数组，这样该辅助空间就可以被释放掉
	for i := 0; i < newSize; i++ {
		data[begin+i] = result[i]
	}
	return data
}