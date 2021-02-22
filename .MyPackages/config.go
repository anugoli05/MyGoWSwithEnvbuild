package util

import (
	"fmt"
	"os"
    // "strings"
	"github.com/spf13/viper"
)

//ENV_VAR_SUFFIX is used to set the enviornment
const ENV_VAR_SUFFIX = "rc"


	
//Config stores all configuration of the application.
//The values are read by viper from a config file ot environment variables
type Config struct {
	TestEnv string `mapstructure:"TEST_ENV"`
	URL string `mapstructure:"URL"`
}

//LoadConfig reads configuration from file or environment variables
func LoadConfig(path string)(config Config, err error){

	os.Setenv("TEST_ENVIRONMENT", ENV_VAR_SUFFIX)
	fmt.Printf("After set from CONFIG FILE , TEST_ENVIRONMENT: %s\n", os.Getenv("TEST_ENVIRONMENT"))
	
	// fmt.Printf("*********************************")

	// for _, env := range os.Environ() {
		
	// 		envPair := strings.SplitN(env, "=", 2)
	// 		key := envPair[0]
	// 		value := envPair[1]
	
	// 		fmt.Printf("%s : %s\n", key, value)
	//   }
	//   fmt.Printf("*********************************")

	configFileName := fmt.Sprintf("%s%s", "app", os.Getenv("TEST_ENVIRONMENT"))
	fmt.Printf("*********************************")

	viper.AddConfigPath(path)

	viper.SetConfigName(configFileName)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err!=nil{
		return
	}
    err=viper.Unmarshal(&config)
	return

}
