package main

import (
	"fmt"
	"ipapi/pkg"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: ipapi <ip>\n")
		return
	}
	ipArgs := os.Args[1:][0]

	data, err := pkg.CallIPAPI(ipArgs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("IP:", data.IP)
	fmt.Println("Version:", data.Version)
	fmt.Println("City:", data.City)
	fmt.Println("Region:", data.Region)
	fmt.Println("Region Code:", data.RegionCode)
	fmt.Println("Region Code Iso3:", data.RegionCodeIso3)
	fmt.Println("Country Name:", data.CountryName)
	fmt.Println("Country Capital:", data.CountryCapital)
	fmt.Println("Country TLD:", data.CountryTLD)
	fmt.Println("Continet Code:", data.ContinetCode)
	fmt.Println("In Europe:", data.InEurope)
	fmt.Println("Postal:", data.Postal)
	fmt.Println("Latitude:", data.Latitude)
	fmt.Println("Logitude:", data.Logitude)
	fmt.Println("Timezone:", data.Timezone)
	fmt.Println("Utc Offset:", data.UtcOffset)
	fmt.Println("Country Calling Code:", data.CountryCallingCode)
	fmt.Println("Currency:", data.Currency)
	fmt.Println("Currency Name:", data.CurrencyName)
	fmt.Println("Languages:", data.Languages)
	fmt.Println("Country Area:", data.CountryArea)
	fmt.Println("Country Population:", data.CountryPopulation)
	fmt.Println("ASN:", data.ASN)
	fmt.Println("Org:", data.Org)
	fmt.Println("Hostname:", data.Hostname)
}
