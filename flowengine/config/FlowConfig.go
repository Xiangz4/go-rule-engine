package config

import "go-rule-engine/flowengine/enums"

type FlowConfig interface {
	Name() string
	GetHandlerNode(flowState enums.FlowState, flowEvent enums.FlowEvent) interface{}
}
