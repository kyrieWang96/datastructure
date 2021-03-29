package main

import "fmt"

/**e
 * @Author KyrieWang
 * @Description //TODO 二分法查找 找到=值 返回下标，否则返回 -1
 * @Date 4:56 下午 2021/3/29
 * @Param
 * @return
 */
func BinarySort(input []int, finddata int) int {
	lowIndex := 0
	hightIndex := len(input)

	for lowIndex <= hightIndex {
		mid := (lowIndex + hightIndex) / 2
		if finddata < input[mid] {
			hightIndex = mid - 1
		} else if finddata > input[mid] {
			lowIndex = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

func main() {
	testData := []int{1, 3, 5, 6, 8, 12, 13, 15, 24}
	fmt.Println(BinarySort(testData, 6))

}
