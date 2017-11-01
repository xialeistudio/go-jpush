package jpush

import (
	"testing"
	"os"
)

var client = NewClient(os.Getenv("APP_KEY"), os.Getenv("MASTER_SECRET"))
var registrationId = os.Getenv("REGISTRATION_ID")

func getMsg() *PushRequest {
	params := make(map[string]interface{})
	params["url"] = "https://www.jpush.cn"
	req := &PushRequest{
		Cid:      "f005f301ce83fa605a832fb2-56773c24-f933-44cb-aadc-0e154be7fcbb",
		Platform: PlatformAndroid,
		Audience: &PushAudience{
			RegistrationId: []string{registrationId},
		},
		Notification: &PushNotification{
			Android: &NotificationAndroid{
				Alert:     "alert",
				Title:     "title",
				BuilderId: 1,
				Priority:  1,
				AlertType: 7,
				Extras:    params,
			},
		},
		Options: &PushOptions{
			TimeToLive:     60,
			ApnsCollapseId: "jiguang_test_201706011100",
		},
	}
	return req
}

func TestPushRequestExecuteWithClient(t *testing.T) {
	t.Logf("test with APP_KEY: %s, MASTER_SECRET: %s", client.AppKey, client.MasterSecret)
	req := getMsg()
	result, err := client.Push(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientGetCidPool(t *testing.T) {
	data, err := client.GetCidPool(0, "")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(data)
}

func TestClientDoPushGroup(t *testing.T) {
	result, err := client.GroupPush(getMsg())
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientValidate(t *testing.T) {
	result, err := client.Validate(getMsg())
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}
