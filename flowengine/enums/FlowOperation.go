package enums

type FlowOperation interface {
	Name() string

	GetMethod() string

	IsCallBack() bool
}
