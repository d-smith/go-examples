package timer
import "time"

type BackendCall struct {
	Name string
	Time time.Duration
}

type Contributor struct {
	Name string
	Time time.Duration
	BackendCalls []BackendCall
}

type APITime struct {
	Name string
	Time time.Duration
	Contributors []Contributor
	Error error
	start time.Time
}

func NewAPITimer(name string) *APITime {
	return &APITime {
		Name: name,
		start: time.Now(),
	}
}

func (t *APITime) Stop(err error) {
	t.Time = time.Now().Sub(t.start)
	t.Error = err
}


