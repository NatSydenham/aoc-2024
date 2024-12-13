package day11

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/NatSydenham/aoc-2024/lib/file"
	"github.com/NatSydenham/aoc-2024/lib/slice"
)

func getStones(nums []int, times int) map[int]int {

	numsToAdd := make(map[int]int)
	counter := make(map[int]int)

	// Populate count of initial value
	for _, num := range nums {
		counter[num] += 1
	}

	for i := 0; i < times; i++ {

		//add X number of results per iteration from previous iteration
		for k, v := range numsToAdd {
			counter[k] += v
		}

		//reset the new additions map
		numsToAdd = make(map[int]int, 0)

		for num, count := range counter {
			asString := strconv.Itoa(num)

			// add total count of 0s to total count of 1s, remove 0 as key from map.
			if num == 0 {
				numsToAdd[1] += count
				delete(counter, 0)
			} else if len(asString)%2 == 0 {

				firstHalf, _ := strconv.Atoi(asString[:len(asString)/2])
				secondHalf, _ := strconv.Atoi(asString[len(asString)/2:])

				numsToAdd[firstHalf] += count
				numsToAdd[secondHalf] += count

				delete(counter, num)
			} else {
				numsToAdd[num*2024] += count
				delete(counter, num)
			}
		}
	}

	for k, v := range numsToAdd {
		counter[k] += v
	}

	return counter

}

func Execute(times int) {
	start := time.Now()
	input := file.Readlines("./data/day11.txt")
	nums := slice.Map(strings.Split(input[0], " "), func(val string) int {
		num, _ := strconv.Atoi(val)
		return num
	})

	count := getStones(nums, times)

	sum := 0

	for _, v := range count {
		sum += v
	}

	fmt.Println(sum, "|", time.Since(start))
}
