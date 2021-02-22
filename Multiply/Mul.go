package multiply

import (
	"log"
)

//Multiplyvalue function provides sum of 2 integers
func Multiplyvalue(a, b int) int {
	/* local variable declaration */
	var result int
	log.Print("The program is in the MULTIPLY function and used log to print this log\n")
	result = a * b
	return result
}
