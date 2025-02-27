package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.Now()
	fmt.Printf("Time1 %v\n", time1)

	time1 = time.Date(2002, 7, 28, 23, 59, 59, 0, time.UTC)
	fmt.Printf("Time2 %v\n", time1)

	today := time.Now()
	fmt.Println("Tahun", today.Year(), "Bulan", today.Month(), "Tanggal", today.Day())
}
