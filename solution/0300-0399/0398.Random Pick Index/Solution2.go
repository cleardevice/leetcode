type Solution struct {
	pos map[int][]int
}

func Constructor(nums []int) Solution {
	pos := make(map[int][]int)
	for i, v := range nums {
		pos[v] = append(pos[v], i)
	}
	return Solution{pos: pos}
}

func (this *Solution) Pick(target int) int {
	arr := this.pos[target]
	return arr[rand.Intn(len(arr))]
}
