package services

import (
	"testing"
	"testing/src/api/utils/sort"
)

func TestCons(t *testing.T) {
	if privateConst != "private" {
		t.Error("privateConst should be 'private'")
	}
}

func TestSort(t *testing.T) {
	// elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	elements := sort.GetElements(10)
	// fmt.Println(elements)
	Sort(elements)
	// fmt.Println(elements)
	if elements[0] != 0 {
		t.Error("first element should be 0")
	}
	if elements[len(elements)-1] != 9 {
		t.Error("last element should be 9")
	}
}

func TestSortMoreThan1000(t *testing.T) {
	elements := sort.GetElements(10001)
	Sort(elements)
	if elements[0] != 0 {
		t.Error("first element should be 0")
	}
	if elements[len(elements)-1] != 10000 {
		t.Error("last element should be 9")
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	// elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	elements := sort.GetElements(10000)
	for i := 0; i < b.N; i++ {
		sort.BubbleShort(elements)
	}

}

func BenchmarkSort(b *testing.B) {
	// elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	elements := sort.GetElements(100000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}

}
