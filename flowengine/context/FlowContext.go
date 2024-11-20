package context

import (
	"context"
	"go-rule-engine/flowengine/enums"
)

type FlowContext struct {
	ChannelName            string
	Event                  enums.FlowEvent
	ConfigName             string
	GateWayCallBackContext interface{}
}

type ExecuteContext interface {
	Context() context.Context

	WithValue(key, value interface{})
}
