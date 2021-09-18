package utils

import (
	"fmt"
	"testing"
)

func TestUniqueArray(t *testing.T) {
	fmt.Println(UniqueArray([]int64{11, 22, 11}))
}

func TestIsInArray(t *testing.T) {
	t.Log(IsInArray(1, []int{2}))
}
