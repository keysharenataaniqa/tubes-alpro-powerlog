package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const NMAX = 100

type Perangkat struct {
	nama, ruangan string
	watt, durasi  float64
	jam, menit    int
}
type TabPerangkat [NMAX]Perangkat

var data TabPerangkat

func clearScreen() { // procedure untuk membersihkan tampilan terminal
	var cmd *exec.Cmd                      // mendeklarasikan variabel cmd, cmd digunakan untuk menyimpan perintah terminal yang akan dijalankan, *exec.Cmd merupakan tipe data untuk menjalankan perintah sistem operasi
	cmd = exec.Command("cmd", "/c", "cls") // membuat perintah terminal, "cmd" digunakan untuk membuka Command Prompt Windows, "/c" digunakan agar cmd menjalankan perintah lalu selesai, "cls" digunakan untuk membersihkan layar terminal
	cmd.Stdout = os.Stdout                 // mengarahkan output perintah ke terminal program, cmd adalah perintah yang telah dibuat sebelumnya, os.Stdout merepresentasikan layar terminal
	cmd.Run()                              // menjalankan perintah yang tersimpan pada cmd, setelah dijalankan layar terminal akan dibersihkan
}
func bacaString() string { // function untuk membaca input string dari keyboard dan mengembalikannya, string setelah kurung menunjukkan tipe nilai yang dikembalikan function
	var reader *bufio.Reader // mendeklarasikan variabel reader, reader digunakan untuk membaca input dari keyboard, *bufio.Reader adalah tipe data yang dapat membaca satu baris input sekaligus termasuk spasi
	var str string           // mendeklarasikan variabel str, str digunakan untuk menyimpan input yang dimasukkan pengguna, string digunakan karena data yang dibaca berupa teks

	reader = bufio.NewReader(os.Stdin) // membuat objek reader baru, bufio.NewReader digunakan untuk membaca input secara penuh, os.Stdin menunjukkan sumber input berasal dari keyboard
	str, _ = reader.ReadString('\n')   // membaca input sampai pengguna menekan Enter, hasil input disimpan ke str, '\n' menandakan akhir input ketika tombol Enter ditekan, _ digunakan untuk mengabaikan nilai error yang tidak digunakan
	str = strings.TrimSpace(str)       // membersihkan input yang dibaca, TrimSpace menghapus spasi, tab, dan Enter di awal maupun akhir string, hasil yang sudah bersih disimpan kembali ke str

	for str == "" { // mengulang proses input selama str kosong, str == "" berarti pengguna belum memasukkan karakter apa pun
		str, _ = reader.ReadString('\n') // membaca ulang input dari pengguna, hasil input baru disimpan kembali ke str, _ digunakan untuk mengabaikan nilai error
		str = strings.TrimSpace(str)
	}
	return str // membersihkan kembali input yang baru dimasukkan, memastikan spasi dan Enter tidak ikut tersimpan pada str
}
func tambahData(n *int) { // procedure untuk menambahkan data perangkat ke dalam array, n menyimpan jumlah data yang sudah terisi, *int digunakan agar nilai n di main dapat berubah
	var hasilPerangkat, hasilRuangan, hasilWatt, hasilJam, hasilMenit string // variabel status validasi, digunakan untuk menentukan apakah input sudah valid atau belum
	var sukses int                                                           // menyimpan jumlah variabel yang berhasil dibaca oleh fmt.Scan, digunakan untuk mendeteksi input angka yang salah
	var buang, kembali string                                                // buang digunakan untuk menampung input yang salah saat membersihkan buffer, kembali digunakan untuk validasi sebelum kembali ke menu

	clearScreen() // membersihkan tampilan terminal agar tampilan program lebih rapi
	fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("  ║                           ➕ TAMBAH DATA ➕                           ║")
	fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝")
	fmt.Println()

	if *n < NMAX { // memeriksa apakah jumlah data masih kurang dari kapasitas maksimum array, *n adalah jumlah data saat ini, NMAX adalah kapasitas maksimum array
		hasilPerangkat = "tidak_valid"        // menginisialisasi status nama perangkat sebagai tidak valid sebelum proses input dimulai
		for hasilPerangkat == "tidak_valid" { // mengulang proses input nama selama status masih tidak valid
			fmt.Print("    🏷️  Masukkan Nama Perangkat : ")
			data[*n].nama = bacaString()  // membaca nama perangkat dari pengguna, data[*n] menunjuk data yang sedang diisi, .nama mengakses field nama pada struct
			if cekString(data[*n].nama) { // memeriksa apakah nama perangkat valid, cekString mengembalikan true jika tidak ditemukan angka
				hasilPerangkat = "valid" // mengubah status menjadi valid agar perulangan berhenti
			} else {
				fmt.Println("  ❌ Nama Perangkat tidak valid!\n")
			}
		}

		hasilWatt = "tidak_valid"        // menginisialisasi status watt sebagai tidak valid
		for hasilWatt == "tidak_valid" { // mengulang proses input watt selama status masih tidak valid
			fmt.Print("    ⚡  Masukkan Watt           : ")
			sukses, _ = fmt.Scan(&data[*n].watt) // membaca input watt dan menyimpannya ke field watt pada data ke-n, sukses menyimpan jumlah input yang berhasil dibaca, _ digunakan untuk mengabaikan nilai error
			if sukses == 0 {                     // memeriksa apakah input gagal dibaca sebagai angka
				fmt.Scanln(&buang) // membersihkan sisa input yang salah dari buffer, buang digunakan sebagai penampung sementara agar input yang salah tidak terbaca kembali
				fmt.Println("  ❌ Input tidak valid! Harap masukkan angka.\n")
			} else if data[*n].watt > 0.0 { // memeriksa apakah watt lebih besar dari nol, watt tidak boleh bernilai 0 atau negatif
				hasilWatt = "valid" // mengubah status watt menjadi valid agar perulangan berhenti
			} else {
				fmt.Println("  ❌ Watt tidak valid! Harus lebih dari 0.\n")
			}
		}

		hasilJam = "tidak_valid"        // menginisialisasi status jam sebagai tidak valid
		for hasilJam == "tidak_valid" { // mengulang proses input jam selama status masih tidak valid
			fmt.Print("    ⏱️  Masukkan Durasi (Jam)   : ")
			sukses, _ = fmt.Scan(&data[*n].jam) // membaca input jam dan menyimpannya ke field jam pada data ke-n
			if sukses == 0 {                    // memeriksa apakah input jam gagal dibaca sebagai angka
				fmt.Scanln(&buang)
				fmt.Println("  ❌ Input tidak valid! Harap masukkan angka.\n")
			} else if data[*n].jam >= 0 { // memeriksa apakah jam tidak bernilai negatif
				hasilJam = "valid" // mengubah status jam menjadi valid agar perulangan berhenti
			} else {
				fmt.Println("  ❌ Jam tidak valid! Tidak boleh minus.\n")
			}
		}

		hasilMenit = "tidak_valid"        // menginisialisasi status menit sebagai tidak valid
		for hasilMenit == "tidak_valid" { // mengulang proses input menit selama status masih tidak valid
			fmt.Print("    ⏱️  Masukkan Durasi (Menit) : ")
			sukses, _ = fmt.Scan(&data[*n].menit) // membaca input menit dan menyimpannya ke field menit pada data ke-n
			if sukses == 0 {                      // memeriksa apakah input menit gagal dibaca sebagai angka
				fmt.Scanln(&buang)
				fmt.Println("  ❌ Input tidak valid! Harap masukkan angka.\n")
			} else if data[*n].menit >= 0 && data[*n].menit < 60 { // memeriksa apakah menit berada pada rentang 0 sampai 59, menit tidak boleh negatif dan tidak boleh mencapai 60
				if data[*n].jam == 0 && data[*n].menit == 0 { // memeriksa apakah total durasi bernilai 0 jam dan 0 menit, durasi penggunaan tidak boleh kosong
					fmt.Println("  ❌ Durasi total tidak boleh 0! Masukkan minimal 1 menit.\n")
				} else {
					hasilMenit = "valid" // mengubah status menit menjadi valid jika seluruh syarat terpenuhi
				}
			} else {
				fmt.Println("  ❌ Menit tidak valid! Harap masukkan angka 0 - 59.\n")
			}
		}

		hasilRuangan = "tidak_valid"        // menginisialisasi status ruangan sebagai tidak valid
		for hasilRuangan == "tidak_valid" { // mengulang proses input ruangan selama status masih tidak valid
			fmt.Print("    🏠  Masukkan Ruangan        : ")
			data[*n].ruangan = bacaString()  // membaca nama ruangan dari pengguna dan menyimpannya ke field ruangan pada data ke-n
			if cekString(data[*n].ruangan) { // memeriksa apakah nama ruangan valid dan tidak mengandung angka
				hasilRuangan = "valid" // mengubah status ruangan menjadi valid agar perulangan berhenti
			} else {
				fmt.Println("  ❌ Nama Ruangan tidak valid!\n")
			}
		}

		*n = *n + 1 // menambah jumlah data yang tersimpan, *n adalah jumlah data saat ini, +1 karena ada satu data baru yang berhasil ditambahkan

		fmt.Println("\n  ╔═══════════════════════════════════════════════════════════════════════╗")
		fmt.Println("  ║                     ✨ Data berhasil ditambahkan! ✨                  ║")
		fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")

	} else {
		fmt.Println("\n  ╔═══════════════════════════════════════════════════════════════════════╗")
		fmt.Println("  ║                       ⚠️  Kapasitas Data Penuh!                       ║")
		fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")
	}

	fmt.Println("Ketik 0 dan tekan [Enter] untuk kembali... ")
	kembali = ""         // mengosongkan variabel kembali sebelum proses validasi dilakukan
	for kembali != "0" { // mengulang selama pengguna belum memasukkan angka 0
		fmt.Scan(&kembali)  // membaca input pengguna dan menyimpannya ke variabel kembali
		if kembali != "0" { // memeriksa apakah input yang dimasukkan bukan angka 0
			fmt.Println("  ❌ Input tidak valid! Harap ketik 0: ")
		}
	}
}
func cekString(input string) bool { // function untuk memeriksa apakah string mengandung angka atau tidak, input berisi string yang akan diperiksa, bool digunakan untuk mengembalikan nilai true atau false
	var cekAngka bool // mendeklarasikan variabel cekAngka, digunakan untuk menyimpan hasil pemeriksaan string, bool hanya dapat bernilai true atau false
	var i int         // mendeklarasikan variabel i, digunakan sebagai pencacah untuk mengakses setiap karakter dalam string
	if input == "" || input == `""` || input == "''" || input == "-" || input == " " {
		return false
	}
	cekAngka = true                  // menginisialisasi hasil pemeriksaan sebagai true, diasumsikan input valid dan tidak mengandung angka sebelum dilakukan pengecekan
	for i = 0; i < len(input); i++ { // mengulang pemeriksaan setiap karakter dalam string, i dimulai dari 0 karena index string dimulai dari 0, len(input) digunakan untuk mendapatkan panjang string, i++ digunakan untuk berpindah ke karakter berikutnya
		if input[i] >= '0' && input[i] <= '9' { // memeriksa apakah karakter pada index i merupakan angka, '0' adalah angka terkecil dan '9' adalah angka terbesar pada karakter ASCII, && digunakan agar kedua syarat harus terpenuhi
			cekAngka = false // mengubah hasil pemeriksaan menjadi false, false menandakan ditemukan angka pada string sehingga input dianggap tidak valid
		}
	}
	return cekAngka // mengembalikan hasil pemeriksaan, true berarti tidak ada angka pada string, false berarti terdapat angka pada string
}
func tampilData(n int) { // procedure untuk menampilkan seluruh data perangkat yang tersimpan, n berisi jumlah data yang sedang tersimpan dalam array
	var i int
	var kembali string

	clearScreen()

	fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("  ║                      📋 TAMPILKAN DATA PERANGKAT 📋                   ║")
	fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝")
	fmt.Println()

	if n == 0 { // memeriksa apakah jumlah data yang tersimpan sama dengan 0, n menyimpan banyaknya data yang ada dalam array
		fmt.Println("  📭 Data perangkat masih kosong. Yuk, tambah data dulu!\n")
	} else { // dijalankan jika terdapat minimal satu data yang tersimpan dalam array
		fmt.Println("  ┌────┬──────────────────────┬─────────┬──────────┬──────────────────────┐")
		fmt.Println("  │ No │ Nama Perangkat       │ Watt    │ Durasi   │ Ruangan              │")
		fmt.Println("  ├────┼──────────────────────┼─────────┼──────────┼──────────────────────┤")

		for i = 0; i < n; i++ { // mengulang dari data pertama sampai data terakhir, i dimulai dari 0 karena index array dimulai dari 0, i++ digunakan untuk berpindah ke data berikutnya
			fmt.Printf("  │ %-2d │ %-20s │ %-5.1f W │ %2dj %02dm  │ %-20s │\n", i+1, data[i].nama, data[i].watt, data[i].jam, data[i].menit, data[i].ruangan)
		} // menampilkan data dalam bentuk tabel, Printf digunakan agar tampilan dapat diatur rapi sesuai format yang ditentukan
		fmt.Println("  └────┴──────────────────────┴─────────┴──────────┴──────────────────────┘")
	}

	fmt.Println("Ketik 0 dan tekan [Enter] untuk kembali... ")
	kembali = ""         // mengosongkan variabel kembali sebelum proses validasi input dimulai
	for kembali != "0" { // mengulang selama pengguna belum memasukkan angka 0
		fmt.Scan(&kembali)
		if kembali != "0" { // memeriksa apakah input yang dimasukkan bukan angka 0
			fmt.Println("  ❌ Input tidak valid! Harap ketik 0: ")
		}
	}
}
func updateData(n int) { // procedure untuk mengubah data perangkat yang sudah ada, n menyimpan jumlah data yang tersimpan dalam array
	var nData, i, sukses, suksesWatt, suksesJam, suksesMenit int                             // nData menyimpan nomor data yang dipilih pengguna, i digunakan untuk perulangan, sukses digunakan untuk validasi input nomor data, suksesWatt suksesJam suksesMenit digunakan untuk validasi input angka
	var hasilPerangkat, hasilRuangan, hasilWatt, hasilJam, hasilMenit, buang, kembali string // hasilPerangkat hasilRuangan hasilWatt hasilJam hasilMenit digunakan sebagai status validasi input, buang digunakan untuk membersihkan buffer, kembali digunakan untuk validasi kembali ke menu
	var validFormat bool                                                                     // menyimpan status apakah format input nomor data sudah benar atau belum, bool hanya bernilai true atau false

	clearScreen()

	fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("  ║                             ✏️  UBAH DATA ✏️                            ║")
	fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝")

	if n == 0 { // memeriksa apakah tidak ada data yang tersimpan, n menyimpan jumlah data dalam array
		fmt.Println("  📭 Data perangkat masih kosong!\n")
	} else { // dijalankan jika terdapat minimal satu data yang dapat diubah
		fmt.Println("  ┌────┬──────────────────────┬─────────┬──────────┬──────────────────────┐")
		fmt.Println("  │ No │ Nama Perangkat       │ Watt    │ Durasi   │ Ruangan              │")
		fmt.Println("  ├────┼──────────────────────┼─────────┼──────────┼──────────────────────┤")

		for i = 0; i < n; i++ { // menampilkan seluruh data yang tersedia, i digunakan untuk mengakses setiap data dalam array
			fmt.Printf("  │ %-2d │ %-20s │ %-5.1f W │ %2dj %02dm  │ %-20s │\n", i+1, data[i].nama, data[i].watt, data[i].jam, data[i].menit, data[i].ruangan)
		}
		fmt.Println("  └────┴──────────────────────┴─────────┴──────────┴──────────────────────┘")

		validFormat = false        // menginisialisasi status validasi nomor data sebagai belum valid
		for validFormat == false { // mengulang input nomor data sampai pengguna memasukkan format yang benar
			fmt.Print("  🔢 Masukkan nomor data yang diubah : ")
			sukses, _ = fmt.Scan(&nData) // membaca nomor data yang dipilih pengguna, nData menyimpan nomor tersebut, sukses menyimpan jumlah input yang berhasil dibaca

			if sukses == 0 { // memeriksa apakah pengguna memasukkan huruf atau format yang salah
				fmt.Scanln(&buang) // membersihkan sisa input yang salah dari buffer agar tidak terbaca lagi pada perulangan berikutnya
				fmt.Println("  ❌ Input tidak valid! Harap masukkan angka, bukan huruf.\n")
			} else {
				validFormat = true // mengubah status validasi menjadi benar karena nomor data berhasil dibaca
			}
		}
		nData = nData - 1 // mengubah nomor urut menjadi index array, -1 dilakukan karena nomor tabel dimulai dari 1 sedangkan index array dimulai dari 0

		if nData >= 0 && nData < n { // memeriksa apakah nomor data yang dipilih berada dalam rentang data yang tersedia
			fmt.Println("\n  ╔═══════════════════════════════════════════════════════════════════════╗")
			fmt.Println("  ║                        ✨ MASUKKAN DATA BARU ✨                       ║")
			fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝")
			fmt.Println()

			hasilPerangkat = "tidak_valid"
			for hasilPerangkat == "tidak_valid" {
				fmt.Print("    🏷️  Masukkan Nama Perangkat : ")
				data[nData].nama = bacaString() // menyimpan nama perangkat baru ke data yang dipilih, nData menunjukkan index data yang akan diubah
				if cekString(data[nData].nama) {
					hasilPerangkat = "valid"
				} else {
					fmt.Println("  ❌ Nama Perangkat tidak valid!\n")
				}
			}

			hasilWatt = "tidak_valid"
			for hasilWatt == "tidak_valid" {
				fmt.Print("    ⚡  Masukkan Watt           : ")
				suksesWatt, _ = fmt.Scan(&data[nData].watt) // membaca nilai watt baru dan menyimpannya ke data yang dipilih, suksesWatt digunakan untuk memvalidasi input angka

				if suksesWatt == 0 {
					fmt.Scanln(&buang)
					fmt.Println("  ❌ Input tidak valid! Harap masukkan angka.\n")
				} else if data[nData].watt > 0.0 {
					hasilWatt = "valid"
				} else {
					fmt.Println("  ❌ Watt tidak valid! Harus lebih dari 0.\n")
				}
			}

			hasilJam = "tidak_valid"
			for hasilJam == "tidak_valid" {
				fmt.Print("    ⏱️  Masukkan Durasi (Jam)   : ")
				suksesJam, _ = fmt.Scan(&data[nData].jam) // membaca nilai jam baru dan menyimpannya ke data yang dipilih

				if suksesJam == 0 {
					fmt.Scanln(&buang)
					fmt.Println("  ❌ Input tidak valid! Harap masukkan angka.\n")
				} else if data[nData].jam >= 0 {
					hasilJam = "valid"
				} else {
					fmt.Println("  ❌ Jam tidak valid! Tidak boleh minus.\n")
				}
			}

			hasilMenit = "tidak_valid"
			for hasilMenit == "tidak_valid" {
				fmt.Print("    ⏱️  Masukkan Durasi (Menit) : ")
				suksesMenit, _ = fmt.Scan(&data[nData].menit) // membaca nilai menit baru dan menyimpannya ke data yang dipilih

				if suksesMenit == 0 {
					fmt.Scanln(&buang)
					fmt.Println("  ❌ Input tidak valid! Harap masukkan angka.\n")
				} else if data[nData].menit >= 0 && data[nData].menit < 60 {
					if data[nData].jam == 0 && data[nData].menit == 0 {
						fmt.Println("  ❌ Durasi total tidak boleh 0! Masukkan minimal 1 menit.\n")
					} else {
						hasilMenit = "valid"
					}
				} else {
					fmt.Println("  ❌ Menit tidak valid! Harap masukkan angka 0 - 59.\n")
				}
			}

			hasilRuangan = "tidak_valid"
			for hasilRuangan == "tidak_valid" {
				fmt.Print("    🏠  Masukkan Ruangan        : ")
				data[nData].ruangan = bacaString() // membaca nama ruangan baru dan menyimpannya ke data yang dipilih
				if cekString(data[nData].ruangan) {
					hasilRuangan = "valid"
				} else {
					fmt.Println("  ❌ Nama Ruangan tidak valid!\n")
				}
			}

			fmt.Println("\n  ╔═══════════════════════════════════════════════════════════════════════╗")
			fmt.Println("  ║                      ✨ Data berhasil diubah! ✨                      ║")
			fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")

		} else { // dijalankan jika nomor data yang dimasukkan tidak ditemukan dalam array
			fmt.Println("\n  ╔═══════════════════════════════════════════════════════════════════════╗")
			fmt.Println("  ║                       ⚠️  Data tidak ditemukan!                        ║")
			fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")
		}
	}

	fmt.Println("Ketik 0 dan tekan [Enter] untuk kembali... ")
	kembali = ""         // mengosongkan variabel kembali sebelum validasi input dilakukan
	for kembali != "0" { // mengulang selama pengguna belum memasukkan angka 0
		fmt.Scan(&kembali)  // membaca input pengguna dan menyimpannya ke variabel kembali
		if kembali != "0" { // memeriksa apakah input yang dimasukkan bukan angka 0
			fmt.Println("  ❌ Input tidak valid! Harap ketik 0: ")
		}
	}
}
func hapusData(n *int) { // procedure untuk menghapus data perangkat dari array, n menyimpan jumlah data yang terisi, *int digunakan agar jumlah data asli dapat berubah setelah penghapusan
	var nData, i, sukses int  // nData menyimpan nomor data yang akan dihapus, i digunakan untuk perulangan, sukses digunakan untuk memvalidasi input angka
	var buang, kembali string // buang digunakan untuk membersihkan input yang salah dari buffer, kembali digunakan untuk validasi kembali ke menu
	var validFormat bool      // menyimpan status apakah format input nomor data sudah benar atau belum, bernilai true atau false

	clearScreen()

	fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("  ║                            🗑️  HAPUS DATA 🗑️                            ║")
	fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝")
	fmt.Println()

	if *n == 0 { // memeriksa apakah jumlah data sama dengan 0, *n menyimpan jumlah data yang terisi dalam array
		fmt.Println("  📭 Data perangkat masih kosong!\n")
	} else { // dijalankan jika masih terdapat data yang dapat dihapus
		fmt.Println("  ┌────┬──────────────────────┬─────────┬──────────┬──────────────────────┐")
		fmt.Println("  │ No │ Nama Perangkat       │ Watt    │ Durasi   │ Ruangan              │")
		fmt.Println("  ├────┼──────────────────────┼─────────┼──────────┼──────────────────────┤")

		for i = 0; i < *n; i++ { // menampilkan seluruh data yang tersimpan, i digunakan untuk mengakses setiap data dalam array
			fmt.Printf("  │ %-2d │ %-20s │ %-5.1f W │ %2dj %02dm  │ %-20s │\n", i+1, data[i].nama, data[i].watt, data[i].jam, data[i].menit, data[i].ruangan)
		}
		fmt.Println("  └────┴──────────────────────┴─────────┴──────────┴──────────────────────┘")

		validFormat = false        // menginisialisasi status validasi nomor data sebagai belum valid
		for validFormat == false { // mengulang input nomor data sampai pengguna memasukkan format yang benar
			fmt.Print("  🔢 Masukkan nomor data yang dihapus : ")
			sukses, _ = fmt.Scan(&nData) // membaca nomor data yang akan dihapus, nData menyimpan nomor tersebut, sukses menyimpan jumlah input yang berhasil dibaca

			if sukses == 0 { // memeriksa apakah pengguna memasukkan huruf atau format yang salah
				fmt.Scanln(&buang) // membersihkan sisa input yang salah dari buffer agar tidak terbaca lagi
				fmt.Println("  ❌ Input tidak valid! Harap masukkan angka, bukan huruf.\n")
			} else {
				validFormat = true // mengubah status validasi menjadi benar karena input berhasil dibaca sebagai angka
			}
		}
		nData = nData - 1 // mengubah nomor urut menjadi index array, -1 dilakukan karena nomor tabel dimulai dari 1 sedangkan index array dimulai dari 0

		if nData >= 0 && nData < *n { // memeriksa apakah index yang dipilih masih berada dalam batas data yang tersedia
			for i = nData; i < *n-1; i++ { // menggeser seluruh data setelah data yang dihapus ke posisi kiri, i dimulai dari data yang akan dihapus sampai data terakhir sebelum ujung array
				data[i] = data[i+1] // menyalin data di sebelah kanan ke posisi sekarang, digunakan untuk menutup celah akibat data yang dihapus
			}
			*n = *n - 1 // mengurangi jumlah data sebanyak satu, karena satu data telah berhasil dihapus

			fmt.Println("\n  ╔═══════════════════════════════════════════════════════════════════════╗")
			fmt.Println("  ║                     ✨ Data berhasil dihapus! ✨                      ║")
			fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")
		} else { // dijalankan jika nomor data yang dimasukkan tidak ditemukan
			fmt.Println("\n  ╔═══════════════════════════════════════════════════════════════════════╗")
			fmt.Println("  ║                       ⚠️  Data tidak ditemukan!                        ║")
			fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")
		}
	}

	fmt.Println("Ketik 0 dan tekan [Enter] untuk kembali... ")
	kembali = ""         // mengosongkan variabel kembali sebelum validasi dilakukan
	for kembali != "0" { // mengulang selama pengguna belum memasukkan angka 0
		fmt.Scan(&kembali)  // membaca input pengguna dan menyimpannya ke variabel kembali
		if kembali != "0" { // memeriksa apakah input yang dimasukkan bukan angka 0
			fmt.Println("  ❌ Input tidak valid! Harap ketik 0: ")
		}
	}
}
func cariPerangkat(n int) { // procedure untuk mencari perangkat berdasarkan nama menggunakan Sequential Search, n menyimpan jumlah data yang tersimpan dalam array
	clearScreen()
	var find, kembali string // find digunakan untuk menyimpan nama perangkat yang dicari, kembali digunakan untuk validasi sebelum kembali ke menu
	var i int
	var ketemu bool // menyimpan status apakah data ditemukan atau tidak, bernilai true jika ditemukan dan false jika tidak ditemukan

	fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("  ║                        🔎 CARI NAMA PERANGKAT 🔎                      ║")
	fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝")
	fmt.Println()

	if n == 0 { // memeriksa apakah data masih kosong, n menyimpan jumlah data yang tersimpan
		fmt.Println("  📭 Data perangkat masih kosong!\n")
	} else { // dijalankan jika terdapat data yang dapat dicari
		fmt.Print("  🏷️  Masukkan nama perangkat yang dicari : ")
		find = bacaString() // membaca nama perangkat yang dicari dan menyimpannya ke variabel find
		fmt.Println()

		ketemu = false          // menginisialisasi status pencarian sebagai belum ditemukan
		for i = 0; i < n; i++ { // melakukan Sequential Search dengan memeriksa data satu per satu dari index pertama sampai terakhir
			if data[i].nama == find { // memeriksa apakah nama perangkat pada data ke-i sama dengan nama yang dicari
				if ketemu == false { // memeriksa apakah ini data pertama yang ditemukan, digunakan agar header tabel hanya tampil satu kali
					fmt.Println("  ┌────┬──────────────────────┬─────────┬──────────┬──────────────────────┐")
					fmt.Println("  │ No │ Nama Perangkat       │ Watt    │ Durasi   │ Ruangan              │")
					fmt.Println("  ├────┼──────────────────────┼─────────┼──────────┼──────────────────────┤")
					ketemu = true // mengubah status menjadi ditemukan karena ada data yang cocok
				}
				fmt.Printf("  │ %-2d │ %-20s │ %-5.1f W │ %2dj %02dm  │ %-20s │\n", i+1, data[i].nama, data[i].watt, data[i].jam, data[i].menit, data[i].ruangan)
			}
		}

		if ketemu { // memeriksa apakah minimal ada satu data yang berhasil ditemukan
			fmt.Println("  └────┴──────────────────────┴─────────┴──────────┴──────────────────────┘\n")
			fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
			fmt.Println("  ║                     ✨ Data berhasil ditemukan! ✨                    ║")
			fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")
		} else { // dijalankan jika tidak ada data yang cocok dengan nama yang dicari
			fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
			fmt.Println("  ║                       ⚠️  Data tidak ditemukan!                        ║")
			fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")
		}
	}

	fmt.Println("Ketik 0 dan tekan [Enter] untuk kembali... ")
	kembali = ""         // mengosongkan variabel kembali sebelum proses validasi dimulai
	for kembali != "0" { // mengulang selama pengguna belum memasukkan angka 0
		fmt.Scan(&kembali)  // membaca input pengguna dan menyimpannya ke variabel kembali
		if kembali != "0" { // memeriksa apakah input yang dimasukkan bukan angka 0
			fmt.Println("  ❌ Input tidak valid! Harap ketik 0: ")
		}
	}
}
func cariRuangan(n int) { // procedure untuk mencari perangkat berdasarkan nama ruangan menggunakan Sequential Search, n menyimpan jumlah data dalam array
	clearScreen()
	var find, kembali string // find digunakan untuk menyimpan nama ruangan yang dicari, kembali digunakan untuk validasi kembali ke menu
	var i int
	var ketemu bool // menyimpan status apakah data ditemukan atau tidak, true berarti ditemukan, false berarti tidak ditemukan

	fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("  ║                           🏠 CARI RUANGAN 🏠                          ║")
	fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝")
	fmt.Println()

	if n == 0 { // memeriksa apakah data masih kosong, n menyimpan jumlah data yang tersimpan
		fmt.Println("  📭 Data perangkat masih kosong!\n")
	} else { // dijalankan jika terdapat data yang dapat dicari
		fmt.Print("  🏠 Masukkan nama ruangan yang dicari : ")
		find = bacaString() // membaca nama ruangan yang dicari dan menyimpannya ke variabel find
		fmt.Println()

		ketemu = false          // menginisialisasi status pencarian sebagai belum ditemukan
		for i = 0; i < n; i++ { // melakukan Sequential Search dengan memeriksa setiap data dari index pertama sampai terakhir
			if data[i].ruangan == find { // memeriksa apakah nama ruangan pada data ke-i sama dengan nama ruangan yang dicari

				if ketemu == false { // memeriksa apakah ini data pertama yang ditemukan, digunakan agar header tabel hanya muncul satu kali
					fmt.Println("  ┌────┬──────────────────────┬─────────┬──────────┬──────────────────────┐")
					fmt.Println("  │ No │ Nama Perangkat       │ Watt    │ Durasi   │ Ruangan              │")
					fmt.Println("  ├────┼──────────────────────┼─────────┼──────────┼──────────────────────┤")
				}

				fmt.Printf("  │ %-2d │ %-20s │ %-5.1f W │ %2dj %02dm  │ %-20s │\n", i+1, data[i].nama, data[i].watt, data[i].jam, data[i].menit, data[i].ruangan)
				ketemu = true // mengubah status menjadi ditemukan karena ada data yang cocok
			}
		}

		if ketemu { // memeriksa apakah minimal ada satu data yang berhasil ditemukan
			fmt.Println("  └────┴──────────────────────┴─────────┴──────────┴──────────────────────┘\n")
			fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
			fmt.Println("  ║                     ✨ Data berhasil ditemukan! ✨                    ║")
			fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")
		} else { // dijalankan jika tidak ada data yang cocok dengan ruangan yang dicari
			fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
			fmt.Println("  ║                       ⚠️  Data tidak ditemukan!                        ║")
			fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")
		}
	}

	fmt.Println("Ketik 0 dan tekan [Enter] untuk kembali... ")
	kembali = ""         // mengosongkan variabel kembali sebelum validasi dilakukan
	for kembali != "0" { // mengulang selama pengguna belum memasukkan angka 0
		fmt.Scan(&kembali)  // membaca input pengguna dan menyimpannya ke variabel kembali
		if kembali != "0" { // memeriksa apakah input yang dimasukkan bukan angka 0
			fmt.Println("  ❌ Input tidak valid! Harap ketik 0: ")
		}
	}
}
func searchPerangkat(n int) { // procedure untuk mencari perangkat berdasarkan nama menggunakan Binary Search, n menyimpan jumlah data dalam array
	var low, high, mid, kiri, kanan, i int // low adalah batas kiri pencarian, high adalah batas kanan pencarian, mid adalah posisi tengah, kiri dan kanan digunakan untuk mencari data duplikat, i digunakan untuk perulangan menampilkan data
	var find, kembali string               // find digunakan untuk menyimpan nama yang dicari, kembali digunakan untuk validasi kembali ke menu
	var ketemu bool                        // menyimpan status apakah data berhasil ditemukan atau tidak

	clearScreen()

	fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("  ║                      🔎 PENCARIAN NAMA PERANGKAT 🔎                   ║")
	fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝")
	fmt.Println()

	if n == 0 { // memeriksa apakah data masih kosong, n menyimpan jumlah data yang tersimpan
		fmt.Println("  📭 Data perangkat masih kosong!\n")
	} else {
		urutPerangkatASC(n) // mengurutkan data berdasarkan nama perangkat secara ascending, Binary Search membutuhkan data yang sudah terurut

		fmt.Print("  🏷️  Masukkan nama perangkat yang dicari : ")
		find = bacaString() // membaca nama perangkat yang ingin dicari dan menyimpannya ke variabel find
		fmt.Println()

		low = 0        // menentukan batas kiri pencarian dimulai dari index pertama array
		high = n - 1   // menentukan batas kanan pencarian dimulai dari index terakhir data yang terisi
		ketemu = false // menginisialisasi status pencarian sebagai belum ditemukan

		for low <= high && ketemu == false { // melakukan pencarian selama area pencarian masih ada dan data belum ditemukan
			mid = (low + high) / 2 // mencari index tengah dari area pencarian saat ini, mid digunakan sebagai titik pembanding

			if data[mid].nama == find { // memeriksa apakah nama pada posisi tengah sama dengan nama yang dicari
				ketemu = true // mengubah status menjadi ditemukan karena data yang dicari berada di posisi tengah
			} else if find < data[mid].nama { // memeriksa apakah data yang dicari berada di sebelah kiri posisi tengah secara alfabet
				high = mid - 1 // menggeser batas kanan ke kiri karena data yang dicari berada sebelum posisi tengah
			} else { // dijalankan jika data yang dicari berada di sebelah kanan posisi tengah
				low = mid + 1 // menggeser batas kiri ke kanan karena data yang dicari berada setelah posisi tengah
			}
		}

		if ketemu == true { // memeriksa apakah data berhasil ditemukan
			kiri = mid                                  // menyimpan posisi data yang ditemukan sebagai titik awal pencarian ke kiri
			for kiri > 0 && data[kiri-1].nama == find { // bergerak ke kiri selama masih ditemukan nama yang sama
				kiri = kiri - 1 // menggeser posisi satu langkah ke kiri untuk mencari duplikat sebelumnya
			}

			kanan = mid                                     // menyimpan posisi data yang ditemukan sebagai titik awal pencarian ke kanan
			for kanan < n-1 && data[kanan+1].nama == find { // bergerak ke kanan selama masih ditemukan nama yang sama
				kanan = kanan + 1 // menggeser posisi satu langkah ke kanan untuk mencari duplikat berikutnya
			}

			fmt.Println("  ┌────┬──────────────────────┬─────────┬──────────┬──────────────────────┐")
			fmt.Println("  │ No │ Nama Perangkat       │ Watt    │ Durasi   │ Ruangan              │")
			fmt.Println("  ├────┼──────────────────────┼─────────┼──────────┼──────────────────────┤")

			for i = kiri; i <= kanan; i++ { // menampilkan seluruh data yang memiliki nama sama mulai dari posisi paling kiri sampai paling kanan
				fmt.Printf("  │ %-2d │ %-20s │ %-5.1f W │ %2dj %02dm  │ %-20s │\n", i+1, data[i].nama, data[i].watt, data[i].jam, data[i].menit, data[i].ruangan)
			}

			fmt.Println("  └────┴──────────────────────┴─────────┴──────────┴──────────────────────┘\n")

			fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
			fmt.Println("  ║                     ✨ Data berhasil ditemukan! ✨                    ║")
			fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")
		} else { // dijalankan jika data tidak ditemukan setelah proses Binary Search selesai
			fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
			fmt.Println("  ║                       ⚠️  Data tidak ditemukan!                        ║")
			fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")
		}
	}

	fmt.Println("Ketik 0 dan tekan [Enter] untuk kembali... ")
	kembali = ""         // mengosongkan variabel kembali sebelum validasi dilakukan
	for kembali != "0" { // mengulang selama pengguna belum memasukkan angka 0
		fmt.Scan(&kembali) // membaca input pengguna dan menyimpannya ke variabel kembali
		if kembali != "0" {
			fmt.Println("  ❌ Input tidak valid! Harap ketik 0: ")
		}
	}
}
func urutRuanganASC(n int) { // procedure untuk mengurutkan data berdasarkan nama ruangan secara ascending menggunakan Selection Sort, n menyimpan jumlah data yang akan diurutkan
	var i, j, idxMin int // i digunakan untuk menentukan posisi yang sedang diisi, j digunakan untuk mencari data berikutnya, idxMin menyimpan index ruangan dengan nilai terkecil yang ditemukan
	var temp Perangkat   // temp digunakan sebagai variabel sementara saat proses pertukaran data

	for i = 0; i < n-1; i++ { // melakukan pengurutan dari data pertama sampai sebelum data terakhir, i menunjukkan posisi yang akan diisi oleh data terkecil
		idxMin = i                  // menganggap data pada posisi i sebagai ruangan terkecil sementara, idxMin menyimpan index ruangan terkecil saat ini
		for j = i + 1; j < n; j++ { // memeriksa seluruh data setelah posisi i untuk mencari ruangan yang lebih kecil secara alfabet
			if data[j].ruangan < data[idxMin].ruangan { // membandingkan nama ruangan pada index j dengan ruangan terkecil sementara, tanda < digunakan untuk urutan ascending berdasarkan alfabet
				idxMin = j // menyimpan posisi ruangan yang lebih kecil sebagai kandidat ruangan terkecil yang baru
			}
		}
		temp = data[i]         // menyimpan data pada posisi i ke variabel sementara agar tidak hilang saat ditukar
		data[i] = data[idxMin] // memindahkan data dengan ruangan terkecil ke posisi i
		data[idxMin] = temp    // memindahkan data lama dari posisi i ke posisi idxMin sehingga proses pertukaran selesai
	}
}
func searchRuangan(n int) { // procedure untuk mencari data berdasarkan nama ruangan menggunakan Binary Search, n menyimpan jumlah data dalam array
	var low, high, mid, kiri, kanan, i int // low adalah batas kiri pencarian, high adalah batas kanan pencarian, mid adalah posisi tengah, kiri dan kanan digunakan untuk mencari data duplikat, i digunakan untuk menampilkan hasil pencarian
	var find, kembali string               // find digunakan untuk menyimpan nama ruangan yang dicari, kembali digunakan untuk validasi kembali ke menu
	var ketemu bool                        // menyimpan status apakah data ditemukan atau tidak, true berarti ditemukan dan false berarti tidak ditemukan

	clearScreen()

	fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("  ║                           🏠 CARI RUANGAN 🏠                          ║")
	fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝")
	fmt.Println()

	if n == 0 { // memeriksa apakah data masih kosong, n menyimpan jumlah data yang tersedia
		fmt.Println("  📭 Data perangkat masih kosong!\n")
	} else {
		urutRuanganASC(n) // mengurutkan data berdasarkan nama ruangan secara ascending, Binary Search membutuhkan data yang sudah terurut

		fmt.Print("  🏠 Masukkan nama ruangan yang dicari : ")
		find = bacaString() // membaca nama ruangan yang dicari dan menyimpannya ke variabel find
		fmt.Println()

		low = 0        // menentukan batas kiri pencarian dimulai dari index pertama array
		high = n - 1   // menentukan batas kanan pencarian dimulai dari index terakhir data yang terisi
		ketemu = false // menginisialisasi status pencarian sebagai belum ditemukan

		for low <= high && ketemu == false { // melakukan Binary Search selama area pencarian masih ada dan data belum ditemukan
			mid = (low + high) / 2 // menghitung posisi tengah dari area pencarian saat ini, mid digunakan sebagai titik pembanding

			if data[mid].ruangan == find { // memeriksa apakah ruangan pada posisi tengah sama dengan ruangan yang dicari
				ketemu = true // mengubah status menjadi ditemukan karena data yang dicari berada pada posisi tengah
			} else if find < data[mid].ruangan { // memeriksa apakah ruangan yang dicari berada di sebelah kiri posisi tengah secara alfabet
				high = mid - 1 // menggeser batas kanan ke kiri karena data yang dicari berada sebelum posisi tengah
			} else { // dijalankan jika data yang dicari berada di sebelah kanan posisi tengah
				low = mid + 1 // menggeser batas kiri ke kanan karena data yang dicari berada setelah posisi tengah
			}
		}

		if ketemu == true { // memeriksa apakah data berhasil ditemukan
			kiri = mid                                     // menyimpan posisi data yang ditemukan sebagai titik awal pencarian ke kiri
			for kiri > 0 && data[kiri-1].ruangan == find { // bergerak ke kiri selama masih ditemukan ruangan yang sama
				kiri = kiri - 1 // menggeser posisi satu langkah ke kiri untuk mencari duplikat sebelumnya
			}
			kanan = mid                                        // menyimpan posisi data yang ditemukan sebagai titik awal pencarian ke kanan
			for kanan < n-1 && data[kanan+1].ruangan == find { // bergerak ke kanan selama masih ditemukan ruangan yang sama
				kanan = kanan + 1 // menggeser posisi satu langkah ke kanan untuk mencari duplikat berikutnya
			}

			fmt.Println("  ┌────┬──────────────────────┬─────────┬──────────┬──────────────────────┐")
			fmt.Println("  │ No │ Nama Perangkat       │ Watt    │ Durasi   │ Ruangan              │")
			fmt.Println("  ├────┼──────────────────────┼─────────┼──────────┼──────────────────────┤")

			for i = kiri; i <= kanan; i++ { // menampilkan seluruh data yang memiliki nama ruangan yang sama mulai dari posisi paling kiri sampai paling kanan
				fmt.Printf("  │ %-2d │ %-20s │ %-5.1f W │ %2dj %02dm  │ %-20s │\n", i+1, data[i].nama, data[i].watt, data[i].jam, data[i].menit, data[i].ruangan)
			}

			fmt.Println("  └────┴──────────────────────┴─────────┴──────────┴──────────────────────┘\n")

			fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
			fmt.Println("  ║                     ✨ Data berhasil ditemukan! ✨                    ║")
			fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")
		} else { // dijalankan jika data tidak ditemukan setelah proses Binary Search selesai
			fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
			fmt.Println("  ║                       ⚠️  Data tidak ditemukan!                        ║")
			fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")
		}
	}

	fmt.Println("Ketik 0 dan tekan [Enter] untuk kembali... ")
	kembali = ""         // mengosongkan variabel kembali sebelum validasi dilakukan
	for kembali != "0" { // mengulang selama pengguna belum memasukkan angka 0
		fmt.Scan(&kembali) // membaca input pengguna dan menyimpannya ke variabel kembali
		if kembali != "0" {
			fmt.Println("  ❌ Input tidak valid! Harap ketik 0: ")
		}
	}
}
func urutWattASC(n int) { // procedure untuk mengurutkan data berdasarkan watt dari yang terkecil ke terbesar menggunakan Selection Sort
	var i, j, idxMin int      // i digunakan untuk menentukan posisi yang sedang diisi, j digunakan untuk mencari data berikutnya, idxMin menyimpan index watt terkecil yang ditemukan
	var temp Perangkat        // temp digunakan sebagai variabel sementara saat proses pertukaran data agar data tidak hilang
	for i = 0; i < n-1; i++ { // melakukan pengurutan dari data pertama sampai sebelum data terakhir, i menunjukkan posisi yang akan diisi oleh watt terkecil berikutnya
		idxMin = i                  // menganggap data pada posisi i sebagai watt terkecil sementara, idxMin menyimpan posisi watt terkecil saat ini
		for j = i + 1; j < n; j++ { // memeriksa seluruh data setelah posisi i untuk mencari watt yang lebih kecil
			if data[j].watt < data[idxMin].watt { // membandingkan watt pada index j dengan watt terkecil sementara, tanda < digunakan untuk mencari nilai watt yang lebih kecil sehingga menghasilkan urutan ascending
				idxMin = j // menyimpan posisi watt terkecil yang baru ditemukan
			}
		}
		temp = data[i]         // menyimpan data pada posisi i ke variabel sementara agar tidak hilang saat proses pertukaran
		data[i] = data[idxMin] // memindahkan data dengan watt terkecil ke posisi i
		data[idxMin] = temp    // memindahkan data lama dari posisi i ke posisi idxMin sehingga proses pertukaran selesai
	}
}
func urutWattDESC(n int) { // procedure untuk mengurutkan data berdasarkan watt dari yang terbesar ke terkecil menggunakan Selection Sort, n menyimpan jumlah data yang akan diurutkan
	var i, j, idxMin int      // i digunakan untuk menentukan posisi yang sedang diisi, j digunakan untuk mencari data berikutnya, idxMin menyimpan index watt terbesar yang ditemukan
	var temp Perangkat        // temp digunakan sebagai variabel sementara saat proses pertukaran data agar data tidak hilang
	for i = 0; i < n-1; i++ { // melakukan pengurutan dari data pertama sampai sebelum data terakhir, i menunjukkan posisi yang akan diisi oleh watt terbesar berikutnya
		idxMin = i                  // menganggap data pada posisi i sebagai watt terbesar sementara, idxMin menyimpan posisi watt terbesar saat ini
		for j = i + 1; j < n; j++ { // memeriksa seluruh data setelah posisi i untuk mencari watt yang lebih besar
			if data[j].watt > data[idxMin].watt { // membandingkan watt pada index j dengan watt terbesar sementara, tanda > digunakan untuk mencari nilai watt yang lebih besar sehingga menghasilkan urutan descending
				idxMin = j // menyimpan posisi watt terbesar yang baru ditemukan
			}
		}
		temp = data[i]         // menyimpan data pada posisi i ke variabel sementara agar tidak hilang saat proses pertukaran
		data[i] = data[idxMin] // memindahkan data dengan watt terbesar ke posisi i
		data[idxMin] = temp    // memindahkan data lama dari posisi i ke posisi idxMin sehingga proses pertukaran selesai
	}
}
func urutPerangkatASC(n int) { // procedure untuk mengurutkan data berdasarkan nama perangkat dari A ke Z menggunakan Selection Sort, n menyimpan jumlah data yang akan diurutkan
	var i, j, idxMin int      // i digunakan untuk menentukan posisi yang sedang diisi, j digunakan untuk mencari data berikutnya, idxMin menyimpan index nama perangkat yang paling kecil secara alfabet
	var temp Perangkat        // temp digunakan sebagai variabel sementara saat proses pertukaran data
	for i = 0; i < n-1; i++ { // melakukan pengurutan dari data pertama sampai sebelum data terakhir, i menunjukkan posisi yang akan diisi oleh nama perangkat paling kecil berikutnya
		idxMin = i                  // menganggap data pada posisi i sebagai nama perangkat paling kecil sementara
		for j = i + 1; j < n; j++ { // memeriksa seluruh data setelah posisi i untuk mencari nama perangkat yang lebih kecil secara alfabet
			if data[j].nama < data[idxMin].nama { // membandingkan nama perangkat pada index j dengan nama perangkat terkecil sementara, tanda < digunakan untuk urutan ascending berdasarkan alfabet
				idxMin = j // menyimpan posisi nama perangkat yang lebih kecil yang baru ditemukan

			}
		}
		temp = data[i]         // menyimpan data pada posisi i ke variabel sementara agar tidak hilang saat ditukar
		data[i] = data[idxMin] // memindahkan data dengan nama perangkat paling kecil ke posisi i
		data[idxMin] = temp    // memindahkan data lama dari posisi i ke posisi idxMin sehingga proses pertukaran selesai
	}
}
func urutPerangkatDESC(n int) { // procedure untuk mengurutkan data berdasarkan nama perangkat dari Z ke A menggunakan Selection Sort, n menyimpan jumlah data yang akan diurutkan
	var i, j, idxMin int      // i digunakan untuk menentukan posisi yang sedang diisi, j digunakan untuk mencari data berikutnya, idxMin menyimpan index nama perangkat terbesar secara alfabet
	var temp Perangkat        // temp digunakan sebagai variabel sementara saat proses pertukaran data
	for i = 0; i < n-1; i++ { // melakukan pengurutan dari data pertama sampai sebelum data terakhir, i menunjukkan posisi yang akan diisi oleh nama perangkat terbesar berikutnya
		idxMin = i                  // menganggap data pada posisi i sebagai nama perangkat terbesar sementara
		for j = i + 1; j < n; j++ { // memeriksa seluruh data setelah posisi i untuk mencari nama perangkat yang lebih besar secara alfabet
			if data[j].nama > data[idxMin].nama { // membandingkan nama perangkat pada index j dengan nama perangkat terbesar sementara, tanda > digunakan untuk urutan descending berdasarkan alfabet
				idxMin = j // menyimpan posisi nama perangkat terbesar yang baru ditemukan
			}
		}
		temp = data[i]         // menyimpan data pada posisi i ke variabel sementara agar tidak hilang saat ditukar
		data[i] = data[idxMin] // memindahkan data dengan nama perangkat terbesar ke posisi i
		data[idxMin] = temp    // memindahkan data lama dari posisi i ke posisi idxMin sehingga proses pertukaran selesai
	}
}
func insertionWattASC(n int) { // procedure untuk mengurutkan data berdasarkan watt dari yang terkecil ke terbesar menggunakan Insertion Sort, n menyimpan jumlah data yang akan diurutkan
	var pass, i int    // pass digunakan untuk menentukan data yang sedang disisipkan, i digunakan untuk mencari posisi penyisipan yang tepat
	var temp Perangkat // temp digunakan untuk menyimpan data yang sedang dipindahkan sementara agar tidak hilang
	pass = 1           // memulai proses dari data kedua, karena data pertama dianggap sudah terurut
	for pass <= n-1 {  // melakukan proses penyisipan sampai data terakhir
		i = pass                                  // i dimulai dari posisi data yang sedang akan disisipkan
		temp = data[pass]                         // menyimpan data yang akan disisipkan ke variabel sementara, data[pass] adalah data yang sedang diproses
		for i > 0 && temp.watt < data[i-1].watt { // memeriksa apakah masih ada data di sebelah kiri dan apakah watt pada temp lebih kecil dari watt sebelumnya, jika iya maka data harus digeser
			data[i] = data[i-1] // menggeser data sebelumnya satu langkah ke kanan untuk memberi ruang bagi temp
			i = i - 1           // berpindah satu posisi ke kiri untuk melanjutkan pencarian posisi yang tepat
		}
		data[i] = temp  // menyisipkan temp ke posisi yang sudah sesuai setelah proses penggeseran selesai
		pass = pass + 1 // berpindah ke data berikutnya yang akan disisipkan
	}
}
func insertionWattDSC(n int) { // procedure untuk mengurutkan data berdasarkan watt dari yang terbesar ke terkecil menggunakan Insertion Sort
	var pass, i int    // pass digunakan untuk menentukan data yang sedang disisipkan, i digunakan untuk mencari posisi penyisipan yang tepat
	var temp Perangkat // temp digunakan untuk menyimpan data yang sedang dipindahkan sementara
	pass = 1           // memulai proses dari data kedua karena data pertama dianggap sudah terurut
	for pass <= n-1 {  // melakukan proses penyisipan sampai data terakhir
		i = pass                                  // i dimulai dari posisi data yang sedang akan disisipkan
		temp = data[pass]                         // menyimpan data yang akan disisipkan ke variabel sementara
		for i > 0 && temp.watt > data[i-1].watt { // memeriksa apakah watt pada temp lebih besar dari watt sebelumnya, tanda > digunakan untuk menghasilkan urutan descending
			data[i] = data[i-1] // menggeser data sebelumnya satu langkah ke kanan
			i = i - 1           // berpindah satu posisi ke kiri untuk mencari posisi yang sesuai
		}
		data[i] = temp  // menyisipkan temp ke posisi yang benar setelah proses penggeseran selesai
		pass = pass + 1 // berpindah ke data berikutnya yang akan disisipkan
	}
}
func insertionPerangkatASC(n int) { // procedure untuk mengurutkan data berdasarkan nama perangkat dari A ke Z menggunakan Insertion Sort, n menyimpan jumlah data yang akan diurutkan
	var pass, i int    // pass digunakan untuk menentukan data yang sedang disisipkan, i digunakan untuk mencari posisi penyisipan yang tepat
	var temp Perangkat // temp digunakan untuk menyimpan data yang sedang dipindahkan sementara agar tidak hilang
	pass = 1           // memulai proses dari data kedua karena data pertama dianggap sudah terurut
	for pass <= n-1 {  // melakukan proses penyisipan sampai data terakhir
		i = pass                                  // i dimulai dari posisi data yang sedang akan disisipkan
		temp = data[pass]                         // menyimpan data yang akan disisipkan ke variabel sementara, data[pass] adalah data yang sedang diproses
		for i > 0 && temp.nama < data[i-1].nama { // membandingkan nama perangkat pada temp dengan nama perangkat sebelumnya, tanda < digunakan untuk menghasilkan urutan A sampai Z
			data[i] = data[i-1] // menggeser data sebelumnya satu langkah ke kanan untuk memberi ruang bagi temp
			i = i - 1           // berpindah satu posisi ke kiri untuk mencari posisi penyisipan yang tepat
		}
		data[i] = temp  // menyisipkan temp ke posisi yang sesuai setelah proses penggeseran selesai
		pass = pass + 1 // berpindah ke data berikutnya yang akan disisipkan
	}
}
func insertionPerangkatDSC(n int) { // procedure untuk mengurutkan data berdasarkan nama perangkat dari Z ke A menggunakan Insertion Sort
	var pass, i int    // pass digunakan untuk menentukan data yang sedang disisipkan, i digunakan untuk mencari posisi penyisipan yang tepat
	var temp Perangkat // temp digunakan untuk menyimpan data yang sedang dipindahkan sementara
	pass = 1           // memulai proses dari data kedua karena data pertama dianggap sudah terurut
	for pass <= n-1 {  // melakukan proses penyisipan sampai data terakhir
		i = pass                                  // i dimulai dari posisi data yang sedang akan disisipkan
		temp = data[pass]                         // menyimpan data yang akan disisipkan ke variabel sementara
		for i > 0 && temp.nama > data[i-1].nama { // membandingkan nama perangkat pada temp dengan nama perangkat sebelumnya, tanda > digunakan untuk menghasilkan urutan Z sampai A
			data[i] = data[i-1] // menggeser data sebelumnya satu langkah ke kanan
			i = i - 1           // berpindah satu posisi ke kiri untuk mencari posisi yang sesuai
		}
		data[i] = temp  // menyisipkan temp ke posisi yang benar setelah proses penggeseran selesai
		pass = pass + 1 // berpindah ke data berikutnya yang akan disisipkan
	}
}
func statistikData(n int) { // procedure untuk menghitung dan menampilkan statistik penggunaan daya seluruh perangkat, n menyimpan jumlah data yang tersimpan dalam array
	var i int
	var totalDaya, durasiTotal float64 // totalDaya digunakan untuk menyimpan total penggunaan daya seluruh perangkat, durasiTotal digunakan untuk menyimpan durasi penggunaan dalam satuan jam desimal
	var kembali string                 // digunakan untuk menyimpan input pengguna sebelum kembali ke menu

	clearScreen()

	fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("  ║                      📊 STATISTIK DAYA HARIAN 📊                      ║")
	fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝")
	fmt.Println()

	if n == 0 { // memeriksa apakah data masih kosong, n menyimpan jumlah data yang tersedia
		fmt.Println("  📭 Data perangkat masih kosong!\n")
	} else { // dijalankan jika terdapat data yang dapat dihitung statistiknya
		totalDaya = 0.0         // menginisialisasi total penggunaan daya menjadi 0 sebelum proses perhitungan dimulai
		for i = 0; i < n; i++ { // mengulang seluruh data dari index pertama sampai terakhir untuk menghitung total penggunaan daya
			durasiTotal = float64(data[i].jam) + (float64(data[i].menit) / 60.0) // mengubah durasi penggunaan menjadi satuan jam desimal, jam tetap dalam satuan jam sedangkan menit dibagi 60 agar menjadi pecahan jam
			// mengubah nilai jam dari integer menjadi float64 agar dapat dihitung bersama bilangan desimal
			// mengubah nilai menit dari integer menjadi float64 agar hasil pembagian dapat menghasilkan bilangan desimal
			// mengubah menit menjadi pecahan jam, 60.0 digunakan agar pembagian menghasilkan nilai desimal
			totalDaya += (data[i].watt * durasiTotal) // menambahkan penggunaan daya perangkat ke totalDaya, penggunaan daya diperoleh dari watt dikali lama penggunaan dalam jam
		}

		fmt.Println("  ┌───────────────────────────────────────────────────────────────────────┐")
		fmt.Printf("  │ ⚡ Total Perangkat Aktif      : %-37d │\n", n)           // menampilkan jumlah perangkat yang tersimpan, n menunjukkan banyaknya data perangkat yang aktif
		fmt.Printf("  │ 🔋 Total Penggunaan Daya (Wh) : %-37.1f │\n", totalDaya) // menampilkan total penggunaan daya seluruh perangkat dalam satuan Wh, .1f digunakan agar hanya menampilkan satu angka di belakang koma
		fmt.Println("  └───────────────────────────────────────────────────────────────────────┘\n")
	}

	fmt.Println("  Ketik 0 dan tekan [Enter] untuk kembali... ")
	kembali = ""         // mengosongkan variabel kembali sebelum validasi input dimulai
	for kembali != "0" { // mengulang selama pengguna belum memasukkan angka 0
		fmt.Scan(&kembali)  // membaca input pengguna dan menyimpannya ke variabel kembali
		if kembali != "0" { // memeriksa apakah input yang dimasukkan bukan angka 0
			fmt.Println("  ❌ Input tidak valid! Harap ketik 0: ")
		}
	}
}
func perangkatTerboros(n int) { // procedure untuk mencari dan menampilkan perangkat dengan konsumsi energi terbesar, n menyimpan jumlah data yang tersimpan dalam array
	var i, idxMax int                           // i digunakan sebagai pencacah untuk mengakses setiap data, idxMax digunakan untuk menyimpan index perangkat dengan konsumsi energi terbesar
	var totalEnergi, durasiI, durasiMax float64 // totalEnergi digunakan untuk menyimpan total energi perangkat terboros, durasiI menyimpan durasi perangkat yang sedang diperiksa, durasiMax menyimpan durasi perangkat yang sementara dianggap paling boros
	var kembali string                          // digunakan untuk menyimpan input pengguna sebelum kembali ke menu

	clearScreen()

	fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("  ║                       🔥 PERANGKAT TERBOROS 🔥                        ║")
	fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝")
	fmt.Println()

	if n == 0 { // memeriksa apakah data masih kosong, n menyimpan jumlah data yang tersedia
		fmt.Println("  📭 Data perangkat masih kosong!\n")
	} else { // dijalankan jika terdapat data yang dapat dicari perangkat terborosnya
		idxMax = 0              // menganggap data pertama sebagai perangkat paling boros sementara, idxMax menyimpan posisi perangkat dengan konsumsi energi terbesar saat ini
		for i = 0; i < n; i++ { // memeriksa seluruh data satu per satu untuk mencari perangkat dengan konsumsi energi terbesar
			durasiI = float64(data[i].jam) + (float64(data[i].menit) / 60.0)             // mengubah durasi perangkat yang sedang diperiksa menjadi jam desimal agar dapat digunakan dalam perhitungan energi
			durasiMax = float64(data[idxMax].jam) + (float64(data[idxMax].menit) / 60.0) // mengubah durasi perangkat yang sementara dianggap paling boros menjadi jam desimal

			if (data[i].watt * durasiI) > (data[idxMax].watt * durasiMax) { // membandingkan konsumsi energi perangkat saat ini dengan perangkat paling boros sementara, energi dihitung menggunakan rumus watt × durasi(jam)
				idxMax = i // menyimpan index perangkat yang memiliki konsumsi energi lebih besar sebagai perangkat terboros yang baru
			}
		}

		durasiMax = float64(data[idxMax].jam) + (float64(data[idxMax].menit) / 60.0) // menghitung kembali durasi perangkat terboros dalam jam desimal untuk digunakan pada perhitungan energi akhir
		totalEnergi = data[idxMax].watt * durasiMax                                  // menghitung total energi perangkat terboros menggunakan rumus watt × durasi(jam)

		fmt.Println("  ┌────┬──────────────────────┬─────────┬──────────┬──────────────────────┐")
		fmt.Println("  │ No │ Nama Perangkat       │ Watt    │ Durasi   │ Ruangan              │")
		fmt.Println("  ├────┼──────────────────────┼─────────┼──────────┼──────────────────────┤")

		fmt.Printf("  │ %-2d │ %-20s │ %-5.1f W │ %2dj %02dm  │ %-20s │\n", 1, data[idxMax].nama, data[idxMax].watt, data[idxMax].jam, data[idxMax].menit, data[idxMax].ruangan) // menampilkan data perangkat yang memiliki konsumsi energi terbesar

		fmt.Println("  └────┴──────────────────────┴─────────┴──────────┴──────────────────────┘\n")

		fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
		fmt.Printf("  ║  💥 Total Konsumsi Energi (Wh) : %-36.1f ║\n", totalEnergi) // menampilkan total konsumsi energi perangkat terboros dalam satuan Wh
		fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")
	}

	fmt.Println("Ketik 0 dan tekan [Enter] untuk kembali... ")
	kembali = ""         // mengosongkan variabel kembali sebelum validasi input dimulai
	for kembali != "0" { // mengulang selama pengguna belum memasukkan angka 0
		fmt.Scan(&kembali)  // membaca input pengguna dan menyimpannya ke variabel kembali
		if kembali != "0" { // memeriksa apakah input yang dimasukkan bukan angka 0
			fmt.Print("  ❌ Input tidak valid! Harap ketik 0: ")
		}
	}
}
func main() {
	var menu, n, pilih, sukses int // menu menyimpan pilihan menu utama, n menyimpan jumlah data, pilih menyimpan pilihan submenu, sukses menyimpan jumlah input yang berhasil dibaca oleh fmt.Scan
	var buang, kembali string      // buang digunakan untuk membersihkan input yang salah dari buffer, kembali digunakan untuk validasi sebelum kembali ke menu
	var validFormat bool           // digunakan untuk mengecek apakah format input pengguna sudah benar atau belum
	generateDummyData(&n)
	for menu != 11 { // mengulang menu utama selama pengguna belum memilih menu 11 (keluar)
		clearScreen()
		fmt.Println("  ╔═════════════════════════════════════════╗")
		fmt.Println("  ║            ⚡ POWERLOG SYSTEM ⚡        ║")
		fmt.Println("  ╠═════════════════════════════════════════╣")
		fmt.Println("  ║                                         ║")
		fmt.Println("  ║  [ 1] ➕ Tambah Data                    ║")
		fmt.Println("  ║  [ 2] 📋 Tampilkan Data                 ║")
		fmt.Println("  ║  [ 3] ✏️  Ubah Data                      ║")
		fmt.Println("  ║  [ 4] ❌ Hapus Data                     ║")
		fmt.Println("  ║  [ 5] 🔎 Cari Nama Perangkat            ║")
		fmt.Println("  ║  [ 6] 🏠 Cari Ruangan                   ║")
		fmt.Println("  ║  [ 7] 🔤 Urutkan Nama                   ║")
		fmt.Println("  ║  [ 8] ⚡ Urutkan Watt                   ║")
		fmt.Println("  ║  [ 9] 📊 Statistik Daya                 ║")
		fmt.Println("  ║  [10] 🔥 Perangkat Terboros             ║")
		fmt.Println("  ║  [11] 🚪 Keluar                         ║")
		fmt.Println("  ║                                         ║")
		fmt.Println("  ╚═════════════════════════════════════════╝")

		validFormat = false
		for validFormat == false {
			fmt.Print("\n  👉 Masukkan pilihan menu (1-11): ")
			sukses, _ = fmt.Scan(&menu) // membaca pilihan menu yang dimasukkan pengguna

			if sukses == 0 { // memeriksa apakah input gagal dibaca sebagai angka
				fmt.Scanln(&buang) // membersihkan sisa input yang salah agar tidak terbaca kembali pada perulangan berikutnya
				fmt.Println("  ❌ Input tidak valid! Harap masukkan angka.")
			} else {
				validFormat = true
			}
		}

		if menu == 1 {
			tambahData(&n)
		} else if menu == 2 {
			tampilData(n)
		} else if menu == 3 {
			updateData(n)
		} else if menu == 4 {
			hapusData(&n)
		} else if menu == 5 {
			clearScreen()
			fmt.Println("  ╔═════════════════════════════════════════╗")
			fmt.Println("  ║       🔎 PILIH METODE PENCARIAN 🔎      ║")
			fmt.Println("  ╠═════════════════════════════════════════╣")
			fmt.Println("  ║  [1] Sequential Search                  ║")
			fmt.Println("  ║  [2] Binary Search                      ║")
			fmt.Println("  ╚═════════════════════════════════════════╝")

			validFormat = false
			for validFormat == false {
				fmt.Print("\n  👉 Pilih metode pencarian: ")
				sukses, _ = fmt.Scan(&pilih) // membaca pilihan metode yang dimasukkan pengguna
				if sukses == 0 {
					fmt.Scanln(&buang)
					fmt.Println("  ❌ Input tidak valid! Harap masukkan angka.")
				} else {
					validFormat = true
				}
			}

			if pilih == 1 {
				cariPerangkat(n)
			} else if pilih == 2 {
				searchPerangkat(n)
			}
		} else if menu == 6 {
			clearScreen()
			fmt.Println("  ╔═════════════════════════════════════════╗")
			fmt.Println("  ║       🏠 PILIH METODE PENCARIAN 🏠      ║")
			fmt.Println("  ╠═════════════════════════════════════════╣")
			fmt.Println("  ║  [1] Sequential Search                  ║")
			fmt.Println("  ║  [2] Binary Search                      ║")
			fmt.Println("  ╚═════════════════════════════════════════╝")

			validFormat = false
			for validFormat == false {
				fmt.Print("\n  👉 Pilih metode pencarian: ")
				sukses, _ = fmt.Scan(&pilih)
				if sukses == 0 {
					fmt.Scanln(&buang)
					fmt.Println("  ❌ Input tidak valid! Harap masukkan angka.")
				} else {
					validFormat = true
				}
			}

			if pilih == 1 {
				cariRuangan(n)
			} else if pilih == 2 {
				searchRuangan(n)
			}
		} else if menu == 7 {
			clearScreen()
			fmt.Println("  ╔═════════════════════════════════════════╗")
			fmt.Println("  ║         📑 PILIH METODE SORTING 📑      ║")
			fmt.Println("  ╠═════════════════════════════════════════╣")
			fmt.Println("  ║  [1] Selection Sort (Ascending)         ║")
			fmt.Println("  ║  [2] Selection Sort (Descending)        ║")
			fmt.Println("  ║  [3] Insertion Sort (Ascending)         ║")
			fmt.Println("  ║  [4] Insertion Sort (Descending)        ║")
			fmt.Println("  ╚═════════════════════════════════════════╝")

			validFormat = false
			for validFormat == false {
				fmt.Print("\n  👉 Pilih metode sorting: ")
				sukses, _ = fmt.Scan(&pilih)
				if sukses == 0 {
					fmt.Scanln(&buang)
					fmt.Println("  ❌ Input tidak valid! Harap masukkan angka.")
				} else {
					validFormat = true
				}
			}

			if pilih == 1 {
				urutPerangkatASC(n)
				tampilData(n)
			} else if pilih == 2 {
				urutPerangkatDESC(n)
				tampilData(n)
			} else if pilih == 3 {
				insertionPerangkatASC(n)
				tampilData(n)
			} else if pilih == 4 {
				insertionPerangkatDSC(n)
				tampilData(n)
			}
		} else if menu == 8 {
			clearScreen()
			fmt.Println("  ╔═════════════════════════════════════════╗")
			fmt.Println("  ║         ⚡ PILIH METODE SORTING ⚡      ║")
			fmt.Println("  ╠═════════════════════════════════════════╣")
			fmt.Println("  ║  [1] Selection Sort (Ascending)         ║")
			fmt.Println("  ║  [2] Selection Sort (Descending)        ║")
			fmt.Println("  ║  [3] Insertion Sort (Ascending)         ║")
			fmt.Println("  ║  [4] Insertion Sort (Descending)        ║")
			fmt.Println("  ╚═════════════════════════════════════════╝")

			validFormat = false
			for validFormat == false {
				fmt.Print("\n  👉 Pilih metode sorting: ")
				sukses, _ = fmt.Scan(&pilih)
				if sukses == 0 {
					fmt.Scanln(&buang)
					fmt.Println("  ❌ Input tidak valid! Harap masukkan angka.")
				} else {
					validFormat = true
				}
			}

			if pilih == 1 {
				urutWattASC(n)
				tampilData(n)
			} else if pilih == 2 {
				urutWattDESC(n)
				tampilData(n)
			} else if pilih == 3 {
				insertionWattASC(n)
				tampilData(n)
			} else if pilih == 4 {
				insertionWattDSC(n)
				tampilData(n)
			}
		} else if menu == 9 {
			statistikData(n)
		} else if menu == 10 {
			perangkatTerboros(n)
		} else if menu != 11 {
			fmt.Println("\n  ⚠️  Menu tidak tersedia!")
			fmt.Println("Ketik 0 dan tekan [Enter] untuk kembali... ")
			kembali = ""         // mengosongkan variabel kembali sebelum validasi input dimulai
			for kembali != "0" { // mengulang selama pengguna belum memasukkan angka 0
				fmt.Scan(&kembali)  // membaca input pengguna dan menyimpannya ke variabel kembali
				if kembali != "0" { // memeriksa apakah input yang dimasukkan bukan angka 0
					fmt.Print("  ❌ Input tidak valid! Harap ketik 0: ")
				}
			}
		}
	}

}
func generateDummyData(n *int) {
	var i int
	var dummy = [20]Perangkat{
		{nama: "AC 1/2 PK", watt: 400.0, jam: 8, menit: 30, ruangan: "Kamar Utama"},
		{nama: "Kulkas 2 Pintu", watt: 150.0, jam: 24, menit: 0, ruangan: "Dapur"},
		{nama: "TV LED 43 Inch", watt: 80.0, jam: 4, menit: 15, ruangan: "Ruang Keluarga"},
		{nama: "Lampu Teras", watt: 15.0, jam: 12, menit: 0, ruangan: "Teras Depan"},
		{nama: "PC Gaming", watt: 550.0, jam: 5, menit: 45, ruangan: "Kamar Anak"},
		{nama: "Mesin Cuci", watt: 350.0, jam: 1, menit: 30, ruangan: "Ruang Cuci"},
		{nama: "Setrika Listrik", watt: 300.0, jam: 1, menit: 0, ruangan: "Kamar Utama"},
		{nama: "Rice Cooker", watt: 400.0, jam: 1, menit: 20, ruangan: "Dapur"},
		{nama: "Dispenser Air", watt: 250.0, jam: 24, menit: 0, ruangan: "Ruang Makan"},
		{nama: "Microwave", watt: 800.0, jam: 0, menit: 15, ruangan: "Dapur"},
		{nama: "Kipas Angin", watt: 45.0, jam: 6, menit: 0, ruangan: "Ruang Keluarga"},
		{nama: "Pompa Air", watt: 200.0, jam: 2, menit: 0, ruangan: "Halaman Belakang"},
		{nama: "Lampu Tidur", watt: 5.0, jam: 8, menit: 0, ruangan: "Kamar Utama"},
		{nama: "Hair Dryer", watt: 600.0, jam: 0, menit: 20, ruangan: "Kamar Mandi"},
		{nama: "TV LED 32 Inch", watt: 50.0, jam: 3, menit: 0, ruangan: "Kamar Anak"},
		{nama: "Laptop Kerja", watt: 65.0, jam: 6, menit: 30, ruangan: "Ruang Kerja"},
		{nama: "Router WiFi", watt: 10.0, jam: 24, menit: 0, ruangan: "Ruang Keluarga"},
		{nama: "Air Purifier", watt: 35.0, jam: 12, menit: 0, ruangan: "Kamar Utama"},
		{nama: "Lampu Dapur", watt: 20.0, jam: 5, menit: 0, ruangan: "Dapur"},
		{nama: "Exhaust Fan", watt: 30.0, jam: 4, menit: 0, ruangan: "Kamar Mandi"},
	}

	// Memasukkan ke-20 data di atas ke dalam array data utama milik program
	for i = 0; i < 20; i++ {
		data[*n] = dummy[i]
		*n = *n + 1
	}
}
