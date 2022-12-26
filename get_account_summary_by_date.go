package exante

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// https://api-demo.exante.eu/md/{version}/summary/{id}/{date}/{currency}

func (client *Client) GetAccountSummaryByDate(accountID, currency string, date time.Time) (AccountSummary, error) {
	client.refreshAccessToken()

	url := fmt.Sprintf("%s/md/3.0/summary/%s/%s/%s", client.serverAddr, accountID, date.Format("2006-01-02"), currency)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AccountSummary{}, err
	}
	req.WithContext(client.ctx)
	req.Header.Add("Authorization", strings.Join([]string{"Bearer", client.accessToken}, " "))
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return AccountSummary{}, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return AccountSummary{}, fmt.Errorf("cannot read response body: %w", err)
	}

	if resp.StatusCode > 399 {
		return AccountSummary{}, fmt.Errorf("bad http response code: %s: %s", resp.Status, string(data))
	}

	var summary AccountSummary
	err = json.Unmarshal(data, &summary)
	if err != nil {
		return AccountSummary{}, fmt.Errorf("cannot parse response: %w", err)
	}

	return summary, nil
}
