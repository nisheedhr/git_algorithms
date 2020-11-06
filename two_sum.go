package algorithms

func twoSum(nums []int, target int) []int {
	 lookupMap := make(map[int] int)
	 for pos, val := range nums {
	 	if compPos, ok := lookupMap[val] ; ok {
	 		return [] int {pos, compPos}
		}
		lookupMap[target - val] = pos
	}
	return [] int {}
}
