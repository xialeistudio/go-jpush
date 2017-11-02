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
	data, err := client.GetCidPool(0, "push")
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

func TestClientReportReceived(t *testing.T) {
	msgId := "1345223734"
	result, err := client.ReportReceived([]string{msgId})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientReportStatusMessage(t *testing.T) {
	msgId := 1345223734
	result, err := client.ReportStatusMessage(&ReportStatusRequest{
		MsgId:           msgId,
		RegistrationIds: []string{registrationId},
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientDeviceView(t *testing.T) {
	result, err := client.DeviceView(registrationId)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientDeviceRequest(t *testing.T) {
	tags := &DeviceSettingRequestTags{
		Add: []string{"test"},
	}
	req := &DeviceSettingRequest{
		Alias:  "xialei",
		Mobile: "13333333333",
		Tags:   tags,
	}
	result, err := client.DeviceRequest(registrationId, req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(result))
}

func TestClientDeviceEmptyTagsRequest(t *testing.T) {
	req := &DeviceSettingEmptyTagsRequest{
		Alias:  "xialei",
		Mobile: "13333333333",
		Tags:   "",
	}
	result, err := client.DeviceEmptyTagsRequest(registrationId, req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(result))
}

func TestClientDeviceGetWithAlias(t *testing.T) {
	result, err := client.DeviceGetWithAlias("xialei", nil)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientDeviceDeleteAlias(t *testing.T) {
	result, err := client.DeviceDeleteAlias("xialei1")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientDeviceGetTags(t *testing.T) {
	result, err := client.DeviceGetTags()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientDeviceCheckDeviceWithTag(t *testing.T) {
	result, err := client.DeviceCheckDeviceWithTag("xialei", registrationId)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}
func TestClientDeviceBindTags(t *testing.T) {
	req := &DeviceBindTagsRequest{
		Add: []string{registrationId},
	}
	result, err := client.DeviceBindTags("xialei", req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientDeviceDeleteTag(t *testing.T) {
	result, err := client.DeviceDeleteTag("xialei", nil)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}
