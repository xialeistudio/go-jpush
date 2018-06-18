package jpush

import "encoding/json"

type Platform string

const (
	PlatformAndroid  Platform = "android"
	PlatformIOS      Platform = "ios"
	PlatformWinPhone Platform = "winphone"
)

type PushAudience struct {
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
	ApnsProduction  bool   `json:"apns_production"`
	ApnsCollapseId  string `json:"apns_collapse_id,omitempty"`
	BigPushDuration int    `json:"big_push_duration,int,omitempty"`
}

type PushRequest struct {
	Cid          string            `json:"cid,omitempty"`
	Platform     Platform          `json:"platform"`
	Audience     *PushAudience     `json:"audience,omitempty"`
	Notification *PushNotification `json:"notification,omitempty"`
	Message      *PushMessage      `json:"message,omitempty"`
	SmsMessage   *SmsMessage       `json:"sms_message,omitempty"`
	Options      *PushOptions      `json:"options,omitempty"`
}

type Response struct {
	data []byte
}

func (res *Response) Array() ([]interface{}, error) {
	list := make([]interface{}, 0)
	err := json.Unmarshal(res.data, &list)
	return list, err
}

func (res *Response) Map() (map[string]interface{}, error) {
	result := make(map[string]interface{})
	err := json.Unmarshal(res.data, &result)
	return result, err
}

func (res *Response) Bytes() []byte {
	return res.data
}
