//+build !mul,!sub,sum

package test

import (
	"os"
	"testing"
	"github.com/anugoli05/MyGoWSwithEnvbuild/.MyPackages"
	"github.com/stretchr/testify/assert"
	"github.com/anugoli05/MyGoWSwithEnvbuild/Sumv"
)

//TestSumvalue function validates Sumvalue function
func TestSumvalue(t *testing.T) {
	
	t.Logf("\n-----------------------------------------------------------------------")
	t.Logf("\n-----------------------------------------------------------------------")
	t.Logf("\nThis is message before calling Sum function the test go file")
	/*Calling max function this is a comment */
	var Sumresultvalue int
	Sumresultvalue = sumvalue.Sumvalue(15, 5)
	//To log any info while test is running the test has to be run in command prompt with go test -v command to show log
	t.Logf("\nThis is message from t.Logf and Maxvalue returned is  %d", Sumresultvalue)
	// assert.Equal(t,Maxvalue,15)

	// if Sumresultvalue != 20 {

	// 	t.Errorf("Expected Max value is  %d and Actual Value is %d", 20, Sumresultvalue)
	// 	// t.FailNow()
	// }
	assert.Equal(t, 20, Sumresultvalue)
	t.Log("Testpassed after the assert logic of max function in test go file.")
	
	config, err := util.LoadConfig("./..")

 if err!=nil{
 t.Fatal("cannot  load config: ", err)
}

t.Log("*******************************")


t.Log("Testenvironment: ", config.TestEnv)
t.Log("\nURL: ", config.URL)
t.Log("\nTEST_ENVIRONMENT from SUMTEST file: ", os.Getenv("TEST_ENVIRONMENT"))
t.Log("*******************************")

}
