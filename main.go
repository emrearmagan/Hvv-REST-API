package main

import (
	"fmt"
	"hvvApi/app"
	"hvvApi/config"
)

func main() {
	c := app.NewClient()
	request := config.GetDeparuteRequest()

	resp, err := c.DepartureList(request)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
