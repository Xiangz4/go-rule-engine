package context

import "go-rule-engine/flowengine/enums"

/*
	用于处理器与服务之间传递需要处理的数据
*/

type GateWayCallBackContext struct {
	State                  enums.FlowState
	CallbackMessage        string
	ChannelResponseCode    string
	ChannelResponseMessage string
}
