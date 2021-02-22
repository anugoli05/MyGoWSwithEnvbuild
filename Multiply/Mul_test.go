package multiply

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/anugoli05/MyGoWSwithEnv/.MyPackages"
)

//TestMultiplyvalue function validates Sumvalue function
func TestMultiplyvalue(t *testing.T) {
	t.Logf("\n-----------------------------------------------------------------------")
	t.Logf("\n-----------------------------------------------------------------------")
	t.Logf("\nThis is message before calling MULTIPLY function the test go file")
	/*Calling max function this is a comment */
	var Mulresultvalue int
	Mulresultvalue = Multiplyvalue(3, 4)
	//To log any info while test is running the test has to be run in command prompt with go test -v command to show log
	t.Logf("\nThis is message from t.Logf and Maxvalue returned is  %d", Mulresultvalue)
	// assert.Equal(t,Maxvalue,15)

	// if Sumresultvalue != 20 {

	// 	t.Errorf("Expected Max value is  %d and Actual Value is %d", 20, Sumresultvalue)
	// 	// t.FailNow()
	// }
	assert.Equal(t, 12, Mulresultvalue)
	t.Log("Testpassed after the assert logic of max function in test go file.")

	config, err := util.LoadConfig("./..")

	if err!=nil{
	t.Fatal("cannot  load config: ", err)
   }
   
   t.Log("*******************************")
   
   
   t.Log("Testenvironment: ", config.TestEnv)
   t.Log("\nURL: ", config.URL)
   t.Log("*******************************")

}
