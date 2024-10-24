package main
import "fmt"

func cetakDeret(n int) {
	fmt.Print(n) 

	for n != 1 {
		if n%2 == 0 {
			n /= 2 
		} else {
			n = n*3 + 1 
		}
		fmt.Print(" ", n) 
	}

	fmt.Println() 
}

func main() {
	var n int

	fmt.Print("Masukkan nilai awal deret: ")
	fmt.Scan(&n)

	if n >= 1000000 {
		fmt.Println("Nilai awal harus lebih kecil dari 1.000.000")
		return
	}

	cetakDeret(n)
}