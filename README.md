# Hvv REST API for Go
A RESTful API example for Hvv. Documentation can be found at: https://geofox.hvv.de/gti/doc/index.jsp

## Requirements
- Key and Username provided by HBT

##Installation
The easiest way to use the HVV API in your Go project is to install it using **go get**:
```go
go get https://github.com/emrearmagan/Hvv-REST-API
```

Before running, you should set the config values in **[config.go]**
(https://github.com/emrearmagan/Hvv-REST-API/config/config.go)
```go
func GetDeparuteRequest() *HVVDepartureListRequest {
	return &HVVDepartureListRequest{
		Origin:        Station{Name: "ORIGIN", Type: "STATION", ID: "Master:41022"},
		DateTime:      DateTime{Date: "13.06.2019", Time: "14:00"},
		Language:      "de",
		MaxList:       30,
		ServiceTypes:  []string{"BUS", "ZUG", "FAEHRE"},
		MaxTimeOffset: 120,
		ApiKey:        "YOUR_KEY",
		Username:      "YOUR_USERNAME",
	}
}


func GetRouteRequest() *HVVRequest {
	return &HVVRequest{
		Origin: Station{Name: "ORIGIN"},
		Destinations: Station{Name: "DESTINATION"},
		DateTime: DateTime{Date: "12.06.2019", Time: "14:00"},
		Language: "de",
		MaxList:  3,
		Apikey:   "YOUR_KEY",
		Username: "YOUR_USERNAME",
	}
}
```

##Example
Here is a quick example on how to get started.
Create a new Client and get your Configs:
````go
import "github.com/messagebird/go-rest-api"

client := app.NewClient()
request := config.GetRouteRequest()
*OR*
request := config.GetDeparuteRequest
````
Now you can make simple API calls:
`````go
request, err := c.GetRoute(request)
	if err != nil {
		panic(err)
	}

*OR*
	
resp, err := c.DepartureList(request)
	if err != nil {
		panic(err)
	}

        .
	.
	.
	.
	.
`````

## Author
Emre Armagan, emre.armagan@hotmail.de