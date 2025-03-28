// Code generated by thriftgo (0.3.18). DO NOT EDIT.

package rpc

import (
	apache_warning "code.byted.org/kitex/apache_monitor"
	"context"
	"fmt"
	"github.com/Yotinx/go_demo/kitex_gen/base"
	thrift "github.com/cloudwego/kitex/pkg/protocol/bthrift/apache"
	"strings"
)

type Item struct {
	Id      int64  `thrift:"id,1,required" frugal:"1,required,i64" json:"id"`
	Title   string `thrift:"title,2,required" frugal:"2,required,string" json:"title"`
	Content string `thrift:"content,3,required" frugal:"3,required,string" json:"content"`
}

func NewItem() *Item {
	return &Item{}
}

func (p *Item) InitDefault() {
}

func (p *Item) GetId() (v int64) {
	return p.Id
}

func (p *Item) GetTitle() (v string) {
	return p.Title
}

func (p *Item) GetContent() (v string) {
	return p.Content
}
func (p *Item) SetId(val int64) {
	p.Id = val
}
func (p *Item) SetTitle(val string) {
	p.Title = val
}
func (p *Item) SetContent(val string) {
	p.Content = val
}

var fieldIDToName_Item = map[int16]string{
	1: "id",
	2: "title",
	3: "content",
}

func (p *Item) Read(iprot thrift.TProtocol) (err error) {

	apache_warning.WarningApache("Item")

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetId bool = false
	var issetTitle bool = false
	var issetContent bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetId = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
				issetTitle = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
				issetContent = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetId {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetTitle {
		fieldId = 2
		goto RequiredFieldNotSetError
	}

	if !issetContent {
		fieldId = 3
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Item[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_Item[fieldId]))
}

func (p *Item) ReadField1(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Id = _field
	return nil
}
func (p *Item) ReadField2(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Title = _field
	return nil
}
func (p *Item) ReadField3(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Content = _field
	return nil
}

func (p *Item) Write(oprot thrift.TProtocol) (err error) {

	apache_warning.WarningApache("Item")

	var fieldId int16
	if err = oprot.WriteStructBegin("Item"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *Item) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("id", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.Id); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *Item) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("title", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Title); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *Item) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("content", thrift.STRING, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Content); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *Item) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Item(%+v)", *p)

}

func (p *Item) DeepEqual(ano *Item) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Id) {
		return false
	}
	if !p.Field2DeepEqual(ano.Title) {
		return false
	}
	if !p.Field3DeepEqual(ano.Content) {
		return false
	}
	return true
}

func (p *Item) Field1DeepEqual(src int64) bool {

	if p.Id != src {
		return false
	}
	return true
}
func (p *Item) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Title, src) != 0 {
		return false
	}
	return true
}
func (p *Item) Field3DeepEqual(src string) bool {

	if strings.Compare(p.Content, src) != 0 {
		return false
	}
	return true
}

type GetItemRequest struct {
	Id   int64      `thrift:"id,1,required" frugal:"1,required,i64" json:"id"`
	Base *base.Base `thrift:"Base,255,optional" frugal:"255,optional,base.Base" json:"Base,omitempty"`
}

func NewGetItemRequest() *GetItemRequest {
	return &GetItemRequest{}
}

func (p *GetItemRequest) InitDefault() {
}

func (p *GetItemRequest) GetId() (v int64) {
	return p.Id
}

var GetItemRequest_Base_DEFAULT *base.Base

func (p *GetItemRequest) GetBase() (v *base.Base) {
	if !p.IsSetBase() {
		return GetItemRequest_Base_DEFAULT
	}
	return p.Base
}
func (p *GetItemRequest) SetId(val int64) {
	p.Id = val
}
func (p *GetItemRequest) SetBase(val *base.Base) {
	p.Base = val
}

var fieldIDToName_GetItemRequest = map[int16]string{
	1:   "id",
	255: "Base",
}

func (p *GetItemRequest) IsSetBase() bool {
	return p.Base != nil
}

func (p *GetItemRequest) Read(iprot thrift.TProtocol) (err error) {

	apache_warning.WarningApache("GetItemRequest")

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetId bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetId = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 255:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField255(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetId {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GetItemRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_GetItemRequest[fieldId]))
}

func (p *GetItemRequest) ReadField1(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Id = _field
	return nil
}
func (p *GetItemRequest) ReadField255(iprot thrift.TProtocol) error {
	_field := base.NewBase()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Base = _field
	return nil
}

func (p *GetItemRequest) Write(oprot thrift.TProtocol) (err error) {

	apache_warning.WarningApache("GetItemRequest")

	var fieldId int16
	if err = oprot.WriteStructBegin("GetItemRequest"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField255(oprot); err != nil {
			fieldId = 255
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *GetItemRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("id", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.Id); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *GetItemRequest) writeField255(oprot thrift.TProtocol) (err error) {
	if p.IsSetBase() {
		if err = oprot.WriteFieldBegin("Base", thrift.STRUCT, 255); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Base.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 end error: ", p), err)
}

func (p *GetItemRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetItemRequest(%+v)", *p)

}

func (p *GetItemRequest) DeepEqual(ano *GetItemRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Id) {
		return false
	}
	if !p.Field255DeepEqual(ano.Base) {
		return false
	}
	return true
}

func (p *GetItemRequest) Field1DeepEqual(src int64) bool {

	if p.Id != src {
		return false
	}
	return true
}
func (p *GetItemRequest) Field255DeepEqual(src *base.Base) bool {

	if !p.Base.DeepEqual(src) {
		return false
	}
	return true
}

type GetItemResponse struct {
	Item     *Item          `thrift:"item,1,required" frugal:"1,required,Item" json:"item"`
	BaseResp *base.BaseResp `thrift:"BaseResp,255,optional" frugal:"255,optional,base.BaseResp" json:"BaseResp,omitempty"`
}

func NewGetItemResponse() *GetItemResponse {
	return &GetItemResponse{}
}

func (p *GetItemResponse) InitDefault() {
}

var GetItemResponse_Item_DEFAULT *Item

func (p *GetItemResponse) GetItem() (v *Item) {
	if !p.IsSetItem() {
		return GetItemResponse_Item_DEFAULT
	}
	return p.Item
}

var GetItemResponse_BaseResp_DEFAULT *base.BaseResp

func (p *GetItemResponse) GetBaseResp() (v *base.BaseResp) {
	if !p.IsSetBaseResp() {
		return GetItemResponse_BaseResp_DEFAULT
	}
	return p.BaseResp
}
func (p *GetItemResponse) SetItem(val *Item) {
	p.Item = val
}
func (p *GetItemResponse) SetBaseResp(val *base.BaseResp) {
	p.BaseResp = val
}

var fieldIDToName_GetItemResponse = map[int16]string{
	1:   "item",
	255: "BaseResp",
}

func (p *GetItemResponse) IsSetItem() bool {
	return p.Item != nil
}

func (p *GetItemResponse) IsSetBaseResp() bool {
	return p.BaseResp != nil
}

func (p *GetItemResponse) Read(iprot thrift.TProtocol) (err error) {

	apache_warning.WarningApache("GetItemResponse")

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetItem bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetItem = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 255:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField255(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetItem {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GetItemResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_GetItemResponse[fieldId]))
}

func (p *GetItemResponse) ReadField1(iprot thrift.TProtocol) error {
	_field := NewItem()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Item = _field
	return nil
}
func (p *GetItemResponse) ReadField255(iprot thrift.TProtocol) error {
	_field := base.NewBaseResp()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.BaseResp = _field
	return nil
}

func (p *GetItemResponse) Write(oprot thrift.TProtocol) (err error) {

	apache_warning.WarningApache("GetItemResponse")

	var fieldId int16
	if err = oprot.WriteStructBegin("GetItemResponse"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField255(oprot); err != nil {
			fieldId = 255
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *GetItemResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("item", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Item.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *GetItemResponse) writeField255(oprot thrift.TProtocol) (err error) {
	if p.IsSetBaseResp() {
		if err = oprot.WriteFieldBegin("BaseResp", thrift.STRUCT, 255); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.BaseResp.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 end error: ", p), err)
}

func (p *GetItemResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetItemResponse(%+v)", *p)

}

func (p *GetItemResponse) DeepEqual(ano *GetItemResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Item) {
		return false
	}
	if !p.Field255DeepEqual(ano.BaseResp) {
		return false
	}
	return true
}

func (p *GetItemResponse) Field1DeepEqual(src *Item) bool {

	if !p.Item.DeepEqual(src) {
		return false
	}
	return true
}
func (p *GetItemResponse) Field255DeepEqual(src *base.BaseResp) bool {

	if !p.BaseResp.DeepEqual(src) {
		return false
	}
	return true
}

type ItemService interface {
	GetItem(ctx context.Context, req *GetItemRequest) (r *GetItemResponse, err error)
}

type ItemServiceGetItemArgs struct {
	Req *GetItemRequest `thrift:"req,1" frugal:"1,default,GetItemRequest" json:"req"`
}

func NewItemServiceGetItemArgs() *ItemServiceGetItemArgs {
	return &ItemServiceGetItemArgs{}
}

func (p *ItemServiceGetItemArgs) InitDefault() {
}

var ItemServiceGetItemArgs_Req_DEFAULT *GetItemRequest

func (p *ItemServiceGetItemArgs) GetReq() (v *GetItemRequest) {
	if !p.IsSetReq() {
		return ItemServiceGetItemArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *ItemServiceGetItemArgs) SetReq(val *GetItemRequest) {
	p.Req = val
}

var fieldIDToName_ItemServiceGetItemArgs = map[int16]string{
	1: "req",
}

func (p *ItemServiceGetItemArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ItemServiceGetItemArgs) Read(iprot thrift.TProtocol) (err error) {

	apache_warning.WarningApache("ItemServiceGetItemArgs")

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ItemServiceGetItemArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ItemServiceGetItemArgs) ReadField1(iprot thrift.TProtocol) error {
	_field := NewGetItemRequest()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Req = _field
	return nil
}

func (p *ItemServiceGetItemArgs) Write(oprot thrift.TProtocol) (err error) {

	apache_warning.WarningApache("ItemServiceGetItemArgs")

	var fieldId int16
	if err = oprot.WriteStructBegin("GetItem_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ItemServiceGetItemArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Req.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *ItemServiceGetItemArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ItemServiceGetItemArgs(%+v)", *p)

}

func (p *ItemServiceGetItemArgs) DeepEqual(ano *ItemServiceGetItemArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Req) {
		return false
	}
	return true
}

func (p *ItemServiceGetItemArgs) Field1DeepEqual(src *GetItemRequest) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

type ItemServiceGetItemResult struct {
	Success *GetItemResponse `thrift:"success,0,optional" frugal:"0,optional,GetItemResponse" json:"success,omitempty"`
}

func NewItemServiceGetItemResult() *ItemServiceGetItemResult {
	return &ItemServiceGetItemResult{}
}

func (p *ItemServiceGetItemResult) InitDefault() {
}

var ItemServiceGetItemResult_Success_DEFAULT *GetItemResponse

func (p *ItemServiceGetItemResult) GetSuccess() (v *GetItemResponse) {
	if !p.IsSetSuccess() {
		return ItemServiceGetItemResult_Success_DEFAULT
	}
	return p.Success
}
func (p *ItemServiceGetItemResult) SetSuccess(x interface{}) {
	p.Success = x.(*GetItemResponse)
}

var fieldIDToName_ItemServiceGetItemResult = map[int16]string{
	0: "success",
}

func (p *ItemServiceGetItemResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ItemServiceGetItemResult) Read(iprot thrift.TProtocol) (err error) {

	apache_warning.WarningApache("ItemServiceGetItemResult")

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ItemServiceGetItemResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ItemServiceGetItemResult) ReadField0(iprot thrift.TProtocol) error {
	_field := NewGetItemResponse()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Success = _field
	return nil
}

func (p *ItemServiceGetItemResult) Write(oprot thrift.TProtocol) (err error) {

	apache_warning.WarningApache("ItemServiceGetItemResult")

	var fieldId int16
	if err = oprot.WriteStructBegin("GetItem_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ItemServiceGetItemResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *ItemServiceGetItemResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ItemServiceGetItemResult(%+v)", *p)

}

func (p *ItemServiceGetItemResult) DeepEqual(ano *ItemServiceGetItemResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *ItemServiceGetItemResult) Field0DeepEqual(src *GetItemResponse) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}
