// Code generated by protoc-gen-go. DO NOT EDIT.
// source: LinuxJobWorker.proto

package LinuxJobWorkerPb

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
	return fileDescriptor_9656cae1b775fd8c, []int{0}
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
	return fileDescriptor_9656cae1b775fd8c, []int{1}
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
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopRequest) Reset()         { *m = StopRequest{} }
func (m *StopRequest) String() string { return proto.CompactTextString(m) }
func (*StopRequest) ProtoMessage()    {}
func (*StopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9656cae1b775fd8c, []int{2}
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

func (m *StopRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type StopResponse struct {
	//output of stdout.log
	Stdout []byte `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	//output of stderr.log
	Stderr []byte `protobuf:"bytes,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
	//time of job being killed
	EndTimeStamp string `protobuf:"bytes,3,opt,name=endTimeStamp,proto3" json:"endTimeStamp,omitempty"`
	//indication of process being killed
	IsKilled             bool     `protobuf:"varint,4,opt,name=isKilled,proto3" json:"isKilled,omitempty"`
	ExitCode             int32    `protobuf:"varint,5,opt,name=exitCode,proto3" json:"exitCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopResponse) Reset()         { *m = StopResponse{} }
func (m *StopResponse) String() string { return proto.CompactTextString(m) }
func (*StopResponse) ProtoMessage()    {}
func (*StopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9656cae1b775fd8c, []int{3}
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

func (m *StopResponse) GetEndTimeStamp() string {
	if m != nil {
		return m.EndTimeStamp
	}
	return ""
}

func (m *StopResponse) GetIsKilled() bool {
	if m != nil {
		return m.IsKilled
	}
	return false
}

func (m *StopResponse) GetExitCode() int32 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

type ProcessInfo struct {
	//process id assigned for process ran in server
	Pid            int32  `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	StartTimeStamp string `protobuf:"bytes,2,opt,name=startTimeStamp,proto3" json:"startTimeStamp,omitempty"`
	EndTimeStamp   string `protobuf:"bytes,3,opt,name=endTimeStamp,proto3" json:"endTimeStamp,omitempty"`
	ProcessName    string `protobuf:"bytes,4,opt,name=processName,proto3" json:"processName,omitempty"`
	//status of whether process is still running
	IsRunning bool `protobuf:"varint,5,opt,name=isRunning,proto3" json:"isRunning,omitempty"`
	//exitCode of process
	ExitCode             int32    `protobuf:"varint,6,opt,name=exitCode,proto3" json:"exitCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessInfo) Reset()         { *m = ProcessInfo{} }
func (m *ProcessInfo) String() string { return proto.CompactTextString(m) }
func (*ProcessInfo) ProtoMessage()    {}
func (*ProcessInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9656cae1b775fd8c, []int{4}
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
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryOneProcessRequest) Reset()         { *m = QueryOneProcessRequest{} }
func (m *QueryOneProcessRequest) String() string { return proto.CompactTextString(m) }
func (*QueryOneProcessRequest) ProtoMessage()    {}
func (*QueryOneProcessRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9656cae1b775fd8c, []int{5}
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
	return fileDescriptor_9656cae1b775fd8c, []int{6}
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
	return fileDescriptor_9656cae1b775fd8c, []int{7}
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
	ProcessTable         map[string]*ProcessInfo `protobuf:"bytes,1,rep,name=processTable,proto3" json:"processTable,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *QueryRunningProcessesResponse) Reset()         { *m = QueryRunningProcessesResponse{} }
func (m *QueryRunningProcessesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryRunningProcessesResponse) ProtoMessage()    {}
func (*QueryRunningProcessesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9656cae1b775fd8c, []int{8}
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

func (m *QueryRunningProcessesResponse) GetProcessTable() map[string]*ProcessInfo {
	if m != nil {
		return m.ProcessTable
	}
	return nil
}

func init() {
	proto.RegisterType((*StartRequest)(nil), "linuxJobWorker.StartRequest")
	proto.RegisterType((*StartResponse)(nil), "linuxJobWorker.StartResponse")
	proto.RegisterType((*StopRequest)(nil), "linuxJobWorker.StopRequest")
	proto.RegisterType((*StopResponse)(nil), "linuxJobWorker.StopResponse")
	proto.RegisterType((*ProcessInfo)(nil), "linuxJobWorker.ProcessInfo")
	proto.RegisterType((*QueryOneProcessRequest)(nil), "linuxJobWorker.QueryOneProcessRequest")
	proto.RegisterType((*QueryOneProcessResponse)(nil), "linuxJobWorker.QueryOneProcessResponse")
	proto.RegisterType((*QueryRunningProcessesRequest)(nil), "linuxJobWorker.QueryRunningProcessesRequest")
	proto.RegisterType((*QueryRunningProcessesResponse)(nil), "linuxJobWorker.QueryRunningProcessesResponse")
	proto.RegisterMapType((map[string]*ProcessInfo)(nil), "linuxJobWorker.QueryRunningProcessesResponse.ProcessTableEntry")
}

func init() { proto.RegisterFile("LinuxJobWorker.proto", fileDescriptor_9656cae1b775fd8c) }

var fileDescriptor_9656cae1b775fd8c = []byte{
	// 570 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x5e, 0x9a, 0xb5, 0xb4, 0xa7, 0x65, 0x0c, 0x8b, 0x8d, 0xa8, 0xeb, 0x50, 0xc9, 0xc5, 0xe8,
	0xc5, 0xa8, 0x44, 0xb9, 0x00, 0x71, 0x83, 0xc4, 0x04, 0x12, 0x3f, 0x82, 0x91, 0x4e, 0x42, 0x42,
	0x20, 0x91, 0x36, 0x87, 0xc9, 0x5a, 0x6b, 0x07, 0xdb, 0xa9, 0xd6, 0xdb, 0x3d, 0x06, 0x2f, 0xc4,
	0x7b, 0xf0, 0x24, 0xc8, 0x8e, 0x9b, 0x35, 0x69, 0xba, 0x8d, 0xbb, 0xe3, 0xe3, 0xcf, 0x9f, 0xbf,
	0xef, 0x9c, 0x63, 0xc3, 0xbd, 0x0f, 0x94, 0x25, 0xe7, 0xef, 0xf8, 0xe8, 0x0b, 0x17, 0x67, 0x28,
	0xfa, 0xb1, 0xe0, 0x8a, 0x93, 0xad, 0x49, 0x2e, 0xeb, 0xff, 0x80, 0xd6, 0x50, 0x85, 0x42, 0x05,
	0xf8, 0x2b, 0x41, 0xa9, 0x88, 0x07, 0xb7, 0xc6, 0x7c, 0x3a, 0x0d, 0x59, 0xe4, 0x39, 0x5d, 0xa7,
	0xd7, 0x08, 0x16, 0x4b, 0x42, 0x60, 0x33, 0x14, 0xa7, 0xd2, 0xab, 0x74, 0xdd, 0x5e, 0x23, 0x30,
	0x31, 0xd9, 0x06, 0x17, 0xd9, 0xcc, 0x73, 0x4d, 0x4a, 0x87, 0x3a, 0x13, 0x51, 0xe1, 0x6d, 0x9a,
	0xb3, 0x3a, 0xf4, 0xbf, 0xc3, 0x6d, 0x7b, 0x83, 0x8c, 0x39, 0x93, 0xa8, 0x21, 0x31, 0x4d, 0xe9,
	0xab, 0x81, 0x0e, 0x35, 0x75, 0x92, 0xd0, 0xc8, 0xab, 0x98, 0x53, 0x26, 0x26, 0x07, 0xb0, 0x25,
	0xf5, 0xb1, 0x13, 0x3a, 0xc5, 0xa1, 0x0a, 0xa7, 0xb1, 0xe7, 0x9a, 0xdd, 0x42, 0xd6, 0x7f, 0x08,
	0xcd, 0xa1, 0xe2, 0xf1, 0x42, 0xff, 0x82, 0xca, 0xb9, 0xa4, 0xf2, 0x7f, 0x3b, 0xda, 0xa4, 0xc6,
	0x58, 0x05, 0xbb, 0x50, 0x93, 0x2a, 0xe2, 0x89, 0x32, 0xb0, 0x56, 0x60, 0x57, 0x36, 0x8f, 0x42,
	0x18, 0x25, 0x69, 0x1e, 0x85, 0x20, 0x3e, 0xb4, 0x90, 0x45, 0x45, 0x25, 0xb9, 0x1c, 0x69, 0x43,
	0x9d, 0xca, 0xf7, 0x74, 0x32, 0xc1, 0xc8, 0xb8, 0xaf, 0x07, 0xd9, 0x5a, 0xef, 0xe1, 0x39, 0x55,
	0x47, 0x3c, 0x42, 0xaf, 0x6a, 0x6c, 0x67, 0x6b, 0xff, 0x8f, 0x03, 0xcd, 0x63, 0xc1, 0xc7, 0x28,
	0xe5, 0x5b, 0xf6, 0x93, 0x97, 0x54, 0x67, 0xb5, 0x12, 0x95, 0xb2, 0x4a, 0xdc, 0x48, 0x65, 0x17,
	0x9a, 0x71, 0x7a, 0xd9, 0xc7, 0x70, 0x8a, 0xb6, 0x4d, 0xcb, 0x29, 0xd2, 0x81, 0x06, 0x95, 0x41,
	0xc2, 0x18, 0x65, 0xa7, 0x46, 0x6c, 0x3d, 0xb8, 0x4c, 0xe4, 0x9c, 0xd4, 0x0a, 0x4e, 0x0e, 0x61,
	0xf7, 0x73, 0x82, 0x62, 0xfe, 0x89, 0xa1, 0x35, 0x74, 0x55, 0x53, 0x2e, 0x1c, 0xb8, 0xbf, 0x02,
	0xb7, 0xfd, 0x79, 0x06, 0x75, 0x2d, 0x49, 0xd7, 0xc3, 0x9c, 0x69, 0x0e, 0xf6, 0xfa, 0xf9, 0xb9,
	0xed, 0x2f, 0x95, 0x2c, 0xc8, 0xc0, 0x4b, 0x8d, 0xad, 0xac, 0x69, 0xac, 0xbb, 0xdc, 0x58, 0xff,
	0x01, 0x74, 0x8c, 0x06, 0x6b, 0xcf, 0x92, 0xe2, 0x42, 0xb8, 0xff, 0xd7, 0x81, 0xfd, 0x35, 0x00,
	0x2b, 0x75, 0x0c, 0x2d, 0x5b, 0xbd, 0x93, 0x70, 0x34, 0x41, 0xcf, 0xe9, 0xba, 0xbd, 0xe6, 0xe0,
	0x65, 0x51, 0xee, 0x95, 0x24, 0x0b, 0x33, 0x86, 0xe1, 0x35, 0x53, 0x62, 0x1e, 0xe4, 0x48, 0xdb,
	0xdf, 0xe0, 0xee, 0x0a, 0x44, 0x0f, 0xca, 0x19, 0xce, 0x6d, 0x4d, 0x75, 0x48, 0x9e, 0x40, 0x75,
	0x16, 0x4e, 0x12, 0x34, 0xe6, 0xaf, 0xa9, 0x59, 0x8a, 0x7c, 0x51, 0x79, 0xee, 0x0c, 0x2e, 0x5c,
	0xd8, 0xc9, 0xff, 0x15, 0x43, 0x14, 0x33, 0x3a, 0x46, 0xf2, 0x06, 0xaa, 0xe6, 0xe9, 0x92, 0x4e,
	0x91, 0x6a, 0xf9, 0xcf, 0x68, 0xef, 0xaf, 0xd9, 0x4d, 0xdd, 0xf9, 0x1b, 0xe4, 0x08, 0x36, 0xf5,
	0xfb, 0x23, 0x7b, 0xab, 0xc0, 0xec, 0xe5, 0xb6, 0x3b, 0xe5, 0x9b, 0x19, 0x49, 0x04, 0x77, 0x0a,
	0xf3, 0x42, 0x0e, 0x4a, 0xcb, 0xbc, 0x32, 0x7f, 0xed, 0x47, 0xd7, 0xe2, 0xb2, 0x5b, 0x66, 0xb0,
	0x53, 0xda, 0x2b, 0x72, 0x78, 0xc3, 0x96, 0xa6, 0x37, 0x3e, 0xfe, 0xaf, 0x01, 0xf0, 0x37, 0x5e,
	0x91, 0xaf, 0xdb, 0xf9, 0x1e, 0x1c, 0x8f, 0x46, 0x35, 0xf3, 0x65, 0x3f, 0xfd, 0x17, 0x00, 0x00,
	0xff, 0xff, 0x74, 0x1d, 0x95, 0xb2, 0xca, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LinuxJobWorkerServiceClient is the client API for LinuxJobWorkerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LinuxJobWorkerServiceClient interface {
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error)
	QueryOneProcess(ctx context.Context, in *QueryOneProcessRequest, opts ...grpc.CallOption) (*QueryOneProcessResponse, error)
	QueryRunningProcesses(ctx context.Context, in *QueryRunningProcessesRequest, opts ...grpc.CallOption) (*QueryRunningProcessesResponse, error)
}

type linuxJobWorkerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLinuxJobWorkerServiceClient(cc grpc.ClientConnInterface) LinuxJobWorkerServiceClient {
	return &linuxJobWorkerServiceClient{cc}
}

func (c *linuxJobWorkerServiceClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, "/linuxJobWorker.LinuxJobWorkerService/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linuxJobWorkerServiceClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := c.cc.Invoke(ctx, "/linuxJobWorker.LinuxJobWorkerService/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linuxJobWorkerServiceClient) QueryOneProcess(ctx context.Context, in *QueryOneProcessRequest, opts ...grpc.CallOption) (*QueryOneProcessResponse, error) {
	out := new(QueryOneProcessResponse)
	err := c.cc.Invoke(ctx, "/linuxJobWorker.LinuxJobWorkerService/QueryOneProcess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linuxJobWorkerServiceClient) QueryRunningProcesses(ctx context.Context, in *QueryRunningProcessesRequest, opts ...grpc.CallOption) (*QueryRunningProcessesResponse, error) {
	out := new(QueryRunningProcessesResponse)
	err := c.cc.Invoke(ctx, "/linuxJobWorker.LinuxJobWorkerService/QueryRunningProcesses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinuxJobWorkerServiceServer is the server API for LinuxJobWorkerService service.
type LinuxJobWorkerServiceServer interface {
	Start(context.Context, *StartRequest) (*StartResponse, error)
	Stop(context.Context, *StopRequest) (*StopResponse, error)
	QueryOneProcess(context.Context, *QueryOneProcessRequest) (*QueryOneProcessResponse, error)
	QueryRunningProcesses(context.Context, *QueryRunningProcessesRequest) (*QueryRunningProcessesResponse, error)
}

// UnimplementedLinuxJobWorkerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLinuxJobWorkerServiceServer struct {
}

func (*UnimplementedLinuxJobWorkerServiceServer) Start(ctx context.Context, req *StartRequest) (*StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (*UnimplementedLinuxJobWorkerServiceServer) Stop(ctx context.Context, req *StopRequest) (*StopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (*UnimplementedLinuxJobWorkerServiceServer) QueryOneProcess(ctx context.Context, req *QueryOneProcessRequest) (*QueryOneProcessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryOneProcess not implemented")
}
func (*UnimplementedLinuxJobWorkerServiceServer) QueryRunningProcesses(ctx context.Context, req *QueryRunningProcessesRequest) (*QueryRunningProcessesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryRunningProcesses not implemented")
}

func RegisterLinuxJobWorkerServiceServer(s *grpc.Server, srv LinuxJobWorkerServiceServer) {
	s.RegisterService(&_LinuxJobWorkerService_serviceDesc, srv)
}

func _LinuxJobWorkerService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinuxJobWorkerServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linuxJobWorker.LinuxJobWorkerService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinuxJobWorkerServiceServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinuxJobWorkerService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinuxJobWorkerServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linuxJobWorker.LinuxJobWorkerService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinuxJobWorkerServiceServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinuxJobWorkerService_QueryOneProcess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryOneProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinuxJobWorkerServiceServer).QueryOneProcess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linuxJobWorker.LinuxJobWorkerService/QueryOneProcess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinuxJobWorkerServiceServer).QueryOneProcess(ctx, req.(*QueryOneProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinuxJobWorkerService_QueryRunningProcesses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRunningProcessesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinuxJobWorkerServiceServer).QueryRunningProcesses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linuxJobWorker.LinuxJobWorkerService/QueryRunningProcesses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinuxJobWorkerServiceServer).QueryRunningProcesses(ctx, req.(*QueryRunningProcessesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LinuxJobWorkerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "linuxJobWorker.LinuxJobWorkerService",
	HandlerType: (*LinuxJobWorkerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _LinuxJobWorkerService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _LinuxJobWorkerService_Stop_Handler,
		},
		{
			MethodName: "QueryOneProcess",
			Handler:    _LinuxJobWorkerService_QueryOneProcess_Handler,
		},
		{
			MethodName: "QueryRunningProcesses",
			Handler:    _LinuxJobWorkerService_QueryRunningProcesses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "LinuxJobWorker.proto",
}
