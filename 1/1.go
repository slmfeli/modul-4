package main
import "fmt"

func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 2; i <= n; i++ {
		*hasil *= i
	}
}

func permutation(n, r int, hasil *int) {
	var nFact, n_rFact int
	factorial(n, &nFact)
	factorial(n-r, &n_rFact)
	*hasil = nFact / n_rFact
}

func combination(n, r int, hasil *int) {
	var nFact, rFact, n_rFact int
	factorial(n, &nFact)
	factorial(r, &rFact)
	factorial(n-r, &n_rFact)
	*hasil = nFact / (rFact * n_rFact)
}

func main() {
	var a, b, c, d int
	fmt.Println("Masukkan nilai a, b, c, d (a >= c dan b >= d):")
	fmt.Scan(&a, &b, &c, &d)

	var p1, c1, p2, c2 int

	permutation(a, c, &p1)
	combination(a, c, &c1)

	permutation(b, d, &p2)
	combination(b, d, &c2)

	fmt.Printf("Baris pertama (P(a, c), C(a, c)): %d %d\n", p1, c1)
	fmt.Printf("Baris kedua (P(b, d), C(b, d)): %d %d\n", p2, c2)
}