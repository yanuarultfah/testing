package sort

import (
	"testing"
	"time"

	// "github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/assert"
)

func SortWithTimeOut() {

}

func TestBubbleSortIncreasingOrder(t *testing.T) {
	// elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	elements := GetElements(10)
	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 9, elements[0])
	assert.EqualValues(t, 0, elements[len(elements)-1])

	timeoutchan := make(chan bool, 1)
	defer close(timeoutchan)
	go func() {
		BubbleShort(elements)
		timeoutchan <- false
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		timeoutchan <- true
	}()

	if <-timeoutchan {
		assert.Fail(t, "buble sort took more than 500 ms")
		return
	}

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 0, elements[0])
	assert.EqualValues(t, 9, elements[len(elements)-1])

	// if elements[0] != 0 {
	// 	t.Error("first element should be 0")
	// }
	// if elements[len(elements)-1] != 9 {
	// 	t.Error("last element should be 9")
	// }
}

func TestSortIncreasingOrder(t *testing.T) {
	// elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	elements := GetElements(10)
	Sort(elements)

	assert.EqualValues(t, 0, elements[0], "first element should be 0")
	assert.EqualValues(t, 9, elements[len(elements)-1], "last element should be 9")
	// if elements[0] != 0 {
	// 	t.Error("first element should be 0")
	// }
	// if elements[len(elements)-1] != 9 {
	// 	t.Error("last element should be 9")
	// }
}

func BenchmarkBubbleSort(b *testing.B) {
	// elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	elements := GetElements(10000)
	for i := 0; i < b.N; i++ {
		BubbleShort(elements)
	}

}

func BenchmarkSort(b *testing.B) {
	// elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	elements := GetElements(100000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}

}
