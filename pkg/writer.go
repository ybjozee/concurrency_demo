package pkg

import (
	"fmt"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"github.com/xuri/excelize/v2"
	"log"
)

func WriteResults(accountDetails []twilioApi.ApiV2010Account, usageRecords []twilioApi.ApiV2010UsageRecordAllTime, messageRecords []twilioApi.ApiV2010Message, callRecords []twilioApi.ApiV2010Call) {
	file := excelize.NewFile()
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	writeAccountDetails(accountDetails, file)
	writeUsageRecords(usageRecords, file)
	writeMessageRecords(messageRecords, file)
	writeCallRecords(callRecords, file)
	err := file.SaveAs("Usage Report.xlsx")
	checkError(err)
}

func writeAccountDetails(accountDetails []twilioApi.ApiV2010Account, file *excelize.File) {
	sheetName := "Accounts"
	index, err := file.NewSheet("Sheet1")
	checkError(err)
	file.SetActiveSheet(index)
	err = file.SetSheetName("Sheet1", sheetName)
	checkError(err)

	writeToCell(file, sheetName, "A1", "S/N")
	writeToCell(file, sheetName, "B1", "Friendly name")
	writeToCell(file, sheetName, "C1", "SID")
	writeToCell(file, sheetName, "D1", "Date created")
	writeToCell(file, sheetName, "E1", "Status")
	writeToCell(file, sheetName, "F1", "Type")

	writeThickBorder(file, sheetName, "A1", "F1")

	for i, account := range accountDetails {
		//A is reserved for header so start from 2
		row := i + 2
		writeToCell(file, sheetName, fmt.Sprintf("A%d", row), i+1)
		writeToCell(file, sheetName, fmt.Sprintf("B%d", row), *account.FriendlyName)
		writeToCell(file, sheetName, fmt.Sprintf("C%d", row), *account.Sid)
		writeToCell(file, sheetName, fmt.Sprintf("D%d", row), *account.DateCreated)
		writeToCell(file, sheetName, fmt.Sprintf("E%d", row), *account.Status)
		writeToCell(file, sheetName, fmt.Sprintf("F%d", row), *account.Type)
		writeThinBorder(file, sheetName, fmt.Sprintf("A%d", row), fmt.Sprintf("F%d", row))
	}

	setCellSize(file, sheetName, "A", "A", 10)
	setCellSize(file, sheetName, "B", "D", 40)
	setCellSize(file, sheetName, "E", "F", 15)
}

func writeUsageRecords(usageRecords []twilioApi.ApiV2010UsageRecordAllTime, file *excelize.File) {

	sheetName := "All time usage"
	index, err := file.NewSheet(sheetName)
	checkError(err)
	file.SetActiveSheet(index)

	writeToCell(file, sheetName, "A1", "S/N")
	writeToCell(file, sheetName, "B1", "Account SID")
	writeToCell(file, sheetName, "C1", "Category")
	writeToCell(file, sheetName, "D1", "Description")
	writeToCell(file, sheetName, "E1", "Usage unit")
	writeToCell(file, sheetName, "F1", "Usage")
	writeToCell(file, sheetName, "G1", "Price")

	writeThickBorder(file, sheetName, "A1", "G1")

	for i, record := range usageRecords {
		//Row 1 is reserved for header so start from 2
		row := i + 2
		writeToCell(file, sheetName, fmt.Sprintf("A%d", row), i+1)
		writeToCell(file, sheetName, fmt.Sprintf("B%d", row), *record.AccountSid)
		writeToCell(file, sheetName, fmt.Sprintf("C%d", row), *record.Category)
		writeToCell(file, sheetName, fmt.Sprintf("D%d", row), *record.Description)
		if record.UsageUnit != nil {
			writeToCell(file, sheetName, fmt.Sprintf("E%d", row), *record.UsageUnit)
		}
		writeToCell(file, sheetName, fmt.Sprintf("F%d", row), *record.Usage)
		writeToCell(file, sheetName, fmt.Sprintf("G%d", row), *record.Price)
		writeThinBorder(file, sheetName, fmt.Sprintf("A%d", row), fmt.Sprintf("G%d", row))
	}

	setCellSize(file, sheetName, "A", "A", 10)
	setCellSize(file, sheetName, "B", "D", 60)
	setCellSize(file, sheetName, "E", "E", 20)
	setCellSize(file, sheetName, "F", "G", 10)
}

func writeMessageRecords(messageRecords []twilioApi.ApiV2010Message, file *excelize.File) {

	sheetName := "Messages"
	index, err := file.NewSheet(sheetName)
	checkError(err)
	file.SetActiveSheet(index)

	//write header
	writeToCell(file, sheetName, "A1", "S/N")
	writeToCell(file, sheetName, "B1", "SID")
	writeToCell(file, sheetName, "C1", "Account SID")
	writeToCell(file, sheetName, "D1", "Date created")
	writeToCell(file, sheetName, "E1", "Date sent")
	writeToCell(file, sheetName, "F1", "From")
	writeToCell(file, sheetName, "G1", "To")
	writeToCell(file, sheetName, "H1", "Status")
	writeToCell(file, sheetName, "I1", "Segments")
	writeToCell(file, sheetName, "J1", "Media")
	writeToCell(file, sheetName, "K1", "Price")

	writeThickBorder(file, sheetName, "A1", "K1")

	for i, record := range messageRecords {
		row := i + 2
		writeToCell(file, sheetName, fmt.Sprintf("A%d", row), i+1)
		writeToCell(file, sheetName, fmt.Sprintf("B%d", row), *record.Sid)
		writeToCell(file, sheetName, fmt.Sprintf("C%d", row), *record.AccountSid)
		writeToCell(file, sheetName, fmt.Sprintf("D%d", row), *record.DateCreated)
		writeToCell(file, sheetName, fmt.Sprintf("E%d", row), *record.DateSent)
		writeToCell(file, sheetName, fmt.Sprintf("F%d", row), *record.From)
		writeToCell(file, sheetName, fmt.Sprintf("G%d", row), *record.To)
		writeToCell(file, sheetName, fmt.Sprintf("H%d", row), *record.Status)
		writeToCell(file, sheetName, fmt.Sprintf("I%d", row), *record.NumSegments)
		writeToCell(file, sheetName, fmt.Sprintf("J%d", row), *record.NumMedia)
		if record.Price != nil {
			writeToCell(file, sheetName, fmt.Sprintf("K%d", row), *record.Price)
		}
		writeThinBorder(file, sheetName, fmt.Sprintf("A%d", row), fmt.Sprintf("K%d", row))

		setCellSize(file, sheetName, "A", "A", 10)
		setCellSize(file, sheetName, "B", "E", 50)
		setCellSize(file, sheetName, "F", "G", 30)
		setCellSize(file, sheetName, "H", "K", 15)
	}
}

func writeCallRecords(callRecords []twilioApi.ApiV2010Call, file *excelize.File) {

	sheetName := "Calls"
	index, err := file.NewSheet(sheetName)
	checkError(err)
	file.SetActiveSheet(index)

	//write header
	writeToCell(file, sheetName, "A1", "S/N")
	writeToCell(file, sheetName, "B1", "SID")
	writeToCell(file, sheetName, "C1", "Account SID")
	writeToCell(file, sheetName, "D1", "Date created")
	writeToCell(file, sheetName, "E1", "From")
	writeToCell(file, sheetName, "F1", "To")
	writeToCell(file, sheetName, "G1", "Status")
	writeToCell(file, sheetName, "H1", "Start time")
	writeToCell(file, sheetName, "I1", "End time")
	writeToCell(file, sheetName, "J1", "Price")

	writeThickBorder(file, sheetName, "A1", "J1")

	for i, record := range callRecords {
		row := i + 2
		writeToCell(file, sheetName, fmt.Sprintf("A%d", row), i+1)
		writeToCell(file, sheetName, fmt.Sprintf("B%d", row), *record.Sid)
		writeToCell(file, sheetName, fmt.Sprintf("C%d", row), *record.AccountSid)
		writeToCell(file, sheetName, fmt.Sprintf("D%d", row), *record.DateCreated)
		writeToCell(file, sheetName, fmt.Sprintf("E%d", row), *record.From)
		writeToCell(file, sheetName, fmt.Sprintf("F%d", row), *record.To)
		writeToCell(file, sheetName, fmt.Sprintf("G%d", row), *record.Status)
		writeToCell(file, sheetName, fmt.Sprintf("H%d", row), *record.StartTime)
		writeToCell(file, sheetName, fmt.Sprintf("I%d", row), *record.EndTime)
		if record.Price != nil {
			writeToCell(file, sheetName, fmt.Sprintf("J%d", row), *record.Price)
		}
		writeThinBorder(file, sheetName, fmt.Sprintf("A%d", row), fmt.Sprintf("J%d", row))
		setCellSize(file, sheetName, "A", "A", 10)
		setCellSize(file, sheetName, "B", "C", 60)
		setCellSize(file, sheetName, "D", "I", 30)
	}
}

func writeToCell(file *excelize.File, sheetName, cell string, value any) {
	err := file.SetCellValue(sheetName, cell, value)
	checkError(err)
}

func writeThinBorder(file *excelize.File, sheetName, start, end string) {
	style, err := file.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
	})
	checkError(err)
	err = file.SetCellStyle(sheetName, start, end, style)
	checkError(err)
}

func writeThickBorder(file *excelize.File, sheetName, start, end string) {
	style, err := file.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 2},
			{Type: "top", Color: "000000", Style: 2},
			{Type: "bottom", Color: "000000", Style: 2},
			{Type: "right", Color: "000000", Style: 2},
		},
	})
	checkError(err)
	err = file.SetCellStyle(sheetName, start, end, style)
	checkError(err)
}

func setCellSize(file *excelize.File, sheetName, start, end string, size float64) {
	err := file.SetColWidth(sheetName, start, end, size)
	checkError(err)
}
