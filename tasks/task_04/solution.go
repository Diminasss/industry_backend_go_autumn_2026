package main

import "math"

type Stats struct {
	Count         int
	Sum, Min, Max int64
}

func Calc(nums []int64) Stats {
	var stats Stats

	if len(nums) < 2 {
		return stats
	}

	stats.Min = int64(math.MaxInt64)
	stats.Max = int64(math.MinInt64)

	for i := 1; i < len(nums); i++ {
		difference := nums[i] - nums[i-1]

		stats.Sum += difference
		stats.Count++

		if difference < stats.Min {
			stats.Min = difference
		}

		if difference > stats.Max {
			stats.Max = difference
		}
	}

	return stats
}
