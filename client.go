package jpush

import (
	"fmt"
	"encoding/base64"
	"runtime"
	"net/http"
	"io/ioutil"
	"io"
)

type Client struct {
	AppKey       string
	MasterSecret string
	pushUrl      string
	reportUrl    string
	deviceUrl    string
}

func NewClient(appKey, masterSecret string) *Client {
	client := &Client{AppKey: appKey, MasterSecret: masterSecret}
	client.pushUrl = "https://api.jpush.cn"
	client.reportUrl = "https://report.jpush.cn"
	client.deviceUrl = "https://device.jpush.cn"
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
