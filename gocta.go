package gocta

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"
)

// StationMapID maps each station name
// (represented as a string) to its map
// ID (integer)
var StationMapID = map[string]int{
	"47th (Red Line)":             41230,
	"63rd (Red Line)":             40910,
	"69th (Red Line)":             40990,
	"79th (Red Line)":             40240,
	"87th (Red Line)":             41430,
	"95th/Dan Ryan (Red Line)":    40450,
	"Addison (Red Line)":          41420,
	"Argyle (Red Line)":           41200,
	"Belmont (Red Line)":          41320,
	"Berwyn (Red Line)":           40340,
	"Bryn Mawr (Red Line)":        41380,
	"Cermak-Chinatown (Red Line)": 41000,
	"Chicago (Red Line)":          41450,
	"Clark/Division (Red Line)":   40630,
	"Fullerton (Red Line)":        41220,
	"Garfield (Red Line)":         41170,
	"Grand (Red Line)":            40330,
	"Granville (Red Line)":        40760,
	"Harrison (Red Line)":         41490,
	"Howard (Red Line)":           40900,
	"Jackson (Red Line)":          40560,
	"Jarvis (Red Line)":           41190,
	"Lake (Red Line)":             41660,
	"Lawrence (Red Line)":         40770,
	"Loyola (Red Line)":           41300,
	"Monroe (Red Line)":           41090,
	"Morse (Red Line)":            40100,
	"North/Clybourn (Red Line)":   40650,
	"Roosevelt (Red Line)":        41400,
	"Sheridan (Red Line)":         40080,
	"Sox-35th (Red Line)":         40190,
	"Thorndale (Red Line)":        40880,
	"Wilson (Red Line)":           40540,
}

// The CTAResponse type holds the top level
// information of an XML response returned by
// an API call.
type CTAResponse struct {
	Timestamp      string `xml:"tmst"`
	ErrorCode      int    `xml:"errCd"`
	ErrorName      string `xml:"errNm"`
	PredictionList []Eta  `xml:"eta"`
}

// An Eta type holds the specific information
// of a single prediction. A single API
// response generally returns multiple
// predictions (thus CTAResponse holds a list
// of Etas)
type Eta struct {
	StationID   int    `xml:"staId"`
	StopID      int    `xml:"stpId"`
	StationName string `xml:"staNm"`
	StopDe      string `xml:"stpDe"`
	RunNo       int    `xml:"rn"`
	Route       string `xml:"rt"`
	DestStation int    `xml:"destSt"`
	DestName    string `xml:"destNm"`
	TrainDir    int    `xml:"trDr"`
	PredDate    string `xml:"prdt"`
	Arrival     string `xml:"arrT"`
	IsApp       bool   `xml:"isApp"`
	IsDly       bool   `xml:"isDly"`
	IsSch       bool   `xml:"isSch"`
	IsFlt       bool   `xml:"isFlt"`
	Flags       string `xml:"flags"`
	Latitude    string `xml:"lat"`
	Longtitude  string `xml:"lon"`
	Heading     string `xml:"heading"`
}

// ParseCTAResponse parses an API response r.
// It returns a CTAResponse struct containing all the parsed
// information and any unmarshalling errors encountered.
func ParseCTAResponse(r []byte) (*CTAResponse, error) {
	var q CTAResponse
	if err := xml.Unmarshal(r, &q); err != nil {
		return nil, err
	}
	return &q, nil
}

// ParseCTATime parses a date and time in the API's
// format accepted as a string s, and returns a Time
// struct together with any parsing errors.
func ParseCTATime(s string) (time.Time, error) {
	const layout = "20060102 15:04:05 MST"
	// Hopefully Chicago won't change timezones :-)
	s += " CDT"
	return time.Parse(layout, s)
}

// GetPredictions makes a call to the CTA's API
// given an API key and a map ID as arguments.
// It returns the response of the call and any
// errors encountered.
func GetPredictions(apiKey, mapID string) ([]byte, error) {
	req := generateRequest(apiKey, mapID)
	resp, err := http.Get(req)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func generateRequest(apiKey, mapID string) string {
	const apiURL = "http://lapi.transitchicago.com/api/1.0/ttarrivals.aspx"
	buf := bytes.NewBufferString(apiURL)
	buf.WriteString("?key=")
	buf.WriteString(apiKey)
	buf.WriteString("&mapid=")
	buf.WriteString(mapID)
	return buf.String()
}
