package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-employment-information-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetEmploymentInformation(personIDExternal string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(personIDExternal)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(personIDExternal string) {
	headerData, err := c.callEmploymentInformationSrvAPIRequirementHeader("EmpEmployment", personIDExternal)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(headerData)

	personNavData, err := c.callPersonNav(headerData[0].PersonNav)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(personNavData)
	
	personalInfoNavData, err := c.callPersonalInfoNav(personNavData.PersonalInfoNav)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(personalInfoNavData)


}

func (c *SAPAPICaller) callEmploymentInformationSrvAPIRequirementHeader(api, personIDExternal string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "odata/v2", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithHeader(req, personIDExternal)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callPersonNav(url string) (*sap_api_output_formatter.PersonNav, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPersonNav(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callPersonalInfoNav(url string) ([]sap_api_output_formatter.PersonalInfoNav, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPersonalInfoNav(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithHeader(req *http.Request, personIDExternal string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("personIdExternal eq '%s'", personIDExternal))
	req.URL.RawQuery = params.Encode()
}
