package enums

type CommonOrderState int

const (
	INIT = iota
	PROCESS
	SUCCESS
	FAIL
)

func (o CommonOrderState) Name() string {
	switch o {
	case INIT:
		return "INIT"
	case PROCESS:
		return "PROCESS"
	case SUCCESS:
		return "SUCCESS"
	case FAIL:
		return "FAIL"
	default:
		return "UNKNOW"
	}
}
