package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-employee-basic-data-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
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

func (c *SAPAPICaller) AsyncGetEmployeeBasicData(objectID, userID, employeeID string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "BusinessUserCollection":
			func() {
				c.BusinessUserCollection(objectID, userID)
				wg.Done()
			}()
		case "EmployeeBasicData":
			func() {
				c.EmployeeBasicData(objectID, userID, employeeID)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) BusinessUserCollection(objectID, userID string) {
	businessUserCollectionData, err := c.callEmployeeBasicDataSrvAPIRequirementBusinessUserCollection("BusinessUserCollection", objectID, userID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(businessUserCollectionData)

	businessUserBusinessRoleAssignmentData, err := c.callToBusinessUserBusinessRoleAssignment(businessUserCollectionData[0].ToBusinessUserBusinessRoleAssignment)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(businessUserBusinessRoleAssignmentData)

}

func (c *SAPAPICaller) callEmployeeBasicDataSrvAPIRequirementBusinessUserCollection(api, objectID, userID string) ([]sap_api_output_formatter.BusinessUserCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithBusinessUserCollection(req, objectID, userID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToBusinessUserCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToBusinessUserBusinessRoleAssignment(url string) ([]sap_api_output_formatter.ToBusinessUserBusinessRoleAssignment, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToBusinessUserBusinessRoleAssignment(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) EmployeeBasicData(objectID, userID, employeeID string) {
	employeeBasicDataData, err := c.callEmployeeBasicDataSrvAPIRequirementEmployeeBasicData("EmployeeBasicData", objectID, userID, employeeID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(employeeBasicDataData)
}

func (c *SAPAPICaller) callEmployeeBasicDataSrvAPIRequirementEmployeeBasicData(api, objectID, userID, employeeID string) ([]sap_api_output_formatter.EmployeeBasicData, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithEmployeeBasicData(req, objectID, userID, employeeID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToEmployeeBasicData(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithBusinessUserCollection(req *http.Request, objectID, userID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("ObjectID eq '%s' and UserID eq '%s'", objectID, userID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithEmployeeBasicData(req *http.Request, objectID, userID, employeeID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("ObjectID eq '%s' and UserID eq '%s' and and EmployeeID eq '%s'", objectID, userID, employeeID))
	req.URL.RawQuery = params.Encode()
}
