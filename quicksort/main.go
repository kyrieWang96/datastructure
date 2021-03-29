package main

import "fmt"

/**
 * @Author KyrieWang
 * @Description //TODO quicksort 时间复杂度为 O(nlgn)
 * @Date 9:48 上午 2021/3/29
 * @Param
 * @return
 */
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	startArr := arr[0]
	lowData := make([]int, 0)
	highData := make([]int, 0)
	sameData := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if arr[i] < startArr {
			lowData = append(lowData, arr[i])
		} else if arr[i] > startArr {
			highData = append(highData, arr[i])
		} else {
			sameData = append(sameData, arr[i])
		}
	}
	lowData, highData = QuickSort(lowData), QuickSort(highData) //递归进行排序

	return append(append(lowData, sameData...), highData...) // append ...追加多个参数
}

func main() {
	data := []int{1, 5, 2, 4, 12, 21, 8, 12, 31, 24, 12, 14, 23}
	listData := QuickSort(data)
	fmt.Println(listData)
}
