package api

type Task struct {
	UUID        string    `json:"uuid"`
	Owner       string    `json:"owner"`
	Op          string    `json:"type"`
	JobUUID     string    `json:"job_uuid"`
	ArchiveUUID string    `json:"archive_uuid"`
	Status      string    `json:"status"`
	StartedAt   Timestamp `json:"started_at"`
	StoppedAt   Timestamp `json:"stopped_at"`
	Log         string    `json:"log"`
}

type TaskFilter struct {
	Status string
	Debug  YesNo
}

func FetchListTasks(status string, debugFlag bool) ([]Task, error) {
	// FIXME: legacy
	return GetTasks(TaskFilter{
		Status: status,
		Debug:  Maybe(debugFlag),
	})
}

func GetTasks(filter TaskFilter) ([]Task, error) {
	uri := ShieldURI("/v1/tasks")
	uri.MaybeAddParameter("status", filter.Status)
	uri.MaybeAddParameter("debug", filter.Debug)

	var data []Task
	return data, uri.Get(&data)
}

func GetTask(uuid string) (Task, error) {
	var data Task
	return data, ShieldURI("v1/task/%s", uuid).Get(&data)
}

func CancelTask(uuid string) error {
	return ShieldURI("/v1/task/%s", uuid).Delete(nil)
}