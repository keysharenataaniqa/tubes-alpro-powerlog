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

func clearScreen() {
	var cmd *exec.Cmd
	cmd = exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func bacaString() string {
	var reader *bufio.Reader
	var str string

	reader = bufio.NewReader(os.Stdin)
	str, _ = reader.ReadString('\n')
	str = strings.TrimSpace(str)

	for str == "" {
		str, _ = reader.ReadString('\n')
		str = strings.TrimSpace(str)
	}
	return str
}
func tambahData(n *int) {
	var hasilPerangkat, hasilRuangan, hasilWatt, hasilJam, hasilMenit string
	var sukses int
	var buang, kembali string

	clearScreen()
	fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("  ║                           ➕ TAMBAH DATA ➕                           ║")
	fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝")
	fmt.Println()

	if *n < NMAX {
		hasilPerangkat = "tidak_valid"
		for hasilPerangkat == "tidak_valid" {
			fmt.Print("    🏷️  Masukkan Nama Perangkat : ")
			data[*n].nama = bacaString()
			if cekString(data[*n].nama) {
				hasilPerangkat = "valid"
			} else {
				fmt.Println("  ❌ Nama Perangkat tidak valid!\n")
			}
		}

		hasilWatt = "tidak_valid"
		for hasilWatt == "tidak_valid" {
			fmt.Print("    ⚡  Masukkan Watt           : ")
			sukses, _ = fmt.Scan(&data[*n].watt)
			if sukses == 0 {
				fmt.Scanln(&buang)
				fmt.Println("  ❌ Input tidak valid! Harap masukkan angka.\n")
			} else if data[*n].watt > 0.0 {
				hasilWatt = "valid"
			} else {
				fmt.Println("  ❌ Watt tidak valid! Harus lebih dari 0.\n")
			}
		}

		hasilJam = "tidak_valid"
		for hasilJam == "tidak_valid" {
			fmt.Print("    ⏱️  Masukkan Durasi (Jam)   : ")
			sukses, _ = fmt.Scan(&data[*n].jam)
			if sukses == 0 {
				fmt.Scanln(&buang)
				fmt.Println("  ❌ Input tidak valid! Harap masukkan angka.\n")
			} else if data[*n].jam >= 0 {
				hasilJam = "valid"
			} else {
				fmt.Println("  ❌ Jam tidak valid! Tidak boleh minus.\n")
			}
		}

		hasilMenit = "tidak_valid"
		for hasilMenit == "tidak_valid" {
			fmt.Print("    ⏱️  Masukkan Durasi (Menit) : ")
			sukses, _ = fmt.Scan(&data[*n].menit)
			if sukses == 0 {
				fmt.Scanln(&buang)
				fmt.Println("  ❌ Input tidak valid! Harap masukkan angka.\n")
			} else if data[*n].menit >= 0 && data[*n].menit < 60 {
				if data[*n].jam == 0 && data[*n].menit == 0 {
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
			data[*n].ruangan = bacaString()
			if cekString(data[*n].ruangan) {
				hasilRuangan = "valid"
			} else {
				fmt.Println("  ❌ Nama Ruangan tidak valid!\n")
			}
		}

		*n = *n + 1

		fmt.Println("\n  ╔═══════════════════════════════════════════════════════════════════════╗")
		fmt.Println("  ║                     ✨ Data berhasil ditambahkan! ✨                  ║")
		fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")

	} else {
		fmt.Println("\n  ╔═══════════════════════════════════════════════════════════════════════╗")
		fmt.Println("  ║                       ⚠️  Kapasitas Data Penuh!                       ║")
		fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")
	}

	fmt.Println("Ketik 0 dan tekan [Enter] untuk kembali... ")
	kembali = ""
	for kembali != "0" {
		fmt.Scan(&kembali)
		if kembali != "0" {
			fmt.Println("  ❌ Input tidak valid! Harap ketik 0: ")
		}
	}
}
func cekString(input string) bool {
	var cekAngka bool
	var i int
	if input == "" || input == `""` || input == "''" || input == "-" || input == " " {
		return false
	}
	cekAngka = true
	for i = 0; i < len(input); i++ {
		if input[i] >= '0' && input[i] <= '9' {
			cekAngka = false
		}
	}
	return cekAngka
}
func tampilData(n int) {
	var i int
	var kembali string

	clearScreen()

	fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("  ║                      📋 TAMPILKAN DATA PERANGKAT 📋                   ║")
	fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝")
	fmt.Println()

	if n == 0 {
		fmt.Println("  📭 Data perangkat masih kosong. Yuk, tambah data dulu!\n")
	} else {
		fmt.Println("  ┌────┬──────────────────────┬─────────┬──────────┬──────────────────────┐")
		fmt.Println("  │ No │ Nama Perangkat       │ Watt    │ Durasi   │ Ruangan              │")
		fmt.Println("  ├────┼──────────────────────┼─────────┼──────────┼──────────────────────┤")

		for i = 0; i < n; i++ {
			fmt.Printf("  │ %-2d │ %-20s │ %-5.1f W │ %2dj %02dm  │ %-20s │\n", i+1, data[i].nama, data[i].watt, data[i].jam, data[i].menit, data[i].ruangan)
		}
		fmt.Println("  └────┴──────────────────────┴─────────┴──────────┴──────────────────────┘")
	}

	fmt.Println("Ketik 0 dan tekan [Enter] untuk kembali... ")
	kembali = ""
	for kembali != "0" {
		fmt.Scan(&kembali)
		if kembali != "0" {
			fmt.Println("  ❌ Input tidak valid! Harap ketik 0: ")
		}
	}
}
func updateData(n int) {
	var nData, i, sukses, suksesWatt, suksesJam, suksesMenit int
	var hasilPerangkat, hasilRuangan, hasilWatt, hasilJam, hasilMenit, buang, kembali string
	var validFormat bool

	clearScreen()

	fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("  ║                             ✏️  UBAH DATA ✏️                            ║")
	fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝")

	if n == 0 {
		fmt.Println("  📭 Data perangkat masih kosong!\n")
	} else {
		fmt.Println("  ┌────┬──────────────────────┬─────────┬──────────┬──────────────────────┐")
		fmt.Println("  │ No │ Nama Perangkat       │ Watt    │ Durasi   │ Ruangan              │")
		fmt.Println("  ├────┼──────────────────────┼─────────┼──────────┼──────────────────────┤")

		for i = 0; i < n; i++ {
			fmt.Printf("  │ %-2d │ %-20s │ %-5.1f W │ %2dj %02dm  │ %-20s │\n", i+1, data[i].nama, data[i].watt, data[i].jam, data[i].menit, data[i].ruangan)
		}
		fmt.Println("  └────┴──────────────────────┴─────────┴──────────┴──────────────────────┘")

		validFormat = false
		for validFormat == false {
			fmt.Print("  🔢 Masukkan nomor data yang diubah : ")
			sukses, _ = fmt.Scan(&nData)

			if sukses == 0 {
				fmt.Scanln(&buang)
				fmt.Println("  ❌ Input tidak valid! Harap masukkan angka, bukan huruf.\n")
			} else {
				validFormat = true
			}
		}
		nData = nData - 1

		if nData >= 0 && nData < n {
			fmt.Println("\n  ╔═══════════════════════════════════════════════════════════════════════╗")
			fmt.Println("  ║                        ✨ MASUKKAN DATA BARU ✨                       ║")
			fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝")
			fmt.Println()

			hasilPerangkat = "tidak_valid"
			for hasilPerangkat == "tidak_valid" {
				fmt.Print("    🏷️  Masukkan Nama Perangkat : ")
				data[nData].nama = bacaString()
				if cekString(data[nData].nama) {
					hasilPerangkat = "valid"
				} else {
					fmt.Println("  ❌ Nama Perangkat tidak valid!\n")
				}
			}

			hasilWatt = "tidak_valid"
			for hasilWatt == "tidak_valid" {
				fmt.Print("    ⚡  Masukkan Watt           : ")
				suksesWatt, _ = fmt.Scan(&data[nData].watt)

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
				suksesJam, _ = fmt.Scan(&data[nData].jam)

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
				suksesMenit, _ = fmt.Scan(&data[nData].menit)

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
				data[nData].ruangan = bacaString()
				if cekString(data[nData].ruangan) {
					hasilRuangan = "valid"
				} else {
					fmt.Println("  ❌ Nama Ruangan tidak valid!\n")
				}
			}

			fmt.Println("\n  ╔═══════════════════════════════════════════════════════════════════════╗")
			fmt.Println("  ║                      ✨ Data berhasil diubah! ✨                      ║")
			fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")

		} else {
			fmt.Println("\n  ╔═══════════════════════════════════════════════════════════════════════╗")
			fmt.Println("  ║                       ⚠️  Data tidak ditemukan!                        ║")
			fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")
		}
	}

	fmt.Println("Ketik 0 dan tekan [Enter] untuk kembali... ")
	kembali = ""
	for kembali != "0" {
		fmt.Scan(&kembali)
		if kembali != "0" {
			fmt.Println("  ❌ Input tidak valid! Harap ketik 0: ")
		}
	}
}
func hapusData(n *int) {
	var nData, i, sukses int
	var buang, kembali string
	var validFormat bool

	clearScreen()

	fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("  ║                            🗑️  HAPUS DATA 🗑️                            ║")
	fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝")
	fmt.Println()

	if *n == 0 {
		fmt.Println("  📭 Data perangkat masih kosong!\n")
	} else {
		fmt.Println("  ┌────┬──────────────────────┬─────────┬──────────┬──────────────────────┐")
		fmt.Println("  │ No │ Nama Perangkat       │ Watt    │ Durasi   │ Ruangan              │")
		fmt.Println("  ├────┼──────────────────────┼─────────┼──────────┼──────────────────────┤")

		for i = 0; i < *n; i++ {
			fmt.Printf("  │ %-2d │ %-20s │ %-5.1f W │ %2dj %02dm  │ %-20s │\n", i+1, data[i].nama, data[i].watt, data[i].jam, data[i].menit, data[i].ruangan)
		}
		fmt.Println("  └────┴──────────────────────┴─────────┴──────────┴──────────────────────┘")

		validFormat = false
		for validFormat == false {
			fmt.Print("  🔢 Masukkan nomor data yang dihapus : ")
			sukses, _ = fmt.Scan(&nData)

			if sukses == 0 {
				fmt.Scanln(&buang)
				fmt.Println("  ❌ Input tidak valid! Harap masukkan angka, bukan huruf.\n")
			} else {
				validFormat = true
			}
		}
		nData = nData - 1

		if nData >= 0 && nData < *n {
			for i = nData; i < *n-1; i++ {
				data[i] = data[i+1]
			}
			*n = *n - 1

			fmt.Println("\n  ╔═══════════════════════════════════════════════════════════════════════╗")
			fmt.Println("  ║                     ✨ Data berhasil dihapus! ✨                      ║")
			fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")
		} else {
			fmt.Println("\n  ╔═══════════════════════════════════════════════════════════════════════╗")
			fmt.Println("  ║                       ⚠️  Data tidak ditemukan!                        ║")
			fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")
		}
	}

	fmt.Println("Ketik 0 dan tekan [Enter] untuk kembali... ")
	kembali = ""
	for kembali != "0" {
		fmt.Scan(&kembali)
		if kembali != "0" {
			fmt.Println("  ❌ Input tidak valid! Harap ketik 0: ")
		}
	}
}
func cariPerangkat(n int) {
	clearScreen()
	var find, kembali string
	var i int
	var ketemu bool

	fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("  ║                        🔎 CARI NAMA PERANGKAT 🔎                      ║")
	fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝")
	fmt.Println()

	if n == 0 {
		fmt.Println("  📭 Data perangkat masih kosong!\n")
	} else {
		fmt.Print("  🏷️  Masukkan nama perangkat yang dicari : ")
		find = bacaString()
		fmt.Println()

		ketemu = false
		for i = 0; i < n; i++ {
			if data[i].nama == find {
				if ketemu == false {
					fmt.Println("  ┌────┬──────────────────────┬─────────┬──────────┬──────────────────────┐")
					fmt.Println("  │ No │ Nama Perangkat       │ Watt    │ Durasi   │ Ruangan              │")
					fmt.Println("  ├────┼──────────────────────┼─────────┼──────────┼──────────────────────┤")
					ketemu = true
				}
				fmt.Printf("  │ %-2d │ %-20s │ %-5.1f W │ %2dj %02dm  │ %-20s │\n", i+1, data[i].nama, data[i].watt, data[i].jam, data[i].menit, data[i].ruangan)
			}
		}

		if ketemu {
			fmt.Println("  └────┴──────────────────────┴─────────┴──────────┴──────────────────────┘\n")
			fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
			fmt.Println("  ║                     ✨ Data berhasil ditemukan! ✨                    ║")
			fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")
		} else {
			fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
			fmt.Println("  ║                       ⚠️  Data tidak ditemukan!                        ║")
			fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")
		}
	}

	fmt.Println("Ketik 0 dan tekan [Enter] untuk kembali... ")
	kembali = ""
	for kembali != "0" {
		fmt.Scan(&kembali)
		if kembali != "0" {
			fmt.Println("  ❌ Input tidak valid! Harap ketik 0: ")
		}
	}
}
func cariRuangan(n int) {
	clearScreen()
	var find, kembali string
	var i int
	var ketemu bool

	fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("  ║                           🏠 CARI RUANGAN 🏠                          ║")
	fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝")
	fmt.Println()

	if n == 0 {
		fmt.Println("  📭 Data perangkat masih kosong!\n")
	} else {
		fmt.Print("  🏠 Masukkan nama ruangan yang dicari : ")
		find = bacaString()
		fmt.Println()

		ketemu = false
		for i = 0; i < n; i++ {
			if data[i].ruangan == find {

				if ketemu == false {
					fmt.Println("  ┌────┬──────────────────────┬─────────┬──────────┬──────────────────────┐")
					fmt.Println("  │ No │ Nama Perangkat       │ Watt    │ Durasi   │ Ruangan              │")
					fmt.Println("  ├────┼──────────────────────┼─────────┼──────────┼──────────────────────┤")
				}

				fmt.Printf("  │ %-2d │ %-20s │ %-5.1f W │ %2dj %02dm  │ %-20s │\n", i+1, data[i].nama, data[i].watt, data[i].jam, data[i].menit, data[i].ruangan)
				ketemu = true
			}
		}

		if ketemu {
			fmt.Println("  └────┴──────────────────────┴─────────┴──────────┴──────────────────────┘\n")
			fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
			fmt.Println("  ║                     ✨ Data berhasil ditemukan! ✨                    ║")
			fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")
		} else {
			fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
			fmt.Println("  ║                       ⚠️  Data tidak ditemukan!                        ║")
			fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")
		}
	}

	fmt.Println("Ketik 0 dan tekan [Enter] untuk kembali... ")
	kembali = ""
	for kembali != "0" {
		fmt.Scan(&kembali)
		if kembali != "0" {
			fmt.Println("  ❌ Input tidak valid! Harap ketik 0: ")
		}
	}
}
func searchPerangkat(n int) {
	var low, high, mid, kiri, kanan, i int
	var find, kembali string
	var ketemu bool

	clearScreen()

	fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("  ║                      🔎 PENCARIAN NAMA PERANGKAT 🔎                   ║")
	fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝")
	fmt.Println()

	if n == 0 {
		fmt.Println("  📭 Data perangkat masih kosong!\n")
	} else {
		urutPerangkatASC(n)

		fmt.Print("  🏷️  Masukkan nama perangkat yang dicari : ")
		find = bacaString()
		fmt.Println()

		low = 0
		high = n - 1
		ketemu = false

		for low <= high && ketemu == false {
			mid = (low + high) / 2

			if data[mid].nama == find {
				ketemu = true
			} else if find < data[mid].nama {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}

		if ketemu == true {
			kiri = mid
			for kiri > 0 && data[kiri-1].nama == find {
				kiri = kiri - 1
			}

			kanan = mid
			for kanan < n-1 && data[kanan+1].nama == find {
				kanan = kanan + 1
			}

			fmt.Println("  ┌────┬──────────────────────┬─────────┬──────────┬──────────────────────┐")
			fmt.Println("  │ No │ Nama Perangkat       │ Watt    │ Durasi   │ Ruangan              │")
			fmt.Println("  ├────┼──────────────────────┼─────────┼──────────┼──────────────────────┤")

			for i = kiri; i <= kanan; i++ {
				fmt.Printf("  │ %-2d │ %-20s │ %-5.1f W │ %2dj %02dm  │ %-20s │\n", i+1, data[i].nama, data[i].watt, data[i].jam, data[i].menit, data[i].ruangan)
			}

			fmt.Println("  └────┴──────────────────────┴─────────┴──────────┴──────────────────────┘\n")

			fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
			fmt.Println("  ║                     ✨ Data berhasil ditemukan! ✨                    ║")
			fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")
		} else {
			fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
			fmt.Println("  ║                       ⚠️  Data tidak ditemukan!                        ║")
			fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")
		}
	}

	fmt.Println("Ketik 0 dan tekan [Enter] untuk kembali... ")
	kembali = ""
	for kembali != "0" {
		fmt.Scan(&kembali)
		if kembali != "0" {
			fmt.Println("  ❌ Input tidak valid! Harap ketik 0: ")
		}
	}
}
func urutRuanganASC(n int) {
	var i, j, idxMin int
	var temp Perangkat

	for i = 0; i < n-1; i++ {
		idxMin = i
		for j = i + 1; j < n; j++ {
			if data[j].ruangan < data[idxMin].ruangan {
				idxMin = j
			}
		}
		temp = data[i]
		data[i] = data[idxMin]
		data[idxMin] = temp
	}
}
func searchRuangan(n int) {
	var low, high, mid, kiri, kanan, i int
	var find, kembali string
	var ketemu bool

	clearScreen()

	fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("  ║                           🏠 CARI RUANGAN 🏠                          ║")
	fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝")
	fmt.Println()

	if n == 0 {
		fmt.Println("  📭 Data perangkat masih kosong!\n")
	} else {
		urutRuanganASC(n)

		fmt.Print("  🏠 Masukkan nama ruangan yang dicari : ")
		find = bacaString()
		fmt.Println()

		low = 0
		high = n - 1
		ketemu = false

		for low <= high && ketemu == false {
			mid = (low + high) / 2

			if data[mid].ruangan == find {
				ketemu = true
			} else if find < data[mid].ruangan {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}

		if ketemu == true {
			kiri = mid
			for kiri > 0 && data[kiri-1].ruangan == find {
				kiri = kiri - 1
			}
			kanan = mid
			for kanan < n-1 && data[kanan+1].ruangan == find {
				kanan = kanan + 1
			}

			fmt.Println("  ┌────┬──────────────────────┬─────────┬──────────┬──────────────────────┐")
			fmt.Println("  │ No │ Nama Perangkat       │ Watt    │ Durasi   │ Ruangan              │")
			fmt.Println("  ├────┼──────────────────────┼─────────┼──────────┼──────────────────────┤")

			for i = kiri; i <= kanan; i++ {
				fmt.Printf("  │ %-2d │ %-20s │ %-5.1f W │ %2dj %02dm  │ %-20s │\n", i+1, data[i].nama, data[i].watt, data[i].jam, data[i].menit, data[i].ruangan)
			}

			fmt.Println("  └────┴──────────────────────┴─────────┴──────────┴──────────────────────┘\n")

			fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
			fmt.Println("  ║                     ✨ Data berhasil ditemukan! ✨                    ║")
			fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")
		} else {
			fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
			fmt.Println("  ║                       ⚠️  Data tidak ditemukan!                        ║")
			fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")
		}
	}

	fmt.Println("Ketik 0 dan tekan [Enter] untuk kembali... ")
	kembali = ""
	for kembali != "0" {
		fmt.Scan(&kembali)
		if kembali != "0" {
			fmt.Println("  ❌ Input tidak valid! Harap ketik 0: ")
		}
	}
}
func urutWattASC(n int) {
	var i, j, idxMin int
	var temp Perangkat
	for i = 0; i < n-1; i++ {
		idxMin = i
		for j = i + 1; j < n; j++ {
			if data[j].watt < data[idxMin].watt {
				idxMin = j
			}
		}
		temp = data[i]
		data[i] = data[idxMin]
		data[idxMin] = temp
	}
}
func urutWattDESC(n int) {
	var i, j, idxMin int
	var temp Perangkat
	for i = 0; i < n-1; i++ {
		idxMin = i
		for j = i + 1; j < n; j++ {
			if data[j].watt > data[idxMin].watt {
				idxMin = j
			}
		}
		temp = data[i]
		data[i] = data[idxMin]
		data[idxMin] = temp
	}
}
func urutPerangkatASC(n int) {
	var i, j, idxMin int
	var temp Perangkat
	for i = 0; i < n-1; i++ {
		idxMin = i
		for j = i + 1; j < n; j++ {
			if data[j].nama < data[idxMin].nama {
				idxMin = j

			}
		}
		temp = data[i]
		data[i] = data[idxMin]
		data[idxMin] = temp
	}
}
func urutPerangkatDESC(n int) {
	var i, j, idxMin int
	var temp Perangkat
	for i = 0; i < n-1; i++ {
		idxMin = i
		for j = i + 1; j < n; j++ {
			if data[j].nama > data[idxMin].nama {
				idxMin = j
			}
		}
		temp = data[i]
		data[i] = data[idxMin]
		data[idxMin] = temp
	}
}
func insertionWattASC(n int) {
	var pass, i int
	var temp Perangkat
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = data[pass]
		for i > 0 && temp.watt < data[i-1].watt {
			data[i] = data[i-1]
			i = i - 1
		}
		data[i] = temp
		pass = pass + 1
	}
}
func insertionWattDSC(n int) {
	var pass, i int
	var temp Perangkat
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = data[pass]
		for i > 0 && temp.watt > data[i-1].watt {
			data[i] = data[i-1]
			i = i - 1
		}
		data[i] = temp
		pass = pass + 1
	}
}
func insertionPerangkatASC(n int) {
	var pass, i int
	var temp Perangkat
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = data[pass]
		for i > 0 && temp.nama < data[i-1].nama {
			data[i] = data[i-1]
			i = i - 1
		}
		data[i] = temp
		pass = pass + 1
	}
}
func insertionPerangkatDSC(n int) {
	var pass, i int
	var temp Perangkat
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = data[pass]
		for i > 0 && temp.nama > data[i-1].nama {
			data[i] = data[i-1]
			i = i - 1
		}
		data[i] = temp
		pass = pass + 1
	}
}
func statistikData(n int) {
	var i int
	var totalDaya, durasiTotal float64
	var kembali string

	clearScreen()

	fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("  ║                      📊 STATISTIK DAYA HARIAN 📊                      ║")
	fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝")
	fmt.Println()

	if n == 0 {
		fmt.Println("  📭 Data perangkat masih kosong!\n")
	} else {
		totalDaya = 0.0
		for i = 0; i < n; i++ {
			durasiTotal = float64(data[i].jam) + (float64(data[i].menit) / 60.0)

			totalDaya += (data[i].watt * durasiTotal)
		}

		fmt.Println("  ┌───────────────────────────────────────────────────────────────────────┐")
		fmt.Printf("  │ ⚡ Total Perangkat Aktif      : %-37d │\n", n)
		fmt.Printf("  │ 🔋 Total Penggunaan Daya (Wh) : %-37.1f │\n", totalDaya)
		fmt.Println("  └───────────────────────────────────────────────────────────────────────┘\n")
	}

	fmt.Println("  Ketik 0 dan tekan [Enter] untuk kembali... ")
	kembali = ""
	for kembali != "0" {
		fmt.Scan(&kembali)
		if kembali != "0" {
			fmt.Println("  ❌ Input tidak valid! Harap ketik 0: ")
		}
	}
}
func perangkatTerboros(n int) {
	var i, idxMax int
	var totalEnergi, durasiI, durasiMax float64
	var kembali string

	clearScreen()

	fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("  ║                       🔥 PERANGKAT TERBOROS 🔥                        ║")
	fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝")
	fmt.Println()

	if n == 0 {
		fmt.Println("  📭 Data perangkat masih kosong!\n")
	} else {
		idxMax = 0
		for i = 0; i < n; i++ {
			durasiI = float64(data[i].jam) + (float64(data[i].menit) / 60.0)
			durasiMax = float64(data[idxMax].jam) + (float64(data[idxMax].menit) / 60.0)

			if (data[i].watt * durasiI) > (data[idxMax].watt * durasiMax) {
				idxMax = i
			}
		}

		durasiMax = float64(data[idxMax].jam) + (float64(data[idxMax].menit) / 60.0)
		totalEnergi = data[idxMax].watt * durasiMax

		fmt.Println("  ┌────┬──────────────────────┬─────────┬──────────┬──────────────────────┐")
		fmt.Println("  │ No │ Nama Perangkat       │ Watt    │ Durasi   │ Ruangan              │")
		fmt.Println("  ├────┼──────────────────────┼─────────┼──────────┼──────────────────────┤")

		fmt.Printf("  │ %-2d │ %-20s │ %-5.1f W │ %2dj %02dm  │ %-20s │\n", 1, data[idxMax].nama, data[idxMax].watt, data[idxMax].jam, data[idxMax].menit, data[idxMax].ruangan)

		fmt.Println("  └────┴──────────────────────┴─────────┴──────────┴──────────────────────┘\n")

		fmt.Println("  ╔═══════════════════════════════════════════════════════════════════════╗")
		fmt.Printf("  ║  💥 Total Konsumsi Energi (Wh) : %-36.1f ║\n", totalEnergi)
		fmt.Println("  ╚═══════════════════════════════════════════════════════════════════════╝\n")
	}

	fmt.Println("Ketik 0 dan tekan [Enter] untuk kembali... ")
	kembali = ""
	for kembali != "0" {
		fmt.Scan(&kembali)
		if kembali != "0" {
			fmt.Print("  ❌ Input tidak valid! Harap ketik 0: ")
		}
	}
}
func main() {
	var menu, n, pilih, sukses int
	var buang, kembali string
	var validFormat bool
	generateDummyData(&n)
	for menu != 11 {
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
			sukses, _ = fmt.Scan(&menu)

			if sukses == 0 {
				fmt.Scanln(&buang)
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
				sukses, _ = fmt.Scan(&pilih)
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
			kembali = ""
			for kembali != "0" {
				fmt.Scan(&kembali)
				if kembali != "0" {
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

	for i = 0; i < 20; i++ {
		data[*n] = dummy[i]
		*n = *n + 1
	}
}
