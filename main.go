package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Carrier struct {
	Name         string
	Status       string
	ResponseTime string
	ThreeMinute  string
	Hour         string
	Day          string
	LastUpdate   string
}

func main() {

	s := spinner.New(spinner.CharSets[39], 100*time.Millisecond)
	s.Start()

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
					e.Status = (sel.Text())
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
				if index != 0 && (index+1)%7 == 0 {
					companies = append(companies, e)
				}
			})
		}
	})

	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Carrier", "Status", "Response Time", "<3min^", "<1h^", "<24h^", "Last Update"})
	for i, Carrier := range companies {
		switch i {
		case 0:
			Carrier.Name = "FedEx"
		case 1:
			Carrier.Name = "UPS"
		case 2:
			Carrier.Name = "USPS"
		case 3:
			Carrier.Name = "CanadaPost"
		}
		if Carrier.Status == "Online" {
			Carrier.Status = color.GreenString(Carrier.Status)
		} else {
			Carrier.Status = color.RedString(Carrier.Status)
		}
		s := []string{
			Carrier.Name,
			Carrier.Status,
			Carrier.ResponseTime,
			Carrier.ThreeMinute,
			strings.Split(Carrier.Day, "%")[0] + "%",
			strings.Split(Carrier.Hour, "%")[0] + "%",
			Carrier.LastUpdate,
		}
		table.Append(s)
	}
	s.Stop()
	table.Render()
}
