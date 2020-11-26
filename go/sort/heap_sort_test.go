package sort

import (
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := make([]int, len(testCase))
	copy(arr, testCase)
	HeapSort(arr)
	t.Log(arr)
}

// 构建最大堆,每次都把堆顶元素与最后一位交换,之后重新下沉
func HeapSort(arr []int) {
	buildHeap(arr)
	k := len(arr) - 1
	for k > 0 {
		arr[0], arr[k] = arr[k], arr[0]
		k--
		heapify(arr, k, 0)
	}
}

// []int{55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9}
// 从最后一个节点的父节点开始
func buildHeap(arr []int) {
	for i := (len(arr) - 2) / 2; i >= 0; i-- {
		heapify(arr, len(arr), i)
	}
}

// index从1开始的话 父节点 p/2 ,左子节点 p*2 ,右子节点 p*2+1
// index从0开始的话 父节点 (p-1)/2,左子节点p*2+1,右子节点 P*2+2
func heapify(arr []int, length int, p int) {
	for {
		maxPos := p
		if p*2+1 < length { // 左子节点存在
			if arr[p*2+1] > arr[p] {
				maxPos = p*2 + 1
			}
		}
		if p*2+2 < length { // 如果右子节点存在
			if arr[p*2+2] > arr[maxPos] {
				maxPos = p*2 + 2
			}
		}
		if maxPos == p {
			break
		}
		arr[maxPos], arr[p] = arr[p], arr[maxPos]
		p = maxPos
	}
}
