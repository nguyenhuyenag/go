package main

import (
	"fmt"
	"slices"
	"sort"
)

// sort.Float64s(f)
func sortInt() {
	a := []int{3, 2, 1}
	fmt.Println("Before:", a)
	sort.Ints(a) 								// Mặc định sort theo thứ tự tăng dần
	fmt.Println("Sorted 1:", a)
	sort.Sort(sort.Reverse(sort.IntSlice(a))) 	// Sắp xếp theo thứ tự giảm dần
	fmt.Println("Sorted 2:", a)
}

func sortString() {
	s := []string{"banana", "apple", "cat"}
	fmt.Println("Before:", s)
	sort.Strings(s)
	fmt.Println("Sorted 1:", s)
	sort.Reverse(sort.StringSlice(s))
}

// Sort với comparator
func sortWithComparator() {
	s := []string{"banana", "apple", "cat"}
	fmt.Println("Before:", s)

	// Cách 1
	//sort.Slice(s, func(i, j int) bool {
	//	return len(s[i]) < len(s[j])
	//})

	// Cách 2
	slices.SortFunc(s, func(s1, s2 string) int {
		return len(s1) - len(s2)
	})

	fmt.Println("Sorted by length:", s)
}

func swap() {
	arr := []int{1, 2, 3}
	fmt.Println("Before:", arr)
	arr[0], arr[2] = arr[2], arr[0]
	fmt.Println("Swap:", arr)
}

func main() {
	swap()
	// sortInt()
	// sortString()
	// sortWithComparator()
}

