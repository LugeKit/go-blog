package errcode

var (
	GetTagListFail = newError(20010001, "get tag list fail")
	CreateTagFail  = newError(20010002, "create tag fail")
	UpdateTagFail  = newError(20010003, "update tag fail")
	DeleteTagFail  = newError(20010004, "delete tag fail")
)
