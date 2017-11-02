package jpush

const (
	ScheduleTimeUnitDay   = "day"
	ScheduleTimeUnitWeek  = "week"
	ScheduleTimeUnitMonth = "month"
)

type ScheduleRequest struct {
	Cid     string           `json:"cid,omitempty"`
	Name    string           `json:"name"`
	Enabled bool             `json:"enabled"`
	Push    *PushRequest     `json:"push"`
	Trigger *ScheduleTrigger `json:"trigger"`
}

type ScheduleTrigger struct {
	Single     *ScheduleTriggerSingle     `json:"single,omitempty"`
	Periodical *ScheduleTriggerPeriodical `json:"periodical,omitempty"`
}

type ScheduleTriggerSingle struct {
	Timer string `json:"time,omitempty"`
}

type ScheduleTriggerPeriodical struct {
	Start     string      `json:"start,omitempty"`
	End       string      `json:"end,omitempty"`
	Time      string      `json:"time,omitempty"`
	TimeUnit  string      `json:"time_unit,omitempty"`
	Frequency int         `json:"frequency,int,omitempty"`
	Point     interface{} `json:"point,omitempty"`
}
