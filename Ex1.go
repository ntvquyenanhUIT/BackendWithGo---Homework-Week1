package main

import "fmt"

func main() {

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
