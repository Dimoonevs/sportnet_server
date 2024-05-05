package data

type CronData struct {
	Id           int32
	CronId       int32
	DaysOfWeek   []int
	TimeZone     string
	TimeTraining []string
}
