package writer

import (
	"fmt"
	"log"
	"os"
	"tg-invite-link-generator/config"
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

	conf, cErr := config.LoadConfig()
	if cErr != nil {
		log.Fatal("Config loading error:", cErr)
	}
	withPrefPostString := fmt.Sprintf("%s%s%s", conf.Prefix, data, conf.Postfix)
	if _, err = f.WriteString(fmt.Sprintf("%s\r\n", withPrefPostString)); err != nil {
		panic(err)
	}
}
