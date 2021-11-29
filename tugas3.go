package main

import "fmt"

func main() {
	buah := []string{"apel" , "jeruk" , "anggur" , "melon"}
	buah = append(buah, "pepaya")

	fmt.Println("jumlah elemen : ", len(buah))
	fmt.Println("isi elemen : ", buah)
	for i := 0; i < len(buah); i++ {
		fmt.Println("elemen ke - : ",i , buah)
	}
}
