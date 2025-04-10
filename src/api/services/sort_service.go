package services

import "testing/src/api/utils/sort"

const (
	privateConst = "private"
)

// func getElements(n int) []int {
// 	result := make([]int, n)
// 	j := 0
// 	for i := n - 1; i > 0; i-- {
// 		result[i] = i
// 		j++
// 	}
// 	return result
// }

func Sort(elements []int) {
	if len(elements) <= 10000 {
		sort.BubbleShort(elements)
		return
	}
	sort.Sort(elements)
}

// func BubbleShort(elements []int) {
// 	keepWorking := true
// 	for keepWorking {
// 		keepWorking = false
// 		for i := 0; i < len(elements)-1; i++ {
// 			if elements[i] > elements[i+1] {
// 				keepWorking = true
// 				elements[i], elements[i+1] = elements[i+1], elements[i]
// 			}
// 		}
// 	}
// }
