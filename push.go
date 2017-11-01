package jpush

import (
	"net/http"
	"encoding/json"
	"bytes"
	"io/ioutil"
)

type Platform string

const (
	PlatformAndroid  Platform = "android"
	PlatformIOS      Platform = "ios"
	PlatformWinPhone Platform = "winphone"
)

type RequestAudience struct {
	Tag            []string `json:"tag,omitempty"`
	TagAnd         []string `json:"tag_and,omitempty"`
	TagNot         []string `json:"tag_not,omitempty"`
	Alias          []string `json:"alias,omitempty"`
	RegistrationId []string `json:"registration_id,omitempty"`
	Segment        []string `json:"segment,omitempty"`
	ABTest         []string `json:"abtest,omitempty"`
}

type PushNotification struct {
	Alert    string                `json:"alert,omitempty"`
	Android  *NotificationAndroid  `json:"android,omitempty"`
	IOS      *NotificationIOS      `json:"ios,omitempty"`
	WinPhone *NotificationWinPhone `json:"winphone,omitempty"`
}

type NotificationAndroid struct {
	Alert      string                 `json:"alert"`
	Title      string                 `json:"title,omitempty"`
	BuilderId  int                    `json:"builder_id,int,omitempty"`
	Priority   int                    `json:"priority,omitempty"`
	Category   string                 `json:"category,omitempty"`
	Style      int                    `json:"style,int,omitempty"`
	AlertType  int                    `json:"alert_type,int,omitempty"`
	BigText    string                 `json:"big_text,omitempty"`
	Inbox      map[string]interface{} `json:"inbox,omitempty"`
	BigPicPath string                 `json:"big_pic_path,omitempty"`
	Extras     map[string]interface{} `json:"extras,omitempty"`
}

type NotificationIOS struct {
	Alert            interface{}            `json:"alert"`
	Sound            string                 `json:"sound,omitempty"`
	Badge            int                    `json:"badge,int,omitempty"`
	ContentAvailable bool                   `json:"content-available,omitempty"`
	MutableContent   bool                   `json:"mutable-content,omitempty"`
	Category         string                 `json:"category,omitempty"`
	Extras           map[string]interface{} `json:"extras,omitempty"`
}

type NotificationWinPhone struct {
	Alert    string                 `json:"alert"`
	Title    string                 `json:"title,omitempty"`
	OpenPage string                 `json:"_open_page,omitempty"`
	Extras   map[string]interface{} `json:"extras,omitempty"`
}

type PushMessage struct {
	MsgContent  string                 `json:"msg_content"`
	Title       string                 `json:"title,omitempty"`
	ContentType string                 `json:"content_type,omitempty"`
	Extras      map[string]interface{} `json:"extras,omitempty"`
}

type SmsMessage struct {
	Content   string `json:"content"`
	DelayTime int    `json:"delay_time,int,omitempty"`
}

type PushOptions struct {
	SendNo          int    `json:"sendno,int,omitempty"`
	TimeToLive      int    `json:"time_to_live,int,omitempty"`
	OverrideMsgId   int64  `json:"override_msg_id,int64,omitempty"`
	ApnsProduction  bool   `json:"apns_production,omitempty"`
	ApnsCollapseId  string `json:"apns_collapse_id,omitempty"`
	BigPushDuration int    `json:"big_push_duration,int,omitempty"`
}

type PushRequest struct {
	Cid          string            `json:"cid"`
	Platform     Platform          `json:"platform"`
	Audience     *RequestAudience  `json:"audience"`
	Notification *PushNotification `json:"notification,omitempty"`
	Message      *PushMessage      `json:"message,omitempty"`
	SmsMessage   *SmsMessage       `json:"sms_message,omitempty"`
	Options      *PushOptions      `json:"options,omitempty"`
}

func (push *PushRequest) ExecuteWithClient(client *Client) (map[string]interface{}, error) {
	link := client.baseUrl + "/v3/push"
	buf, err := json.Marshal(push)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", link, bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", client.getAuthorization())
	req.Header.Set("User-Agent", client.getUserAgent())
	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buf, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result := make(map[string]interface{})
	err = json.Unmarshal(buf, &result)
	return result, err
}
