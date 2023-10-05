package main

import (
	"app/pkg"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	setup()
	fmt.Println("Starting synchronous execution")
	start := time.Now()
	generateReportSynchronously()
	fmt.Printf("Synchronous Execution Time: %s\n", time.Since(start).String())

	fmt.Println("Starting concurrent execution")
	start = time.Now()
	generateReportConcurrently()
	fmt.Printf("Concurrent Execution Time: %s\n", time.Since(start).String())
}

func generateReportSynchronously() {

	accountDetails := pkg.GetAccountDetails()
	usageRecords := pkg.GetUsageRecords()
	messageRecords := pkg.GetMessageRecords()
	callRecords := pkg.GetCallRecords()
	pkg.WriteResults(accountDetails, usageRecords, messageRecords, callRecords)
}

func generateReportConcurrently() {

	var results = struct {
		AccountDetails []twilioApi.ApiV2010Account
		UsageRecords   []twilioApi.ApiV2010UsageRecordAllTime
		MessageRecords []twilioApi.ApiV2010Message
		CallRecords    []twilioApi.ApiV2010Call
	}{}

	wg := sync.WaitGroup{}
	wg.Add(4)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		results.AccountDetails = pkg.GetAccountDetails()
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		results.UsageRecords = pkg.GetUsageRecords()
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		results.MessageRecords = pkg.GetMessageRecords()
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		results.CallRecords = pkg.GetCallRecords()
	}(&wg)

	wg.Wait()

	pkg.WriteResults(results.AccountDetails, results.UsageRecords, results.MessageRecords, results.CallRecords)
}

func setup() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	pkg.TwilioClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
}
