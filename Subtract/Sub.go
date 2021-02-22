package subtract

import (
	"github.com/anugoli05/MyGoWSwithEnvbuild/.MyPackages/mathutils"
	
)

//CallSub is a function to call subtract function of the mygopackage to maintain my own package
func CallSub(a, b int) int {
	var Subresult int
	Subresult = mathutils.Subtract(a, b)
	return Subresult
}
