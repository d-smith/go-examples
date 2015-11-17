package timing

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	cc "github.com/d-smith/go-examples/custom-handler/customctx"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"time"
)

func genCorrID() string {
	cid := make([]byte, 16)
	_, err := rand.Read(cid)
	if err != nil {
		log.Println("WARNING - error creating a cid using rand - will degrade to timestamp")
		return time.Now().Format(time.RFC3339Nano)
	}

	return hex.EncodeToString(cid)

}

type timer struct {
	Name      string
	StartTime time.Time
	EndTime   time.Time
	Error     string
}

type timings struct {
	CorrId     string
	RequestURI string
	Timers     []*timer
}

type timingContributor struct {
	Name     string
	Duration time.Duration
}

type timingSummary struct {
	CorrId           string
	RequestURI       string
	EndToEndDuration time.Duration
	Contributors     []timingContributor
}

func newTimingSummary(endToEnd time.Duration, t *timings) *timingSummary {
	summary := &timingSummary{
		CorrId:           t.CorrId,
		RequestURI:       t.RequestURI,
		EndToEndDuration: endToEnd,
		Contributors:     make([]timingContributor, 0),
	}

	for _, v := range t.Timers {
		timingContributor := timingContributor{
			Name:     v.Name,
			Duration: v.EndTime.Sub(v.StartTime),
		}
		summary.Contributors = append(summary.Contributors, timingContributor)
	}

	return summary
}

func (t *timings) dumpTimings(duration time.Duration) ([]byte, error) {
	return json.Marshal(newTimingSummary(duration, t))
}

type key int

const timeKey = 1

func newContextWithTimings(ctx context.Context, req *http.Request) context.Context {
	timings := timings{
		CorrId:     genCorrID(),
		RequestURI: req.RequestURI,
		Timers:     make([]*timer, 0),
	}
	return context.WithValue(ctx, timeKey, &timings)
}

func timingsFromContext(ctx context.Context) *timings {
	return ctx.Value(timeKey).(*timings)
}

func TimerMiddleware(h cc.ContextHandler) cc.ContextHandler {
	return cc.ContextHandlerFunc(func(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
		ctx = newContextWithTimings(ctx, req)
		start := time.Now()

		println("-->Timer mgmt serve http")
		h.ServeHTTPContext(ctx, rw, req)
		println("<--Timer mgmt http served")

		stop := time.Now()
		timings := ctx.Value(timeKey).(*timings)

		timingStr, err := timings.dumpTimings(stop.Sub(start))
		if err != nil {
			log.Println("Unable to dump timings", err.Error())
		} else {
			log.Println(string(timingStr))
		}

	})
}

func StartTimer(ctx context.Context, timerName string) {
	now := time.Now()
	timer := timer{
		Name:      timerName,
		StartTime: now,
		EndTime:   now,
	}

	timings := timingsFromContext(ctx)
	timings.Timers = append(timings.Timers, &timer)
	log.Println("timings post start timer", timings)
}

func grabTiming(timings []*timer, timingName string) (*timer, bool) {
	for _, v := range timings {
		if v.Name == timingName {
			return v, true
		}
	}

	return nil, false
}

func EndTimer(ctx context.Context, timerName string, err error) {
	timings := timingsFromContext(ctx)
	log.Println(timings.Timers)
	timer, ok := grabTiming(timings.Timers, timerName)
	if !ok {
		log.Println("WARNING: no start timing for ", timerName)
		return
	}

	timer.EndTime = time.Now()
	if err != nil {
		timer.Error = err.Error()
	}
}
