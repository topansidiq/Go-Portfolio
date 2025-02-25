package main

import (
	"fmt"
	"os"
)

var path = "C:/Users/topan/OneDrive/Dokumen/temp/edit.html"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func createFile() {
	// Cek apakah file sudah ada?
	var _, err = os.Stat(path)

	// Buat file jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("==> file berhasil dibuat", path)
}

func writeFile(data string) {
	// buka file dengan level akses READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Tulis data ke file
	_, err = file.WriteString(data)
	if isError(err) {
		return
	}

	// Simpan perubahan
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("==> File berhasil di edit")
}

func readFile() {
	// Buka file
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Baca file
	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != nil {
			if isError(err) {
				break
			}
		}
		if n == 0 {
			break
		}
	}
	if isError(err) {
		return
	}

	fmt.Println("==> File berhasil dibaca")
	fmt.Println(string(text))
}

func deleteFile() {
	var err = os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("==> File berhasil di delete")
}

func main() {
	createFile()
	writeFile("Hello\nI'm human from man!")
	readFile()
	deleteFile()
}
