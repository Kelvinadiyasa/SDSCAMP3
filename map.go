package main

import "fmt"

func main() {
	// Membuat map string ke int
	umur := make(map[string]int)

	// Menambahkan data ke dalam map
	umur["Gibran"] = 30
	umur["Kelvin"] = 28
	umur["Adiyasa"] = 35

	// Mencetak umur
	fmt.Println("Umur Gibran:", umur["Gibran"])
	fmt.Println("Umur Kelvin:", umur["Kelvin"])
	fmt.Println("Umur Adiyasa:", umur["Adiyasa"])

	// Mengubah data di dalam map
	umur["Gibran"] = 31
	fmt.Println("Umur Gibran setelah diubah:", umur["Gibran"])

	// Menghapus data dari map
	delete(umur, "Adiyasa")
	fmt.Println("Umur Adiyasa setelah dihapus:", umur["Adiyasa"]) // Output akan menjadi 0 karena Bob tidak ada dalam map

	// Memeriksa apakah kunci ada dalam map
	umurAdiyasa, exist := umur["Adiyasa"]
	fmt.Println("Umur Adiyasa setelah dihapus (dengan pemeriksaan eksistensi):", umurAdiyasa, "Exist:", exist) // Output akan menjadi 0 false karena Bob tidak ada dalam map
}
