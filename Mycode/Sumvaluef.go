package sumvalue

import (
	"log"
)

//Sumvalue function provides sum of 2 integers
func Sumvalue(a, b int) int {
	/* local variable declaration */
	var result int
	log.Print("The program is in the Sum function and used log to print this log\n")
	result = a + b
	return result
}
