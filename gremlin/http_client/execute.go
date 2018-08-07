package http_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/creativesoftwarefdn/weaviate/gremlin"
)

type gremlin_http_query struct {
	Gremlin string `json:"gremlin"`
}

type gremlinResponseStatus struct {
	Message string `json:"string"`
	Code    int    `json:"code"`
}

type gremlinResponseResult struct {
	Data []interface{} `json:"data"`
	Meta interface{}   `json:"meta"`
}

type gremlinResponse struct {
	Status gremlinResponseStatus `json:"status"`
	Result gremlinResponseResult `json:"result"`
}

func (c *Client) Execute(query *gremlin.Query) (*gremlin.Response, error) {
	log := c.logger.WithField("query", query.Query())
	log.Debugf("Sending query")

	q := gremlin_http_query{
		Gremlin: query.Query(),
	}

	json_bytes, err := json.Marshal(&q)
	if err != nil {
		log.Errorf("Could not create query, because %v", err)
		return nil, fmt.Errorf("Could not create query because %v", err)
	}

	bytes_reader := bytes.NewReader(json_bytes)
	req, err := http.NewRequest("POST", c.endpoint, bytes_reader)
	if err != nil {
		log.Errorf("Could not create HTTP request, because %v", err)
		return nil, fmt.Errorf("Could not create HTTP request to resolve a Gremlin query; %v", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	http_response, err := c.client.Do(req)

	defer http_response.Body.Close()

	buf, err := ioutil.ReadAll(http_response.Body)
	var response_data gremlinResponse
	json.Unmarshal(buf, &response_data)

	log.WithField("status_code", http_response.StatusCode).WithField("response", string(buf)).Debugf("Received reply")
	switch http_response.StatusCode {
	case 200:
		client_response := gremlin.Response{Data: response_data.Result.Data}
		return &client_response, nil
	case 500:
		return nil, fmt.Errorf("Server error: %s", string(buf))
	default:
		return nil, fmt.Errorf("Unexpected status code %v", http_response.StatusCode)
	}

	// should not reach this; default case in switch should handle everything.
	return nil, nil
}