// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.3
// source: tensorflow/core/framework/attr_value.proto

package attr_value_go_proto

import (
	tensor_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_go_proto"
	tensor_shape_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_shape_go_proto"
	types_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/types_go_proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Protocol buffer representing the value for an attr used to configure an Op.
// Comment indicates the corresponding attr type.  Only the field matching the
// attr type may be filled.
type AttrValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*AttrValue_S
	//	*AttrValue_I
	//	*AttrValue_F
	//	*AttrValue_B
	//	*AttrValue_Type
	//	*AttrValue_Shape
	//	*AttrValue_Tensor
	//	*AttrValue_List
	//	*AttrValue_Func
	//	*AttrValue_Placeholder
	Value isAttrValue_Value `protobuf_oneof:"value"`
}

func (x *AttrValue) Reset() {
	*x = AttrValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_attr_value_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttrValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttrValue) ProtoMessage() {}

func (x *AttrValue) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_attr_value_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttrValue.ProtoReflect.Descriptor instead.
func (*AttrValue) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_attr_value_proto_rawDescGZIP(), []int{0}
}

func (m *AttrValue) GetValue() isAttrValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *AttrValue) GetS() []byte {
	if x, ok := x.GetValue().(*AttrValue_S); ok {
		return x.S
	}
	return nil
}

func (x *AttrValue) GetI() int64 {
	if x, ok := x.GetValue().(*AttrValue_I); ok {
		return x.I
	}
	return 0
}

func (x *AttrValue) GetF() float32 {
	if x, ok := x.GetValue().(*AttrValue_F); ok {
		return x.F
	}
	return 0
}

func (x *AttrValue) GetB() bool {
	if x, ok := x.GetValue().(*AttrValue_B); ok {
		return x.B
	}
	return false
}

func (x *AttrValue) GetType() types_go_proto.DataType {
	if x, ok := x.GetValue().(*AttrValue_Type); ok {
		return x.Type
	}
	return types_go_proto.DataType_DT_INVALID
}

func (x *AttrValue) GetShape() *tensor_shape_go_proto.TensorShapeProto {
	if x, ok := x.GetValue().(*AttrValue_Shape); ok {
		return x.Shape
	}
	return nil
}

func (x *AttrValue) GetTensor() *tensor_go_proto.TensorProto {
	if x, ok := x.GetValue().(*AttrValue_Tensor); ok {
		return x.Tensor
	}
	return nil
}

func (x *AttrValue) GetList() *AttrValue_ListValue {
	if x, ok := x.GetValue().(*AttrValue_List); ok {
		return x.List
	}
	return nil
}

func (x *AttrValue) GetFunc() *NameAttrList {
	if x, ok := x.GetValue().(*AttrValue_Func); ok {
		return x.Func
	}
	return nil
}

func (x *AttrValue) GetPlaceholder() string {
	if x, ok := x.GetValue().(*AttrValue_Placeholder); ok {
		return x.Placeholder
	}
	return ""
}

type isAttrValue_Value interface {
	isAttrValue_Value()
}

type AttrValue_S struct {
	S []byte `protobuf:"bytes,2,opt,name=s,proto3,oneof"` // "string"
}

type AttrValue_I struct {
	I int64 `protobuf:"varint,3,opt,name=i,proto3,oneof"` // "int"
}

type AttrValue_F struct {
	F float32 `protobuf:"fixed32,4,opt,name=f,proto3,oneof"` // "float"
}

type AttrValue_B struct {
	B bool `protobuf:"varint,5,opt,name=b,proto3,oneof"` // "bool"
}

type AttrValue_Type struct {
	Type types_go_proto.DataType `protobuf:"varint,6,opt,name=type,proto3,enum=tensorflow.DataType,oneof"` // "type"
}

type AttrValue_Shape struct {
	Shape *tensor_shape_go_proto.TensorShapeProto `protobuf:"bytes,7,opt,name=shape,proto3,oneof"` // "shape"
}

type AttrValue_Tensor struct {
	Tensor *tensor_go_proto.TensorProto `protobuf:"bytes,8,opt,name=tensor,proto3,oneof"` // "tensor"
}

type AttrValue_List struct {
	List *AttrValue_ListValue `protobuf:"bytes,1,opt,name=list,proto3,oneof"` // any "list(...)"
}

type AttrValue_Func struct {
	// "func" represents a function. func.name is a function's name or
	// a primitive op's name. func.attr.first is the name of an attr
	// defined for that function. func.attr.second is the value for
	// that attr in the instantiation.
	Func *NameAttrList `protobuf:"bytes,10,opt,name=func,proto3,oneof"`
}

type AttrValue_Placeholder struct {
	// This is a placeholder only used in nodes defined inside a
	// function.  It indicates the attr value will be supplied when
	// the function is instantiated.  For example, let us suppose a
	// node "N" in function "FN". "N" has an attr "A" with value
	// placeholder = "foo". When FN is instantiated with attr "foo"
	// set to "bar", the instantiated node N's attr A will have been
	// given the value "bar".
	Placeholder string `protobuf:"bytes,9,opt,name=placeholder,proto3,oneof"`
}

func (*AttrValue_S) isAttrValue_Value() {}

func (*AttrValue_I) isAttrValue_Value() {}

func (*AttrValue_F) isAttrValue_Value() {}

func (*AttrValue_B) isAttrValue_Value() {}

func (*AttrValue_Type) isAttrValue_Value() {}

func (*AttrValue_Shape) isAttrValue_Value() {}

func (*AttrValue_Tensor) isAttrValue_Value() {}

func (*AttrValue_List) isAttrValue_Value() {}

func (*AttrValue_Func) isAttrValue_Value() {}

func (*AttrValue_Placeholder) isAttrValue_Value() {}

// A list of attr names and their values. The whole list is attached
// with a string name.  E.g., MatMul[T=float].
type NameAttrList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Attr map[string]*AttrValue `protobuf:"bytes,2,rep,name=attr,proto3" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *NameAttrList) Reset() {
	*x = NameAttrList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_attr_value_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameAttrList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameAttrList) ProtoMessage() {}

func (x *NameAttrList) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_attr_value_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameAttrList.ProtoReflect.Descriptor instead.
func (*NameAttrList) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_attr_value_proto_rawDescGZIP(), []int{1}
}

func (x *NameAttrList) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NameAttrList) GetAttr() map[string]*AttrValue {
	if x != nil {
		return x.Attr
	}
	return nil
}

// LINT.IfChange
type AttrValue_ListValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	S      [][]byte                                  `protobuf:"bytes,2,rep,name=s,proto3" json:"s,omitempty"`                                        // "list(string)"
	I      []int64                                   `protobuf:"varint,3,rep,packed,name=i,proto3" json:"i,omitempty"`                                // "list(int)"
	F      []float32                                 `protobuf:"fixed32,4,rep,packed,name=f,proto3" json:"f,omitempty"`                               // "list(float)"
	B      []bool                                    `protobuf:"varint,5,rep,packed,name=b,proto3" json:"b,omitempty"`                                // "list(bool)"
	Type   []types_go_proto.DataType                 `protobuf:"varint,6,rep,packed,name=type,proto3,enum=tensorflow.DataType" json:"type,omitempty"` // "list(type)"
	Shape  []*tensor_shape_go_proto.TensorShapeProto `protobuf:"bytes,7,rep,name=shape,proto3" json:"shape,omitempty"`                                // "list(shape)"
	Tensor []*tensor_go_proto.TensorProto            `protobuf:"bytes,8,rep,name=tensor,proto3" json:"tensor,omitempty"`                              // "list(tensor)"
	Func   []*NameAttrList                           `protobuf:"bytes,9,rep,name=func,proto3" json:"func,omitempty"`                                  // "list(attr)"
}

func (x *AttrValue_ListValue) Reset() {
	*x = AttrValue_ListValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_attr_value_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttrValue_ListValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttrValue_ListValue) ProtoMessage() {}

func (x *AttrValue_ListValue) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_attr_value_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttrValue_ListValue.ProtoReflect.Descriptor instead.
func (*AttrValue_ListValue) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_attr_value_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AttrValue_ListValue) GetS() [][]byte {
	if x != nil {
		return x.S
	}
	return nil
}

func (x *AttrValue_ListValue) GetI() []int64 {
	if x != nil {
		return x.I
	}
	return nil
}

func (x *AttrValue_ListValue) GetF() []float32 {
	if x != nil {
		return x.F
	}
	return nil
}

func (x *AttrValue_ListValue) GetB() []bool {
	if x != nil {
		return x.B
	}
	return nil
}

func (x *AttrValue_ListValue) GetType() []types_go_proto.DataType {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *AttrValue_ListValue) GetShape() []*tensor_shape_go_proto.TensorShapeProto {
	if x != nil {
		return x.Shape
	}
	return nil
}

func (x *AttrValue_ListValue) GetTensor() []*tensor_go_proto.TensorProto {
	if x != nil {
		return x.Tensor
	}
	return nil
}

func (x *AttrValue_ListValue) GetFunc() []*NameAttrList {
	if x != nil {
		return x.Func
	}
	return nil
}

var File_tensorflow_core_framework_attr_value_proto protoreflect.FileDescriptor

var file_tensorflow_core_framework_attr_value_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x74, 0x74, 0x72,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x26, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2c, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x5f, 0x73, 0x68, 0x61, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x05, 0x0a, 0x09, 0x41, 0x74, 0x74, 0x72, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x01, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00,
	0x52, 0x01, 0x73, 0x12, 0x0e, 0x0a, 0x01, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00,
	0x52, 0x01, 0x69, 0x12, 0x0e, 0x0a, 0x01, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00,
	0x52, 0x01, 0x66, 0x12, 0x0e, 0x0a, 0x01, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00,
	0x52, 0x01, 0x62, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x34, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x53, 0x68, 0x61, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x00, 0x52, 0x05,
	0x73, 0x68, 0x61, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x00,
	0x52, 0x06, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x12, 0x35, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x2e, 0x0a, 0x04, 0x66, 0x75, 0x6e, 0x63, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x41,
	0x74, 0x74, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x04, 0x66, 0x75, 0x6e, 0x63, 0x12,
	0x22, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x1a, 0x90, 0x02, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x0c, 0x0a, 0x01, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x01, 0x73, 0x12,
	0x10, 0x0a, 0x01, 0x69, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x42, 0x02, 0x10, 0x01, 0x52, 0x01,
	0x69, 0x12, 0x10, 0x0a, 0x01, 0x66, 0x18, 0x04, 0x20, 0x03, 0x28, 0x02, 0x42, 0x02, 0x10, 0x01,
	0x52, 0x01, 0x66, 0x12, 0x10, 0x0a, 0x01, 0x62, 0x18, 0x05, 0x20, 0x03, 0x28, 0x08, 0x42, 0x02,
	0x10, 0x01, 0x52, 0x01, 0x62, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x42, 0x02, 0x10, 0x01, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53, 0x68, 0x61, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x52, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x52, 0x06, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x12, 0x2c, 0x0a, 0x04, 0x66, 0x75, 0x6e, 0x63,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x41, 0x74, 0x74, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x04, 0x66, 0x75, 0x6e, 0x63, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0xaa, 0x01, 0x0a, 0x0c, 0x4e, 0x61, 0x6d, 0x65, 0x41, 0x74, 0x74, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x61, 0x74, 0x74, 0x72, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x4e, 0x61, 0x6d, 0x65, 0x41, 0x74, 0x74, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x41, 0x74, 0x74,
	0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x61, 0x74, 0x74, 0x72, 0x1a, 0x4e, 0x0a, 0x09,
	0x41, 0x74, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x83, 0x01, 0x0a,
	0x18, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a, 0x51, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x74, 0x74, 0x72,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8,
	0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_framework_attr_value_proto_rawDescOnce sync.Once
	file_tensorflow_core_framework_attr_value_proto_rawDescData = file_tensorflow_core_framework_attr_value_proto_rawDesc
)

func file_tensorflow_core_framework_attr_value_proto_rawDescGZIP() []byte {
	file_tensorflow_core_framework_attr_value_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_framework_attr_value_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_framework_attr_value_proto_rawDescData)
	})
	return file_tensorflow_core_framework_attr_value_proto_rawDescData
}

var file_tensorflow_core_framework_attr_value_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tensorflow_core_framework_attr_value_proto_goTypes = []interface{}{
	(*AttrValue)(nil),            // 0: tensorflow.AttrValue
	(*NameAttrList)(nil),         // 1: tensorflow.NameAttrList
	(*AttrValue_ListValue)(nil),  // 2: tensorflow.AttrValue.ListValue
	nil,                          // 3: tensorflow.NameAttrList.AttrEntry
	(types_go_proto.DataType)(0), // 4: tensorflow.DataType
	(*tensor_shape_go_proto.TensorShapeProto)(nil), // 5: tensorflow.TensorShapeProto
	(*tensor_go_proto.TensorProto)(nil),            // 6: tensorflow.TensorProto
}
var file_tensorflow_core_framework_attr_value_proto_depIdxs = []int32{
	4,  // 0: tensorflow.AttrValue.type:type_name -> tensorflow.DataType
	5,  // 1: tensorflow.AttrValue.shape:type_name -> tensorflow.TensorShapeProto
	6,  // 2: tensorflow.AttrValue.tensor:type_name -> tensorflow.TensorProto
	2,  // 3: tensorflow.AttrValue.list:type_name -> tensorflow.AttrValue.ListValue
	1,  // 4: tensorflow.AttrValue.func:type_name -> tensorflow.NameAttrList
	3,  // 5: tensorflow.NameAttrList.attr:type_name -> tensorflow.NameAttrList.AttrEntry
	4,  // 6: tensorflow.AttrValue.ListValue.type:type_name -> tensorflow.DataType
	5,  // 7: tensorflow.AttrValue.ListValue.shape:type_name -> tensorflow.TensorShapeProto
	6,  // 8: tensorflow.AttrValue.ListValue.tensor:type_name -> tensorflow.TensorProto
	1,  // 9: tensorflow.AttrValue.ListValue.func:type_name -> tensorflow.NameAttrList
	0,  // 10: tensorflow.NameAttrList.AttrEntry.value:type_name -> tensorflow.AttrValue
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_tensorflow_core_framework_attr_value_proto_init() }
func file_tensorflow_core_framework_attr_value_proto_init() {
	if File_tensorflow_core_framework_attr_value_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_framework_attr_value_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttrValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tensorflow_core_framework_attr_value_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameAttrList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tensorflow_core_framework_attr_value_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttrValue_ListValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_tensorflow_core_framework_attr_value_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AttrValue_S)(nil),
		(*AttrValue_I)(nil),
		(*AttrValue_F)(nil),
		(*AttrValue_B)(nil),
		(*AttrValue_Type)(nil),
		(*AttrValue_Shape)(nil),
		(*AttrValue_Tensor)(nil),
		(*AttrValue_List)(nil),
		(*AttrValue_Func)(nil),
		(*AttrValue_Placeholder)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_core_framework_attr_value_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_framework_attr_value_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_framework_attr_value_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_framework_attr_value_proto_msgTypes,
	}.Build()
	File_tensorflow_core_framework_attr_value_proto = out.File
	file_tensorflow_core_framework_attr_value_proto_rawDesc = nil
	file_tensorflow_core_framework_attr_value_proto_goTypes = nil
	file_tensorflow_core_framework_attr_value_proto_depIdxs = nil
}
