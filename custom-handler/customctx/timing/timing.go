package timing
import (
	"time"
	"golang.org/x/net/context"
	"net/http"
	"encoding/hex"
	"crypto/rand"
	"log"
	cc "github.com/d-smith/go-examples/custom-handler/customctx"
	"encoding/json"
)

func genCorrID()(string) {
	cid := make([]byte, 16)
	_,err := rand.Read(cid)
	if err != nil {
		log.Println("WARNING - error creating a cid using rand - will degrade to timestamp")
		return time.Now().Format(time.RFC3339Nano)
	}

	return hex.EncodeToString(cid)

}

type Timer struct {
	Name string
	StartTime time.Time
	EndTime time.Time
	Error string
}

type Timings struct {
	CorrId string
	RequestURI string
	Timers map[string]Timer
}

type  TimingContributor struct {
	Name string
	Duration time.Duration
}

type TimingSummary struct {
	CorrId string
	RequestURI string
	EndToEndDuration time.Duration
	Contributors []TimingContributor
}

func NewTimingSummary(endToEnd time.Duration, t *Timings) *TimingSummary {
	summary := &TimingSummary{
		CorrId: t.CorrId,
		RequestURI: t.RequestURI,
		EndToEndDuration: endToEnd,
		Contributors: make([]TimingContributor,0),
	}

	for _,v := range t.Timers {
		timingContributor := TimingContributor{
			Name: v.Name,
			Duration: v.EndTime.Sub(v.StartTime),
		}
		summary.Contributors = append(summary.Contributors, timingContributor)
	}

	return summary
}

func (t *Timings) DumpTimings(duration time.Duration) ([]byte,error) {
	return json.Marshal(NewTimingSummary(duration,t))
}

type key int

const timeKey = 1

func NewContextWithTimings(ctx context.Context, req *http.Request) context.Context {
	timings := Timings {
		CorrId: genCorrID(),
		RequestURI: req.RequestURI,
		Timers: make(map[string]Timer),
	}
	return context.WithValue(ctx, timeKey, timings)
}

func TimingsFromContext(ctx context.Context) Timings {
	return ctx.Value(timeKey).(Timings)
}

func TimerMiddleware(h cc.ContextHandler) cc.ContextHandler {
	return cc.ContextHandlerFunc(func(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
		ctx = NewContextWithTimings(ctx, req)
		start := time.Now()


		println("-->Timer mgmt serve http")
		h.ServeHTTPContext(ctx, rw, req)
		println("<--Timer mgmt http served")

		stop := time.Now()
		timings := ctx.Value(timeKey).(Timings)

		timingStr, err := timings.DumpTimings(stop.Sub(start))
		if err != nil {
			log.Println("Unable to dump timings", err.Error())
		} else {
			log.Println(string(timingStr))
		}

	})
}
