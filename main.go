package main

import (
	sap_api_caller "sap-api-integrations-employment-information-reads/SAP_API_Caller"
	sap_api_input_reader "sap-api-integrations-employment-information-reads/SAP_API_Input_Reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Employment_Information_Header_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/successfactors/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"Header",
		}
	}

	caller.AsyncGetEmploymentInformation(
		inoutSDC.EmploymentInformation.PersonIDExternal,
		accepter,
	)
}
