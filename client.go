package jpush

import (
	"fmt"
	"encoding/base64"
	"runtime"
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

func (c *Client) getAuthorization() string {
	buf := []byte(c.AppKey + ":" + c.MasterSecret)
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString(buf))
}

func (c *Client) getUserAgent() string {
	return fmt.Sprintf("(%s) go/%s", runtime.GOOS, runtime.Version())
}
