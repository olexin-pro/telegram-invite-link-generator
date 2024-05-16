package main

import (
	"fmt"
	"log"
	"tg-invite-link-generator/client"
	"tg-invite-link-generator/config"
	"tg-invite-link-generator/limiter"
	"tg-invite-link-generator/writer"
	"time"
)

func main() {
	fmt.Println("Start generate invite links")
	conf, cErr := config.LoadConfig()
	if cErr != nil {
		log.Fatal("Config loading error:", cErr)
	}

	limit := limiter.NewLimiter(time.Minute, conf.RequestsPerMinute)

	for {
		limit.Wait()

		invite, err := client.GetInvite()

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(invite.Link)
			writer.WriteStringToFile("links.csv", invite.Link)
		}
	}
}
