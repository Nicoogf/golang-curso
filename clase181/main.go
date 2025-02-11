//Package mymath prvides math solutions
package mymath


//Suma un numero
func Sum(xi ...int) int {
	sum := 0 
	for _ , v := range xi {
		sum += v
	}
	return sum
}