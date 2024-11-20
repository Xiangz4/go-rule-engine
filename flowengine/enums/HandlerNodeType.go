package enums

type HandlerNode struct{
	CurrentState FlowState

	NextState FlowState

	HandlerFunc func()

}

var (
	TRANSITION()
	REQUEST()
	NOTIFY()
)