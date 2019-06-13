package config

import (
	"encoding/json"
)

type Request struct {
}

/* HVVRequest is the request struct for HVV GetRoute APi
Example of Request Body
{"start":{"name":"START"},"dest":{"name":"DESTINATION"},"time":{"date":"12.06.2019","time":"14:00"},"language":"de","schedulesAfter":3,"timeIsDeparture":true,"schedulesBefore":0,"realtime":"REALTIME"}
*/
type HVVRequest struct {
	// Origins is a list of addresses and/or textual values. (example = Station{Name: "MyStation"})
	// Required.
	Origin Station `json:"start"`
	// Destinations is a list of addresses and/or textual values.(example = Station{Name: "MyDestionation"})
	// to which to calculate distance and time.
	// Required.
	Destinations Station `json:"dest,omitempty"`
	//Date and Time of the Request. (example = DateTime{Date: "11.06.2019", Time: "14:00"})
	//Required.
	DateTime DateTime `json:"time"`
	// Language in which to return results.
	// Optional. (default is english)
	Language string `json:"language"`
	// MaxList of routes to return.
	// Required
	MaxList int `json:"schedulesAfter"`
	// apiKey from HBT
	// Required
	Apikey string `json:"-"`
	// Username from HBT.
	//Required
	Username string `json:"-"`
}

/* HVVDepartureListRequest is the request struct for HVV departureList APi
*/
type HVVDepartureListRequest struct {
	// Origins is a list of addresses and/or textual values. (example = Station{Name: "MyStation"})
	// Required.
	Origin Station `json:"station"`
	//Date and Time of the Request. (example = DateTime{Date: "11.06.2019", Time: "14:00"})
	//Required.
	DateTime     DateTime `json:"time"`
	MaxList      int      `json:"maxList"`
	ServiceTypes []string `json:"serviceTypes"`
	// Language in which to return results.
	// Optional. (default is english)
	Language      string `json:"language"`
	MaxTimeOffset int    `json:"maxTimeOffset"`
	// apiKey from HBT
	// Required
	ApiKey   string `json:"-"`
	Username string `json:"-"`
}

type Station struct {
	Name string `json:"name"`
	Type string `json:"type"`
	ID   string `json:"id"`
}
type DateTime struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

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

// Implements the Json Interface and
// adds additional information to the body when called with json.Marshal(r)
func (r *HVVRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		HVVRequest
		TimeIsDeparture bool `json:"timeIsDeparture"`
		SchedulesBefore int  `json:"schedulesBefore"`
		// Provides you with Realtime data
		// Optional
		RealTime string `json:"realtime"`
	}{
		HVVRequest:      HVVRequest(*r),
		RealTime:        "REALTIME",
		TimeIsDeparture: true,
		SchedulesBefore: 0,
	})
}

// Implements the Json Interface and
// adds additional information to the body when called with json.Marshal(r)
func (r *HVVDepartureListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		HVVDepartureListRequest
		RealTime bool `json:"useRealtime"`
	}{
		HVVDepartureListRequest: HVVDepartureListRequest(*r),
		RealTime:                true,
	})
}
