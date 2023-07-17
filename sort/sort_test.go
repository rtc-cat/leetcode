package sort

import (
	"fmt"
	"os"
	"testing"
)

var (
	testCase = []int{55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9}
)

func TestMain(m *testing.M) {
	fmt.Printf("input: %v\n", testCase)
	os.Exit(m.Run())
}
