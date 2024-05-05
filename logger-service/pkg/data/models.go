package data

type LogRequest struct {
	Name  string
	Data  any
	Field string
}

type LogsPayload struct {
	Name string
	Data any
	Time string
}
