package dbexport

import (
	"fmt"
	"os"

	"github.com/subosito/gotenv"
)

func GetConfig() {
	env := os.Getenv("DB_ENV")

	if env == "" {
		env = "local"
	}

	err := gotenv.Load(".env." + env)

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
