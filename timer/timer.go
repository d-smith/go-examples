package timer

import (
	"encoding/json"
	"sync"
	"time"
)

type ServiceCall struct {
	Name  string
	Time  time.Duration
	start time.Time
}

type Contributor struct {
	sync.Mutex
	Name         string
	Time         time.Duration
	Error        error
	start        time.Time
	ServiceCalls []*ServiceCall
}

type APITime struct {
	sync.Mutex
	Name         string
	Time         time.Duration
	Contributors []*Contributor
	ErrorFree    bool
	Error        error
	start        time.Time
}

func NewAPITimer(name string) *APITime {
	return &APITime{
		Name:  name,
		start: time.Now(),
	}
}

func (t *APITime) Stop(err error) {
	t.Time = time.Now().Sub(t.start)
	t.Error = err
	t.ErrorFree = len(t.ContributorErrors()) == 0 && t.Error == nil
}

func (t *APITime) StartContributor(name string) *Contributor {
	contributor := &Contributor{
		Name:  name,
		start: time.Now(),
	}

	t.Lock()
	t.Contributors = append(t.Contributors, contributor)
	t.Unlock()

	return contributor
}

func (t *APITime) ContributorErrors() []error {
	var errs []error
	for _, c := range t.Contributors {
		if c.Error != nil {
			errs = append(errs, c.Error)
		}
	}
	return errs
}

func (t *APITime) ToJSONString() string {
	s, err := json.Marshal(t)
	if err != nil {
		s = []byte("{}")
	}
	return string(s)
}

func (c *Contributor) End(err error) {
	c.Time = time.Now().Sub(c.start)
	c.Error = err
}

func (c *Contributor) StartServiceCall(name string) *ServiceCall {
	svcCall := &ServiceCall{
		start: time.Now(),
		Name:  name,
	}

	c.Lock()
	c.ServiceCalls = append(c.ServiceCalls, svcCall)
	c.Unlock()

	return svcCall
}

func (sc *ServiceCall) End() {
	sc.Time = time.Now().Sub(sc.start)
}
