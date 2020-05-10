package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jeffalyanak/check_namecheap/logger"
	"github.com/jeffalyanak/check_namecheap/model"
)

func main() {
	logger, err := logger.Get()
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	// Struct for holding data
	var d model.ApiResponse

	// Handle cli arguments
	domain := flag.String("domain", "", "domain to check")
	apiuser := flag.String("apiuser", "", "API Username")
	username := flag.String("username", "", "Username")
	key := flag.String("key", "", "API Key")
	clientip := flag.String("clientip", "", "Client IP")

	warn := flag.Int64("warn", 15, "days until warning")
	crit := flag.Int64("crit", 7, "days until critical")

	// Create warn and crit durations
	warning := time.Duration(int64(time.Hour) * int64(24**warn))
	critical := time.Duration(int64(time.Hour) * int64(24**crit))

	flag.Parse()

	if *domain == "" {
		logger.Println("No domain provided")
		fmt.Println("No domain provided")
		os.Exit(3)
	}
	if *apiuser == "" {
		logger.Println("No API User provided")
		fmt.Println("No API User provided")
		os.Exit(3)
	}
	if *username == "" {
		logger.Println("No Username provided")
		fmt.Println("No Username provided")
		os.Exit(3)
	}
	if *domain == "" {
		logger.Println("No domain provided")
		fmt.Println("No domain provided")
		os.Exit(3)
	}
	if *key == "" {
		logger.Println("No API key provided")
		fmt.Println("No API key provided")
		os.Exit(3)
	}
	if *clientip == "" {
		logger.Println("No Client IP provided")
		fmt.Println("No Client IP provided")
		os.Exit(3)
	}

	// Build strings for request
	apicall := "https://api.namecheap.com/xml.response?ApiUser=" +
		*apiuser +
		"&ApiKey=" +
		*key +
		"&UserName=" +
		*username +
		"&Command=namecheap.domains.getinfo&ClientIp=" +
		*clientip +
		"&DomainName=" +
		*domain

	// Build request
	client := &http.Client{}
	req, _ := http.NewRequest("GET", apicall, nil)

	// Make Request
	resp, err := client.Do(req)
	if err != nil {
		logger.Println("Error!")
		fmt.Println("Error!")
		os.Exit(3)
	}
	defer resp.Body.Close()

	// Check for rate limiting
	if resp.StatusCode == 429 {
		retry, _ := strconv.Atoi(resp.Header.Get("Retry-After"))
		if err != nil {
			logger.Println(err)
			fmt.Println(err)
			os.Exit(3)
		}

		// Wait for a bit
		delay := time.Duration(int64(time.Second)*int64(retry) + 1)
		time.Sleep(time.Duration(delay))

		logger.Println("Rate limit reached, waiting " + strconv.Itoa(retry) + "s")
	}

	// Marshal json data into struct
	body, err := ioutil.ReadAll(resp.Body)
	if err := xml.Unmarshal(body, &d); err != nil {
		logger.Println(err)
		fmt.Println(err)
		os.Exit(3)
	}

	// Attempt to parse the time from string
	expiration, err := time.Parse("01/02/2006", d.CommandResponse.DomainGetInfoResult.DomainDetails.ExpiredDate)
	if err != nil {
		logger.Println(err)
		fmt.Println(err)
		os.Exit(3)
	}

	// Differential between now and expiry
	diff := expiration.Sub(time.Now())

	// Exit status and string
	exitstatus := 0
	exitstring := ""

	// Determine status
	if diff < 0 {
		exitstatus = 2
		exitstring += "CRITICAL - [" + *domain + "] Expired "
	} else if diff < warning {
		exitstatus = 2
		exitstring += "CRITICAL - [" + *domain + "] Expires "
	} else if diff < critical {
		exitstatus = 1
		exitstring += "WARNING - [" + *domain + "] Expires "
	} else {
		exitstring += "OK - [" + *domain + "] Expires "
	}

	exitstring += "in " + durationDays(diff) + ", at " + expiration.String()

	logger.Println(exitstring)
	fmt.Println(exitstring)
	os.Exit(exitstatus)
}

func durationDays(diff time.Duration) string {
	if float64(diff) < 86400000000000 {
		return durationHours(diff)
	}
	return strconv.FormatFloat(float64(diff)/86400000000000, 'f', 0, 64) + " day(s)"
}

func durationHours(diff time.Duration) string {
	return strconv.FormatFloat(float64(diff)/3600000000000, 'f', 0, 64) + " hours(s)"
}
