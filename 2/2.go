package main
import "fmt"

func hitungSkor(soal []int) (int, int) {
	totalSkor := 0
	totalWaktu := 0

	for _, waktu := range soal {
		if waktu <= 300 {
			totalSkor++
			totalWaktu += waktu
		}
	}

	return totalSkor, totalWaktu
}

func main() {
	var jumlahPeserta, jumlahSoal int

	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&jumlahPeserta)

	fmt.Print("Masukkan jumlah soal: ")
	fmt.Scan(&jumlahSoal)

	namaPeserta := make([]string, jumlahPeserta)
	waktuPeserta := make([][]int, jumlahPeserta)

	for i := 0; i < jumlahPeserta; i++ {
		fmt.Printf("Masukkan nama peserta %d: ", i+1)
		fmt.Scan(&namaPeserta[i])

		waktuPeserta[i] = make([]int, jumlahSoal)
		fmt.Printf("Masukkan waktu pengerjaan soal (dalam menit) untuk %s:\n", namaPeserta[i])
		for j := 0; j < jumlahSoal; j++ {
			fmt.Printf("Soal %d: ", j+1)
			fmt.Scan(&waktuPeserta[i][j])
		}
	}

	pemenang := ""
	maxSkor := 0
	minWaktu := 301 

	for i, soal := range waktuPeserta {
		skor, waktu := hitungSkor(soal)
		fmt.Printf("%s menyelesaikan %d soal dengan waktu total %d menit\n", namaPeserta[i], skor, waktu)

		if skor > maxSkor || (skor == maxSkor && waktu < minWaktu) {
			pemenang = namaPeserta[i]
			maxSkor = skor
			minWaktu = waktu
		}
	}

	fmt.Printf("Pemenang adalah %s dengan skor %d dan waktu %d menit\n", pemenang, maxSkor, minWaktu)
}