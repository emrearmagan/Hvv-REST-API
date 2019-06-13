package app

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"hvvApi/config"
)

var hvvConfig = &config.ApiConfig{
	Host: "https://api-test.geofox.de",
}

func (c *Client) GetRoute(r *config.HVVRequest) (*GetRouteResponse, error) {
	hvvConfig.Path = "/gti/public/getRoute"

	if err := checkHVVParams(r); err != nil {
		return nil, err
	}

	reqBody, err := json.Marshal(r)

	if err != nil {
		return nil, errors.New("couldn't Marshal HVV-Request")
	}
	signature := ComputeHmac256(reqBody, string(r.Apikey))

	header := map[string]string{"Content-Type": "application/json", "geofox-auth-signature": signature, "geofox-auth-user": r.Username}
	var response struct {
		GetRouteResponse
		HVVCommonResponse
	}

	if err := c.post(hvvConfig, r, &response, reqBody, header); err != nil {
		return nil, err
	}

	if err := response.StatusError(); err != nil {
		return nil, err
	}

	return &response.GetRouteResponse, nil
}

func (c *Client) DepartureList(r *config.HVVDepartureListRequest) (*HVVDepartureListResponse, error) {
	hvvConfig.Path = "/gti/public/departureList"

	//if err := checkHVVParams(r); err != nil {
	//	return nil, err
	//}

	reqBody, err := json.Marshal(r)
	if err != nil {
		return nil, errors.New("couldn't marshal HVV-Request")
	}
	signature := ComputeHmac256(reqBody, r.ApiKey)

	header := map[string]string{"Content-Type": "application/json", "geofox-auth-signature": signature, "geofox-auth-user": r.Username}
	var response struct {
		HVVDepartureListResponse
		HVVCommonResponse
	}

	if err := c.post(hvvConfig, r, &response, reqBody, header); err != nil {
		return nil, err
	}

	if err := response.StatusError(); err != nil {
		return nil, err
	}
	return &response.HVVDepartureListResponse, nil
}

//Generates the Signature for the Header
func ComputeHmac256(message []byte, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha1.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func checkHVVParams(r *config.HVVRequest) error {
	if len(r.Origin.Name) == 0 {
		return errors.New("origins empty")
	}
	if len(r.Destinations.Name) == 0 {
		return errors.New("destinations empty")
	}
	if len(r.Apikey) == 0 {
		return errors.New("no apiKey selected")
	}
	//@todo make sure date is now or in the future
	if len(r.DateTime.Time) == 0 || len(r.DateTime.Date) == 0 {
		return errors.New("no Date or Time selected")
	}
	if len(r.Username) == 0 {
		return errors.New("no username selected")
	}
	if r.MaxList <= 0 {
		return errors.New("no MaxList set")
	}
	if len(r.Language) == 0 {
		r.Language = "en"
	}
	return nil
}

// Expected Response from HBT API
type GetRouteResponse struct {
	RealtimeSchedules []struct {
		Start            config.Station `json:"start"`
		Dest             config.Station `json:"dest"`
		Time             float64        `json:"time"`
		FootpathTime     float64        `json:"footpathTime"`
		ScheduleElements []struct {
			From struct {
				Name    string          `json:"name"`
				DepTime config.DateTime `json:"depTime"`
			} `json:"from,omitempty"`
			To struct {
				Name    string          `json:"name"`
				ArrTime config.DateTime `json:"arrTime"`
			} `json:"to,omitempty"`
			Line struct {
				BusLine   string `json:"name"`
				Direction string `json:"direction"`
				Origin    string `json:"origin"`
				Type      struct {
					SimpleType string `json:"simpleType"`
					ShortInfo  string `json:"shortInfo"`
				} `json:"type"`
			} `json:"line,omitempty"`
		} `json:"scheduleElements"`
	} `json:"realtimeSchedules"`
}

type HVVCommonResponse struct {
	Status       config.Status       `json:"returnCode"`
	ErrorMessage config.ErrorMessage `json:"errorDevInfo,omitempty"`
	ErrorText    string              `json:"errorText,omitempty"`
}

//StatusError returns an error if this object has a Status different
func (c *HVVCommonResponse) StatusError() error {
	if c.Status != "OK" && c.Status != "ok" && c.Status != "200" {
		return fmt.Errorf("maps: %s - %s -%s", c.Status, c.ErrorText, c.ErrorMessage)
	}
	return nil
}

/*
Expected Response for DepatureList from HBT API
*/
type HVVDepartureListResponse struct {
	Time struct {
		Date string `json:"date"`
		Time string `json:"time"`
	} `json:"time"`
	Departures []struct {
		Line struct {
			Name      string `json:"name"`
			Direction string `json:"direction"`
			Origin    string `json:"origin"`
			Type      struct {
				SimpleType string `json:"simpleType"`
				ShortInfo  string `json:"shortInfo"`
			} `json:"type"`
		} `json:"line"`
		TimeOffset int `json:"timeOffset"`
	} `json:"departures"`
}
