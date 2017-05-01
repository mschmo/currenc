package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

type FixerApiResponse struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

func catchErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	fromPtr := flag.String("from", "USD", "Currency base to be quoted against.")
	toPtr := flag.String("to", "EUR", "Currency to convert to.")
	amountPtr := flag.Float64("amount", 1.0, "Amount of base currency to convert.")
	flag.Parse()

	if *fromPtr == *toPtr {
		fmt.Println("Base and To currencies are the same. No exchange.")
		return
	}

	apiBase := "http://api.fixer.io"
	resp, err := http.Get(apiBase + "/latest?base=" + *fromPtr)
	catchErr(err)

	defer resp.Body.Close()
	apiResponse := FixerApiResponse{}
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	catchErr(err)

	rates := apiResponse.Rates
	if toRate, ok := rates[*toPtr]; !ok {
		fmt.Printf("%s is not a valid currency.\n", *toPtr)
		fmt.Println("Valid currencies are:")
		for curr := range rates {
		    fmt.Println(curr)
		}
	} else {
		convertedAmount := *amountPtr * toRate
		fmt.Println("Exchange rate for", apiResponse.Date)
		fmt.Printf("%f %s = %f %s\n", *amountPtr, *fromPtr, convertedAmount, *toPtr)
	}
}
