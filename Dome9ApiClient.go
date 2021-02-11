package dome9api

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	baseURL = "https://api.dome9.com/v2"
)

//Dome9 core API details
type Dome9 struct {
	BaseURL    string
	apiKey     string
	apiSecret  string
	HTTPClient *http.Client
}

// NewDome9 - Create a new API Client
func NewDome9(apiKey string, apiSecret string) *Dome9 {

	return &Dome9{
		BaseURL:    baseURL,
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		HTTPClient: &http.Client{Timeout: time.Minute},
	}

}

func (c *Dome9) sendRequest(req *http.Request, v interface{}) error {

	req.Header.Set("Authorization", "Basic "+basicAuth(c.apiKey, c.apiSecret))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	HTTPClient := &http.Client{Timeout: time.Minute}

	res, err := HTTPClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("HTTP Error, Status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}

func basicAuth(username string, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

//GetAWSAccounts - Get all AWS Accounts from Dome9
func (c *Dome9) GetAWSAccounts(ctx context.Context) (AWSCloudAccounts, error) {
	accounts := AWSCloudAccounts{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/cloudaccounts", c.BaseURL), nil)

	if err != nil {
		return accounts, err
	}

	req = req.WithContext(ctx)

	if err := c.sendRequest(req, &accounts); err != nil {
		return accounts, err
	}

	return accounts, nil
}

//GetAzureAccounts - Get all AWS Accounts from Dome9
func (c *Dome9) GetAzureAccounts(ctx context.Context) (AzureCloudAccounts, error) {
	accounts := AzureCloudAccounts{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/azurecloudaccount", c.BaseURL), nil)

	if err != nil {
		return accounts, err
	}

	req = req.WithContext(ctx)

	if err := c.sendRequest(req, &accounts); err != nil {
		return accounts, err
	}

	return accounts, nil
}

//GetProtectedAssets based on filter parameters
func (c *Dome9) GetProtectedAssets(ctx context.Context, pageSize int, includedEntityTypes []string, excludedEntityTypes []string) (SearchResult, error) {

	result := SearchResult{}
	payload := ProtectedAssetRequest{}

	if excludedEntityTypes != nil {
		payload.Filter.ExcludedEntityTypes = excludedEntityTypes
	}

	if includedEntityTypes != nil {
		payload.Filter.IncludedEntityTypes = includedEntityTypes
	}

	payload.PageSize = pageSize

	jsonPayload, err := json.Marshal(payload)

	if err != nil {
		return result, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/protected-asset/search", c.BaseURL), bytes.NewBuffer(jsonPayload))

	if err != nil {
		return result, err
	}

	req = req.WithContext(ctx)

	if err := c.sendRequest(req, &result); err != nil {
		return result, err
	}

	return result, nil
}

//GetBundleResults - Returns the results on a ruleset bundle assessment on an account
func (c *Dome9) GetBundleResults(ctx context.Context, bundleID int, cloudAccountids string, fromTime time.Time) (BundleResult, error) {
	testResults := BundleResult{}

	params := url.Values{}
	params.Add("bundleId", fmt.Sprintf("%d", bundleID))
	params.Add("cloudAccountIds", cloudAccountids)
	params.Add("fromTime", fromTime.Format(time.RFC3339))

	req, err := http.NewRequest("GET", fmt.Sprint(c.BaseURL+"/AssessmentHistoryV2/bundleResults?"+params.Encode()), nil)

	if err != nil {
		return testResults, err
	}

	req = req.WithContext(ctx)

	if err := c.sendRequest(req, &testResults); err != nil {
		return testResults, err
	}

	return testResults, nil

}

//GetAssessmentResults - Returns the assessment results for the given time period for the specified bundles and cloudaccounts - Not Working yet
func (c *Dome9) GetAssessmentResults(ctx context.Context, bundleIds []int, cloudAccountIDs []string, fromDate time.Time, toDate time.Time) (TestResults, error) {

	payload := BundleFilters{}
	testResults := TestResults{}
	filter := CloudAccountBundleFilters{}
	filters := []CloudAccountBundleFilters{}

	filter.BundleIds = bundleIds
	filter.CloudAccountIds = cloudAccountIDs
	filter.CloudAccountType = "aws"

	filters = append(filters, filter)

	payload.CloudAccountBundleFilters = filters
	payload.From = fromDate
	payload.To = toDate

	jsonPayload, err := json.Marshal(payload)

	fmt.Println(string(jsonPayload))

	if err != nil {
		return testResults, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/AssessmentHistoryV2/LastAssessmentResults", c.BaseURL), bytes.NewBuffer(jsonPayload))

	if err != nil {
		return testResults, err
	}

	req = req.WithContext(ctx)
	req.Close = true

	if err := c.sendRequest(req, &testResults); err != nil {
		return testResults, err
	}

	return testResults, nil
}
