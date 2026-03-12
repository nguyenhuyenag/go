package main

import "fmt"

/*
	- append()

		func append(slice []Type, elems ...Type) []Type
			// pass

		Khi gọi append:
		  - Nếu capacity còn đủ → Go chỉ thêm vào array hiện tại.
		  - Nếu capacity không đủ → Go tạo array mới lớn hơn, copy dữ liệu sang.
		  - Hàm sẽ trả về slide mới, không thay đổi slide ban đầu.

	- cap()

		func cap(v Type) int
			// pass

		Trả về capacity (dung lượng) của v, tùy theo kiểu dữ liệu của nó:
			- Array (mảng): Số lượng phần tử trong v (giống như len(v)).
			- Con trỏ tới mảng (pointer to array): Số lượng phần tử trong *v (cũng giống len(v)).
			- Slice: Độ dài tối đa mà slice có thể đạt được khi reslice; nếu v là nil thì cap(v) bằng 0.
			- Channel: Dung lượng buffer của channel, tính theo số phần tử; nếu v là nil thì cap(v) bằng 0.
*/
func main() {
	arr1 := []int{1, 2, 3}

	fmt.Printf("len: %d, cap: %d\n", len(arr1), cap(arr1))

	// Thêm phần tử vào slide
	arr2 := append(arr1, 4, 5, 6)
	fmt.Println("arr1:", arr1)
	fmt.Println("arr2:", arr2)

	// Nối 2 slide
	arr3 := append(arr1, arr2...)
	fmt.Println("arr3:", arr3)
}
