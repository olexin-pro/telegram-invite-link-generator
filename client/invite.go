package client

import (
	"encoding/json"
	"log"
)

type Invite struct {
	Link string `json:"invite_link"`
}

type InviteResponse struct {
	Status      bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
	Link        Invite `json:"result"`
}

type TelegramError struct {
	ErrorDescription string
}

func (m *TelegramError) Error() string {
	return m.ErrorDescription
}

func GetInvite() (Invite, error) {

	resp, err := fetch()

	if err != nil {
		return Invite{}, err
	}
	defer resp.Body.Close()

	var cResp InviteResponse

	//Decode the data
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		log.Fatal("ooopsss! an error occurred, please try again")
	}

	if cResp.Status {
		return cResp.Link, nil
	} else {
		return Invite{}, &TelegramError{ErrorDescription: cResp.Description}
	}
}
