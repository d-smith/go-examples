package service

import (
	"github.com/d-smith/go-examples/ctx-hierarchy/wrappers"
	"golang.org/x/net/context"
	"net/http"
)

//Yes - this is inspired by the bro app on silicon valley

func pretendServiceCall() {

}

func NewBroHandler() func(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
		//Grab the context
		timer := wrappers.TimerFromContext(ctx)

		//Make sure you have a timer
		if timer == nil {
			http.Error(rw, "System error - no timer in context", http.StatusInternalServerError)
			return
		}

		//It's likely there's been several layers of middleware wrapped around
		//the start of the service code. Set the service name on the top level timer
		//once we're in the right place to do so.
		timer.Name = "bro"

		//Start the contributor  timing for the service implementation
		pluginTime := timer.StartContributor("bro service plugin")

		//Record calls to other services as contributors
		svcTime := pluginTime.StartServiceCall("bro backend")
		pretendServiceCall()
		svcTime.End()

		//End the contributor call, noting any errors that might have happened
		pluginTime.End(nil)

		rw.Write([]byte("Bro\n"))
	}
}