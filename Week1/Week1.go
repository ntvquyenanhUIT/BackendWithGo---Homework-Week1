package main

import (
	"fmt"
	"math"
)

func Exercise1() {
	var ChieuDai, ChieuRong int

	fmt.Print("Nhập chiều dài HCN: ")
	fmt.Scan(&ChieuDai)
	fmt.Print("Nhập chiều rộng HCN: ")
	fmt.Scan(&ChieuRong)
	var ChuVi int = (ChieuDai + ChieuRong) * 2
	var DienTich int = ChieuDai * ChieuRong

	fmt.Println("Chu vi HCN là: ", ChuVi)
	fmt.Println("Diện tích HCN là: ", DienTich)
}

func Exercise2() {
	fmt.Print("Nhập 1 chuỗi: ")
	var s string
	fmt.Scan(&s)
	if len(s)%2 == 0 {
		fmt.Println("True")
	}
	fmt.Println("False")
}

func Exercise3() {
	var numOfNumber int
	fmt.Print("Nhập vào số lượng số nguyên mảng có: ")
	fmt.Scan(&numOfNumber)
	mySlice := make([]int, numOfNumber)
	sliceSum, sliceAverage := 0, 0
	sliceMax, sliceMin := math.MinInt64, math.MaxInt64

	for i := 0; i < numOfNumber; i++ {
		fmt.Scan(&mySlice[i])
		if mySlice[i] < sliceMin {
			sliceMin = mySlice[i]
		}
		if mySlice[i] > sliceMax {
			sliceMax = mySlice[i]
		}
		sliceSum += mySlice[i]
		sliceAverage += mySlice[i]

	}
	sliceAverage /= numOfNumber
	mySlice = SelectionSort(mySlice)
	fmt.Printf("Sum: %v - Max: %v - Min: %v - Avg: %v - SortedSlice: %v", sliceSum, sliceMax, sliceMin, sliceAverage, mySlice)
}

func SelectionSort(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		currentMin := slice[i]
		newMinIndex := i
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < currentMin {
				currentMin = slice[j]
				newMinIndex = j
			}
		}
		if newMinIndex != i {
			temp := slice[i]
			slice[i] = slice[newMinIndex]
			slice[newMinIndex] = temp
		}
	}
	return slice
}

func Exercise4(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				indices := []int{i, j}
				return indices
			}
		}
	}
	return nil
}

func main() {

	Exercise1()
	Exercise2()
	Exercise3()
}
