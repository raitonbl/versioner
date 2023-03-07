package common

import (
	"fmt"
	"os"
)

func DoExit(err error) {
	fmt.Println(err)
	os.Exit(402)
}
