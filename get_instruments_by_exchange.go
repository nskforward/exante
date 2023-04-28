package exante_http

import (
	"fmt"
	"io"
	"net/http"
)

func (client *Client) GetInstrumentsByExchange(exchangeID string) ([]byte, error) {
	url := fmt.Sprintf("%s/md/3.0/exchanges/%s", client.serverAddr, exchangeID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.executeHTTPRequest(req)
	if err != nil {
		return nil, err
	}
	defer client.closeResponse(resp.Body)
	return io.ReadAll(resp.Body)
}
