package main

import "fmt"

func main() {
	var nim []string
	var nama []string
	var nilai []float64
	var lulus []bool
	var jumlahMahasiswa int = 0

	for {
		fmt.Println("=== MENU DATA MAHASISWA ===")
		fmt.Println("1. Tambah Mahasiswa")
		fmt.Println("2. Tampilkan Data")
		fmt.Println("3. Cari Mahasiswa")
		fmt.Println("4. Hitung Rata-rata")
		fmt.Println("5. Jumlah Mahasiswa Lulus")
		fmt.Println("6. Keluar")

		fmt.Print("Pilihan: ")
		pilih := 0
		fmt.Scan(&pilih)

		if pilih == 1 {
			// Input mahasiswa
			var tempNama string
			var tempNIM string
			var tempNilai float64

			fmt.Print("Nama   : ")
			fmt.Scan(&tempNama)

			fmt.Print("NIM    : ")
			fmt.Scan(&tempNIM)

			fmt.Print("Nilai  : ")
			fmt.Scan(&tempNilai)

			// Simpan ke slice
			nama = append(nama, tempNama)
			nim = append(nim, tempNIM)
			nilai = append(nilai, tempNilai)

			// Tentukan kelulusan
			if tempNilai >= 60 {
				lulus = append(lulus, true)
			} else {
				lulus = append(lulus, false)
			}

			jumlahMahasiswa = jumlahMahasiswa + 1
			fmt.Println("Data mahasiswa berhasil ditambahkan.\n")

		} else if pilih == 2 {
			fmt.Println("\n=== DATA MAHASISWA ===")
			for i := 0; i < jumlahMahasiswa; i++ {
				fmt.Printf("Nama : %s | NIM : %s | Nilai : %.2f | Lulus: %t\n",
					nama[i], nim[i], nilai[i], lulus[i])
			}
			fmt.Println()

		} else if pilih == 3 {
			var cari string
			fmt.Print("Masukkan NIM yang dicari: ")
			fmt.Scan(&cari)

			found := false
			for i := 0; i < jumlahMahasiswa; i++ {
				if nim[i] == cari {
					fmt.Println("Data ditemukan!")
					fmt.Printf("Nama : %s | NIM : %s | Nilai : %.2f | Lulus : %t\n",
						nama[i], nim[i], nilai[i], lulus[i])
					found = true
					break
				}
			}

			if !found {
				fmt.Println("Data tidak ditemukan.")
			}
			fmt.Println()

		} else if pilih == 4 {
			if jumlahMahasiswa == 0 {
				fmt.Println("Belum ada data.")
			} else {
				total := 0.0
				for i := 0; i < jumlahMahasiswa; i++ {
					total = total + nilai[i]
				}
				avg := total / float64(jumlahMahasiswa)
				fmt.Println("Rata-rata nilai seluruh mahasiswa:", avg)
			}
			fmt.Println()

		} else if pilih == 5 {
			count := 0
			for i := 0; i < jumlahMahasiswa; i++ {
				if lulus[i] {
					count++
				}
			}
			fmt.Println("Jumlah mahasiswa yang lulus:", count)
			fmt.Println()

		} else if pilih == 6 {
			fmt.Println("Program selesai.")
			break

		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
