package main

import (
	"fmt"

	systemutil "example.com/http_dashboard/system_util"
	"example.com/http_dashboard/webserver"
)

func main() {

	_, err := systemutil.BuildDashboardData()

	if err != nil {
		fmt.Println(err)
		return
	}

	webserver.StartWebServer()
}
