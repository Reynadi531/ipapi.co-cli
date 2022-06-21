package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type IPAPIResponse struct {
	IP                 string  `json:"ip"`
	Version            string  `json:"version"`
	City               string  `json:"city"`
	Region             string  `json:"region"`
	RegionCode         string  `json:"region_code"`
	RegionCodeIso3     string  `json:"region_code_iso3"`
	CountryName        string  `json:"country_name"`
	CountryCapital     string  `json:"country_capital"`
	CountryTLD         string  `json:"country_tld"`
	ContinetCode       string  `json:"continent_code"`
	InEurope           bool    `json:"in_eu"`
	Postal             string  `json:"postal"`
	Latitude           float32 `json:"latitude"`
	Logitude           float32 `json:"longitude"`
	Timezone           string  `json:"timezone"`
	UtcOffset          string  `json:"utc_offset"`
	CountryCallingCode string  `json:"country_calling_code"`
	Currency           string  `json:"currency"`
	CurrencyName       string  `json:"currency_name"`
	Languages          string  `json:"languages"`
	CountryArea        float32 `json:"country_area"`
	CountryPopulation  uint    `json:"country_population"`
	ASN                string  `json:"asn"`
	Org                string  `json:"org"`
	Hostname           string  `json:"hostname"`
}

func CallIPAPI(ip string) (IPAPIResponse, error) {
	var result IPAPIResponse
	var err error
	var client = &http.Client{}
	url := fmt.Sprintf("http://ipapi.co/%s/json", ip)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return IPAPIResponse{}, err
	}

	request.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36")

	response, err := client.Do(request)
	if err != nil {
		return IPAPIResponse{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return IPAPIResponse{}, fmt.Errorf("Status code is %d", response.StatusCode)
	}

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return IPAPIResponse{}, err
	}

	return result, err
}
