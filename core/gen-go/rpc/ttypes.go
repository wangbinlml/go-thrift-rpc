// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package rpc

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

// A specific column was requested that does not exist.
type NotFoundException struct {
}

func NewNotFoundException() *NotFoundException {
	return &NotFoundException{}
}

func (p *NotFoundException) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *NotFoundException) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("NotFoundException"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *NotFoundException) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("NotFoundException(%+v)", *p)
}

func (p *NotFoundException) Error() string {
	return p.String()
}

// Invalid request could mean keyspace or column family does not exist, required parameters are missing, or a parameter is malformed.
// why contains an associated error message.
//
// Attributes:
//  - Why
type InvalidRequestException struct {
	Why string `thrift:"why,1,required" json:"why"`
}

func NewInvalidRequestException() *InvalidRequestException {
	return &InvalidRequestException{}
}

func (p *InvalidRequestException) GetWhy() string {
	return p.Why
}
func (p *InvalidRequestException) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetWhy bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetWhy = true
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetWhy {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Why is not set"))
	}
	return nil
}

func (p *InvalidRequestException) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Why = v
	}
	return nil
}

func (p *InvalidRequestException) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("InvalidRequestException"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *InvalidRequestException) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("why", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:why: ", p), err)
	}
	if err := oprot.WriteString(string(p.Why)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.why (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:why: ", p), err)
	}
	return err
}

func (p *InvalidRequestException) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("InvalidRequestException(%+v)", *p)
}

func (p *InvalidRequestException) Error() string {
	return p.String()
}

// Not all the replicas required could be created and/or read.
type UnavailableException struct {
}

func NewUnavailableException() *UnavailableException {
	return &UnavailableException{}
}

func (p *UnavailableException) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *UnavailableException) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("UnavailableException"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *UnavailableException) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UnavailableException(%+v)", *p)
}

func (p *UnavailableException) Error() string {
	return p.String()
}

// RPC timeout was exceeded.  either a node failed mid-operation, or load was too high, or the requested op was too large.
type TimedOutException struct {
}

func NewTimedOutException() *TimedOutException {
	return &TimedOutException{}
}

func (p *TimedOutException) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TimedOutException) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TimedOutException"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TimedOutException) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TimedOutException(%+v)", *p)
}

func (p *TimedOutException) Error() string {
	return p.String()
}

// Attributes:
//  - Protocol
//  - Tid
//  - MsgType
//  - InvokeMode
//  - ConnectionId
//  - MsgName
//  - RpcId
//  - RelayState
//  - URL
//  - ComeFrom
//  - To
//  - ResultCode
//  - ApId
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

func NewHeader() *Header {
	return &Header{}
}

func (p *Header) GetProtocol() string {
	return p.Protocol
}

func (p *Header) GetTid() string {
	return p.Tid
}

func (p *Header) GetMsgType() int32 {
	return p.MsgType
}

var Header_InvokeMode_DEFAULT int32

func (p *Header) GetInvokeMode() int32 {
	if !p.IsSetInvokeMode() {
		return Header_InvokeMode_DEFAULT
	}
	return *p.InvokeMode
}

var Header_ConnectionId_DEFAULT string

func (p *Header) GetConnectionId() string {
	if !p.IsSetConnectionId() {
		return Header_ConnectionId_DEFAULT
	}
	return *p.ConnectionId
}

var Header_MsgName_DEFAULT string

func (p *Header) GetMsgName() string {
	if !p.IsSetMsgName() {
		return Header_MsgName_DEFAULT
	}
	return *p.MsgName
}

var Header_RpcId_DEFAULT string

func (p *Header) GetRpcId() string {
	if !p.IsSetRpcId() {
		return Header_RpcId_DEFAULT
	}
	return *p.RpcId
}

var Header_RelayState_DEFAULT string

func (p *Header) GetRelayState() string {
	if !p.IsSetRelayState() {
		return Header_RelayState_DEFAULT
	}
	return *p.RelayState
}

var Header_URL_DEFAULT string

func (p *Header) GetURL() string {
	if !p.IsSetURL() {
		return Header_URL_DEFAULT
	}
	return *p.URL
}

var Header_ComeFrom_DEFAULT string

func (p *Header) GetComeFrom() string {
	if !p.IsSetComeFrom() {
		return Header_ComeFrom_DEFAULT
	}
	return *p.ComeFrom
}

var Header_To_DEFAULT string

func (p *Header) GetTo() string {
	if !p.IsSetTo() {
		return Header_To_DEFAULT
	}
	return *p.To
}

var Header_ResultCode_DEFAULT int32

func (p *Header) GetResultCode() int32 {
	if !p.IsSetResultCode() {
		return Header_ResultCode_DEFAULT
	}
	return *p.ResultCode
}

var Header_ApId_DEFAULT string

func (p *Header) GetApId() string {
	if !p.IsSetApId() {
		return Header_ApId_DEFAULT
	}
	return *p.ApId
}
func (p *Header) IsSetInvokeMode() bool {
	return p.InvokeMode != nil
}

func (p *Header) IsSetConnectionId() bool {
	return p.ConnectionId != nil
}

func (p *Header) IsSetMsgName() bool {
	return p.MsgName != nil
}

func (p *Header) IsSetRpcId() bool {
	return p.RpcId != nil
}

func (p *Header) IsSetRelayState() bool {
	return p.RelayState != nil
}

func (p *Header) IsSetURL() bool {
	return p.URL != nil
}

func (p *Header) IsSetComeFrom() bool {
	return p.ComeFrom != nil
}

func (p *Header) IsSetTo() bool {
	return p.To != nil
}

func (p *Header) IsSetResultCode() bool {
	return p.ResultCode != nil
}

func (p *Header) IsSetApId() bool {
	return p.ApId != nil
}

func (p *Header) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetProtocol bool = false
	var issetTid bool = false
	var issetMsgType bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetProtocol = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetTid = true
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
			issetMsgType = true
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.readField6(iprot); err != nil {
				return err
			}
		case 7:
			if err := p.readField7(iprot); err != nil {
				return err
			}
		case 8:
			if err := p.readField8(iprot); err != nil {
				return err
			}
		case 9:
			if err := p.readField9(iprot); err != nil {
				return err
			}
		case 10:
			if err := p.readField10(iprot); err != nil {
				return err
			}
		case 11:
			if err := p.readField11(iprot); err != nil {
				return err
			}
		case 12:
			if err := p.readField12(iprot); err != nil {
				return err
			}
		case 13:
			if err := p.readField13(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetProtocol {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Protocol is not set"))
	}
	if !issetTid {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Tid is not set"))
	}
	if !issetMsgType {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field MsgType is not set"))
	}
	return nil
}

func (p *Header) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Protocol = v
	}
	return nil
}

func (p *Header) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Tid = v
	}
	return nil
}

func (p *Header) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.MsgType = v
	}
	return nil
}

func (p *Header) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.InvokeMode = &v
	}
	return nil
}

func (p *Header) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.ConnectionId = &v
	}
	return nil
}

func (p *Header) readField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.MsgName = &v
	}
	return nil
}

func (p *Header) readField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		p.RpcId = &v
	}
	return nil
}

func (p *Header) readField8(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 8: ", err)
	} else {
		p.RelayState = &v
	}
	return nil
}

func (p *Header) readField9(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 9: ", err)
	} else {
		p.URL = &v
	}
	return nil
}

func (p *Header) readField10(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 10: ", err)
	} else {
		p.ComeFrom = &v
	}
	return nil
}

func (p *Header) readField11(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 11: ", err)
	} else {
		p.To = &v
	}
	return nil
}

func (p *Header) readField12(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 12: ", err)
	} else {
		p.ResultCode = &v
	}
	return nil
}

func (p *Header) readField13(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 13: ", err)
	} else {
		p.ApId = &v
	}
	return nil
}

func (p *Header) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Header"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := p.writeField6(oprot); err != nil {
		return err
	}
	if err := p.writeField7(oprot); err != nil {
		return err
	}
	if err := p.writeField8(oprot); err != nil {
		return err
	}
	if err := p.writeField9(oprot); err != nil {
		return err
	}
	if err := p.writeField10(oprot); err != nil {
		return err
	}
	if err := p.writeField11(oprot); err != nil {
		return err
	}
	if err := p.writeField12(oprot); err != nil {
		return err
	}
	if err := p.writeField13(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Header) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("protocol", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:protocol: ", p), err)
	}
	if err := oprot.WriteString(string(p.Protocol)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.protocol (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:protocol: ", p), err)
	}
	return err
}

func (p *Header) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("tid", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:tid: ", p), err)
	}
	if err := oprot.WriteString(string(p.Tid)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.tid (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:tid: ", p), err)
	}
	return err
}

func (p *Header) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("msgType", thrift.I32, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:msgType: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.MsgType)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.msgType (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:msgType: ", p), err)
	}
	return err
}

func (p *Header) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetInvokeMode() {
		if err := oprot.WriteFieldBegin("invokeMode", thrift.I32, 4); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:invokeMode: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.InvokeMode)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.invokeMode (4) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 4:invokeMode: ", p), err)
		}
	}
	return err
}

func (p *Header) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetConnectionId() {
		if err := oprot.WriteFieldBegin("connectionId", thrift.STRING, 5); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:connectionId: ", p), err)
		}
		if err := oprot.WriteString(string(*p.ConnectionId)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.connectionId (5) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 5:connectionId: ", p), err)
		}
	}
	return err
}

func (p *Header) writeField6(oprot thrift.TProtocol) (err error) {
	if p.IsSetMsgName() {
		if err := oprot.WriteFieldBegin("msgName", thrift.STRING, 6); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:msgName: ", p), err)
		}
		if err := oprot.WriteString(string(*p.MsgName)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.msgName (6) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 6:msgName: ", p), err)
		}
	}
	return err
}

func (p *Header) writeField7(oprot thrift.TProtocol) (err error) {
	if p.IsSetRpcId() {
		if err := oprot.WriteFieldBegin("rpcId", thrift.STRING, 7); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:rpcId: ", p), err)
		}
		if err := oprot.WriteString(string(*p.RpcId)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.rpcId (7) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 7:rpcId: ", p), err)
		}
	}
	return err
}

func (p *Header) writeField8(oprot thrift.TProtocol) (err error) {
	if p.IsSetRelayState() {
		if err := oprot.WriteFieldBegin("relayState", thrift.STRING, 8); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 8:relayState: ", p), err)
		}
		if err := oprot.WriteString(string(*p.RelayState)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.relayState (8) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 8:relayState: ", p), err)
		}
	}
	return err
}

func (p *Header) writeField9(oprot thrift.TProtocol) (err error) {
	if p.IsSetURL() {
		if err := oprot.WriteFieldBegin("url", thrift.STRING, 9); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 9:url: ", p), err)
		}
		if err := oprot.WriteString(string(*p.URL)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.url (9) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 9:url: ", p), err)
		}
	}
	return err
}

func (p *Header) writeField10(oprot thrift.TProtocol) (err error) {
	if p.IsSetComeFrom() {
		if err := oprot.WriteFieldBegin("comeFrom", thrift.STRING, 10); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 10:comeFrom: ", p), err)
		}
		if err := oprot.WriteString(string(*p.ComeFrom)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.comeFrom (10) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 10:comeFrom: ", p), err)
		}
	}
	return err
}

func (p *Header) writeField11(oprot thrift.TProtocol) (err error) {
	if p.IsSetTo() {
		if err := oprot.WriteFieldBegin("to", thrift.STRING, 11); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 11:to: ", p), err)
		}
		if err := oprot.WriteString(string(*p.To)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.to (11) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 11:to: ", p), err)
		}
	}
	return err
}

func (p *Header) writeField12(oprot thrift.TProtocol) (err error) {
	if p.IsSetResultCode() {
		if err := oprot.WriteFieldBegin("resultCode", thrift.I32, 12); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 12:resultCode: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.ResultCode)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.resultCode (12) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 12:resultCode: ", p), err)
		}
	}
	return err
}

func (p *Header) writeField13(oprot thrift.TProtocol) (err error) {
	if p.IsSetApId() {
		if err := oprot.WriteFieldBegin("apId", thrift.STRING, 13); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 13:apId: ", p), err)
		}
		if err := oprot.WriteString(string(*p.ApId)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.apId (13) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 13:apId: ", p), err)
		}
	}
	return err
}

func (p *Header) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Header(%+v)", *p)
}

// Attributes:
//  - Header
//  - Body
type Msg struct {
	Header *Header `thrift:"header,1,required" json:"header"`
	Body   string  `thrift:"body,2,required" json:"body"`
}

func NewMsg() *Msg {
	return &Msg{}
}

var Msg_Header_DEFAULT *Header

func (p *Msg) GetHeader() *Header {
	if !p.IsSetHeader() {
		return Msg_Header_DEFAULT
	}
	return p.Header
}

func (p *Msg) GetBody() string {
	return p.Body
}
func (p *Msg) IsSetHeader() bool {
	return p.Header != nil
}

func (p *Msg) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetHeader bool = false
	var issetBody bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetHeader = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetBody = true
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetHeader {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Header is not set"))
	}
	if !issetBody {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Body is not set"))
	}
	return nil
}

func (p *Msg) readField1(iprot thrift.TProtocol) error {
	p.Header = &Header{}
	if err := p.Header.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Header), err)
	}
	return nil
}

func (p *Msg) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Body = v
	}
	return nil
}

func (p *Msg) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Msg"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Msg) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("header", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:header: ", p), err)
	}
	if err := p.Header.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Header), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:header: ", p), err)
	}
	return err
}

func (p *Msg) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("body", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:body: ", p), err)
	}
	if err := oprot.WriteString(string(p.Body)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.body (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:body: ", p), err)
	}
	return err
}

func (p *Msg) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Msg(%+v)", *p)
}
