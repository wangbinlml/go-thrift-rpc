package core

type Header struct {
	Protocol     string  `thrift:"protocol,1,required" json:"protocol"`
	Tid          string  `thrift:"tid,2,required" json:"tid"`
	MsgType      int32   `thrift:"msgType,3,required" json:"msgType"`
	InvokeMode   *int32  `thrift:"invokeMode,4" json:"invokeMode,omitempty"`
	ConnectionId *string `thrift:"connectionId,5" json:"connectionId,omitempty"`
	MsgName      *string `thrift:"msgName,6" json:"msgName,omitempty"`
	RpcId        *string `thrift:"rpcId,7" json:"rpcId,omitempty"`
	RelayState   *string `thrift:"relayState,8" json:"relayState,omitempty"`
	URL          *string `thrift:"url,9" json:"url,omitempty"`
	ComeFrom     *string `thrift:"comeFrom,10" json:"comeFrom,omitempty"`
	To           *string `thrift:"to,11" json:"to,omitempty"`
	ResultCode   *int32  `thrift:"resultCode,12" json:"resultCode,omitempty"`
	ApId         *string `thrift:"apId,13" json:"apId,omitempty"`
}