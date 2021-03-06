// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package account

import (
	"bytes"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type Account interface {
	// Parameters:
	//  - Request
	DoAction(request *Request) (r string, err error)
}

type AccountClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewAccountClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *AccountClient {
	return &AccountClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewAccountClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *AccountClient {
	return &AccountClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - Request
func (p *AccountClient) DoAction(request *Request) (r string, err error) {
	if err = p.sendDoAction(request); err != nil {
		return
	}
	return p.recvDoAction()
}

func (p *AccountClient) sendDoAction(request *Request) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("doAction", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := AccountDoActionArgs{
		Request: request,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *AccountClient) recvDoAction() (value string, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "doAction" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "doAction failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "doAction failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error1 error
		error1, err = error0.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error1
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "doAction failed: invalid message type")
		return
	}
	result := AccountDoActionResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.E != nil {
		err = result.E
		return
	}
	value = result.GetSuccess()
	return
}

type AccountProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      Account
}

func (p *AccountProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *AccountProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *AccountProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewAccountProcessor(handler Account) *AccountProcessor {

	self2 := &AccountProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self2.processorMap["doAction"] = &accountProcessorDoAction{handler: handler}
	return self2
}

func (p *AccountProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x3.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x3

}

type accountProcessorDoAction struct {
	handler Account
}

func (p *accountProcessorDoAction) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := AccountDoActionArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("doAction", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := AccountDoActionResult{}
	var retval string
	var err2 error
	if retval, err2 = p.handler.DoAction(args.Request); err2 != nil {
		switch v := err2.(type) {
		case *InvalidOperation:
			result.E = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing doAction: "+err2.Error())
			oprot.WriteMessageBegin("doAction", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("doAction", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Request
type AccountDoActionArgs struct {
	Request *Request `thrift:"request,1" json:"request"`
}

func NewAccountDoActionArgs() *AccountDoActionArgs {
	return &AccountDoActionArgs{}
}

var AccountDoActionArgs_Request_DEFAULT *Request

func (p *AccountDoActionArgs) GetRequest() *Request {
	if !p.IsSetRequest() {
		return AccountDoActionArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *AccountDoActionArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *AccountDoActionArgs) Read(iprot thrift.TProtocol) error {
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
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
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
	return nil
}

func (p *AccountDoActionArgs) readField1(iprot thrift.TProtocol) error {
	p.Request = &Request{}
	if err := p.Request.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Request), err)
	}
	return nil
}

func (p *AccountDoActionArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("doAction_args"); err != nil {
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

func (p *AccountDoActionArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("request", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request: ", p), err)
	}
	if err := p.Request.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Request), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request: ", p), err)
	}
	return err
}

func (p *AccountDoActionArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AccountDoActionArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type AccountDoActionResult struct {
	Success *string           `thrift:"success,0" json:"success,omitempty"`
	E       *InvalidOperation `thrift:"e,1" json:"e,omitempty"`
}

func NewAccountDoActionResult() *AccountDoActionResult {
	return &AccountDoActionResult{}
}

var AccountDoActionResult_Success_DEFAULT string

func (p *AccountDoActionResult) GetSuccess() string {
	if !p.IsSetSuccess() {
		return AccountDoActionResult_Success_DEFAULT
	}
	return *p.Success
}

var AccountDoActionResult_E_DEFAULT *InvalidOperation

func (p *AccountDoActionResult) GetE() *InvalidOperation {
	if !p.IsSetE() {
		return AccountDoActionResult_E_DEFAULT
	}
	return p.E
}
func (p *AccountDoActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountDoActionResult) IsSetE() bool {
	return p.E != nil
}

func (p *AccountDoActionResult) Read(iprot thrift.TProtocol) error {
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
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		case 1:
			if err := p.readField1(iprot); err != nil {
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
	return nil
}

func (p *AccountDoActionResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *AccountDoActionResult) readField1(iprot thrift.TProtocol) error {
	p.E = &InvalidOperation{}
	if err := p.E.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *AccountDoActionResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("doAction_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
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

func (p *AccountDoActionResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *AccountDoActionResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin("e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *AccountDoActionResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AccountDoActionResult(%+v)", *p)
}
