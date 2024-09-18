package main

import (
	"fmt"
	"os"
	"regexp"
	"unsafe"
)

var gopherRegexp = regexp.MustCompile("gopher")

func RangeKeywordInvolveCopying() {
	// Don't do this
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// arr is copied
	for key, value := range arr {
		fmt.Print(key, value)
	}
	fmt.Println()

	// Do this instead
	for i := 0; i < len(arr); i++ {
		fmt.Print(i, arr[i])
	}
	fmt.Println()

	// arr is copied
	for key, value := range slice {
		fmt.Println(key, value)
	}
	fmt.Println()
	// Do this instead
	for i := 0; i < len(slice); i++ {
		fmt.Println(i, slice[i])
	}
}

func FindGopherWithEatMemory(filename string) []byte {
	//Reading a very huge file  1,000,000,000 bytes (1GB)
	b, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error while reading file:", err)
		return nil
	}

	//Taking a just 6 byte sub slice
	gopherSlice := gopherRegexp.Find(b)
	return gopherSlice
}

func FindGopherWithNoEatUp(filename string) []byte {
	//Reading a very huge file  1,000,000,000 bytes (1GB)
	b, _ := os.ReadFile(filename)

	//Taking a just 6 byte sub slice
	gopherSlice := make([]byte, len("gopher"))

	// Make a deep copy
	copy(gopherSlice, gopherRegexp.Find(b))

	return gopherSlice
}

func GarbageCollectionNuances() {
	// From the above example we read a very huge file (1GB) and returned a sub slice of it (just 6 bytes), since the gopherSlice still reference the same underlying array as the huge file, which means that 1GB of memory can not be garbage collected even though we are not using it anymore.
	// If you call the FindGopher function multiples times, you program can eat all the computers memory. To fix this, like I said earlier we make a deep copy so gopherSlice doesn't share the same underlying array as the huge slice
	filename := "C:\\pp\\Passport\\NameChange\\test.txt"

	FindGopherWithEatMemory(filename)

	//To fix this, like I said earlier we make a deep copy so gopherSlice doesn't share the same underlying array as the huge slice
	FindGopherWithNoEatUp(filename)
}

func DeepCopyUsingCopy() {
	s := []int{1, 2, 3}
	t := make([]int, len(s))

	fmt.Println("Original Slice", s)
	copy(t, s)
	fmt.Println("copied slice", t) // => [1, 2, 3], and is a deep copy of s

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x := make([]int, len(arr))
	copy(x, arr)
	fmt.Println(x) // => [1, 2, 3], and is a deep copy of s

	// copy a range of items
	t = make([]int, len(s)-1)
	copy(t, s[0:2])
	fmt.Println(t) // => [1, 2], and is a deep copy of s
}

func DeepCopySliceUsingAppend() {
	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice2 := []int{}
	slice2 = append(slice2, slice1...)

	fmt.Println(slice1) // => [1 2 3 4 5 6]
	fmt.Println(slice2) // => [1 2 3 4 5 6]

	//Modifying slice2 doesn't affect slice1
	slice2[0] = 100
	fmt.Println(slice1) // => [1 2 3 4 5 6]
	fmt.Println(slice2) // => [100 2 3 4 5 6]

	// copying a range of items
	slice3 := []int{}
	slice3 = append(slice3, slice1[3:5]...)
	fmt.Println(slice3) // => [4 5]

	//Again slice3 and slice1 doesn't share underlying array

	slice3[0] = -10
	fmt.Println(slice1) // => [1 2 3 4 5 6]
	fmt.Println(slice3) // => [-10 5]
}

func AppendBeyondCapacity() {
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n1 := n[:6]
	n2 := n[3:8]
	n3 := n[4:10]

	fmt.Println("slice n1:", n1, len(n1), cap(n1)) // => [1 2 3 4 5 6]  6 10
	fmt.Println("slice n2:", n2, len(n2), cap(n2)) // => [4 5 6 7 8]  5 7
	fmt.Println("slice n3:", n3, len(n3), cap(n3)) // => [5 6 7 8 9 10] 6 6

	n2 = append(n2, 100)
	fmt.Println(n)  // => [1 2 3 4 5 6 7 8 100 10]
	fmt.Println(n1) // => [1 2 3 4 5 6]
	fmt.Println(n2) // => [4 5 6 7 8, 100]
	fmt.Println(n3) // => [5 6 7 8 100 10]

	n2 = append(n2, 101)
	fmt.Println(n)  // => [1 2 3 4 5 6 7 8 100 101]
	fmt.Println(n1) // => [1 2 3 4 5 6]
	fmt.Println(n2) // => [4 5 6 7 8 100 101]
	fmt.Println(n3) // => [5 6 7 8 100 101]

	// Check the capacity and length of n2
	fmt.Println(cap(n2), len(n2)) // =>  7 7
	//from the above code n2 has a length of 5 and a capacity of 7 which means we can append two more items
	// without a new array being created and it'll continue so share the same underlying array with the other sub slices

	//As we appended more items it was affecting n and n3 but now the n2 slice is full (capacity == length).
	//Now that capacity of n2 is equal to it's length so appending a new item will cause new array to created for n2 and
	//it will no longer share the same underlying array with the other sub slices

	fmt.Println("New array is created if we append beyond capacity")
	n2 = append(n2, 102)
	fmt.Println(n)  // => [1 2 3 4 5 6 7 8 100 101]
	fmt.Println(n1) // => [1 2 3 4 5 6]
	fmt.Println(n2) // => [4 5 6 7 8 100 101 102]
	fmt.Println(n3) // => [5 6 7 8 100 101]

	fmt.Println("appending multiple items capacity") //append multiple items to slice
	s := []int{10, 20, 30, 40, 50, 60}
	s2 := []int{70, 80, 90}

	// Appending slice to slice
	s = append(s, s2...)
	fmt.Println(s) // => [10 20 30 40 50 60 70 80 90]

	// Appending multiple values
	s = append(s, 100, 110, 120)
	fmt.Println(s) // => [10 20 30 40 50 60 70 80 90 100 110 120]
}

func main() {
	RangeKeywordInvolveCopying()
	DeepCopyUsingCopy()
	AppendBeyondCapacity()

	var nums [5]int
	fmt.Println(nums) // => [0 0 0 0 0]

	//Array of 10 strings
	var strs [10]string
	fmt.Println(strs) // => [         ]

	// Nested arrays
	var nested = [3][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 13, 15},
	}

	fmt.Println(nested) // => [[1 2 3 4 5] [6 7 8 9 10] [11 12 13 13 15]]

	// var name = [L]T{...} where ... represents the array items of type T
	//Intializing an array containing 10 intergers
	var nums2 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(nums2) // => [1 2 3 4 5 6 7 8 9 10]

	//Intializing an array containing 10 strings
	var strs2 = [10]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	fmt.Println(strs2) // => [one two three four five six seven eight nine ten]

	type Car struct {
		Brand string
		Color string
		Price float32
	}

	//Array of 5 items of type Car
	var arrayOfCars = [5]Car{
		{Brand: "Porsche", Color: "Black", Price: 20_000.00},
		{Brand: "Volvo", Color: "White", Price: 8_000.00},
		{Brand: "Honda", Color: "Blue", Price: 7_000.00},
		{Brand: "Tesla", Color: "Black", Price: 50_000.00},
		{Brand: "Kia", Color: "Red", Price: 5_000.98},
	}
	fmt.Println(arrayOfCars) // => [{Porsche Black 20000} {Volvo White 8000} {Honda Blue 7000} {Tesla Black 50000} {Kia Red 5000.98}]

	//Array containing 5 items of different type
	var randomsArray = [5]interface{}{"Hello world!", 35, false, 33.33, 'A'}
	fmt.Println(randomsArray) // => [Hello world! 35 false 33.33 65]

	bytes := []byte{104, 101, 108, 108, 111}

	p := unsafe.Pointer(&bytes)
	str := *(*string)(p) //cast it to a string pointer and assign the value of this pointer
	fmt.Println(str)     //prints "hello"

	x := [5]int{1, 2, 3, 4, 5} //array copy
	fmt.Println(x)
	y := x
	fmt.Println(y)

	// create a slice pointing to array
	//The slice expression signature is s[start:end:cap]
	//If the start index is zero, you can omit it s[:end] similarly if the end index is the end of the array you can omit it like so s[start:]
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n1 := n[:6]
	n2 := n[3:8]
	n3 := n[4:10]

	fmt.Println(n1, len(n1), cap(n1)) // => [1 2 3 4 5 6]  6 10
	fmt.Println(n2, len(n2), cap(n2)) // => [4 5 6 7 8]  5 7
	fmt.Println(n3, len(n3), cap(n3)) // => [5 6 7 8 9 10] 6 6

	// change n1 at index 4 to 15
	n1[4] = 15

	fmt.Println(n)  // => [1 2 3 4 15 6 7 8 9 10]
	fmt.Println(n1) // => [1 2 3 4 15 6]
	fmt.Println(n2) // => [4 15 6 7 8]
	fmt.Println(n3) // => [15 6 7 8 9 10]

	//func append(s []T, x ...T) []T
	//The Go append function allows you to add elements to the end of a slice.
	s := []int{1, 2, 3}
	s = append(s, 4, 5, 6)
	fmt.Println(s) // => s is now [1, 2, 3, 4, 5, 6]
}
