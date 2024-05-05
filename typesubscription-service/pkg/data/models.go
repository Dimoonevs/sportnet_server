package data

type SchedulerData struct {
	Id         int32
	Time       []string
	CronId     int32
	DaysOfWeek []int
	TimeZone   string
}
type TypeSub int32

const (
	FIXED_COUNT TypeSub = iota
	TIME_LIMITED
)

type TimeLimited int32

const (
	WEEK TimeLimited = iota
	MONTH
	YEAR
	CUSTOM
)

type SubscriptionResponse struct {
	Data string
}

type SubscriptionData struct {
	Id                      int32
	CoachId                 int32
	Name                    string
	Description             string
	TypeSub                 int32
	TimeLimited             int32
	CustomTimeLimited       int32
	Price                   int32
	Currency                string
	DaysOfWeek              []string
	AutomaticallyManagement bool
}
