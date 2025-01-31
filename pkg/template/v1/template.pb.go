// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: pkg/template/v1/template.proto

package v1

import (
	v1 "github.com/GuanceCloud/iacker/pkg/resource/v1"
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

// Manifest is a template definitions that generates files from templates.
type Manifest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Inputs      *Inputs       `protobuf:"bytes,2,opt,name=inputs,proto3" json:"inputs,omitempty"`
	Outputs     *Outputs      `protobuf:"bytes,3,opt,name=outputs,proto3" json:"outputs,omitempty"`
	Diagnostics []*Diagnostic `protobuf:"bytes,4,rep,name=diagnostics,proto3" json:"diagnostics,omitempty"`
}

func (x *Manifest) Reset() {
	*x = Manifest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_template_v1_template_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Manifest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Manifest) ProtoMessage() {}

func (x *Manifest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_template_v1_template_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Manifest.ProtoReflect.Descriptor instead.
func (*Manifest) Descriptor() ([]byte, []int) {
	return file_pkg_template_v1_template_proto_rawDescGZIP(), []int{0}
}

func (x *Manifest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Manifest) GetInputs() *Inputs {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *Manifest) GetOutputs() *Outputs {
	if x != nil {
		return x.Outputs
	}
	return nil
}

func (x *Manifest) GetDiagnostics() []*Diagnostic {
	if x != nil {
		return x.Diagnostics
	}
	return nil
}

// Inputs is a set of inputs for a template.
type Inputs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vars      map[string]string       `protobuf:"bytes,1,rep,name=vars,proto3" json:"vars,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Resources map[string]*v1.Resource `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Errors    map[string]*v1.Error    `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Inputs) Reset() {
	*x = Inputs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_template_v1_template_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Inputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inputs) ProtoMessage() {}

func (x *Inputs) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_template_v1_template_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inputs.ProtoReflect.Descriptor instead.
func (*Inputs) Descriptor() ([]byte, []int) {
	return file_pkg_template_v1_template_proto_rawDescGZIP(), []int{1}
}

func (x *Inputs) GetVars() map[string]string {
	if x != nil {
		return x.Vars
	}
	return nil
}

func (x *Inputs) GetResources() map[string]*v1.Resource {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *Inputs) GetErrors() map[string]*v1.Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

// Outputs is a set of outputs for a template.
type Outputs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files map[string]*File `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Outputs) Reset() {
	*x = Outputs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_template_v1_template_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Outputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outputs) ProtoMessage() {}

func (x *Outputs) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_template_v1_template_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outputs.ProtoReflect.Descriptor instead.
func (*Outputs) Descriptor() ([]byte, []int) {
	return file_pkg_template_v1_template_proto_rawDescGZIP(), []int{2}
}

func (x *Outputs) GetFiles() map[string]*File {
	if x != nil {
		return x.Files
	}
	return nil
}

// Diagnostic is a diagnostic message.
type Diagnostic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message  string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Severity int32  `protobuf:"varint,2,opt,name=severity,proto3" json:"severity,omitempty"`
}

func (x *Diagnostic) Reset() {
	*x = Diagnostic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_template_v1_template_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Diagnostic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Diagnostic) ProtoMessage() {}

func (x *Diagnostic) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_template_v1_template_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Diagnostic.ProtoReflect.Descriptor instead.
func (*Diagnostic) Descriptor() ([]byte, []int) {
	return file_pkg_template_v1_template_proto_rawDescGZIP(), []int{3}
}

func (x *Diagnostic) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Diagnostic) GetSeverity() int32 {
	if x != nil {
		return x.Severity
	}
	return 0
}

// Output is a template output.
type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_template_v1_template_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_template_v1_template_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_pkg_template_v1_template_proto_rawDescGZIP(), []int{4}
}

func (x *File) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_pkg_template_v1_template_proto protoreflect.FileDescriptor

var file_pkg_template_v1_template_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x70, 0x6b, 0x67, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x08, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x52, 0x06, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x52, 0x07,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x3d, 0x0a, 0x0b, 0x64, 0x69, 0x61, 0x67, 0x6e,
	0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70,
	0x6b, 0x67, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x52, 0x0b, 0x64, 0x69, 0x61, 0x67, 0x6e,
	0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x22, 0xa7, 0x03, 0x0a, 0x06, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x73, 0x12, 0x35, 0x0a, 0x04, 0x76, 0x61, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x2e, 0x56, 0x61, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x04, 0x76, 0x61, 0x72, 0x73, 0x12, 0x44, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x6b,
	0x67, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x3b,
	0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x56,
	0x61, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x57, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x51, 0x0a,
	0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x70, 0x6b, 0x67, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x95, 0x01, 0x0a, 0x07, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x39, 0x0a, 0x05,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x6b,
	0x67, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x1a, 0x4f, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x42, 0x0a, 0x0a, 0x44, 0x69, 0x61, 0x67,
	0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x22, 0x20, 0x0a, 0x04,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x32,
	0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x75, 0x61,
	0x6e, 0x63, 0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x69, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_template_v1_template_proto_rawDescOnce sync.Once
	file_pkg_template_v1_template_proto_rawDescData = file_pkg_template_v1_template_proto_rawDesc
)

func file_pkg_template_v1_template_proto_rawDescGZIP() []byte {
	file_pkg_template_v1_template_proto_rawDescOnce.Do(func() {
		file_pkg_template_v1_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_template_v1_template_proto_rawDescData)
	})
	return file_pkg_template_v1_template_proto_rawDescData
}

var file_pkg_template_v1_template_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_template_v1_template_proto_goTypes = []interface{}{
	(*Manifest)(nil),    // 0: pkg.template.v1.Manifest
	(*Inputs)(nil),      // 1: pkg.template.v1.Inputs
	(*Outputs)(nil),     // 2: pkg.template.v1.Outputs
	(*Diagnostic)(nil),  // 3: pkg.template.v1.Diagnostic
	(*File)(nil),        // 4: pkg.template.v1.File
	nil,                 // 5: pkg.template.v1.Inputs.VarsEntry
	nil,                 // 6: pkg.template.v1.Inputs.ResourcesEntry
	nil,                 // 7: pkg.template.v1.Inputs.ErrorsEntry
	nil,                 // 8: pkg.template.v1.Outputs.FilesEntry
	(*v1.Resource)(nil), // 9: pkg.resource.v1.Resource
	(*v1.Error)(nil),    // 10: pkg.resource.v1.Error
}
var file_pkg_template_v1_template_proto_depIdxs = []int32{
	1,  // 0: pkg.template.v1.Manifest.inputs:type_name -> pkg.template.v1.Inputs
	2,  // 1: pkg.template.v1.Manifest.outputs:type_name -> pkg.template.v1.Outputs
	3,  // 2: pkg.template.v1.Manifest.diagnostics:type_name -> pkg.template.v1.Diagnostic
	5,  // 3: pkg.template.v1.Inputs.vars:type_name -> pkg.template.v1.Inputs.VarsEntry
	6,  // 4: pkg.template.v1.Inputs.resources:type_name -> pkg.template.v1.Inputs.ResourcesEntry
	7,  // 5: pkg.template.v1.Inputs.errors:type_name -> pkg.template.v1.Inputs.ErrorsEntry
	8,  // 6: pkg.template.v1.Outputs.files:type_name -> pkg.template.v1.Outputs.FilesEntry
	9,  // 7: pkg.template.v1.Inputs.ResourcesEntry.value:type_name -> pkg.resource.v1.Resource
	10, // 8: pkg.template.v1.Inputs.ErrorsEntry.value:type_name -> pkg.resource.v1.Error
	4,  // 9: pkg.template.v1.Outputs.FilesEntry.value:type_name -> pkg.template.v1.File
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_pkg_template_v1_template_proto_init() }
func file_pkg_template_v1_template_proto_init() {
	if File_pkg_template_v1_template_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_template_v1_template_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Manifest); i {
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
		file_pkg_template_v1_template_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Inputs); i {
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
		file_pkg_template_v1_template_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Outputs); i {
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
		file_pkg_template_v1_template_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Diagnostic); i {
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
		file_pkg_template_v1_template_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_template_v1_template_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_template_v1_template_proto_goTypes,
		DependencyIndexes: file_pkg_template_v1_template_proto_depIdxs,
		MessageInfos:      file_pkg_template_v1_template_proto_msgTypes,
	}.Build()
	File_pkg_template_v1_template_proto = out.File
	file_pkg_template_v1_template_proto_rawDesc = nil
	file_pkg_template_v1_template_proto_goTypes = nil
	file_pkg_template_v1_template_proto_depIdxs = nil
}
