package jpush

import (
	"strings"
	"encoding/json"
	"errors"
	"bytes"
)

func (c *Client) ReportReceived(msgIds []string) ([]interface{}, error) {
	if len(msgIds) == 0 {
		return nil, errors.New("msgIds不能为空")
	}
	link := c.reportUrl + "/v3/received?msg_ids=" + strings.Join(msgIds, ",")
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Array()
}

func (c *Client) ReportStatusMessage(req *ReportStatusRequest) (map[string]interface{}, error) {
	link := c.reportUrl + "/v3/status/message"
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := c.request("POST", link, bytes.NewReader(buf), false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}
