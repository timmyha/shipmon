package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/olekukonko/tablewriter"
	"log"
	"net/http"
	"os"
)

type Carrier struct {
	Name      string
	Status       string
	ResponseTime string
	ThreeMinute  string
	Hour         string
	Day          string
	LastUpdate   string
}

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.shippingapimonitor.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var companies []Carrier

	doc.Find("table").Each(func(i int, sel *goquery.Selection) {
		if i == 0 {
			e := Carrier{}
			sel.Find("td").Each(func(index int, sel *goquery.Selection) {
				if index%7 == 1 {
					e.Status = sel.Text()
				}
				if index%7 == 2 {
					e.ResponseTime = sel.Text()
				}
				if index%7 == 3 {
					e.ThreeMinute = sel.Text()
				}
				if index%7 == 4 {
					e.Hour = sel.Text()
				}
				if index%7 == 5 {
					e.Day = sel.Text()
				}
				if index%7 == 6 {
					e.LastUpdate = sel.Text()
				}

				// Add the element to our array
				if index != 0 && (index+1)%7 == 0 {
					companies = append(companies, e)
				}
			})
		}
	})

	table := tablewriter.NewWriter(os.Stdout)

	// Setting our headers
	table.SetHeader([]string{"Carrier", "Status", "Response Time", "Three Minute^", "Hour^", "Day^", "Last Update"})
	for i, Carrier := range companies {
		if i == 0 {
			Carrier.Name = "FedEx"
		} else if i == 1 {
  		Carrier.Name = "UPS"
		} else if i == 2 {
			Carrier.Name = "USPS"
		} else if i == 3 {
			Carrier.Name = "CanadaPost"
		}
		s := []string{
			Carrier.Name,
			Carrier.Status,
			Carrier.ResponseTime,
			Carrier.ThreeMinute,
			Carrier.Day,
			Carrier.Hour,
			Carrier.LastUpdate,
		}
		table.Append(s)
	}
	table.Render()
}

