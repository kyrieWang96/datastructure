package main

/**
 * @Author KyrisWang
 * @Description //TODO https://leetcode-cn.com/problems/two-sum/
 * @Date 11:22 下午 2021/3/29
 * @Param
 * @return
 */

//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//
//你可以按任意顺序返回答案。
// Todo 暴力破解
func towSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// TODO  hash数组实现
func towSumHash(nums []int, target int) []int {
	hashMap := make(map[int]int, 0)
	for index, data := range nums {
		if i, ok := hashMap[target-data]; ok {
			return []int{i, index}
		}
		hashMap[data] = index
	}
	return nil
}
