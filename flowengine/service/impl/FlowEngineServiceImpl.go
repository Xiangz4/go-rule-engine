package impl

import (
	"go-rule-engine/flowengine/config"
	"go-rule-engine/flowengine/context"
	"log"
)

type FlowEngineServiceImpl struct {
	ConfigMap map[string]config.FlowConfig

	configList []config.FlowConfig
}

func (f *FlowEngineServiceImpl) Init() {
	for _, v := range f.configList {
		f.ConfigMap[v.Name()] = v
	}
}

func (f *FlowEngineServiceImpl) SetConfigList(configList []config.FlowConfig) {
	for _, v := range configList {
		f.configList = append(f.configList, v)
	}
}

func (f *FlowEngineServiceImpl) Execute(flowContext context.FlowContext) {
	config := f.ConfigMap[flowContext.ConfigName]

	if config == nil {
		log.Fatal("No configuration found for name: " + flowContext.ConfigName)
	}
}
