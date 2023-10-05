package pkg

import (
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"log"
)

var TwilioClient *twilio.RestClient

func GetAccountDetails() []twilioApi.ApiV2010Account {
	parameters := &twilioApi.ListAccountParams{}
	parameters.SetPageSize(1000)
	accountDetails, err := TwilioClient.Api.ListAccount(parameters)
	checkError(err)
	return accountDetails
}

func GetUsageRecords() []twilioApi.ApiV2010UsageRecordAllTime {
	parameters := &twilioApi.ListUsageRecordAllTimeParams{}
	parameters.SetPageSize(1000)
	usageRecords, err := TwilioClient.Api.ListUsageRecordAllTime(parameters)
	checkError(err)
	return usageRecords
}

func GetMessageRecords() []twilioApi.ApiV2010Message {
	parameters := &twilioApi.ListMessageParams{}
	parameters.SetPageSize(1000)
	messageRecords, err := TwilioClient.Api.ListMessage(parameters)
	checkError(err)
	return messageRecords
}

func GetCallRecords() []twilioApi.ApiV2010Call {
	parameters := &twilioApi.ListCallParams{}
	parameters.SetPageSize(1000)
	callRecords, err := TwilioClient.Api.ListCall(parameters)
	checkError(err)
	return callRecords
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
