package gocta

import (
	"fmt"
	"testing"
)

var sampleResponse = `
<ctatt>
	<tmst>20160525 09:15:33</tmst>
	<errCd>0</errCd><errNm/>
	<eta>
		<staId>40190</staId>
		<stpId>30036</stpId>
		<staNm>Sox-35th</staNm>
		<stpDe>Service toward Howard</stpDe>
		<rn>811</rn>
		<rt>Red</rt>
		<destSt>30173</destSt>
		<destNm>Howard</destNm>
		<trDr>1</trDr>
		<prdt>20160525 09:14:59</prdt>
		<arrT>20160525 09:16:59</arrT>
		<isApp>0</isApp>
		<isSch>0</isSch>
		<isDly>0</isDly>
		<isFlt>0</isFlt>
		<flags/>
		<lat>41.81264</lat>
		<lon>-87.63019</lon>
		<heading>1</heading>
	</eta>
	<eta>
		<staId>40190</staId>
		<stpId>30037</stpId>
		<staNm>Sox-35th</staNm>
		<stpDe>Service toward 95th</stpDe>
		<rn>913</rn>
		<rt>Red</rt>
		<destSt>30089</destSt>
		<destNm>95th/Dan Ryan</destNm>
		<trDr>5</trDr>
		<prdt>20160525 09:15:17</prdt>
		<arrT>20160525 09:16:17</arrT>
		<isApp>1</isApp>
		<isSch>0</isSch>
		<isDly>0</isDly>
		<isFlt>0</isFlt>
		<flags/>
		<lat>41.83896</lat>
		<lon>-87.63085</lon>
		<heading>178</heading>
	</eta>
</ctatt>
`

func TestResponseParser(t *testing.T) {
	var tests = []struct {
		input  string
		tmstmp string
		predno int
	}{
		{sampleResponse, "20160525 09:15:33", 2},
	}
	for _, test := range tests {
		got, err := ParseCTAResponse(test.input)
		if err != nil {
			t.Errorf("ParseCTAResponse(%q) - could not parse! - %v",
				test.input, err)
		}
		if got.Timestamp != test.tmstmp {
			t.Errorf("ParseCTAResponse(%q) = {..Timestamp = %v..}",
				test.input, got.Timestamp)
		}
		if preds := len(got.PredictionList); preds != test.predno {
			t.Errorf("ParseCTAResponse(%q) gives %d predictions",
				test.input, preds)
		}
	}
}

func BenchmarkResponseParser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseCTAResponse(sampleResponse)
	}
}

func ExampleParseCTAResponse() {
	res, err := ParseCTAResponse(sampleResponse)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Println(res.Timestamp)
	fmt.Println(len(res.PredictionList))
	fmt.Println(res.PredictionList[0].StopDe)
	// Output:
	// 20160525 09:15:33
	// 2
	// Service toward Howard
}
