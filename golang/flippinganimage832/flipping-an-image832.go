package main

func main() {
	A := [][]int{{1, 1, 0}, {0, 0, 1}, {0, 1, 0}}
	flipAndInvertImage(A[0:len(A)])
}

func flipAndInvertImage(A [][]int) [][]int {
	for _, V := range A {
		for i := 0; i < len(V)-i-1; i++ {
			V[i] ^= V[len(V)-i-1]
			V[len(V)-i-1] ^= V[i]
			V[i] ^= V[len(V)-i-1]
		}
		for i := 0; i < len(V); i++ {
			V[i] = -V[i] + 1
		}
	}
	return A
}
