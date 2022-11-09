// v1.0.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.10.0
// source: logs.proto

package model

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

type LogLevel int32

const (
	LogLevel_Verbose    LogLevel = 0
	LogLevel_Debug      LogLevel = 1
	LogLevel_Infomation LogLevel = 2
	LogLevel_Warning    LogLevel = 3
	LogLevel_Error      LogLevel = 4
	LogLevel_Fatal      LogLevel = 5
)

// Enum value maps for LogLevel.
var (
	LogLevel_name = map[int32]string{
		0: "Verbose",
		1: "Debug",
		2: "Infomation",
		3: "Warning",
		4: "Error",
		5: "Fatal",
	}
	LogLevel_value = map[string]int32{
		"Verbose":    0,
		"Debug":      1,
		"Infomation": 2,
		"Warning":    3,
		"Error":      4,
		"Fatal":      5,
	}
)

func (x LogLevel) Enum() *LogLevel {
	p := new(LogLevel)
	*p = x
	return p
}

func (x LogLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_logs_proto_enumTypes[0].Descriptor()
}

func (LogLevel) Type() protoreflect.EnumType {
	return &file_logs_proto_enumTypes[0]
}

func (x LogLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogLevel.Descriptor instead.
func (LogLevel) EnumDescriptor() ([]byte, []int) {
	return file_logs_proto_rawDescGZIP(), []int{0}
}

type LogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: db:"ID"
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty" db:"ID"`
	// @gotags: db:"TraceNo"
	TraceNo string `protobuf:"bytes,2,opt,name=TraceNo,proto3" json:"TraceNo,omitempty" db:"TraceNo"`
	// @gotags: db:"User"
	User string `protobuf:"bytes,3,opt,name=User,proto3" json:"User,omitempty" db:"User"`
	// @gotags: db:"Message"
	Message string `protobuf:"bytes,4,opt,name=Message,proto3" json:"Message,omitempty" db:"Message"`
	// @gotags: db:"Error"
	Error string `protobuf:"bytes,5,opt,name=Error,proto3" json:"Error,omitempty" db:"Error"`
	// @gotags: db:"StackTrace"
	StackTrace string `protobuf:"bytes,6,opt,name=StackTrace,proto3" json:"StackTrace,omitempty" db:"StackTrace"`
	// @gotags: db:"Payload"
	Payload string `protobuf:"bytes,7,opt,name=Payload,proto3" json:"Payload,omitempty" db:"Payload"`
	// @gotags: db:"Level"
	Level LogLevel `protobuf:"varint,8,opt,name=Level,proto3,enum=model.LogLevel" json:"Level,omitempty" db:"Level"`
	// @gotags: db:"CreatedOnUtc"
	CreatedOnUtc int64 `protobuf:"varint,9,opt,name=CreatedOnUtc,proto3" json:"CreatedOnUtc,omitempty" db:"CreatedOnUtc"`
}

func (x *LogEntry) Reset() {
	*x = LogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntry) ProtoMessage() {}

func (x *LogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_logs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntry.ProtoReflect.Descriptor instead.
func (*LogEntry) Descriptor() ([]byte, []int) {
	return file_logs_proto_rawDescGZIP(), []int{0}
}

func (x *LogEntry) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *LogEntry) GetTraceNo() string {
	if x != nil {
		return x.TraceNo
	}
	return ""
}

func (x *LogEntry) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *LogEntry) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LogEntry) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *LogEntry) GetStackTrace() string {
	if x != nil {
		return x.StackTrace
	}
	return ""
}

func (x *LogEntry) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *LogEntry) GetLevel() LogLevel {
	if x != nil {
		return x.Level
	}
	return LogLevel_Verbose
}

func (x *LogEntry) GetCreatedOnUtc() int64 {
	if x != nil {
		return x.CreatedOnUtc
	}
	return 0
}

type WriteLogCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientID string    `protobuf:"bytes,1,opt,name=ClientID,proto3" json:"ClientID,omitempty"`
	LogEntry *LogEntry `protobuf:"bytes,2,opt,name=LogEntry,proto3" json:"LogEntry,omitempty"`
}

func (x *WriteLogCommand) Reset() {
	*x = WriteLogCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteLogCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteLogCommand) ProtoMessage() {}

func (x *WriteLogCommand) ProtoReflect() protoreflect.Message {
	mi := &file_logs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteLogCommand.ProtoReflect.Descriptor instead.
func (*WriteLogCommand) Descriptor() ([]byte, []int) {
	return file_logs_proto_rawDescGZIP(), []int{1}
}

func (x *WriteLogCommand) GetClientID() string {
	if x != nil {
		return x.ClientID
	}
	return ""
}

func (x *WriteLogCommand) GetLogEntry() *LogEntry {
	if x != nil {
		return x.LogEntry
	}
	return nil
}

type LogEntryQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DBName    string `protobuf:"bytes,1,opt,name=DBName,proto3" json:"DBName,omitempty"`
	TableName string `protobuf:"bytes,2,opt,name=TableName,proto3" json:"TableName,omitempty"`
	ID        string `protobuf:"bytes,3,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *LogEntryQuery) Reset() {
	*x = LogEntryQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntryQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntryQuery) ProtoMessage() {}

func (x *LogEntryQuery) ProtoReflect() protoreflect.Message {
	mi := &file_logs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntryQuery.ProtoReflect.Descriptor instead.
func (*LogEntryQuery) Descriptor() ([]byte, []int) {
	return file_logs_proto_rawDescGZIP(), []int{2}
}

func (x *LogEntryQuery) GetDBName() string {
	if x != nil {
		return x.DBName
	}
	return ""
}

func (x *LogEntryQuery) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *LogEntryQuery) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type LogEntryResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message  string    `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	LogEntry *LogEntry `protobuf:"bytes,2,opt,name=LogEntry,proto3" json:"LogEntry,omitempty"`
}

func (x *LogEntryResult) Reset() {
	*x = LogEntryResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntryResult) ProtoMessage() {}

func (x *LogEntryResult) ProtoReflect() protoreflect.Message {
	mi := &file_logs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntryResult.ProtoReflect.Descriptor instead.
func (*LogEntryResult) Descriptor() ([]byte, []int) {
	return file_logs_proto_rawDescGZIP(), []int{3}
}

func (x *LogEntryResult) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LogEntryResult) GetLogEntry() *LogEntry {
	if x != nil {
		return x.LogEntry
	}
	return nil
}

type LogEntriesQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DBName     string   `protobuf:"bytes,1,opt,name=DBName,proto3" json:"DBName,omitempty"`
	TableName  string   `protobuf:"bytes,2,opt,name=TableName,proto3" json:"TableName,omitempty"`
	TraceNo    string   `protobuf:"bytes,3,opt,name=TraceNo,proto3" json:"TraceNo,omitempty"`
	User       string   `protobuf:"bytes,4,opt,name=User,proto3" json:"User,omitempty"`
	Message    string   `protobuf:"bytes,5,opt,name=Message,proto3" json:"Message,omitempty"`
	Error      string   `protobuf:"bytes,6,opt,name=Error,proto3" json:"Error,omitempty"`
	StackTrace string   `protobuf:"bytes,7,opt,name=StackTrace,proto3" json:"StackTrace,omitempty"`
	OrderDir   string   `protobuf:"bytes,8,opt,name=OrderDir,proto3" json:"OrderDir,omitempty"`
	StartTime  string   `protobuf:"bytes,9,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime    string   `protobuf:"bytes,10,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	Level      LogLevel `protobuf:"varint,11,opt,name=Level,proto3,enum=model.LogLevel" json:"Level,omitempty"`
	PageSize   int32    `protobuf:"varint,12,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	PageIndex  int32    `protobuf:"varint,13,opt,name=PageIndex,proto3" json:"PageIndex,omitempty"`
	OrderBy    int32    `protobuf:"varint,14,opt,name=OrderBy,proto3" json:"OrderBy,omitempty"`
}

func (x *LogEntriesQuery) Reset() {
	*x = LogEntriesQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntriesQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntriesQuery) ProtoMessage() {}

func (x *LogEntriesQuery) ProtoReflect() protoreflect.Message {
	mi := &file_logs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntriesQuery.ProtoReflect.Descriptor instead.
func (*LogEntriesQuery) Descriptor() ([]byte, []int) {
	return file_logs_proto_rawDescGZIP(), []int{4}
}

func (x *LogEntriesQuery) GetDBName() string {
	if x != nil {
		return x.DBName
	}
	return ""
}

func (x *LogEntriesQuery) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *LogEntriesQuery) GetTraceNo() string {
	if x != nil {
		return x.TraceNo
	}
	return ""
}

func (x *LogEntriesQuery) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *LogEntriesQuery) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LogEntriesQuery) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *LogEntriesQuery) GetStackTrace() string {
	if x != nil {
		return x.StackTrace
	}
	return ""
}

func (x *LogEntriesQuery) GetOrderDir() string {
	if x != nil {
		return x.OrderDir
	}
	return ""
}

func (x *LogEntriesQuery) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *LogEntriesQuery) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *LogEntriesQuery) GetLevel() LogLevel {
	if x != nil {
		return x.Level
	}
	return LogLevel_Verbose
}

func (x *LogEntriesQuery) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *LogEntriesQuery) GetPageIndex() int32 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *LogEntriesQuery) GetOrderBy() int32 {
	if x != nil {
		return x.OrderBy
	}
	return 0
}

type LogEntriesResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message    string      `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	TotalCount int64       `protobuf:"varint,2,opt,name=TotalCount,proto3" json:"TotalCount,omitempty"`
	LogEntries []*LogEntry `protobuf:"bytes,3,rep,name=LogEntries,proto3" json:"LogEntries,omitempty"`
}

func (x *LogEntriesResult) Reset() {
	*x = LogEntriesResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntriesResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntriesResult) ProtoMessage() {}

func (x *LogEntriesResult) ProtoReflect() protoreflect.Message {
	mi := &file_logs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntriesResult.ProtoReflect.Descriptor instead.
func (*LogEntriesResult) Descriptor() ([]byte, []int) {
	return file_logs_proto_rawDescGZIP(), []int{5}
}

func (x *LogEntriesResult) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LogEntriesResult) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *LogEntriesResult) GetLogEntries() []*LogEntry {
	if x != nil {
		return x.LogEntries
	}
	return nil
}

var File_logs_proto protoreflect.FileDescriptor

var file_logs_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x22, 0xfd, 0x01, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x18, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x63, 0x65, 0x4e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x54, 0x72, 0x61, 0x63, 0x65, 0x4e, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1e,
	0x0a, 0x0a, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x25, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x22, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x55, 0x74, 0x63, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e,
	0x55, 0x74, 0x63, 0x22, 0x5a, 0x0a, 0x0f, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x12, 0x2b, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x6f, 0x67,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22,
	0x55, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x44, 0x42, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x44, 0x42, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x57, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x6f, 0x67,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22,
	0x94, 0x03, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x42, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x44, 0x42, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x4e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x4e, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x63, 0x6b,
	0x54, 0x72, 0x61, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x44, 0x69, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x44, 0x69, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x22, 0x7d, 0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x4c, 0x6f, 0x67, 0x45, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x2a, 0x55, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x0b, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x65, 0x10, 0x00, 0x12, 0x09,
	0x0a, 0x05, 0x44, 0x65, 0x62, 0x75, 0x67, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x6e, 0x66,
	0x6f, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x61, 0x72,
	0x6e, 0x69, 0x6e, 0x67, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10,
	0x04, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x61, 0x74, 0x61, 0x6c, 0x10, 0x05, 0x32, 0xcf, 0x01, 0x0a,
	0x0f, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3e, 0x0a, 0x0d, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4c,
	0x6f, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x3a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x6f,
	0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x40, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x16, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x6f,
	0x67, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x1e,
	0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x75, 0x6b,
	0x69, 0x79, 0x61, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_logs_proto_rawDescOnce sync.Once
	file_logs_proto_rawDescData = file_logs_proto_rawDesc
)

func file_logs_proto_rawDescGZIP() []byte {
	file_logs_proto_rawDescOnce.Do(func() {
		file_logs_proto_rawDescData = protoimpl.X.CompressGZIP(file_logs_proto_rawDescData)
	})
	return file_logs_proto_rawDescData
}

var file_logs_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_logs_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_logs_proto_goTypes = []interface{}{
	(LogLevel)(0),            // 0: model.LogLevel
	(*LogEntry)(nil),         // 1: model.LogEntry
	(*WriteLogCommand)(nil),  // 2: model.WriteLogCommand
	(*LogEntryQuery)(nil),    // 3: model.LogEntryQuery
	(*LogEntryResult)(nil),   // 4: model.LogEntryResult
	(*LogEntriesQuery)(nil),  // 5: model.LogEntriesQuery
	(*LogEntriesResult)(nil), // 6: model.LogEntriesResult
}
var file_logs_proto_depIdxs = []int32{
	0, // 0: model.LogEntry.Level:type_name -> model.LogLevel
	1, // 1: model.WriteLogCommand.LogEntry:type_name -> model.LogEntry
	1, // 2: model.LogEntryResult.LogEntry:type_name -> model.LogEntry
	0, // 3: model.LogEntriesQuery.Level:type_name -> model.LogLevel
	1, // 4: model.LogEntriesResult.LogEntries:type_name -> model.LogEntry
	2, // 5: model.LogEntryService.WriteLogEntry:input_type -> model.WriteLogCommand
	3, // 6: model.LogEntryService.GetLogEntry:input_type -> model.LogEntryQuery
	5, // 7: model.LogEntryService.GetLogEntries:input_type -> model.LogEntriesQuery
	4, // 8: model.LogEntryService.WriteLogEntry:output_type -> model.LogEntryResult
	4, // 9: model.LogEntryService.GetLogEntry:output_type -> model.LogEntryResult
	6, // 10: model.LogEntryService.GetLogEntries:output_type -> model.LogEntriesResult
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_logs_proto_init() }
func file_logs_proto_init() {
	if File_logs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_logs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEntry); i {
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
		file_logs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteLogCommand); i {
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
		file_logs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEntryQuery); i {
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
		file_logs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEntryResult); i {
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
		file_logs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEntriesQuery); i {
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
		file_logs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEntriesResult); i {
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
			RawDescriptor: file_logs_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_logs_proto_goTypes,
		DependencyIndexes: file_logs_proto_depIdxs,
		EnumInfos:         file_logs_proto_enumTypes,
		MessageInfos:      file_logs_proto_msgTypes,
	}.Build()
	File_logs_proto = out.File
	file_logs_proto_rawDesc = nil
	file_logs_proto_goTypes = nil
	file_logs_proto_depIdxs = nil
}
