package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type akun struct {
	username string
	password string
}

type destinasi struct {
	nama   string
	lokasi string
	jarak  float64
	biaya  int
	rating float64
}

var users [10]akun
var numUser int

var spots = [20]destinasi{
	{nama: "Taman Mini Indonesia Indah", lokasi: "Jakarta Timur", jarak: 11.45, biaya: 10000, rating: 4.2},
	{nama: "Museum Nasional Indonesia", lokasi: "Jakarta Pusat", jarak: 3.54, biaya: 5000, rating: 4.3},
	{nama: "Ancol Dreamland", lokasi: "Jakarta Utara", jarak: 8.96, biaya: 25000, rating: 4.0},
	{nama: "Kota Tua Jakarta", lokasi: "Jakarta Barat", jarak: 7.80, biaya: 0, rating: 4.4},
	{nama: "Museum Bank Indonesia", lokasi: "Jakarta Barat", jarak: 7.98, biaya: 5000, rating: 4.5},
	{nama: "Jakarta Aquarium", lokasi: "Jakarta Barat", jarak: 11.42, biaya: 200000, rating: 4.6},
	{nama: "Dufan", lokasi: "Jakarta Utara", jarak: 9.03, biaya: 30000, rating: 4.1},
	{nama: "Ragunan Zoo", lokasi: "Jakarta Selatan", jarak: 10.73, biaya: 40000, rating: 3.9},
	{nama: "Setu Babakan", lokasi: "Jakarta Selatan", jarak: 17.42, biaya: 0, rating: 4.0},
	{nama: "Monas", lokasi: "Jakarta Pusat", jarak: 3.47, biaya: 20000, rating: 4.5},
}

var scanner = bufio.NewScanner(os.Stdin)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
	fmt.Print("\033[3J")
}

func main() {
	clearScreen()
	for {
		fmt.Println("       Selamat datang di aplikasi pariwisata      ")
		fmt.Println("====================================================")
		fmt.Println("Silakan buat akun baru atau login")
		fmt.Println("1. Buat akun baru")
		fmt.Println("2. Login")
		fmt.Println("0. Exit")
		fmt.Println("====================================================")
		fmt.Print("Masukkan pilihan (angka) : ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			register()
		case "2":
			login()
		case "0":
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			os.Exit(0)
		default:
			fmt.Println("")
			fmt.Println("Pilihan tidak valid.")
			fmt.Println("")
		}
	}
}

func register() {

	clearScreen()

	fmt.Println("               Silakan buat akun baru    ")
	fmt.Println("====================================================")
	fmt.Print("Masukkan username baru : ")
	scanner.Scan()
	username := scanner.Text()

	fmt.Print("Masukkan password baru : ")
	scanner.Scan()
	password := scanner.Text()

	for i := 0; i < numUser; i++ {
		if users[i].username == username {
			fmt.Println("Username sudah ada, silakan gunakan username lain.")
			return
		}
	}

	if numUser >= len(users) {
		fmt.Println("Tidak dapat mendaftarkan lebih banyak pengguna, kapasitas penuh.")
		return
	}

	users[numUser] = akun{username: username, password: password}
	numUser++

	login()
	fmt.Println("")
	clearScreen()
}

func login() {
	clearScreen()
	fmt.Println("                   Silakan login                 ")
	fmt.Println("====================================================")
	fmt.Print("Masukkan username : ")
	scanner.Scan()
	username := scanner.Text()

	fmt.Print("Masukkan password : ")
	scanner.Scan()
	password := scanner.Text()

	for _, user := range users {
		if user.username == username && user.password == password {
			fmt.Println("Login berhasil!")
			userMenu()
			return
		}
	}
	if username == "admin_kece" && password == "admin135" {
		adminMenu()
		return
	}

	fmt.Println("Login gagal!")
	fmt.Println("Username atau password salah")
	fmt.Println("====================================================")
	fmt.Println("1. Coba lagi")
	fmt.Println("2. Keluar ke menu awal")
	fmt.Print("Pilih opsi: ")
	scanner.Scan()
	choice := scanner.Text()

	switch choice {
	case "1":
		login()
	case "2":
		main()
	default:
		fmt.Println("")
		fmt.Println("Pilihan tidak valid.")
		fmt.Println("")
	}
}

func userMenu() {
	clearScreen()
	fmt.Println("        Selamat datang di aplikasi pariwisata   ")
	for {
		fmt.Println("                  Menu Utama User   ")
		fmt.Println("====================================================")
		fmt.Println("Silakan pilih opsi berikut:")
		fmt.Println("1. Lihat daftar destinasi wisata")
		fmt.Println("2. Cari destinasi wisata")
		fmt.Println("3. Urutkan daftar destinasi wisata")
		fmt.Println("0. Keluar")
		fmt.Println("====================================================")
		fmt.Print("Masukkan pilihan (angka) : ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			displayTouristSpots()
		case "2":
			searchTouristSpot()
		case "3":
			sortTouristSpots()
		case "0":
			clearScreen()
			return
		default:
			fmt.Println("")
			fmt.Println("Pilihan tidak valid.")
			fmt.Println("")
		}
	}
}

func displayTouristSpots() {
	clearScreen()
	fmt.Println("              Daftar destinasi wisata         ")
	fmt.Println("====================================================")
	for _, spot := range spots {
		if spot.nama != "" {
			fmt.Printf("Nama destinasi wisata : %s\n", spot.nama)
			fmt.Printf("Lokasi: %s\n", spot.lokasi)
			fmt.Printf("Jarak: %.1f km\n", spot.jarak)
			fmt.Printf("Biaya: Rp %d\n", spot.biaya)
			fmt.Printf("Rating: %.1f\n", spot.rating)
			fmt.Println("====================================================")
		}
	}
	fmt.Println("Melihat daftar destinasi wisata")
	fmt.Println("")
}

func compare(a, b destinasi, sortBy string, ascending bool) bool {
	switch sortBy {
	case "nama":
		if ascending {
			return a.nama < b.nama
		} else {
			return a.nama > b.nama
		}
	case "rating":
		if ascending {
			return a.rating < b.rating
		} else {
			return a.rating > b.rating
		}
	case "biaya":
		if ascending {
			return a.biaya < b.biaya
		} else {
			return a.biaya > b.biaya
		}
	case "jarak":
		if ascending {
			return a.jarak < b.jarak
		} else {
			return a.jarak > b.jarak
		}
	}
	return false
}

func insertionSort(spots *[20]destinasi, sortBy string, ascending bool) {
	for i := 1; i < len(spots); i++ {
		key := spots[i]
		j := i - 1
		if ascending {
			for j >= 0 && compare(spots[j], key, sortBy, false) {
				spots[j+1] = spots[j]
				j--
			}
		} else {
			for j >= 0 && compare(spots[j], key, sortBy, true) {
				spots[j+1] = spots[j]
				j--
			}
		}
		spots[j+1] = key
	}
}
func selectionSort(spots *[20]destinasi, sortBy string, ascending bool) {
	for i := 0; i < len(spots); i++ {
		extremeIndex := i
		for j := i + 1; j < len(spots); j++ {
			if compare(spots[j], spots[extremeIndex], sortBy, ascending) {
				extremeIndex = j
			}
		}
		spots[i], spots[extremeIndex] = spots[extremeIndex], spots[i]
	}
}

func binarySearch(spots *[20]destinasi, nama string) int {
	low, high := 0, 19
	nama = strings.ToLower(nama)

	for low <= high {
		mid := low + (high-low)/2
		spotName := strings.ToLower(spots[mid].nama)

		if spotName < nama {
			low = mid + 1
		} else if spotName > nama {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func searchTouristSpot() {
	clearScreen()
	fmt.Println("               Cari destinasi wisata         ")
	fmt.Println("====================================================")
	fmt.Print("Masukkan nama destinasi wisata : ")
	scanner.Scan()
	fmt.Println("====================================================")

	nama := strings.TrimSpace(scanner.Text())

	if nama == "" {
		fmt.Println("Nama destinasi tidak boleh kosong.")
		return
	}

	//Pencarian Binary
	selectionSort(&spots, "nama", true)
	index := binarySearch(&spots, nama)
	if index != -1 {
		spot := spots[index]
		fmt.Println("Pencarian Ditemukan")
		fmt.Println("Hasil pencarian dengan konsep Binary Search:")
		fmt.Println("")
		fmt.Printf("Nama destinasi wisata: %s\n", spot.nama)
		fmt.Printf("Lokasi: %s\n", spot.lokasi)
		fmt.Printf("Jarak: %.1f km\n", spot.jarak)
		fmt.Printf("Biaya: Rp %d\n", spot.biaya)
		fmt.Printf("Rating: %.1f\n", spot.rating)
		fmt.Println("====================================================")
	} else {
		fmt.Println("Destinasi wisata tidak ditemukan dengan metode Binary Search")
		fmt.Println("====================================================")
		fmt.Println("")
	}

	//Pencarian Sequential
	found := false
	for i := 0; i < len(spots); i++ {
		if strings.ToLower(spots[i].nama) == strings.ToLower(nama) {
			fmt.Println("Pencarian Ditemukan")
			fmt.Println("Hasil pencarian dengan konsep Sequential Search:")
			fmt.Println("")
			fmt.Printf("Nama destinasi wisata: %s\n", spots[i].nama)
			fmt.Printf("Lokasi: %s\n", spots[i].lokasi)
			fmt.Printf("Jarak: %.1f km\n", spots[i].jarak)
			fmt.Printf("Biaya: Rp %d\n", spots[i].biaya)
			fmt.Printf("Rating: %.1f\n", spots[i].rating)
			found = true
		}
	}
	if !found {
		fmt.Println("Destinasi wisata tidak ditemukan dengan metode Sequential Search")
		fmt.Println("====================================================")
		fmt.Println("")
	}
	fmt.Println("====================================================")
}

func sortTouristSpots() {
	clearScreen()
	fmt.Println("         Urutkan destinasi wisata berdasarkan:    ")
	fmt.Println("====================================================")
	fmt.Println("1. Rating")
	fmt.Println("2. Biaya")
	fmt.Println("3. Jarak")
	fmt.Println("0. Kembali ke menu user")
	fmt.Print("Masukkan pilihan (angka) : ")
	scanner.Scan()
	sortChoice := scanner.Text()

	var sortKey string
	var ascending bool = true

	switch sortChoice {
	case "1":
		sortKey = "rating"
	case "2":
		sortKey = "biaya"
	case "3":
		sortKey = "jarak"
	case "0":
		clearScreen()
		return
	default:
		fmt.Println("")
		fmt.Println("Pilihan tidak valid.")
		fmt.Println("Kembali ke menu user")
		fmt.Println("")
		return
	}

	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	fmt.Println("0. Kembali ke menu user")
	fmt.Print("Pilih urutan : ")
	scanner.Scan()
	orderChoice := scanner.Text()

	switch orderChoice {
	case "1":
		insertionSort(&spots, sortKey, true)
		ascending = true
		fmt.Println("Order: Ascending")
	case "2":
		insertionSort(&spots, sortKey, false)
		ascending = false
		fmt.Println("Order: Descending")
	case "0":
		sortTouristSpots()
	default:
		fmt.Println("")
		fmt.Println("Pilihan tidak valid.")
		fmt.Println("Kembali ke menu user")
		fmt.Println("")
		return
	}

	clearScreen()
	fmt.Println("          Daftar destinasi wisata terurut    ")
	fmt.Println("====================================================")
	fmt.Println("Sorting by: ", sortKey)
	if ascending {
		fmt.Println("Order: Ascending")
	} else {
		fmt.Println("Order: Descending")
	}
	fmt.Println()
	for _, spot := range spots {
		if spot.nama != "" {
			fmt.Printf("Nama destinasi wisata : %s\n", spot.nama)
			fmt.Printf("Lokasi: %s\n", spot.lokasi)
			fmt.Printf("Jarak: %.1f km\n", spot.jarak)
			fmt.Printf("Biaya: Rp %d\n", spot.biaya)
			fmt.Printf("Rating: %.1f\n", spot.rating)
			fmt.Println("====================================================")
		}
	}
	fmt.Println("Berhasil mengurutkan destinasi wisata!")
	fmt.Println("")
}

func adminMenu() {
	clearScreen()
	for {
		fmt.Println("                Menu Utama Admin!   ")
		fmt.Println("====================================================")
		fmt.Println("Silakan pilih opsi berikut:")
		fmt.Println("1. Tambah destinasi wisata")
		fmt.Println("2. Ubah destinasi wisata")
		fmt.Println("3. Hapus destinasi wisata")
		fmt.Println("4. Lihat daftar destinasi wisata")
		fmt.Println("5. Ke Menu User")
		fmt.Println("0. Keluar")
		fmt.Println("====================================================")
		fmt.Print("Masukkan pilihan (angka) : ")
		scanner.Scan()

		choice := scanner.Text()

		switch choice {
		case "1":
			addTouristSpot()
		case "2":
			editTouristSpot()
		case "3":
			removeTouristSpot()
		case "4":
			displayTouristSpots()
		case "5":
			userMenu()
		case "0":
			clearScreen()
			return
		default:
			fmt.Println("")
			fmt.Println("Pilihan tidak valid.")
			fmt.Println("")
		}
	}
}

func addTouristSpot() {
	clearScreen()
	fmt.Println("             Tambah destinasi wisata         ")
	fmt.Println("====================================================")

	fmt.Print("Masukkan nama destinasi wisata : ")
	scanner.Scan()
	nama := scanner.Text()
	if nama == "" {
		fmt.Println("Nama destinasi wisata tidak boleh kosong.")
		fmt.Println("Kembali ke menu admin")
		fmt.Println("")
		return
	}

	fmt.Print("Masukkan lokasi destinasi wisata : ")
	scanner.Scan()
	lokasi := scanner.Text()
	if lokasi == "" {
		fmt.Println("Lokasi destinasi wisata tidak boleh kosong.")
		fmt.Println("Kembali ke menu admin")
		fmt.Println("")
		return
	}

	fmt.Print("Masukkan jarak destinasi wisata (dalam km) : ")
	scanner.Scan()
	jarakInput := scanner.Text()
	jarak, err := strconv.ParseFloat(jarakInput, 64)
	if err != nil {
		fmt.Println("Input jarak tidak valid. Silakan masukkan angka.")
		fmt.Println("Kembali ke menu admin")
		fmt.Println("")
		return
	}

	fmt.Print("Masukkan biaya destinasi wisata : ")
	scanner.Scan()
	biayaInput := scanner.Text()
	biaya, err := strconv.Atoi(biayaInput)
	if err != nil {
		fmt.Println("Input biaya tidak valid. Silakan masukkan angka.")
		fmt.Println("Kembali ke menu admin")
		fmt.Println("")
		return
	}

	fmt.Print("Masukkan rating destinasi wisata (1.0 - 5.0) : ")
	scanner.Scan()
	ratingInput := scanner.Text()
	rating, err := strconv.ParseFloat(ratingInput, 64)
	if rating < 1.0 || rating > 5.0 {
		fmt.Println("Input rating tidak valid. Silakan masukkan angka antara 1.0 dan 5.0.")
		fmt.Println("Kembali ke menu admin")
		fmt.Println("")
		return
	}

	added := false
	for i := 0; i < len(spots) && !added; i++ {
		if spots[i].nama == "" {
			spots[i] = destinasi{nama: nama, lokasi: lokasi, jarak: jarak, biaya: biaya, rating: rating}
			added = true
		}
	}

	if !added {
		fmt.Println("Tidak ada ruang tersisa untuk menambahkan destinasi baru.")
	} else {
		displayTouristSpots()
		fmt.Println("Destinasi wisata berhasil ditambahkan!")
	}
	fmt.Println("")
}

func editTouristSpot() {
	clearScreen()
	fmt.Println("             Ubah destinasi wisata         ")
	fmt.Println("====================================================")
	fmt.Println("Daftar destinasi wisata:")
	displayTouristSpots()
	if len(spots) == 0 {
		fmt.Println("Belum ada data tempat wisata.")
		return
	}

	fmt.Print("Masukkan nama destinasi wisata yang ingin diubah : ")
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	selectionSort(&spots, "nama", true) // Sort spots by name in ascending order

	index := -1
	for i, spot := range spots {
		if strings.ToLower(spot.nama) == strings.ToLower(input) {
			index = i
		}
		if index != -1 {
		}
	}
	if index == -1 {
		fmt.Println("")
		fmt.Println("Destinasi tidak ditemukan. Harap masukkan nama yang sesuai dengan daftar.")
		fmt.Println("Kembali ke menu admin")
		fmt.Println("")
		return
	}

	clearScreen()
	destinasiLama := spots[index]
	fmt.Printf("Mengubah data '%s'\n", destinasiLama.nama)
	fmt.Println("====================================================")
	fmt.Print("Nama Tempat Wisata Baru : ")
	scanner.Scan()
	newName := scanner.Text()
	if newName == "" {
		newName = destinasiLama.nama
	}

	fmt.Print("Lokasi Tempat Wisata Baru : ")
	scanner.Scan()
	newLocation := scanner.Text()
	if newLocation == "" {
		newLocation = destinasiLama.lokasi
	}

	fmt.Print("Jarak Tempat Wisata Baru (dalam km) : ")
	scanner.Scan()
	newDistanceInput := scanner.Text()
	var newDistance float64
	var err error
	if newDistanceInput == "" {
		newDistance = destinasiLama.jarak
	} else {
		newDistance, err = strconv.ParseFloat(newDistanceInput, 64)
		if err != nil {
			fmt.Println("Input jarak tidak valid. Menggunakan jarak lama.")
			newDistance = destinasiLama.jarak
		}
	}

	fmt.Print("Biaya Tempat Wisata Baru (dalam IDR) : ")
	scanner.Scan()
	newCostInput := scanner.Text()
	var newCost int
	if newCostInput == "" {
		newCost = destinasiLama.biaya
	} else {
		newCost, err = strconv.Atoi(newCostInput)
		if err != nil {
			fmt.Println("Input biaya tidak valid. Menggunakan biaya lama.")
			newCost = destinasiLama.biaya
		}
	}

	fmt.Print("Rating Tempat Wisata Baru (1.0 - 5.0) : ")
	scanner.Scan()
	newRatingInput := scanner.Text()
	var newRating float64
	if newRatingInput == "" {
		newRating = destinasiLama.rating
	} else {
		newRating, err = strconv.ParseFloat(newRatingInput, 64)
		if err != nil {
			fmt.Println("Input rating tidak valid. Menggunakan rating lama.")
			newRating = destinasiLama.rating
		}
	}

	spots[index] = destinasi{nama: newName, lokasi: newLocation, jarak: newDistance, biaya: newCost, rating: newRating}
	fmt.Println("====================================================")
	displayTouristSpots()
	fmt.Println("Data destinasi wisata berhasil diubah!")
	fmt.Println("")
}

func removeTouristSpot() {
	clearScreen()
	fmt.Println("             Hapus destinasi wisata         ")
	fmt.Println("====================================================")
	fmt.Println("Daftar destinasi wisata:")
	displayTouristSpots()

	fmt.Print("Masukkan nama destinasi wisata yang ingin dihapus : ")
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	selectionSort(&spots, "nama", true)

	index := -1
	for i, spot := range spots {
		if strings.ToLower(spot.nama) == strings.ToLower(input) {
			index = i
		}
		if index != -1 {
		}
	}
	if index == -1 {
		fmt.Println("")
		fmt.Println("Destinasi tidak ditemukan. Harap masukkan nama yang sesuai dengan daftar.")
		fmt.Println("Kembali ke menu admin")
		fmt.Println("")
		return
	} else {
		for i := index; i < len(spots)-1; i++ {
			spots[i] = spots[i+1]
		}
		spots[len(spots)-1] = destinasi{}

		displayTouristSpots()
		fmt.Println("Destinasi wisata berhasil dihapus!")
	}
	fmt.Println("")
}
