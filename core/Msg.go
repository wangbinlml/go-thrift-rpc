package core

type Msg struct {
	Header *Header `thrift:"header,1,required" json:"header"`
	Body   string  `thrift:"body,2,required" json:"body"`
}
