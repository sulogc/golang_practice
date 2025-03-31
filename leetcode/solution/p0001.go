package solution

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type P0001 struct{} 

func (p *P0001) Solve(inputFile string) (string, error) {
	nums, target, err := readStandardInput(inputFile)
	if err != nil {
		return "", err
	}

	result := twoSum(nums, target)
	return fmt.Sprintf("%v", result), nil
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, exists := numMap[complement]; exists {
			return []int{j, i}
		}
		numMap[num] = i
	}
	return []int{}
}

func readStandardInput(filePath string) ([]int, int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	numsStr := scanner.Text()
	numStrSlice := strings.Fields(numsStr)
	nums := make([]int, len(numStrSlice))
	for i, str := range numStrSlice {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, 0, err
		}
		nums[i] = num
	}

	scanner.Scan()
	targetStr := scanner.Text()
	target, err := strconv.Atoi(targetStr)
	if err != nil {
		return nil, 0, err
	}

	return nums, target, nil
}