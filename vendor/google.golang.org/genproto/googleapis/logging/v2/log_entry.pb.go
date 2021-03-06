// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/logging/v2/log_entry.proto

/*
Package logging is a generated protocol buffer package.

It is generated from these files:
	google/logging/v2/log_entry.proto
	google/logging/v2/logging.proto
	google/logging/v2/logging_config.proto
	google/logging/v2/logging_metrics.proto

It has these top-level messages:
	LogEntry
	LogEntryOperation
	LogEntrySourceLocation
	DeleteLogRequest
	WriteLogEntriesRequest
	WriteLogEntriesResponse
	WriteLogEntriesPartialErrors
	ListLogEntriesRequest
	ListLogEntriesResponse
	ListMonitoredResourceDescriptorsRequest
	ListMonitoredResourceDescriptorsResponse
	ListLogsRequest
	ListLogsResponse
	LogSink
	ListSinksRequest
	ListSinksResponse
	GetSinkRequest
	CreateSinkRequest
	UpdateSinkRequest
	DeleteSinkRequest
	LogExclusion
	ListExclusionsRequest
	ListExclusionsResponse
	GetExclusionRequest
	CreateExclusionRequest
	UpdateExclusionRequest
	DeleteExclusionRequest
	LogMetric
	ListLogMetricsRequest
	ListLogMetricsResponse
	GetLogMetricRequest
	CreateLogMetricRequest
	UpdateLogMetricRequest
	DeleteLogMetricRequest
*/
package logging

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_api3 "google.golang.org/genproto/googleapis/api/monitoredres"
import google_logging_type "google.golang.org/genproto/googleapis/logging/type"
import google_logging_type1 "google.golang.org/genproto/googleapis/logging/type"
import google_protobuf2 "github.com/golang/protobuf/ptypes/any"
import google_protobuf3 "github.com/golang/protobuf/ptypes/struct"
import google_protobuf4 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// An individual entry in a log.
type LogEntry struct {
	// Required. The resource name of the log to which this log entry belongs:
	//
	//     "projects/[PROJECT_ID]/logs/[LOG_ID]"
	//     "organizations/[ORGANIZATION_ID]/logs/[LOG_ID]"
	//     "billingAccounts/[BILLING_ACCOUNT_ID]/logs/[LOG_ID]"
	//     "folders/[FOLDER_ID]/logs/[LOG_ID]"
	//
	//  A project number may optionally be used in place of PROJECT_ID. The
	//  project number is translated to its corresponding PROJECT_ID internally
	//  and the `log_name` field will contain PROJECT_ID in queries and exports.
	//
	// `[LOG_ID]` must be URL-encoded within `log_name`. Example:
	// `"organizations/1234567890/logs/cloudresourcemanager.googleapis.com%2Factivity"`.
	// `[LOG_ID]` must be less than 512 characters long and can only include the
	// following characters: upper and lower case alphanumeric characters,
	// forward-slash, underscore, hyphen, and period.
	//
	// For backward compatibility, if `log_name` begins with a forward-slash, such
	// as `/projects/...`, then the log entry is ingested as usual but the
	// forward-slash is removed. Listing the log entry will not show the leading
	// slash and filtering for a log name with a leading slash will never return
	// any results.
	LogName string `protobuf:"bytes,12,opt,name=log_name,json=logName" json:"log_name,omitempty"`
	// Required. The monitored resource associated with this log entry.
	// Example: a log entry that reports a database error would be
	// associated with the monitored resource designating the particular
	// database that reported the error.
	Resource *google_api3.MonitoredResource `protobuf:"bytes,8,opt,name=resource" json:"resource,omitempty"`
	// Optional. The log entry payload, which can be one of multiple types.
	//
	// Types that are valid to be assigned to Payload:
	//	*LogEntry_ProtoPayload
	//	*LogEntry_TextPayload
	//	*LogEntry_JsonPayload
	Payload isLogEntry_Payload `protobuf_oneof:"payload"`
	// Optional. The time the event described by the log entry occurred.
	// This time is used to compute the log entry's age and to enforce
	// the logs retention period. If this field is omitted in a new log
	// entry, then Stackdriver Logging assigns it the current time.
	//
	// Incoming log entries should have timestamps that are no more than
	// the [logs retention period](/logging/quota-policy) in the past,
	// and no more than 24 hours in the future.
	// See the `entries.write` API method for more information.
	Timestamp *google_protobuf4.Timestamp `protobuf:"bytes,9,opt,name=timestamp" json:"timestamp,omitempty"`
	// Output only. The time the log entry was received by Stackdriver Logging.
	ReceiveTimestamp *google_protobuf4.Timestamp `protobuf:"bytes,24,opt,name=receive_timestamp,json=receiveTimestamp" json:"receive_timestamp,omitempty"`
	// Optional. The severity of the log entry. The default value is
	// `LogSeverity.DEFAULT`.
	Severity google_logging_type1.LogSeverity `protobuf:"varint,10,opt,name=severity,enum=google.logging.type.LogSeverity" json:"severity,omitempty"`
	// Optional. A unique identifier for the log entry. If you provide a value,
	// then Stackdriver Logging considers other log entries in the same project,
	// with the same `timestamp`, and with the same `insert_id` to be duplicates
	// which can be removed.  If omitted in new log entries, then Stackdriver
	// Logging assigns its own unique identifier. The `insert_id` is also used
	// to order log entries that have the same `timestamp` value.
	InsertId string `protobuf:"bytes,4,opt,name=insert_id,json=insertId" json:"insert_id,omitempty"`
	// Optional. Information about the HTTP request associated with this
	// log entry, if applicable.
	HttpRequest *google_logging_type.HttpRequest `protobuf:"bytes,7,opt,name=http_request,json=httpRequest" json:"http_request,omitempty"`
	// Optional. A set of user-defined (key, value) data that provides additional
	// information about the log entry.
	Labels map[string]string `protobuf:"bytes,11,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Optional. Information about an operation associated with the log entry, if
	// applicable.
	Operation *LogEntryOperation `protobuf:"bytes,15,opt,name=operation" json:"operation,omitempty"`
	// Optional. Resource name of the trace associated with the log entry, if any.
	// If it contains a relative resource name, the name is assumed to be relative
	// to `//tracing.googleapis.com`. Example:
	// `projects/my-projectid/traces/06796866738c859f2f19b7cfb3214824`
	Trace string `protobuf:"bytes,22,opt,name=trace" json:"trace,omitempty"`
	// Optional. Id of the span within the trace associated with the log entry.
	// e.g. "0000000000000042"
	// For Stackdriver trace spans, this is the same format that the Stackdriver
	// trace API uses.
	// The ID is a 16-character hexadecimal encoding of an 8-byte array.
	SpanId string `protobuf:"bytes,27,opt,name=span_id,json=spanId" json:"span_id,omitempty"`
	// Optional. Source code location information associated with the log entry,
	// if any.
	SourceLocation *LogEntrySourceLocation `protobuf:"bytes,23,opt,name=source_location,json=sourceLocation" json:"source_location,omitempty"`
}

func (m *LogEntry) Reset()                    { *m = LogEntry{} }
func (m *LogEntry) String() string            { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()               {}
func (*LogEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isLogEntry_Payload interface {
	isLogEntry_Payload()
}

type LogEntry_ProtoPayload struct {
	ProtoPayload *google_protobuf2.Any `protobuf:"bytes,2,opt,name=proto_payload,json=protoPayload,oneof"`
}
type LogEntry_TextPayload struct {
	TextPayload string `protobuf:"bytes,3,opt,name=text_payload,json=textPayload,oneof"`
}
type LogEntry_JsonPayload struct {
	JsonPayload *google_protobuf3.Struct `protobuf:"bytes,6,opt,name=json_payload,json=jsonPayload,oneof"`
}

func (*LogEntry_ProtoPayload) isLogEntry_Payload() {}
func (*LogEntry_TextPayload) isLogEntry_Payload()  {}
func (*LogEntry_JsonPayload) isLogEntry_Payload()  {}

func (m *LogEntry) GetPayload() isLogEntry_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *LogEntry) GetLogName() string {
	if m != nil {
		return m.LogName
	}
	return ""
}

func (m *LogEntry) GetResource() *google_api3.MonitoredResource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *LogEntry) GetProtoPayload() *google_protobuf2.Any {
	if x, ok := m.GetPayload().(*LogEntry_ProtoPayload); ok {
		return x.ProtoPayload
	}
	return nil
}

func (m *LogEntry) GetTextPayload() string {
	if x, ok := m.GetPayload().(*LogEntry_TextPayload); ok {
		return x.TextPayload
	}
	return ""
}

func (m *LogEntry) GetJsonPayload() *google_protobuf3.Struct {
	if x, ok := m.GetPayload().(*LogEntry_JsonPayload); ok {
		return x.JsonPayload
	}
	return nil
}

func (m *LogEntry) GetTimestamp() *google_protobuf4.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *LogEntry) GetReceiveTimestamp() *google_protobuf4.Timestamp {
	if m != nil {
		return m.ReceiveTimestamp
	}
	return nil
}

func (m *LogEntry) GetSeverity() google_logging_type1.LogSeverity {
	if m != nil {
		return m.Severity
	}
	return google_logging_type1.LogSeverity_DEFAULT
}

func (m *LogEntry) GetInsertId() string {
	if m != nil {
		return m.InsertId
	}
	return ""
}

func (m *LogEntry) GetHttpRequest() *google_logging_type.HttpRequest {
	if m != nil {
		return m.HttpRequest
	}
	return nil
}

func (m *LogEntry) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *LogEntry) GetOperation() *LogEntryOperation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *LogEntry) GetTrace() string {
	if m != nil {
		return m.Trace
	}
	return ""
}

func (m *LogEntry) GetSpanId() string {
	if m != nil {
		return m.SpanId
	}
	return ""
}

func (m *LogEntry) GetSourceLocation() *LogEntrySourceLocation {
	if m != nil {
		return m.SourceLocation
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LogEntry) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LogEntry_OneofMarshaler, _LogEntry_OneofUnmarshaler, _LogEntry_OneofSizer, []interface{}{
		(*LogEntry_ProtoPayload)(nil),
		(*LogEntry_TextPayload)(nil),
		(*LogEntry_JsonPayload)(nil),
	}
}

func _LogEntry_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LogEntry)
	// payload
	switch x := m.Payload.(type) {
	case *LogEntry_ProtoPayload:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ProtoPayload); err != nil {
			return err
		}
	case *LogEntry_TextPayload:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.TextPayload)
	case *LogEntry_JsonPayload:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.JsonPayload); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("LogEntry.Payload has unexpected type %T", x)
	}
	return nil
}

func _LogEntry_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LogEntry)
	switch tag {
	case 2: // payload.proto_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf2.Any)
		err := b.DecodeMessage(msg)
		m.Payload = &LogEntry_ProtoPayload{msg}
		return true, err
	case 3: // payload.text_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Payload = &LogEntry_TextPayload{x}
		return true, err
	case 6: // payload.json_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf3.Struct)
		err := b.DecodeMessage(msg)
		m.Payload = &LogEntry_JsonPayload{msg}
		return true, err
	default:
		return false, nil
	}
}

func _LogEntry_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LogEntry)
	// payload
	switch x := m.Payload.(type) {
	case *LogEntry_ProtoPayload:
		s := proto.Size(x.ProtoPayload)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LogEntry_TextPayload:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.TextPayload)))
		n += len(x.TextPayload)
	case *LogEntry_JsonPayload:
		s := proto.Size(x.JsonPayload)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Additional information about a potentially long-running operation with which
// a log entry is associated.
type LogEntryOperation struct {
	// Optional. An arbitrary operation identifier. Log entries with the
	// same identifier are assumed to be part of the same operation.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Optional. An arbitrary producer identifier. The combination of
	// `id` and `producer` must be globally unique.  Examples for `producer`:
	// `"MyDivision.MyBigCompany.com"`, `"github.com/MyProject/MyApplication"`.
	Producer string `protobuf:"bytes,2,opt,name=producer" json:"producer,omitempty"`
	// Optional. Set this to True if this is the first log entry in the operation.
	First bool `protobuf:"varint,3,opt,name=first" json:"first,omitempty"`
	// Optional. Set this to True if this is the last log entry in the operation.
	Last bool `protobuf:"varint,4,opt,name=last" json:"last,omitempty"`
}

func (m *LogEntryOperation) Reset()                    { *m = LogEntryOperation{} }
func (m *LogEntryOperation) String() string            { return proto.CompactTextString(m) }
func (*LogEntryOperation) ProtoMessage()               {}
func (*LogEntryOperation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LogEntryOperation) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LogEntryOperation) GetProducer() string {
	if m != nil {
		return m.Producer
	}
	return ""
}

func (m *LogEntryOperation) GetFirst() bool {
	if m != nil {
		return m.First
	}
	return false
}

func (m *LogEntryOperation) GetLast() bool {
	if m != nil {
		return m.Last
	}
	return false
}

// Additional information about the source code location that produced the log
// entry.
type LogEntrySourceLocation struct {
	// Optional. Source file name. Depending on the runtime environment, this
	// might be a simple name or a fully-qualified name.
	File string `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	// Optional. Line within the source file. 1-based; 0 indicates no line number
	// available.
	Line int64 `protobuf:"varint,2,opt,name=line" json:"line,omitempty"`
	// Optional. Human-readable name of the function or method being invoked, with
	// optional context such as the class or package name. This information may be
	// used in contexts such as the logs viewer, where a file and line number are
	// less meaningful. The format can vary by language. For example:
	// `qual.if.ied.Class.method` (Java), `dir/package.func` (Go), `function`
	// (Python).
	Function string `protobuf:"bytes,3,opt,name=function" json:"function,omitempty"`
}

func (m *LogEntrySourceLocation) Reset()                    { *m = LogEntrySourceLocation{} }
func (m *LogEntrySourceLocation) String() string            { return proto.CompactTextString(m) }
func (*LogEntrySourceLocation) ProtoMessage()               {}
func (*LogEntrySourceLocation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LogEntrySourceLocation) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *LogEntrySourceLocation) GetLine() int64 {
	if m != nil {
		return m.Line
	}
	return 0
}

func (m *LogEntrySourceLocation) GetFunction() string {
	if m != nil {
		return m.Function
	}
	return ""
}

func init() {
	proto.RegisterType((*LogEntry)(nil), "google.logging.v2.LogEntry")
	proto.RegisterType((*LogEntryOperation)(nil), "google.logging.v2.LogEntryOperation")
	proto.RegisterType((*LogEntrySourceLocation)(nil), "google.logging.v2.LogEntrySourceLocation")
}

func init() { proto.RegisterFile("google/logging/v2/log_entry.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 729 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x36, 0x25, 0x57, 0xa2, 0x56, 0xf2, 0xdf, 0xc2, 0xb5, 0x68, 0xd9, 0x45, 0x55, 0xbb, 0x68,
	0xd5, 0x0b, 0x05, 0xa8, 0x17, 0xbb, 0x36, 0x50, 0x54, 0x46, 0x61, 0x1b, 0x50, 0x5b, 0x63, 0x5d,
	0xf8, 0x10, 0x08, 0x10, 0xd6, 0xe4, 0x8a, 0xde, 0x84, 0xda, 0x65, 0x96, 0x4b, 0x21, 0x7a, 0x94,
	0xbc, 0x42, 0x1e, 0x25, 0x6f, 0x93, 0x5b, 0x8e, 0xc1, 0x0e, 0x97, 0x92, 0x22, 0x19, 0xce, 0x6d,
	0x66, 0xe7, 0xfb, 0xe6, 0x9b, 0x19, 0xce, 0x10, 0xfd, 0x14, 0x49, 0x19, 0xc5, 0xac, 0x1b, 0xcb,
	0x28, 0xe2, 0x22, 0xea, 0x4e, 0x7b, 0xc6, 0x1c, 0x31, 0xa1, 0xd5, 0xcc, 0x4f, 0x94, 0xd4, 0x12,
	0xef, 0xe5, 0x10, 0xdf, 0x42, 0xfc, 0x69, 0xaf, 0x75, 0x6c, 0x59, 0x34, 0xe1, 0x5d, 0x2a, 0x84,
	0xd4, 0x54, 0x73, 0x29, 0xd2, 0x9c, 0xd0, 0x3a, 0x5d, 0x8a, 0x4e, 0xa4, 0xe0, 0x5a, 0x2a, 0x16,
	0x8e, 0x14, 0x4b, 0x65, 0xa6, 0x02, 0x66, 0x41, 0xbf, 0xac, 0x08, 0xeb, 0x59, 0xc2, 0xba, 0x4f,
	0x5a, 0x27, 0x23, 0xc5, 0xde, 0x66, 0x2c, 0xd5, 0x2f, 0xe1, 0x4c, 0x89, 0x29, 0x9b, 0x32, 0xc5,
	0xb5, 0xad, 0xb2, 0x75, 0x68, 0x71, 0xe0, 0x3d, 0x66, 0xe3, 0x2e, 0x15, 0x45, 0xe8, 0x78, 0x35,
	0x94, 0x6a, 0x95, 0x05, 0x85, 0xc0, 0x8f, 0xab, 0x51, 0xcd, 0x27, 0x2c, 0xd5, 0x74, 0x92, 0xe4,
	0x80, 0x93, 0x4f, 0x15, 0xe4, 0x0e, 0x64, 0xf4, 0xb7, 0x19, 0x09, 0x3e, 0x44, 0xae, 0x11, 0x17,
	0x74, 0xc2, 0xbc, 0x46, 0xdb, 0xe9, 0xd4, 0x48, 0x35, 0x96, 0xd1, 0xbf, 0x74, 0xc2, 0xf0, 0x39,
	0x72, 0x8b, 0x1e, 0x3d, 0xb7, 0xed, 0x74, 0xea, 0xbd, 0x1f, 0x7c, 0x3b, 0x3a, 0x9a, 0x70, 0xff,
	0x9f, 0x62, 0x12, 0xc4, 0x82, 0xc8, 0x1c, 0x8e, 0x2f, 0xd0, 0x16, 0x68, 0x8d, 0x12, 0x3a, 0x8b,
	0x25, 0x0d, 0xbd, 0x12, 0xf0, 0xf7, 0x0b, 0x7e, 0x51, 0x9b, 0xff, 0x97, 0x98, 0xdd, 0x6c, 0x90,
	0x06, 0xf8, 0x77, 0x39, 0x16, 0x9f, 0xa2, 0x86, 0x66, 0xef, 0xf4, 0x9c, 0x5b, 0x36, 0x65, 0xdd,
	0x6c, 0x90, 0xba, 0x79, 0x2d, 0x40, 0x97, 0xa8, 0xf1, 0x3a, 0x95, 0x62, 0x0e, 0xaa, 0x80, 0x40,
	0x73, 0x4d, 0xe0, 0x1e, 0x46, 0x63, 0xd8, 0x06, 0x5e, 0xb0, 0xcf, 0x50, 0x6d, 0x3e, 0x15, 0xaf,
	0x06, 0xd4, 0xd6, 0x1a, 0xf5, 0xff, 0x02, 0x41, 0x16, 0x60, 0x7c, 0x8d, 0xf6, 0x14, 0x0b, 0x18,
	0x9f, 0xb2, 0xd1, 0x22, 0x83, 0xf7, 0xcd, 0x0c, 0xbb, 0x96, 0x34, 0x7f, 0xc1, 0x97, 0xc8, 0x2d,
	0xbe, 0xb8, 0x87, 0xda, 0x4e, 0x67, 0xbb, 0xd7, 0xf6, 0x57, 0x16, 0xd3, 0xac, 0x86, 0x3f, 0x90,
	0xd1, 0xbd, 0xc5, 0x91, 0x39, 0x03, 0x1f, 0xa1, 0x1a, 0x17, 0x29, 0x53, 0x7a, 0xc4, 0x43, 0x6f,
	0x13, 0xbe, 0x9b, 0x9b, 0x3f, 0xdc, 0x86, 0xf8, 0x0a, 0x35, 0x96, 0x17, 0xcf, 0xab, 0x42, 0x79,
	0xcf, 0xa7, 0xbf, 0xd1, 0x3a, 0x21, 0x39, 0x8e, 0xd4, 0x9f, 0x16, 0x0e, 0xfe, 0x13, 0x55, 0x62,
	0xfa, 0xc8, 0xe2, 0xd4, 0xab, 0xb7, 0xcb, 0x9d, 0x7a, 0xef, 0x57, 0x7f, 0xed, 0x6c, 0xfc, 0x62,
	0x8b, 0xfc, 0x01, 0x20, 0xc1, 0x26, 0x96, 0x86, 0xfb, 0xa8, 0x26, 0x13, 0xa6, 0xe0, 0x92, 0xbc,
	0x1d, 0x28, 0xe1, 0xe7, 0x17, 0x72, 0xfc, 0x57, 0x60, 0xc9, 0x82, 0x86, 0xf7, 0xd1, 0x77, 0x5a,
	0xd1, 0x80, 0x79, 0x07, 0xd0, 0x62, 0xee, 0xe0, 0x26, 0xaa, 0xa6, 0x09, 0x15, 0xa6, 0xf5, 0x23,
	0x78, 0xaf, 0x18, 0xf7, 0x36, 0xc4, 0x04, 0xed, 0xe4, 0x0b, 0x38, 0x8a, 0x65, 0x90, 0x0b, 0x37,
	0x41, 0xf8, 0xb7, 0x17, 0x84, 0xef, 0x81, 0x31, 0xb0, 0x04, 0xb2, 0x9d, 0x7e, 0xe5, 0xb7, 0xce,
	0x51, 0x7d, 0xa9, 0x3b, 0xbc, 0x8b, 0xca, 0x6f, 0xd8, 0xcc, 0x73, 0x40, 0xd7, 0x98, 0xa6, 0xc6,
	0x29, 0x8d, 0x33, 0x06, 0x3b, 0x5e, 0x23, 0xb9, 0xf3, 0x47, 0xe9, 0xcc, 0xe9, 0xd7, 0x50, 0xd5,
	0xae, 0xe7, 0x09, 0x47, 0x7b, 0x6b, 0x8d, 0xe2, 0x6d, 0x54, 0xe2, 0xa1, 0x4d, 0x55, 0xe2, 0x21,
	0x6e, 0x21, 0x37, 0x51, 0x32, 0xcc, 0x02, 0xa6, 0x6c, 0xb2, 0xb9, 0x6f, 0x54, 0xc6, 0x5c, 0xa5,
	0x1a, 0xae, 0xc1, 0x25, 0xb9, 0x83, 0x31, 0xda, 0x8c, 0x69, 0xaa, 0x61, 0x03, 0x5c, 0x02, 0xf6,
	0xc9, 0x10, 0x1d, 0x3c, 0xdf, 0x9a, 0x41, 0x8f, 0x79, 0xcc, 0xac, 0x22, 0xd8, 0x90, 0x81, 0x8b,
	0xbc, 0xf8, 0x32, 0x01, 0xdb, 0xd4, 0x31, 0xce, 0x44, 0x00, 0xf3, 0x2b, 0xe7, 0x75, 0x14, 0x7e,
	0xff, 0xbd, 0x83, 0xbe, 0x0f, 0xe4, 0x64, 0x7d, 0x9e, 0xfd, 0xad, 0x42, 0xf5, 0x0e, 0x8e, 0xd9,
	0x79, 0x75, 0x66, 0x31, 0x91, 0x8c, 0xa9, 0x88, 0x7c, 0xa9, 0xa2, 0x6e, 0xc4, 0x04, 0x1c, 0x47,
	0x37, 0x0f, 0xd1, 0x84, 0xa7, 0x4b, 0x7f, 0xea, 0x0b, 0x6b, 0x7e, 0x76, 0x9c, 0x0f, 0xa5, 0xe6,
	0x75, 0xce, 0xbe, 0x8a, 0x65, 0x16, 0x9a, 0x8f, 0x05, 0x3a, 0x0f, 0xbd, 0x8f, 0x45, 0x64, 0x08,
	0x91, 0xa1, 0x8d, 0x0c, 0x1f, 0x7a, 0x8f, 0x15, 0xc8, 0xfd, 0xfb, 0x97, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x80, 0x53, 0xd3, 0xff, 0x04, 0x06, 0x00, 0x00,
}
