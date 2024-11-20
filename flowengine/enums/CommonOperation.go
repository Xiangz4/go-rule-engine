package enums

type CommonOperation struct {
	name   string
	method string
	async  bool
}

func (o CommonOperation) String() string {
	return o.name
}

func (o CommonOperation) Method() string {
	return o.method
}

func (o CommonOperation) IsAsync() bool {
	return o.async
}

var (
	Parse = CommonOperation{"Parse", "parse", true}
	//"pay"发送给网关，网关根据false判断是否是一个回调
	Pay      = CommonOperation{"Pay", "pay", false}
	PayQuery = CommonOperation{"PAY_QUERY", "payQuery", false}
	Refund   = CommonOperation{"REFUND", "refund", false}
)

func AllCommonOperations() []CommonOperation {
	return []CommonOperation{
		Parse,
		Pay,
		PayQuery,
		Refund,
	}
}

func ParseOperation(method string) (CommonOperation, bool) {
	for _, op := range AllCommonOperations() {
		if op.method == method {
			return op, true
		}
	}
	return CommonOperation{}, false
}
