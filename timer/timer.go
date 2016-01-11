package timer

import "time"

type BackendCall struct {
	Name string
	Time time.Duration
}

type Contributor struct {
	Name         string
	Time         time.Duration
	Error        error
	start        time.Time
	BackendCalls []BackendCall
}

type APITime struct {
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

	t.Contributors = append(t.Contributors, contributor)

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

func (c *Contributor) End(err error) {
	c.Time = time.Now().Sub(c.start)
	c.Error = err
}
