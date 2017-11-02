package jpush

type DeviceSettingRequest struct {
	Tags   *DeviceSettingRequestTags `json:"tags"`
	Alias  string                    `json:"alias"`
	Mobile string                    `json:"mobile"`
}
type DeviceSettingEmptyTagsRequest struct {
	Tags   string `json:"tags"`
	Alias  string `json:"alias"`
	Mobile string `json:"mobile"`
}
type DeviceSettingRequestTags struct {
	Add    []string `json:"add,omitempty"`
	Remove []string `json:"remove,omitempty"`
}
type DeviceBindTagsRequest struct {
	Add    []string `json:"add,omitempty"`
	Remove []string `json:"remove,omitempty"`
}
