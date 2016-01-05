package mapctx

import (
	cc "github.com/d-smith/go-examples/custom-handler/customctx"
)



func GenWrapperFunc(ctxmap map[string]string) func(cc.ContextHandler)cc.ContextHandler {

}