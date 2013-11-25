package main

import "fmt"

func main() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5] // item at index 5 not included!

	// load the array with integers: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {

		arr1[i] = i
	}
	// print the slice: 2,3,4
	for i := 0; i < len(slice1); i++ {

		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice:
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {

		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
	// grow the slice beyond capacity:
	// slice1 = slice1[0:7 ] // panic: runtime error: slice bounds out of range

	// array'da slice'ta değişti
	slice1[0] = 4
	for i := 0; i < len(slice1); i++ {

		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("Array at %d is %d\n", i, arr1[i])
	}

	// type, len, cap
	var ss []int = make([]int, 5, 10)
	// short (len = cap)
	// s := make([]int, 5)

	// make([]int, 50, 100)
	// new([100]int)[0:50]

	slic := []int{1, 2, 3}
	fmt.Printf("%d-%d\n", len(slic), cap(slic)) // 3-3
	slic = append(slic, 1, 2)
	fmt.Printf("%d-%d\n", len(slic), cap(slic)) // 5-6
	copy(ss[2:5], slic[2:5])
	fmt.Printf("%d\n", ss)

	slic = append(slic, slic...)

	// delete 3. item
	fmt.Printf("%d\n", slic)
	slic = append(slic[:3], slic[4:]...)
	fmt.Printf("%d\n", slic)

	// delete 4. and 5. items
	slic = append(slic[:3], slic[5:]...)
	fmt.Printf("%d\n", slic)

	// belirtilen index'e ekle
	slic = append(slic[:3], append([]int{11, 22, 33}, slic[3:]...)...)
	fmt.Printf("%d\n", slic)

	// pop
	slic = slic[:len(slic)-1]
	fmt.Printf("%d\n", slic)

}
