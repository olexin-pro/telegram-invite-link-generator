package writer

import (
	"fmt"
	"os"
)

func WriteStringToFile(filepath string, data string) {

	if data == "" || data == " " {
		return
	}

	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(fmt.Sprintf("%s\r\n", data)); err != nil {
		panic(err)
	}
}
