package commons

import (
	"fmt"
	"os"
)

func Exit(err error) {
	fmt.Println(err)
	os.Exit(402)
}
