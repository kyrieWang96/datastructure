package main

import (
	"fmt"
)

/**
 * @Author KyrieWang
 * @Description //TODO bubblesort 循环一次，把最大或者最小的数交换到最后面
 * @Date 1:26 下午 2021/3/29
 * @Param
 * @return
 */

// TODO 冒泡返回最大值
func BubbleMaxSort(input []int) int {
	for i := 1; i < len(input); i++ {
		if input[i] < input[i-1] {
			input[i-1], input[i] = input[i], input[i-1] // 交换位置
		}
	}
	return input[len(input)-1]
}

// TODO 冒泡排序算法： 每次循环把第一个数字和后面的数进行对比，把大的数据往前排 排在第一个
func BubbleSort(input []int) []int {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ { // 精髓在把最大的冒泡上来之后，再选第二个冒泡大的
			if input[i] < input[j] {
				input[i], input[j] = input[j], input[i]
			}
		}
	}
	return input
}
func main() {
	inputData := []int{61, 12, 123, 41, 12, 23, 51, 67, 65}
	// 冒泡排序求出最大值
	fmt.Println(BubbleMaxSort(inputData))

	// 冒泡法排序
	fmt.Println(BubbleSort(inputData))
}
