// Code generated by protoc-gen-go. DO NOT EDIT.
// source: LinuxWorker.proto

package LinuxWorkerPb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type StartRequest struct {
	//path of a program or executable
	Command string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	//arguments to invoke program or executable with
	Args []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	// environment variables for the execution
	//OPTIONAL variable
	Env []string `protobuf:"bytes,3,rep,name=env,proto3" json:"env,omitempty"`
	// current working directory of the execution
	//OPTIONAL variable, default will be the working directory of server
	Dir                  string   `protobuf:"bytes,4,opt,name=dir,proto3" json:"dir,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartRequest) Reset()         { *m = StartRequest{} }
func (m *StartRequest) String() string { return proto.CompactTextString(m) }
func (*StartRequest) ProtoMessage()    {}
func (*StartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d96209d9625940c, []int{0}
}

func (m *StartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartRequest.Unmarshal(m, b)
}
func (m *StartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartRequest.Marshal(b, m, deterministic)
}
func (m *StartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartRequest.Merge(m, src)
}
func (m *StartRequest) XXX_Size() int {
	return xxx_messageInfo_StartRequest.Size(m)
}
func (m *StartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartRequest proto.InternalMessageInfo

func (m *StartRequest) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *StartRequest) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *StartRequest) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *StartRequest) GetDir() string {
	if m != nil {
		return m.Dir
	}
	return ""
}

type StartResponse struct {
	//process Id of the running process executed by the command
	//will be 0 if process failed to execute
	Pid int32 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	//univeral unique identifier that tags each unique request made to server
	Uuid string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	//starting time of start request
	StartTimeStamp       string   `protobuf:"bytes,3,opt,name=startTimeStamp,proto3" json:"startTimeStamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartResponse) Reset()         { *m = StartResponse{} }
func (m *StartResponse) String() string { return proto.CompactTextString(m) }
func (*StartResponse) ProtoMessage()    {}
func (*StartResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d96209d9625940c, []int{1}
}

func (m *StartResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartResponse.Unmarshal(m, b)
}
func (m *StartResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartResponse.Marshal(b, m, deterministic)
}
func (m *StartResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartResponse.Merge(m, src)
}
func (m *StartResponse) XXX_Size() int {
	return xxx_messageInfo_StartResponse.Size(m)
}
func (m *StartResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartResponse proto.InternalMessageInfo

func (m *StartResponse) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *StartResponse) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *StartResponse) GetStartTimeStamp() string {
	if m != nil {
		return m.StartTimeStamp
	}
	return ""
}

type StopRequest struct {
	Pid                  int32    `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopRequest) Reset()         { *m = StopRequest{} }
func (m *StopRequest) String() string { return proto.CompactTextString(m) }
func (*StopRequest) ProtoMessage()    {}
func (*StopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d96209d9625940c, []int{2}
}

func (m *StopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopRequest.Unmarshal(m, b)
}
func (m *StopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopRequest.Marshal(b, m, deterministic)
}
func (m *StopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopRequest.Merge(m, src)
}
func (m *StopRequest) XXX_Size() int {
	return xxx_messageInfo_StopRequest.Size(m)
}
func (m *StopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopRequest proto.InternalMessageInfo

func (m *StopRequest) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

type StopResponse struct {
	Stdout               []byte   `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr               []byte   `protobuf:"bytes,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
	IsKilled             bool     `protobuf:"varint,3,opt,name=isKilled,proto3" json:"isKilled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopResponse) Reset()         { *m = StopResponse{} }
func (m *StopResponse) String() string { return proto.CompactTextString(m) }
func (*StopResponse) ProtoMessage()    {}
func (*StopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d96209d9625940c, []int{3}
}

func (m *StopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopResponse.Unmarshal(m, b)
}
func (m *StopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopResponse.Marshal(b, m, deterministic)
}
func (m *StopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopResponse.Merge(m, src)
}
func (m *StopResponse) XXX_Size() int {
	return xxx_messageInfo_StopResponse.Size(m)
}
func (m *StopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StopResponse proto.InternalMessageInfo

func (m *StopResponse) GetStdout() []byte {
	if m != nil {
		return m.Stdout
	}
	return nil
}

func (m *StopResponse) GetStderr() []byte {
	if m != nil {
		return m.Stderr
	}
	return nil
}

func (m *StopResponse) GetIsKilled() bool {
	if m != nil {
		return m.IsKilled
	}
	return false
}

type ProcessInfo struct {
	Pid                  int32    `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	StartTimeStamp       string   `protobuf:"bytes,2,opt,name=startTimeStamp,proto3" json:"startTimeStamp,omitempty"`
	EndTimeStamp         string   `protobuf:"bytes,3,opt,name=endTimeStamp,proto3" json:"endTimeStamp,omitempty"`
	ProcessName          string   `protobuf:"bytes,4,opt,name=processName,proto3" json:"processName,omitempty"`
	Uuid                 string   `protobuf:"bytes,5,opt,name=uuid,proto3" json:"uuid,omitempty"`
	IsRunning            bool     `protobuf:"varint,6,opt,name=isRunning,proto3" json:"isRunning,omitempty"`
	ExitCode             int32    `protobuf:"varint,7,opt,name=exitCode,proto3" json:"exitCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessInfo) Reset()         { *m = ProcessInfo{} }
func (m *ProcessInfo) String() string { return proto.CompactTextString(m) }
func (*ProcessInfo) ProtoMessage()    {}
func (*ProcessInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d96209d9625940c, []int{4}
}

func (m *ProcessInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessInfo.Unmarshal(m, b)
}
func (m *ProcessInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessInfo.Marshal(b, m, deterministic)
}
func (m *ProcessInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessInfo.Merge(m, src)
}
func (m *ProcessInfo) XXX_Size() int {
	return xxx_messageInfo_ProcessInfo.Size(m)
}
func (m *ProcessInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessInfo proto.InternalMessageInfo

func (m *ProcessInfo) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *ProcessInfo) GetStartTimeStamp() string {
	if m != nil {
		return m.StartTimeStamp
	}
	return ""
}

func (m *ProcessInfo) GetEndTimeStamp() string {
	if m != nil {
		return m.EndTimeStamp
	}
	return ""
}

func (m *ProcessInfo) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *ProcessInfo) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *ProcessInfo) GetIsRunning() bool {
	if m != nil {
		return m.IsRunning
	}
	return false
}

func (m *ProcessInfo) GetExitCode() int32 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

type QueryOneProcessRequest struct {
	Pid                  int32    `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	StartTimeStamp       string   `protobuf:"bytes,2,opt,name=startTimeStamp,proto3" json:"startTimeStamp,omitempty"`
	Uuid                 string   `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryOneProcessRequest) Reset()         { *m = QueryOneProcessRequest{} }
func (m *QueryOneProcessRequest) String() string { return proto.CompactTextString(m) }
func (*QueryOneProcessRequest) ProtoMessage()    {}
func (*QueryOneProcessRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d96209d9625940c, []int{5}
}

func (m *QueryOneProcessRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryOneProcessRequest.Unmarshal(m, b)
}
func (m *QueryOneProcessRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryOneProcessRequest.Marshal(b, m, deterministic)
}
func (m *QueryOneProcessRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryOneProcessRequest.Merge(m, src)
}
func (m *QueryOneProcessRequest) XXX_Size() int {
	return xxx_messageInfo_QueryOneProcessRequest.Size(m)
}
func (m *QueryOneProcessRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryOneProcessRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryOneProcessRequest proto.InternalMessageInfo

func (m *QueryOneProcessRequest) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *QueryOneProcessRequest) GetStartTimeStamp() string {
	if m != nil {
		return m.StartTimeStamp
	}
	return ""
}

func (m *QueryOneProcessRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type QueryOneProcessResponse struct {
	ProcInfo             *ProcessInfo `protobuf:"bytes,1,opt,name=procInfo,proto3" json:"procInfo,omitempty"`
	Stdout               []byte       `protobuf:"bytes,2,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr               []byte       `protobuf:"bytes,3,opt,name=stderr,proto3" json:"stderr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *QueryOneProcessResponse) Reset()         { *m = QueryOneProcessResponse{} }
func (m *QueryOneProcessResponse) String() string { return proto.CompactTextString(m) }
func (*QueryOneProcessResponse) ProtoMessage()    {}
func (*QueryOneProcessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d96209d9625940c, []int{6}
}

func (m *QueryOneProcessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryOneProcessResponse.Unmarshal(m, b)
}
func (m *QueryOneProcessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryOneProcessResponse.Marshal(b, m, deterministic)
}
func (m *QueryOneProcessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryOneProcessResponse.Merge(m, src)
}
func (m *QueryOneProcessResponse) XXX_Size() int {
	return xxx_messageInfo_QueryOneProcessResponse.Size(m)
}
func (m *QueryOneProcessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryOneProcessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryOneProcessResponse proto.InternalMessageInfo

func (m *QueryOneProcessResponse) GetProcInfo() *ProcessInfo {
	if m != nil {
		return m.ProcInfo
	}
	return nil
}

func (m *QueryOneProcessResponse) GetStdout() []byte {
	if m != nil {
		return m.Stdout
	}
	return nil
}

func (m *QueryOneProcessResponse) GetStderr() []byte {
	if m != nil {
		return m.Stderr
	}
	return nil
}

type QueryRunningProcessesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryRunningProcessesRequest) Reset()         { *m = QueryRunningProcessesRequest{} }
func (m *QueryRunningProcessesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRunningProcessesRequest) ProtoMessage()    {}
func (*QueryRunningProcessesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d96209d9625940c, []int{7}
}

func (m *QueryRunningProcessesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRunningProcessesRequest.Unmarshal(m, b)
}
func (m *QueryRunningProcessesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRunningProcessesRequest.Marshal(b, m, deterministic)
}
func (m *QueryRunningProcessesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRunningProcessesRequest.Merge(m, src)
}
func (m *QueryRunningProcessesRequest) XXX_Size() int {
	return xxx_messageInfo_QueryRunningProcessesRequest.Size(m)
}
func (m *QueryRunningProcessesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRunningProcessesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRunningProcessesRequest proto.InternalMessageInfo

type QueryRunningProcessesResponse struct {
	ProcessTable         []*ProcessInfo `protobuf:"bytes,1,rep,name=processTable,proto3" json:"processTable,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *QueryRunningProcessesResponse) Reset()         { *m = QueryRunningProcessesResponse{} }
func (m *QueryRunningProcessesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryRunningProcessesResponse) ProtoMessage()    {}
func (*QueryRunningProcessesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d96209d9625940c, []int{8}
}

func (m *QueryRunningProcessesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRunningProcessesResponse.Unmarshal(m, b)
}
func (m *QueryRunningProcessesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRunningProcessesResponse.Marshal(b, m, deterministic)
}
func (m *QueryRunningProcessesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRunningProcessesResponse.Merge(m, src)
}
func (m *QueryRunningProcessesResponse) XXX_Size() int {
	return xxx_messageInfo_QueryRunningProcessesResponse.Size(m)
}
func (m *QueryRunningProcessesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRunningProcessesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRunningProcessesResponse proto.InternalMessageInfo

func (m *QueryRunningProcessesResponse) GetProcessTable() []*ProcessInfo {
	if m != nil {
		return m.ProcessTable
	}
	return nil
}

func init() {
	proto.RegisterType((*StartRequest)(nil), "linuxWorker.StartRequest")
	proto.RegisterType((*StartResponse)(nil), "linuxWorker.StartResponse")
	proto.RegisterType((*StopRequest)(nil), "linuxWorker.StopRequest")
	proto.RegisterType((*StopResponse)(nil), "linuxWorker.StopResponse")
	proto.RegisterType((*ProcessInfo)(nil), "linuxWorker.ProcessInfo")
	proto.RegisterType((*QueryOneProcessRequest)(nil), "linuxWorker.QueryOneProcessRequest")
	proto.RegisterType((*QueryOneProcessResponse)(nil), "linuxWorker.QueryOneProcessResponse")
	proto.RegisterType((*QueryRunningProcessesRequest)(nil), "linuxWorker.QueryRunningProcessesRequest")
	proto.RegisterType((*QueryRunningProcessesResponse)(nil), "linuxWorker.QueryRunningProcessesResponse")
}

func init() { proto.RegisterFile("LinuxWorker.proto", fileDescriptor_1d96209d9625940c) }

var fileDescriptor_1d96209d9625940c = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0xa5, 0xcd, 0x7e, 0xb4, 0x93, 0x2c, 0x1f, 0x3e, 0x14, 0x6f, 0x54, 0xa0, 0x32, 0x08, 0x2d,
	0x1c, 0x7a, 0x58, 0x38, 0x72, 0xe2, 0x43, 0x08, 0x81, 0x60, 0x49, 0x57, 0x42, 0x5a, 0x69, 0x25,
	0xd2, 0x66, 0x76, 0x65, 0xd1, 0xda, 0xc1, 0x76, 0x96, 0x72, 0xe2, 0x87, 0xf2, 0x5b, 0x90, 0x90,
	0x5d, 0x37, 0x4d, 0xd3, 0x74, 0xb5, 0xb7, 0xf1, 0xf3, 0xf8, 0xcd, 0x7b, 0x33, 0x99, 0xc0, 0xbd,
	0x4f, 0x5c, 0x14, 0xf3, 0x6f, 0x52, 0xfd, 0x40, 0x35, 0xcc, 0x95, 0x34, 0x92, 0x84, 0xd3, 0x15,
	0xc4, 0xbe, 0x43, 0x34, 0x32, 0xa9, 0x32, 0x09, 0xfe, 0x2c, 0x50, 0x1b, 0x42, 0x61, 0x7f, 0x22,
	0x67, 0xb3, 0x54, 0x64, 0xb4, 0x35, 0x68, 0x1d, 0x75, 0x93, 0xe5, 0x91, 0x10, 0xd8, 0x49, 0xd5,
	0xa5, 0xa6, 0xed, 0x41, 0x70, 0xd4, 0x4d, 0x5c, 0x4c, 0xee, 0x42, 0x80, 0xe2, 0x8a, 0x06, 0x0e,
	0xb2, 0xa1, 0x45, 0x32, 0xae, 0xe8, 0x8e, 0x7b, 0x6b, 0x43, 0x76, 0x0e, 0x07, 0xbe, 0x82, 0xce,
	0xa5, 0xd0, 0x68, 0x53, 0x72, 0xbe, 0xa0, 0xdf, 0x4d, 0x6c, 0x68, 0xa9, 0x8b, 0x82, 0x67, 0xb4,
	0xed, 0x5e, 0xb9, 0x98, 0x3c, 0x85, 0xdb, 0xda, 0x3e, 0x3b, 0xe5, 0x33, 0x1c, 0x99, 0x74, 0x96,
	0xd3, 0xc0, 0xdd, 0xd6, 0x50, 0xf6, 0x08, 0xc2, 0x91, 0x91, 0xf9, 0x52, 0xff, 0x06, 0x39, 0x3b,
	0xb3, 0x0e, 0x6d, 0x82, 0x2f, 0xdf, 0x83, 0x3d, 0x6d, 0x32, 0x59, 0x18, 0x97, 0x14, 0x25, 0xfe,
	0xe4, 0x71, 0x54, 0xca, 0xc9, 0x58, 0xe0, 0xa8, 0x14, 0x89, 0xa1, 0xc3, 0xf5, 0x47, 0x3e, 0x9d,
	0x62, 0xe6, 0x24, 0x74, 0x92, 0xf2, 0xcc, 0xfe, 0xb6, 0x20, 0x3c, 0x51, 0x72, 0x82, 0x5a, 0x7f,
	0x10, 0x17, 0xb2, 0xc1, 0xda, 0xa6, 0x8d, 0x76, 0x93, 0x0d, 0xc2, 0x20, 0x42, 0x91, 0xd5, 0xcd,
	0xae, 0x61, 0x64, 0x00, 0x61, 0xbe, 0x28, 0xf6, 0x39, 0x9d, 0xa1, 0xef, 0x71, 0x15, 0x2a, 0x1b,
	0xb9, 0x5b, 0x69, 0x64, 0x1f, 0xba, 0x5c, 0x27, 0x85, 0x10, 0x5c, 0x5c, 0xd2, 0x3d, 0x67, 0x60,
	0x05, 0x58, 0x77, 0x38, 0xe7, 0xe6, 0x8d, 0xcc, 0x90, 0xee, 0x3b, 0xd9, 0xe5, 0x99, 0x5d, 0x40,
	0xef, 0x6b, 0x81, 0xea, 0xf7, 0x17, 0x81, 0xde, 0xe4, 0xd6, 0x2e, 0xdf, 0xd8, 0xe7, 0x52, 0x61,
	0xb0, 0x52, 0xc8, 0xfe, 0xc0, 0xfd, 0x8d, 0x3a, 0x7e, 0x58, 0x2f, 0xa1, 0x63, 0xfd, 0xd9, 0xe6,
	0xba, 0x6a, 0xe1, 0x31, 0x1d, 0x56, 0x3e, 0xdf, 0x61, 0xa5, 0xf9, 0x49, 0x99, 0x59, 0x19, 0x71,
	0x7b, 0xcb, 0x88, 0x83, 0xea, 0x88, 0xd9, 0x43, 0xe8, 0x3b, 0x01, 0xbe, 0x29, 0x9e, 0x14, 0x97,
	0x76, 0xd9, 0x39, 0x3c, 0xd8, 0x72, 0xef, 0x65, 0xbe, 0x82, 0xc8, 0x8f, 0xe1, 0x34, 0x1d, 0x4f,
	0x91, 0xb6, 0x06, 0xc1, 0xb5, 0x52, 0xd7, 0xb2, 0x8f, 0xff, 0xb5, 0x21, 0x72, 0x6b, 0x3a, 0x42,
	0x75, 0xc5, 0x27, 0x48, 0xde, 0x43, 0xf4, 0x6e, 0x8e, 0x93, 0xc2, 0xa0, 0xdb, 0x1c, 0x72, 0xb8,
	0x46, 0x54, 0xdd, 0xd7, 0x38, 0x6e, 0xba, 0x5a, 0xa8, 0x62, 0xb7, 0xc8, 0x5b, 0x08, 0x4b, 0x22,
	0x99, 0x13, 0x5a, 0x4b, 0x2e, 0xd7, 0x26, 0x3e, 0x6c, 0xb8, 0x29, 0x59, 0x10, 0x7a, 0x9e, 0xa5,
	0x36, 0x26, 0xf2, 0x78, 0xed, 0x59, 0xf3, 0xc7, 0x12, 0x3f, 0xb9, 0x3e, 0xa9, 0x2c, 0xf3, 0x0b,
	0xfa, 0xd5, 0x32, 0xf5, 0x66, 0x93, 0x67, 0x9b, 0x3c, 0x5b, 0x06, 0x16, 0x3f, 0xbf, 0x49, 0xea,
	0xb2, 0xf0, 0xeb, 0x3b, 0x67, 0x07, 0x95, 0xbf, 0xe4, 0xc9, 0x78, 0xbc, 0xe7, 0x7e, 0x94, 0x2f,
	0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0xa9, 0xf3, 0xfc, 0x72, 0x3d, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LinuxServiceClient is the client API for LinuxService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LinuxServiceClient interface {
	ExecuteStart(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	ExecuteStop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error)
	ExecuteQueryOneProcess(ctx context.Context, in *QueryOneProcessRequest, opts ...grpc.CallOption) (*QueryOneProcessResponse, error)
	ExecuteQueryRunningProcesses(ctx context.Context, in *QueryRunningProcessesRequest, opts ...grpc.CallOption) (*QueryRunningProcessesResponse, error)
}

type linuxServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLinuxServiceClient(cc grpc.ClientConnInterface) LinuxServiceClient {
	return &linuxServiceClient{cc}
}

func (c *linuxServiceClient) ExecuteStart(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, "/linuxWorker.LinuxService/ExecuteStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linuxServiceClient) ExecuteStop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := c.cc.Invoke(ctx, "/linuxWorker.LinuxService/ExecuteStop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linuxServiceClient) ExecuteQueryOneProcess(ctx context.Context, in *QueryOneProcessRequest, opts ...grpc.CallOption) (*QueryOneProcessResponse, error) {
	out := new(QueryOneProcessResponse)
	err := c.cc.Invoke(ctx, "/linuxWorker.LinuxService/ExecuteQueryOneProcess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linuxServiceClient) ExecuteQueryRunningProcesses(ctx context.Context, in *QueryRunningProcessesRequest, opts ...grpc.CallOption) (*QueryRunningProcessesResponse, error) {
	out := new(QueryRunningProcessesResponse)
	err := c.cc.Invoke(ctx, "/linuxWorker.LinuxService/ExecuteQueryRunningProcesses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinuxServiceServer is the server API for LinuxService service.
type LinuxServiceServer interface {
	ExecuteStart(context.Context, *StartRequest) (*StartResponse, error)
	ExecuteStop(context.Context, *StopRequest) (*StopResponse, error)
	ExecuteQueryOneProcess(context.Context, *QueryOneProcessRequest) (*QueryOneProcessResponse, error)
	ExecuteQueryRunningProcesses(context.Context, *QueryRunningProcessesRequest) (*QueryRunningProcessesResponse, error)
}

// UnimplementedLinuxServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLinuxServiceServer struct {
}

func (*UnimplementedLinuxServiceServer) ExecuteStart(ctx context.Context, req *StartRequest) (*StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteStart not implemented")
}
func (*UnimplementedLinuxServiceServer) ExecuteStop(ctx context.Context, req *StopRequest) (*StopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteStop not implemented")
}
func (*UnimplementedLinuxServiceServer) ExecuteQueryOneProcess(ctx context.Context, req *QueryOneProcessRequest) (*QueryOneProcessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteQueryOneProcess not implemented")
}
func (*UnimplementedLinuxServiceServer) ExecuteQueryRunningProcesses(ctx context.Context, req *QueryRunningProcessesRequest) (*QueryRunningProcessesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteQueryRunningProcesses not implemented")
}

func RegisterLinuxServiceServer(s *grpc.Server, srv LinuxServiceServer) {
	s.RegisterService(&_LinuxService_serviceDesc, srv)
}

func _LinuxService_ExecuteStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinuxServiceServer).ExecuteStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linuxWorker.LinuxService/ExecuteStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinuxServiceServer).ExecuteStart(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinuxService_ExecuteStop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinuxServiceServer).ExecuteStop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linuxWorker.LinuxService/ExecuteStop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinuxServiceServer).ExecuteStop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinuxService_ExecuteQueryOneProcess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryOneProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinuxServiceServer).ExecuteQueryOneProcess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linuxWorker.LinuxService/ExecuteQueryOneProcess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinuxServiceServer).ExecuteQueryOneProcess(ctx, req.(*QueryOneProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinuxService_ExecuteQueryRunningProcesses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRunningProcessesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinuxServiceServer).ExecuteQueryRunningProcesses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linuxWorker.LinuxService/ExecuteQueryRunningProcesses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinuxServiceServer).ExecuteQueryRunningProcesses(ctx, req.(*QueryRunningProcessesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LinuxService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "linuxWorker.LinuxService",
	HandlerType: (*LinuxServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteStart",
			Handler:    _LinuxService_ExecuteStart_Handler,
		},
		{
			MethodName: "ExecuteStop",
			Handler:    _LinuxService_ExecuteStop_Handler,
		},
		{
			MethodName: "ExecuteQueryOneProcess",
			Handler:    _LinuxService_ExecuteQueryOneProcess_Handler,
		},
		{
			MethodName: "ExecuteQueryRunningProcesses",
			Handler:    _LinuxService_ExecuteQueryRunningProcesses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "LinuxWorker.proto",
}
