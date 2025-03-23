// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: internal/ffmpeg/pb/ffmpeg.proto

package pb

import (
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

type PrepareConversionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Filename               string  `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Processor              int32   `protobuf:"varint,3,opt,name=processor,proto3" json:"processor,omitempty"`
	Crf                    *int32  `protobuf:"varint,4,opt,name=crf,proto3,oneof" json:"crf,omitempty"`
	Preset                 *int32  `protobuf:"varint,5,opt,name=preset,proto3,oneof" json:"preset,omitempty"`
	Quality                *int32  `protobuf:"varint,6,opt,name=quality,proto3,oneof" json:"quality,omitempty"`
	AdditionalFfmpegParams *string `protobuf:"bytes,7,opt,name=additional_ffmpeg_params,json=additionalFfmpegParams,proto3,oneof" json:"additional_ffmpeg_params,omitempty"`
}

func (x *PrepareConversionRequest) Reset() {
	*x = PrepareConversionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_ffmpeg_pb_ffmpeg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepareConversionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepareConversionRequest) ProtoMessage() {}

func (x *PrepareConversionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_ffmpeg_pb_ffmpeg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepareConversionRequest.ProtoReflect.Descriptor instead.
func (*PrepareConversionRequest) Descriptor() ([]byte, []int) {
	return file_internal_ffmpeg_pb_ffmpeg_proto_rawDescGZIP(), []int{0}
}

func (x *PrepareConversionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PrepareConversionRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *PrepareConversionRequest) GetProcessor() int32 {
	if x != nil {
		return x.Processor
	}
	return 0
}

func (x *PrepareConversionRequest) GetCrf() int32 {
	if x != nil && x.Crf != nil {
		return *x.Crf
	}
	return 0
}

func (x *PrepareConversionRequest) GetPreset() int32 {
	if x != nil && x.Preset != nil {
		return *x.Preset
	}
	return 0
}

func (x *PrepareConversionRequest) GetQuality() int32 {
	if x != nil && x.Quality != nil {
		return *x.Quality
	}
	return 0
}

func (x *PrepareConversionRequest) GetAdditionalFfmpegParams() string {
	if x != nil && x.AdditionalFfmpegParams != nil {
		return *x.AdditionalFfmpegParams
	}
	return ""
}

type PrepareConversionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Processor              int32  `protobuf:"varint,2,opt,name=processor,proto3" json:"processor,omitempty"`
	AdditionalFfmpegParams string `protobuf:"bytes,3,opt,name=additional_ffmpeg_params,json=additionalFfmpegParams,proto3" json:"additional_ffmpeg_params,omitempty"`
	TempFilename           string `protobuf:"bytes,4,opt,name=temp_filename,json=tempFilename,proto3" json:"temp_filename,omitempty"`
}

func (x *PrepareConversionResponse) Reset() {
	*x = PrepareConversionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_ffmpeg_pb_ffmpeg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepareConversionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepareConversionResponse) ProtoMessage() {}

func (x *PrepareConversionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_ffmpeg_pb_ffmpeg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepareConversionResponse.ProtoReflect.Descriptor instead.
func (*PrepareConversionResponse) Descriptor() ([]byte, []int) {
	return file_internal_ffmpeg_pb_ffmpeg_proto_rawDescGZIP(), []int{1}
}

func (x *PrepareConversionResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PrepareConversionResponse) GetProcessor() int32 {
	if x != nil {
		return x.Processor
	}
	return 0
}

func (x *PrepareConversionResponse) GetAdditionalFfmpegParams() string {
	if x != nil {
		return x.AdditionalFfmpegParams
	}
	return ""
}

func (x *PrepareConversionResponse) GetTempFilename() string {
	if x != nil {
		return x.TempFilename
	}
	return ""
}

type ConversionContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FileContent []byte `protobuf:"bytes,2,opt,name=file_content,json=fileContent,proto3" json:"file_content,omitempty"`
}

func (x *ConversionContent) Reset() {
	*x = ConversionContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_ffmpeg_pb_ffmpeg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversionContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionContent) ProtoMessage() {}

func (x *ConversionContent) ProtoReflect() protoreflect.Message {
	mi := &file_internal_ffmpeg_pb_ffmpeg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionContent.ProtoReflect.Descriptor instead.
func (*ConversionContent) Descriptor() ([]byte, []int) {
	return file_internal_ffmpeg_pb_ffmpeg_proto_rawDescGZIP(), []int{2}
}

func (x *ConversionContent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ConversionContent) GetFileContent() []byte {
	if x != nil {
		return x.FileContent
	}
	return nil
}

type ConversionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FfmpegOutput []byte `protobuf:"bytes,1,opt,name=ffmpeg_output,json=ffmpegOutput,proto3" json:"ffmpeg_output,omitempty"`
	ElapsedTime  int32  `protobuf:"varint,2,opt,name=elapsed_time,json=elapsedTime,proto3" json:"elapsed_time,omitempty"`
}

func (x *ConversionResponse) Reset() {
	*x = ConversionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_ffmpeg_pb_ffmpeg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionResponse) ProtoMessage() {}

func (x *ConversionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_ffmpeg_pb_ffmpeg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionResponse.ProtoReflect.Descriptor instead.
func (*ConversionResponse) Descriptor() ([]byte, []int) {
	return file_internal_ffmpeg_pb_ffmpeg_proto_rawDescGZIP(), []int{3}
}

func (x *ConversionResponse) GetFfmpegOutput() []byte {
	if x != nil {
		return x.FfmpegOutput
	}
	return nil
}

func (x *ConversionResponse) GetElapsedTime() int32 {
	if x != nil {
		return x.ElapsedTime
	}
	return 0
}

type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_ffmpeg_pb_ffmpeg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_internal_ffmpeg_pb_ffmpeg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_internal_ffmpeg_pb_ffmpeg_proto_rawDescGZIP(), []int{4}
}

func (x *Query) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Progress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BitRate string  `protobuf:"bytes,2,opt,name=bit_rate,json=bitRate,proto3" json:"bit_rate,omitempty"`
	Ratio   float64 `protobuf:"fixed64,3,opt,name=ratio,proto3" json:"ratio,omitempty"`
	Fps     int32   `protobuf:"varint,4,opt,name=fps,proto3" json:"fps,omitempty"`
	Q       int32   `protobuf:"varint,5,opt,name=q,proto3" json:"q,omitempty"`
}

func (x *Progress) Reset() {
	*x = Progress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_ffmpeg_pb_ffmpeg_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Progress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Progress) ProtoMessage() {}

func (x *Progress) ProtoReflect() protoreflect.Message {
	mi := &file_internal_ffmpeg_pb_ffmpeg_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Progress.ProtoReflect.Descriptor instead.
func (*Progress) Descriptor() ([]byte, []int) {
	return file_internal_ffmpeg_pb_ffmpeg_proto_rawDescGZIP(), []int{5}
}

func (x *Progress) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Progress) GetBitRate() string {
	if x != nil {
		return x.BitRate
	}
	return ""
}

func (x *Progress) GetRatio() float64 {
	if x != nil {
		return x.Ratio
	}
	return 0
}

func (x *Progress) GetFps() int32 {
	if x != nil {
		return x.Fps
	}
	return 0
}

func (x *Progress) GetQ() int32 {
	if x != nil {
		return x.Q
	}
	return 0
}

var File_internal_ffmpeg_pb_ffmpeg_proto protoreflect.FileDescriptor

var file_internal_ffmpeg_pb_ffmpeg_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x66, 0x66, 0x6d, 0x70, 0x65,
	0x67, 0x2f, 0x70, 0x62, 0x2f, 0x66, 0x66, 0x6d, 0x70, 0x65, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x66, 0x66, 0x6d, 0x70, 0x65, 0x67, 0x22, 0xb2, 0x02, 0x0a, 0x18, 0x50, 0x72,
	0x65, 0x70, 0x61, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72,
	0x12, 0x15, 0x0a, 0x03, 0x63, 0x72, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52,
	0x03, 0x63, 0x72, 0x66, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x06, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x07, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a, 0x18, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x66, 0x66, 0x6d, 0x70, 0x65, 0x67, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x16, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x46, 0x66, 0x6d, 0x70, 0x65, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x88,
	0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x63, 0x72, 0x66, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x42, 0x1b, 0x0a, 0x19, 0x5f, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x5f, 0x66, 0x66, 0x6d, 0x70, 0x65, 0x67, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0xa8,
	0x01, 0x0a, 0x19, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x12, 0x38, 0x0a, 0x18, 0x61, 0x64,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x66, 0x66, 0x6d, 0x70, 0x65, 0x67, 0x5f,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x61, 0x64,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x46, 0x66, 0x6d, 0x70, 0x65, 0x67, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d,
	0x70, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x46, 0x0a, 0x11, 0x43, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0x5c, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x66, 0x6d, 0x70, 0x65,
	0x67, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c,
	0x66, 0x66, 0x6d, 0x70, 0x65, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x17, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6b, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x69, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x69, 0x74, 0x52, 0x61, 0x74, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x70, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x66, 0x70, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x71, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x01, 0x71, 0x32, 0x92, 0x02, 0x0a, 0x06, 0x46, 0x46, 0x6d, 0x70, 0x65, 0x67,
	0x12, 0x58, 0x0a, 0x11, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x66, 0x66, 0x6d, 0x70, 0x65, 0x67, 0x2e, 0x50,
	0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x66, 0x66, 0x6d, 0x70, 0x65, 0x67,
	0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0f, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e,
	0x66, 0x66, 0x6d, 0x70, 0x65, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x1a, 0x2e, 0x66, 0x66, 0x6d, 0x70, 0x65,
	0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x2e, 0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x70,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x2e, 0x66, 0x66, 0x6d,
	0x70, 0x65, 0x67, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0d, 0x2e, 0x66, 0x66, 0x6d, 0x70,
	0x65, 0x67, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x30, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0d, 0x2e, 0x66, 0x66, 0x6d, 0x70, 0x65, 0x67,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x10, 0x2e, 0x66, 0x66, 0x6d, 0x70, 0x65, 0x67, 0x2e,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x30, 0x01, 0x42, 0x14, 0x5a, 0x12, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x66, 0x66, 0x6d, 0x70, 0x65, 0x67, 0x3b, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_ffmpeg_pb_ffmpeg_proto_rawDescOnce sync.Once
	file_internal_ffmpeg_pb_ffmpeg_proto_rawDescData = file_internal_ffmpeg_pb_ffmpeg_proto_rawDesc
)

func file_internal_ffmpeg_pb_ffmpeg_proto_rawDescGZIP() []byte {
	file_internal_ffmpeg_pb_ffmpeg_proto_rawDescOnce.Do(func() {
		file_internal_ffmpeg_pb_ffmpeg_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_ffmpeg_pb_ffmpeg_proto_rawDescData)
	})
	return file_internal_ffmpeg_pb_ffmpeg_proto_rawDescData
}

var file_internal_ffmpeg_pb_ffmpeg_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internal_ffmpeg_pb_ffmpeg_proto_goTypes = []interface{}{
	(*PrepareConversionRequest)(nil),  // 0: ffmpeg.PrepareConversionRequest
	(*PrepareConversionResponse)(nil), // 1: ffmpeg.PrepareConversionResponse
	(*ConversionContent)(nil),         // 2: ffmpeg.ConversionContent
	(*ConversionResponse)(nil),        // 3: ffmpeg.ConversionResponse
	(*Query)(nil),                     // 4: ffmpeg.Query
	(*Progress)(nil),                  // 5: ffmpeg.Progress
}
var file_internal_ffmpeg_pb_ffmpeg_proto_depIdxs = []int32{
	0, // 0: ffmpeg.FFmpeg.PrepareConversion:input_type -> ffmpeg.PrepareConversionRequest
	2, // 1: ffmpeg.FFmpeg.StartConversion:input_type -> ffmpeg.ConversionContent
	4, // 2: ffmpeg.FFmpeg.StopConversion:input_type -> ffmpeg.Query
	4, // 3: ffmpeg.FFmpeg.GetProgress:input_type -> ffmpeg.Query
	1, // 4: ffmpeg.FFmpeg.PrepareConversion:output_type -> ffmpeg.PrepareConversionResponse
	3, // 5: ffmpeg.FFmpeg.StartConversion:output_type -> ffmpeg.ConversionResponse
	4, // 6: ffmpeg.FFmpeg.StopConversion:output_type -> ffmpeg.Query
	5, // 7: ffmpeg.FFmpeg.GetProgress:output_type -> ffmpeg.Progress
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_ffmpeg_pb_ffmpeg_proto_init() }
func file_internal_ffmpeg_pb_ffmpeg_proto_init() {
	if File_internal_ffmpeg_pb_ffmpeg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_ffmpeg_pb_ffmpeg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrepareConversionRequest); i {
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
		file_internal_ffmpeg_pb_ffmpeg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrepareConversionResponse); i {
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
		file_internal_ffmpeg_pb_ffmpeg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversionContent); i {
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
		file_internal_ffmpeg_pb_ffmpeg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversionResponse); i {
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
		file_internal_ffmpeg_pb_ffmpeg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query); i {
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
		file_internal_ffmpeg_pb_ffmpeg_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Progress); i {
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
	file_internal_ffmpeg_pb_ffmpeg_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_ffmpeg_pb_ffmpeg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_ffmpeg_pb_ffmpeg_proto_goTypes,
		DependencyIndexes: file_internal_ffmpeg_pb_ffmpeg_proto_depIdxs,
		MessageInfos:      file_internal_ffmpeg_pb_ffmpeg_proto_msgTypes,
	}.Build()
	File_internal_ffmpeg_pb_ffmpeg_proto = out.File
	file_internal_ffmpeg_pb_ffmpeg_proto_rawDesc = nil
	file_internal_ffmpeg_pb_ffmpeg_proto_goTypes = nil
	file_internal_ffmpeg_pb_ffmpeg_proto_depIdxs = nil
}
