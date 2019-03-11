package sort

// 快速排序
// 时间复杂度：O(n*logn)

func QuickSort(data []int) []int {
	
	if len(data) <= 1 {
		
		return data
		
	}
	
	left := make([]int, 0, len(data))
	right := make([]int, 0, len(data))
	
	for i := 1; i < len(data); i++ {
		
		if data[i] < data[0] {
			
			left = append(left, data[i])
			
		}
		
		if data[i] > data[0] {
			
			right = append(right, data[i])
			
		}
	}
	
	left = QuickSort(left)
	right = QuickSort(right)
	
	left = append(left, data[0])
	
	return append(left, right...)
}