package jpush

import (
	"fmt"
	"encoding/base64"
	"runtime"
	"encoding/json"
	"net/http"
	"bytes"
	"io/ioutil"
	"io"
	"strconv"
	"errors"
	"strings"
)

type Client struct {
	AppKey       string
	MasterSecret string
	pushUrl      string
	reportUrl    string
}

func NewClient(appKey, masterSecret string) *Client {
	client := &Client{AppKey: appKey, MasterSecret: masterSecret}
	client.pushUrl = "https://api.jpush.cn"
	client.reportUrl = "https://report.jpush.cn"
	return client
}

func (c *Client) getAuthorization(isGroup bool) string {
	str := c.AppKey + ":" + c.MasterSecret
	if isGroup {
		str = "group-" + str
	}
	buf := []byte(str)
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString(buf))
}

func (c *Client) getUserAgent() string {
	return fmt.Sprintf("(%s) go/%s", runtime.GOOS, runtime.Version())
}

func (c *Client) request(method, link string, body io.Reader, isGroup bool) (*Response, error) {
	req, err := http.NewRequest(method, link, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", c.getAuthorization(isGroup))
	req.Header.Set("User-Agent", c.getUserAgent())
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &Response{data: buf}, nil
}
func (c *Client) Push(push *PushRequest) (map[string]interface{}, error) {
	link := c.pushUrl + "/v3/push"
	buf, err := json.Marshal(push)
	if err != nil {
		return nil, err
	}
	resp, err := c.request("POST", link, bytes.NewReader(buf), false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}

func (c *Client) GetCidPool(count int, cidType string) (map[string]interface{}, error) {
	link := c.pushUrl + "/v3/push/cid?"
	if count > 0 {
		link += "count=" + strconv.Itoa(count)
	}
	if cidType != "" {
		link += "type=" + cidType
	}
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}

func (c *Client) GroupPush(push *PushRequest) (map[string]interface{}, error) {
	link := c.pushUrl + "/v3/grouppush"
	buf, err := json.Marshal(push)
	if err != nil {
		return nil, err
	}
	resp, err := c.request("POST", link, bytes.NewReader(buf), true)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}

func (c *Client) Validate(req *PushRequest) (map[string]interface{}, error) {
	link := c.pushUrl + "/v3/push/validate"
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

