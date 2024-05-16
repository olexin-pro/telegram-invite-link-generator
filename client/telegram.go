package client

import (
	"fmt"
	"log"
	"net/http"
	"tg-invite-link-generator/config"
)

func fetch() (resp *http.Response, err error) {
	//Build The URL string
	URL, urlErr := buildUrl()
	if urlErr != nil {
		log.Fatal("ooopsss! an error occurred, please try again", urlErr)
	}
	return http.Get(URL)
}

func buildUrl() (string, error) {

	// loading config
	configuration, cErr := config.LoadConfig()
	if cErr != nil {
		log.Fatal("Build url error:", cErr)
	}

	return fmt.Sprintf("https://api.telegram.org/bot%s/createChatInviteLink?chat_id=%s&member_limit=%d", configuration.Token, configuration.Group, configuration.MembersLimit), nil
}
