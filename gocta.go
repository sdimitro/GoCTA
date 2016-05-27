package gocta

import (
	"encoding/xml"
	"time"
)

// A Query type holds the top level information
// of an XML response returned by an API call.
type Query struct {
	Timestamp      string `xml:"tmst"`
	ErrorCode      int    `xml:"errCd"`
	ErrorName      string `xml:"errNm"`
	PredictionList []Eta  `xml:"eta"`
}

// An Eta type holds the specific information
// of a single prediction. A single API
// response generally returns multiple
// predictions (thus a Query holds a list of
// Etas)
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

// ParseCTAResponse parses an API response s.
// It returns a Query struct containing all the parsed
// information and any unmarshalling errors encountered.
func ParseCTAResponse(resp string) (*Query, error) {
	var q Query
	b := []byte(resp)
	if err := xml.Unmarshal(b, &q); err != nil {
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
