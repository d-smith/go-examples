package timing
import (
	"time"
	"golang.org/x/net/context"
	"net/http"
	"encoding/hex"
	"crypto/rand"
	"log"
	cc "github.com/d-smith/go-examples/custom-handler/customctx"
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
	CorrID string
	Name string
	StartTime time.Time
	EndTime time.Time
	Error string
}

type Timings struct {
	CorrId string
	Timers map[string]Timer
}

type key int

const timeKey = 1

func NewContextWithTimings(ctx context.Context, req *http.Request) context.Context {
	timings := Timings {
		CorrId: genCorrID(),
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
		println("-->Timer mgmt serve http")
		h.ServeHTTPContext(ctx, rw, req)
		println("<--Timer mgmt http served")
	})
}
