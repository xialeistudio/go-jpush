package jpush

import (
	"fmt"
	"encoding/base64"
	"runtime"
	"encoding/json"
	"net/http"
	"bytes"
	"io/ioutil"
	"net/url"
	"strings"
	"io"
)

type Client struct {
	AppKey       string
	MasterSecret string
	baseUrl      string
}

func NewClient(appKey, masterSecret string) *Client {
	client := &Client{AppKey: appKey, MasterSecret: masterSecret}
	client.baseUrl = "https://api.jpush.cn"
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

func (c *Client) request(method, link string, body io.Reader, isGroup bool) (map[string]interface{}, error) {
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
	result := make(map[string]interface{})
	err = json.Unmarshal(buf, &result)
	return result, err
}
func (c *Client) Push(push *PushRequest) (map[string]interface{}, error) {
	link := c.baseUrl + "/v3/push"
	buf, err := json.Marshal(push)
	if err != nil {
		return nil, err
	}
	return c.request("POST", link, bytes.NewReader(buf), false)
}

func (c *Client) GetCidPool(count int, cidType string) (map[string]interface{}, error) {
	params := url.Values{}
	if count > 0 {
		params["count"] = []string{"0"}
	}
	if cidType != "" {
		params["type"] = []string{cidType}
	}
	link := c.baseUrl + "/v3/push/cid"

	return c.request("GET", link, strings.NewReader(params.Encode()), false)
}

func (c *Client) GroupPush(push *PushRequest) (map[string]interface{}, error) {
	link := c.baseUrl + "/v3/grouppush"
	buf, err := json.Marshal(push)
	if err != nil {
		return nil, err
	}
	return c.request("POST", link, bytes.NewReader(buf), true)
}

func (c *Client) Validate(req *PushRequest) (map[string]interface{}, error) {
	link := c.baseUrl + "/v3/push/validate"
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	return c.request("POST", link, bytes.NewReader(buf), false)
}
