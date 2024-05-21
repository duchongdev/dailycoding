package test

import "testing"

func TestSlice(t *testing.T) {

	array := []int{1, 2, 3, 4, 5}

	// 2 是切片开始的索引，表示 s 将从 array 的第三个元素（索引为2的元素，因为索引是从0开始的）开始。
	// 4 是切片结束的索引（不包含），表示 s 将包含从 array[2] 开始到 array[3] 结束（不包括 array[4]）的元素。因此，s 将包含 array 中的第四个元素。
	// 5 是切片的容量（capacity），表示 s 的底层数组从 array[2] 开始，直到 array 的末尾（索引为4的元素）。这意味着 s 可以增长到包含从 array[2] 到 array[4] 的元素，而不需要分配新的底层数组。
	s := array[2:4:5]

	t.Logf("s[0]=%v", s[0])
	t.Logf("cap(s)=%v", cap(s))
	t.Logf("len(s)=%v", len(s))

}

func TestArrayEqual(t *testing.T) {
	arr1 := [2]int{1, 2}
	arr2 := [2]int{1, 3}
	//false：
	t.Logf("arr1 == arr2 = %v", arr1 == arr2)

	// arr3 := [3]int{1, 2}
	// t.Logf("arr1 == arr3 = %v", arr1 == arr3)
	// invalid operation: arr1 == arr3 (mismatched types [2]int and [3]int)

}
