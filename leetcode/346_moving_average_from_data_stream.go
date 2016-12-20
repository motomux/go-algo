package leetcode

// 346. Moving Average from Data Stream
type MovingAverage struct {
	nums          []int
	length, point int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		nums:   make([]int, size),
		length: 0,
		point:  0,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	this.nums[this.point] = val
	this.point++
	if this.point == len(this.nums) {
		this.point = 0
	}
	if this.length < len(this.nums) {
		this.length++
	}
	var sum int
	for _, num := range this.nums {
		sum += num
	}
	return float64(sum) / float64(this.length)
}
