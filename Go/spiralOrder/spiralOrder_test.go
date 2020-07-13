package spiralOrder

import "testing"

// 测试空切片
func TestEmptySpiralOrder(t *testing.T) {
	testMaxtrix := [][]int{}
	want := []int{}

	var got = SpiralOrder(testMaxtrix)
	arrCompare(t, got, want)
}

// 测试二维1*1切片
func TestOneElement(t *testing.T) {
	testMaxtrix := [][]int{
		{1},
	}
	want := []int{1}

	var got = SpiralOrder(testMaxtrix)
	arrCompare(t, got, want)
}

// 测试二维1*3切片
func TestOneRowArr(t *testing.T) {
	testMaxtrix := [][]int{
		{1, 2, 3},
	}
	want := []int{
		1, 2, 3,
	}

	var got = SpiralOrder(testMaxtrix)
	arrCompare(t, got, want)
}

// 测试二维3*1切片
func TestOneColArr(t *testing.T) {
	testMaxtrix := [][]int{
		{1},
		{2},
		{3},
	}
	want := []int{
		1, 2, 3,
	}

	var got = SpiralOrder(testMaxtrix)
	arrCompare(t, got, want)
}

// 测试二维3*3切片
func TestSameShapeArr(t *testing.T) {
	testMaxtrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	want := []int{
		1, 2, 3, 6, 9, 8, 7, 4, 5,
	}

	var got = SpiralOrder(testMaxtrix)
	arrCompare(t, got, want)
}


// 测试二维3*4切片
func TestARowsLessThanColsShapeArr(t *testing.T) {
	testMaxtrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	want := []int{
		1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7,
	}

	var got = SpiralOrder(testMaxtrix)
	arrCompare(t, got, want)
}

// 测试二维4*3切片
func TestAColsLessThanRowsShapeArr(t *testing.T) {
	testMaxtrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}
	want := []int{
		1, 2, 3, 6, 9, 12, 11, 10, 7, 4, 5, 8,
	}

	var got = SpiralOrder(testMaxtrix)
	arrCompare(t, got, want)
}


func arrCompare(t *testing.T, got []int, want []int) {
	if len(got) != len(want) {
		t.Errorf("got %v wat %v", got, want)
	}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got %v wat %v", got, want)
		}
	}
}