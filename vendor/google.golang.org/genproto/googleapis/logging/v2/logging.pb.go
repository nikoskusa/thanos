// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/logging/v2/logging.proto

package logging

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_api3 "google.golang.org/genproto/googleapis/api/monitoredres"
import _ "github.com/golang/protobuf/ptypes/duration"
import google_protobuf5 "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import google_rpc "google.golang.org/genproto/googleapis/rpc/status"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The parameters to DeleteLog.
type DeleteLogRequest struct {
	// Required. The resource name of the log to delete:
	//
	//     "projects/[PROJECT_ID]/logs/[LOG_ID]"
	//     "organizations/[ORGANIZATION_ID]/logs/[LOG_ID]"
	//     "billingAccounts/[BILLING_ACCOUNT_ID]/logs/[LOG_ID]"
	//     "folders/[FOLDER_ID]/logs/[LOG_ID]"
	//
	// `[LOG_ID]` must be URL-encoded. For example,
	// `"projects/my-project-id/logs/syslog"`,
	// `"organizations/1234567890/logs/cloudresourcemanager.googleapis.com%2Factivity"`.
	// For more information about log names, see
	// [LogEntry][google.logging.v2.LogEntry].
	LogName string `protobuf:"bytes,1,opt,name=log_name,json=logName" json:"log_name,omitempty"`
}

func (m *DeleteLogRequest) Reset()                    { *m = DeleteLogRequest{} }
func (m *DeleteLogRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteLogRequest) ProtoMessage()               {}
func (*DeleteLogRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *DeleteLogRequest) GetLogName() string {
	if m != nil {
		return m.LogName
	}
	return ""
}

// The parameters to WriteLogEntries.
type WriteLogEntriesRequest struct {
	// Optional. A default log resource name that is assigned to all log entries
	// in `entries` that do not specify a value for `log_name`:
	//
	//     "projects/[PROJECT_ID]/logs/[LOG_ID]"
	//     "organizations/[ORGANIZATION_ID]/logs/[LOG_ID]"
	//     "billingAccounts/[BILLING_ACCOUNT_ID]/logs/[LOG_ID]"
	//     "folders/[FOLDER_ID]/logs/[LOG_ID]"
	//
	// `[LOG_ID]` must be URL-encoded. For example,
	// `"projects/my-project-id/logs/syslog"` or
	// `"organizations/1234567890/logs/cloudresourcemanager.googleapis.com%2Factivity"`.
	// For more information about log names, see
	// [LogEntry][google.logging.v2.LogEntry].
	LogName string `protobuf:"bytes,1,opt,name=log_name,json=logName" json:"log_name,omitempty"`
	// Optional. A default monitored resource object that is assigned to all log
	// entries in `entries` that do not specify a value for `resource`. Example:
	//
	//     { "type": "gce_instance",
	//       "labels": {
	//         "zone": "us-central1-a", "instance_id": "00000000000000000000" }}
	//
	// See [LogEntry][google.logging.v2.LogEntry].
	Resource *google_api3.MonitoredResource `protobuf:"bytes,2,opt,name=resource" json:"resource,omitempty"`
	// Optional. Default labels that are added to the `labels` field of all log
	// entries in `entries`. If a log entry already has a label with the same key
	// as a label in this parameter, then the log entry's label is not changed.
	// See [LogEntry][google.logging.v2.LogEntry].
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Required. The log entries to send to Stackdriver Logging. The order of log
	// entries in this list does not matter. Values supplied in this method's
	// `log_name`, `resource`, and `labels` fields are copied into those log
	// entries in this list that do not include values for their corresponding
	// fields. For more information, see the [LogEntry][google.logging.v2.LogEntry] type.
	//
	// If the `timestamp` or `insert_id` fields are missing in log entries, then
	// this method supplies the current time or a unique identifier, respectively.
	// The supplied values are chosen so that, among the log entries that did not
	// supply their own values, the entries earlier in the list will sort before
	// the entries later in the list. See the `entries.list` method.
	//
	// Log entries with timestamps that are more than the
	// [logs retention period](/logging/quota-policy) in the past or more than
	// 24 hours in the future might be discarded. Discarding does not return
	// an error.
	//
	// To improve throughput and to avoid exceeding the
	// [quota limit](/logging/quota-policy) for calls to `entries.write`,
	// you should try to include several log entries in this list,
	// rather than calling this method for each individual log entry.
	Entries []*LogEntry `protobuf:"bytes,4,rep,name=entries" json:"entries,omitempty"`
	// Optional. Whether valid entries should be written even if some other
	// entries fail due to INVALID_ARGUMENT or PERMISSION_DENIED errors. If any
	// entry is not written, then the response status is the error associated
	// with one of the failed entries and the response includes error details
	// keyed by the entries' zero-based index in the `entries.write` method.
	PartialSuccess bool `protobuf:"varint,5,opt,name=partial_success,json=partialSuccess" json:"partial_success,omitempty"`
}

func (m *WriteLogEntriesRequest) Reset()                    { *m = WriteLogEntriesRequest{} }
func (m *WriteLogEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*WriteLogEntriesRequest) ProtoMessage()               {}
func (*WriteLogEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *WriteLogEntriesRequest) GetLogName() string {
	if m != nil {
		return m.LogName
	}
	return ""
}

func (m *WriteLogEntriesRequest) GetResource() *google_api3.MonitoredResource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *WriteLogEntriesRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *WriteLogEntriesRequest) GetEntries() []*LogEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *WriteLogEntriesRequest) GetPartialSuccess() bool {
	if m != nil {
		return m.PartialSuccess
	}
	return false
}

// Result returned from WriteLogEntries.
// empty
type WriteLogEntriesResponse struct {
}

func (m *WriteLogEntriesResponse) Reset()                    { *m = WriteLogEntriesResponse{} }
func (m *WriteLogEntriesResponse) String() string            { return proto.CompactTextString(m) }
func (*WriteLogEntriesResponse) ProtoMessage()               {}
func (*WriteLogEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

// Error details for WriteLogEntries with partial success.
type WriteLogEntriesPartialErrors struct {
	// When `WriteLogEntriesRequest.partial_success` is true, records the error
	// status for entries that were not written due to a permanent error, keyed
	// by the entry's zero-based index in `WriteLogEntriesRequest.entries`.
	//
	// Failed requests for which no entries are written will not include
	// per-entry errors.
	LogEntryErrors map[int32]*google_rpc.Status `protobuf:"bytes,1,rep,name=log_entry_errors,json=logEntryErrors" json:"log_entry_errors,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *WriteLogEntriesPartialErrors) Reset()                    { *m = WriteLogEntriesPartialErrors{} }
func (m *WriteLogEntriesPartialErrors) String() string            { return proto.CompactTextString(m) }
func (*WriteLogEntriesPartialErrors) ProtoMessage()               {}
func (*WriteLogEntriesPartialErrors) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *WriteLogEntriesPartialErrors) GetLogEntryErrors() map[int32]*google_rpc.Status {
	if m != nil {
		return m.LogEntryErrors
	}
	return nil
}

// The parameters to `ListLogEntries`.
type ListLogEntriesRequest struct {
	// Deprecated. Use `resource_names` instead.  One or more project identifiers
	// or project numbers from which to retrieve log entries.  Example:
	// `"my-project-1A"`. If present, these project identifiers are converted to
	// resource name format and added to the list of resources in
	// `resource_names`.
	ProjectIds []string `protobuf:"bytes,1,rep,name=project_ids,json=projectIds" json:"project_ids,omitempty"`
	// Required. Names of one or more parent resources from which to
	// retrieve log entries:
	//
	//     "projects/[PROJECT_ID]"
	//     "organizations/[ORGANIZATION_ID]"
	//     "billingAccounts/[BILLING_ACCOUNT_ID]"
	//     "folders/[FOLDER_ID]"
	//
	// Projects listed in the `project_ids` field are added to this list.
	ResourceNames []string `protobuf:"bytes,8,rep,name=resource_names,json=resourceNames" json:"resource_names,omitempty"`
	// Optional. A filter that chooses which log entries to return.  See [Advanced
	// Logs Filters](/logging/docs/view/advanced_filters).  Only log entries that
	// match the filter are returned.  An empty filter matches all log entries in
	// the resources listed in `resource_names`. Referencing a parent resource
	// that is not listed in `resource_names` will cause the filter to return no
	// results.
	// The maximum length of the filter is 20000 characters.
	Filter string `protobuf:"bytes,2,opt,name=filter" json:"filter,omitempty"`
	// Optional. How the results should be sorted.  Presently, the only permitted
	// values are `"timestamp asc"` (default) and `"timestamp desc"`. The first
	// option returns entries in order of increasing values of
	// `LogEntry.timestamp` (oldest first), and the second option returns entries
	// in order of decreasing timestamps (newest first).  Entries with equal
	// timestamps are returned in order of their `insert_id` values.
	OrderBy string `protobuf:"bytes,3,opt,name=order_by,json=orderBy" json:"order_by,omitempty"`
	// Optional. The maximum number of results to return from this request.
	// Non-positive values are ignored.  The presence of `next_page_token` in the
	// response indicates that more results might be available.
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// Optional. If present, then retrieve the next batch of results from the
	// preceding call to this method.  `page_token` must be the value of
	// `next_page_token` from the previous response.  The values of other method
	// parameters should be identical to those in the previous call.
	PageToken string `protobuf:"bytes,5,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListLogEntriesRequest) Reset()                    { *m = ListLogEntriesRequest{} }
func (m *ListLogEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListLogEntriesRequest) ProtoMessage()               {}
func (*ListLogEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *ListLogEntriesRequest) GetProjectIds() []string {
	if m != nil {
		return m.ProjectIds
	}
	return nil
}

func (m *ListLogEntriesRequest) GetResourceNames() []string {
	if m != nil {
		return m.ResourceNames
	}
	return nil
}

func (m *ListLogEntriesRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListLogEntriesRequest) GetOrderBy() string {
	if m != nil {
		return m.OrderBy
	}
	return ""
}

func (m *ListLogEntriesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListLogEntriesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// Result returned from `ListLogEntries`.
type ListLogEntriesResponse struct {
	// A list of log entries.  If `entries` is empty, `nextPageToken` may still be
	// returned, indicating that more entries may exist.  See `nextPageToken` for
	// more information.
	Entries []*LogEntry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
	// If there might be more results than those appearing in this response, then
	// `nextPageToken` is included.  To get the next set of results, call this
	// method again using the value of `nextPageToken` as `pageToken`.
	//
	// If a value for `next_page_token` appears and the `entries` field is empty,
	// it means that the search found no log entries so far but it did not have
	// time to search all the possible log entries.  Retry the method with this
	// value for `page_token` to continue the search.  Alternatively, consider
	// speeding up the search by changing your filter to specify a single log name
	// or resource type, or to narrow the time range of the search.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListLogEntriesResponse) Reset()                    { *m = ListLogEntriesResponse{} }
func (m *ListLogEntriesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListLogEntriesResponse) ProtoMessage()               {}
func (*ListLogEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *ListLogEntriesResponse) GetEntries() []*LogEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *ListLogEntriesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The parameters to ListMonitoredResourceDescriptors
type ListMonitoredResourceDescriptorsRequest struct {
	// Optional. The maximum number of results to return from this request.
	// Non-positive values are ignored.  The presence of `nextPageToken` in the
	// response indicates that more results might be available.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// Optional. If present, then retrieve the next batch of results from the
	// preceding call to this method.  `pageToken` must be the value of
	// `nextPageToken` from the previous response.  The values of other method
	// parameters should be identical to those in the previous call.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListMonitoredResourceDescriptorsRequest) Reset() {
	*m = ListMonitoredResourceDescriptorsRequest{}
}
func (m *ListMonitoredResourceDescriptorsRequest) String() string { return proto.CompactTextString(m) }
func (*ListMonitoredResourceDescriptorsRequest) ProtoMessage()    {}
func (*ListMonitoredResourceDescriptorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{6}
}

func (m *ListMonitoredResourceDescriptorsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListMonitoredResourceDescriptorsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// Result returned from ListMonitoredResourceDescriptors.
type ListMonitoredResourceDescriptorsResponse struct {
	// A list of resource descriptors.
	ResourceDescriptors []*google_api3.MonitoredResourceDescriptor `protobuf:"bytes,1,rep,name=resource_descriptors,json=resourceDescriptors" json:"resource_descriptors,omitempty"`
	// If there might be more results than those appearing in this response, then
	// `nextPageToken` is included.  To get the next set of results, call this
	// method again using the value of `nextPageToken` as `pageToken`.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListMonitoredResourceDescriptorsResponse) Reset() {
	*m = ListMonitoredResourceDescriptorsResponse{}
}
func (m *ListMonitoredResourceDescriptorsResponse) String() string { return proto.CompactTextString(m) }
func (*ListMonitoredResourceDescriptorsResponse) ProtoMessage()    {}
func (*ListMonitoredResourceDescriptorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{7}
}

func (m *ListMonitoredResourceDescriptorsResponse) GetResourceDescriptors() []*google_api3.MonitoredResourceDescriptor {
	if m != nil {
		return m.ResourceDescriptors
	}
	return nil
}

func (m *ListMonitoredResourceDescriptorsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The parameters to ListLogs.
type ListLogsRequest struct {
	// Required. The resource name that owns the logs:
	//
	//     "projects/[PROJECT_ID]"
	//     "organizations/[ORGANIZATION_ID]"
	//     "billingAccounts/[BILLING_ACCOUNT_ID]"
	//     "folders/[FOLDER_ID]"
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// Optional. The maximum number of results to return from this request.
	// Non-positive values are ignored.  The presence of `nextPageToken` in the
	// response indicates that more results might be available.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// Optional. If present, then retrieve the next batch of results from the
	// preceding call to this method.  `pageToken` must be the value of
	// `nextPageToken` from the previous response.  The values of other method
	// parameters should be identical to those in the previous call.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListLogsRequest) Reset()                    { *m = ListLogsRequest{} }
func (m *ListLogsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListLogsRequest) ProtoMessage()               {}
func (*ListLogsRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *ListLogsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListLogsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListLogsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// Result returned from ListLogs.
type ListLogsResponse struct {
	// A list of log names. For example,
	// `"projects/my-project/syslog"` or
	// `"organizations/123/cloudresourcemanager.googleapis.com%2Factivity"`.
	LogNames []string `protobuf:"bytes,3,rep,name=log_names,json=logNames" json:"log_names,omitempty"`
	// If there might be more results than those appearing in this response, then
	// `nextPageToken` is included.  To get the next set of results, call this
	// method again using the value of `nextPageToken` as `pageToken`.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListLogsResponse) Reset()                    { *m = ListLogsResponse{} }
func (m *ListLogsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListLogsResponse) ProtoMessage()               {}
func (*ListLogsResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *ListLogsResponse) GetLogNames() []string {
	if m != nil {
		return m.LogNames
	}
	return nil
}

func (m *ListLogsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*DeleteLogRequest)(nil), "google.logging.v2.DeleteLogRequest")
	proto.RegisterType((*WriteLogEntriesRequest)(nil), "google.logging.v2.WriteLogEntriesRequest")
	proto.RegisterType((*WriteLogEntriesResponse)(nil), "google.logging.v2.WriteLogEntriesResponse")
	proto.RegisterType((*WriteLogEntriesPartialErrors)(nil), "google.logging.v2.WriteLogEntriesPartialErrors")
	proto.RegisterType((*ListLogEntriesRequest)(nil), "google.logging.v2.ListLogEntriesRequest")
	proto.RegisterType((*ListLogEntriesResponse)(nil), "google.logging.v2.ListLogEntriesResponse")
	proto.RegisterType((*ListMonitoredResourceDescriptorsRequest)(nil), "google.logging.v2.ListMonitoredResourceDescriptorsRequest")
	proto.RegisterType((*ListMonitoredResourceDescriptorsResponse)(nil), "google.logging.v2.ListMonitoredResourceDescriptorsResponse")
	proto.RegisterType((*ListLogsRequest)(nil), "google.logging.v2.ListLogsRequest")
	proto.RegisterType((*ListLogsResponse)(nil), "google.logging.v2.ListLogsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LoggingServiceV2 service

type LoggingServiceV2Client interface {
	// Deletes all the log entries in a log.
	// The log reappears if it receives new entries.
	// Log entries written shortly before the delete operation might not be
	// deleted.
	DeleteLog(ctx context.Context, in *DeleteLogRequest, opts ...grpc.CallOption) (*google_protobuf5.Empty, error)
	// ## Log entry resources
	//
	// Writes log entries to Stackdriver Logging. This API method is the
	// only way to send log entries to Stackdriver Logging. This method
	// is used, directly or indirectly, by the Stackdriver Logging agent
	// (fluentd) and all logging libraries configured to use Stackdriver
	// Logging.
	WriteLogEntries(ctx context.Context, in *WriteLogEntriesRequest, opts ...grpc.CallOption) (*WriteLogEntriesResponse, error)
	// Lists log entries.  Use this method to retrieve log entries from
	// Stackdriver Logging.  For ways to export log entries, see
	// [Exporting Logs](/logging/docs/export).
	ListLogEntries(ctx context.Context, in *ListLogEntriesRequest, opts ...grpc.CallOption) (*ListLogEntriesResponse, error)
	// Lists the descriptors for monitored resource types used by Stackdriver
	// Logging.
	ListMonitoredResourceDescriptors(ctx context.Context, in *ListMonitoredResourceDescriptorsRequest, opts ...grpc.CallOption) (*ListMonitoredResourceDescriptorsResponse, error)
	// Lists the logs in projects, organizations, folders, or billing accounts.
	// Only logs that have entries are listed.
	ListLogs(ctx context.Context, in *ListLogsRequest, opts ...grpc.CallOption) (*ListLogsResponse, error)
}

type loggingServiceV2Client struct {
	cc *grpc.ClientConn
}

func NewLoggingServiceV2Client(cc *grpc.ClientConn) LoggingServiceV2Client {
	return &loggingServiceV2Client{cc}
}

func (c *loggingServiceV2Client) DeleteLog(ctx context.Context, in *DeleteLogRequest, opts ...grpc.CallOption) (*google_protobuf5.Empty, error) {
	out := new(google_protobuf5.Empty)
	err := grpc.Invoke(ctx, "/google.logging.v2.LoggingServiceV2/DeleteLog", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggingServiceV2Client) WriteLogEntries(ctx context.Context, in *WriteLogEntriesRequest, opts ...grpc.CallOption) (*WriteLogEntriesResponse, error) {
	out := new(WriteLogEntriesResponse)
	err := grpc.Invoke(ctx, "/google.logging.v2.LoggingServiceV2/WriteLogEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggingServiceV2Client) ListLogEntries(ctx context.Context, in *ListLogEntriesRequest, opts ...grpc.CallOption) (*ListLogEntriesResponse, error) {
	out := new(ListLogEntriesResponse)
	err := grpc.Invoke(ctx, "/google.logging.v2.LoggingServiceV2/ListLogEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggingServiceV2Client) ListMonitoredResourceDescriptors(ctx context.Context, in *ListMonitoredResourceDescriptorsRequest, opts ...grpc.CallOption) (*ListMonitoredResourceDescriptorsResponse, error) {
	out := new(ListMonitoredResourceDescriptorsResponse)
	err := grpc.Invoke(ctx, "/google.logging.v2.LoggingServiceV2/ListMonitoredResourceDescriptors", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggingServiceV2Client) ListLogs(ctx context.Context, in *ListLogsRequest, opts ...grpc.CallOption) (*ListLogsResponse, error) {
	out := new(ListLogsResponse)
	err := grpc.Invoke(ctx, "/google.logging.v2.LoggingServiceV2/ListLogs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LoggingServiceV2 service

type LoggingServiceV2Server interface {
	// Deletes all the log entries in a log.
	// The log reappears if it receives new entries.
	// Log entries written shortly before the delete operation might not be
	// deleted.
	DeleteLog(context.Context, *DeleteLogRequest) (*google_protobuf5.Empty, error)
	// ## Log entry resources
	//
	// Writes log entries to Stackdriver Logging. This API method is the
	// only way to send log entries to Stackdriver Logging. This method
	// is used, directly or indirectly, by the Stackdriver Logging agent
	// (fluentd) and all logging libraries configured to use Stackdriver
	// Logging.
	WriteLogEntries(context.Context, *WriteLogEntriesRequest) (*WriteLogEntriesResponse, error)
	// Lists log entries.  Use this method to retrieve log entries from
	// Stackdriver Logging.  For ways to export log entries, see
	// [Exporting Logs](/logging/docs/export).
	ListLogEntries(context.Context, *ListLogEntriesRequest) (*ListLogEntriesResponse, error)
	// Lists the descriptors for monitored resource types used by Stackdriver
	// Logging.
	ListMonitoredResourceDescriptors(context.Context, *ListMonitoredResourceDescriptorsRequest) (*ListMonitoredResourceDescriptorsResponse, error)
	// Lists the logs in projects, organizations, folders, or billing accounts.
	// Only logs that have entries are listed.
	ListLogs(context.Context, *ListLogsRequest) (*ListLogsResponse, error)
}

func RegisterLoggingServiceV2Server(s *grpc.Server, srv LoggingServiceV2Server) {
	s.RegisterService(&_LoggingServiceV2_serviceDesc, srv)
}

func _LoggingServiceV2_DeleteLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceV2Server).DeleteLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.LoggingServiceV2/DeleteLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceV2Server).DeleteLog(ctx, req.(*DeleteLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoggingServiceV2_WriteLogEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteLogEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceV2Server).WriteLogEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.LoggingServiceV2/WriteLogEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceV2Server).WriteLogEntries(ctx, req.(*WriteLogEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoggingServiceV2_ListLogEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLogEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceV2Server).ListLogEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.LoggingServiceV2/ListLogEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceV2Server).ListLogEntries(ctx, req.(*ListLogEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoggingServiceV2_ListMonitoredResourceDescriptors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMonitoredResourceDescriptorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceV2Server).ListMonitoredResourceDescriptors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.LoggingServiceV2/ListMonitoredResourceDescriptors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceV2Server).ListMonitoredResourceDescriptors(ctx, req.(*ListMonitoredResourceDescriptorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoggingServiceV2_ListLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceV2Server).ListLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.LoggingServiceV2/ListLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceV2Server).ListLogs(ctx, req.(*ListLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoggingServiceV2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.logging.v2.LoggingServiceV2",
	HandlerType: (*LoggingServiceV2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteLog",
			Handler:    _LoggingServiceV2_DeleteLog_Handler,
		},
		{
			MethodName: "WriteLogEntries",
			Handler:    _LoggingServiceV2_WriteLogEntries_Handler,
		},
		{
			MethodName: "ListLogEntries",
			Handler:    _LoggingServiceV2_ListLogEntries_Handler,
		},
		{
			MethodName: "ListMonitoredResourceDescriptors",
			Handler:    _LoggingServiceV2_ListMonitoredResourceDescriptors_Handler,
		},
		{
			MethodName: "ListLogs",
			Handler:    _LoggingServiceV2_ListLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/logging/v2/logging.proto",
}

func init() { proto.RegisterFile("google/logging/v2/logging.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 991 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0xc6, 0x4a, 0xb1, 0x23, 0x8d, 0x1a, 0x5b, 0xd9, 0xc4, 0xb2, 0x22, 0xd9, 0xb5, 0x4a, 0x23,
	0xb5, 0x22, 0x20, 0x24, 0xca, 0x22, 0x40, 0xe2, 0x20, 0x17, 0x27, 0x46, 0x51, 0xc0, 0x29, 0x0c,
	0xba, 0x75, 0x80, 0xc0, 0x80, 0x40, 0x49, 0x1b, 0x62, 0x1b, 0x8a, 0xcb, 0xee, 0xae, 0xe4, 0x2a,
	0x41, 0x2e, 0x39, 0xf4, 0x05, 0x7a, 0xe9, 0x33, 0xf4, 0xd0, 0xb7, 0xe8, 0xa5, 0x87, 0x5e, 0x8a,
	0x02, 0x7d, 0x80, 0x3c, 0x44, 0x8f, 0x05, 0x77, 0x97, 0x32, 0xf5, 0x13, 0x59, 0xee, 0x4d, 0x3b,
	0xf3, 0xed, 0xcc, 0x7c, 0xc3, 0x6f, 0x66, 0x05, 0x3b, 0x01, 0x63, 0x41, 0x48, 0x9c, 0x90, 0x05,
	0x01, 0x8d, 0x02, 0x67, 0xe8, 0xa6, 0x3f, 0xed, 0x98, 0x33, 0xc9, 0xf0, 0x4d, 0x0d, 0xb0, 0x53,
	0xeb, 0xd0, 0xad, 0x6d, 0x99, 0x3b, 0x7e, 0x4c, 0x1d, 0x3f, 0x8a, 0x98, 0xf4, 0x25, 0x65, 0x91,
	0xd0, 0x17, 0x6a, 0xbb, 0x19, 0x6f, 0x9f, 0x45, 0x54, 0x32, 0x4e, 0x7a, 0x6d, 0x4e, 0x04, 0x1b,
	0xf0, 0x2e, 0x31, 0xa0, 0xcf, 0xe6, 0xa6, 0x6d, 0x93, 0x48, 0xf2, 0x91, 0x81, 0x7c, 0x6a, 0x20,
	0xea, 0xd4, 0x19, 0xbc, 0x72, 0x7a, 0x03, 0xae, 0x12, 0x19, 0x7f, 0x7d, 0xda, 0x4f, 0xfa, 0xb1,
	0x4c, 0x2f, 0xef, 0x4c, 0x3b, 0x25, 0xed, 0x13, 0x21, 0xfd, 0x7e, 0x6c, 0x00, 0x9b, 0x06, 0xc0,
	0xe3, 0xae, 0x23, 0xa4, 0x2f, 0x07, 0xa6, 0x7c, 0xeb, 0x3e, 0x94, 0x9f, 0x91, 0x90, 0x48, 0x72,
	0xc4, 0x02, 0x8f, 0xfc, 0x30, 0x20, 0x42, 0xe2, 0x3b, 0x50, 0x48, 0xaa, 0x8b, 0xfc, 0x3e, 0xa9,
	0xa2, 0x06, 0x6a, 0x16, 0xbd, 0xeb, 0x21, 0x0b, 0xbe, 0xf1, 0xfb, 0xc4, 0xfa, 0x27, 0x07, 0x95,
	0x17, 0x9c, 0x2a, 0xf8, 0x61, 0x24, 0x39, 0x25, 0xe2, 0xf2, 0x5b, 0xf8, 0x11, 0x14, 0xd2, 0x86,
	0x54, 0x73, 0x0d, 0xd4, 0x2c, 0xb9, 0xdb, 0xb6, 0xe9, 0xb3, 0x1f, 0x53, 0xfb, 0x79, 0xda, 0x36,
	0xcf, 0x80, 0xbc, 0x31, 0x1c, 0x3f, 0x87, 0xd5, 0xd0, 0xef, 0x90, 0x50, 0x54, 0xf3, 0x8d, 0x7c,
	0xb3, 0xe4, 0x3e, 0xb0, 0x67, 0x3e, 0x90, 0x3d, 0xbf, 0x20, 0xfb, 0x48, 0xdd, 0x4b, 0x8c, 0x23,
	0xcf, 0x04, 0xc1, 0x0f, 0xe0, 0x3a, 0xd1, 0xa8, 0xea, 0x35, 0x15, 0xaf, 0x3e, 0x27, 0x9e, 0x09,
	0x35, 0xf2, 0x52, 0x2c, 0xde, 0x83, 0xf5, 0xd8, 0xe7, 0x92, 0xfa, 0x61, 0x5b, 0x0c, 0xba, 0x5d,
	0x22, 0x44, 0x75, 0xa5, 0x81, 0x9a, 0x05, 0x6f, 0xcd, 0x98, 0x4f, 0xb4, 0xb5, 0xf6, 0x08, 0x4a,
	0x99, 0xb4, 0xb8, 0x0c, 0xf9, 0xd7, 0x64, 0x64, 0xda, 0x91, 0xfc, 0xc4, 0xb7, 0x61, 0x65, 0xe8,
	0x87, 0x03, 0xdd, 0x87, 0xa2, 0xa7, 0x0f, 0xfb, 0xb9, 0x87, 0xc8, 0xba, 0x03, 0x9b, 0x33, 0x44,
	0x44, 0xcc, 0x22, 0x41, 0xac, 0x0f, 0x08, 0xb6, 0xa6, 0x7c, 0xc7, 0x3a, 0xef, 0x21, 0xe7, 0x8c,
	0x0b, 0xdc, 0x87, 0xf2, 0x58, 0x4f, 0x6d, 0xa2, 0x6c, 0x55, 0xa4, 0xf8, 0x3d, 0xbd, 0xbc, 0x5f,
	0x13, 0xa1, 0xc6, 0xe4, 0xf5, 0x51, 0xf7, 0x61, 0x2d, 0x9c, 0x30, 0xd6, 0xbe, 0x83, 0x5b, 0x73,
	0x60, 0x59, 0xb6, 0x2b, 0x9a, 0x6d, 0x33, 0xcb, 0xb6, 0xe4, 0xe2, 0xb4, 0x18, 0x1e, 0x77, 0xed,
	0x13, 0x25, 0xc3, 0x6c, 0x07, 0xfe, 0x44, 0xb0, 0x71, 0x44, 0x85, 0x9c, 0xd5, 0xd6, 0x0e, 0x94,
	0x62, 0xce, 0xbe, 0x27, 0x5d, 0xd9, 0xa6, 0x3d, 0x4d, 0xad, 0xe8, 0x81, 0x31, 0x7d, 0xdd, 0x13,
	0xf8, 0x2e, 0xac, 0xa5, 0x92, 0x51, 0x0a, 0x14, 0xd5, 0x82, 0xc2, 0xdc, 0x48, 0xad, 0x89, 0x0e,
	0x05, 0xae, 0xc0, 0xea, 0x2b, 0x1a, 0x4a, 0xc2, 0x4d, 0xfb, 0xcd, 0x29, 0xd1, 0x2e, 0xe3, 0x3d,
	0xc2, 0xdb, 0x9d, 0x51, 0x35, 0xaf, 0xb5, 0xab, 0xce, 0x07, 0x23, 0x5c, 0x87, 0x62, 0xec, 0x07,
	0xa4, 0x2d, 0xe8, 0x1b, 0x52, 0xbd, 0xa6, 0xa8, 0x15, 0x12, 0xc3, 0x09, 0x7d, 0x43, 0xf0, 0x36,
	0x80, 0x72, 0x4a, 0xf6, 0x9a, 0x44, 0x4a, 0x12, 0x45, 0x4f, 0xc1, 0xbf, 0x4d, 0x0c, 0xd6, 0x39,
	0x54, 0xa6, 0xf9, 0xe8, 0x2f, 0x9a, 0xd5, 0x21, 0xba, 0x82, 0x0e, 0x3f, 0x87, 0xf5, 0x88, 0xfc,
	0x28, 0xdb, 0x99, 0xa4, 0x9a, 0xc8, 0x8d, 0xc4, 0x7c, 0x3c, 0x4e, 0x4c, 0x60, 0x2f, 0x49, 0x3c,
	0x33, 0x58, 0xcf, 0x88, 0xe8, 0x72, 0x1a, 0x4b, 0xc6, 0xc7, 0xad, 0x9d, 0xe0, 0x87, 0x16, 0xf2,
	0xcb, 0x4d, 0xf3, 0xfb, 0x0d, 0x41, 0xf3, 0xf2, 0x3c, 0x86, 0xf2, 0x4b, 0xb8, 0x3d, 0xfe, 0x44,
	0xbd, 0x0b, 0xbf, 0xe1, 0xbf, 0xb7, 0x70, 0x21, 0x5c, 0xc4, 0xf3, 0x6e, 0xf1, 0xd9, 0x1c, 0x57,
	0xe8, 0xcb, 0xba, 0xf9, 0x20, 0x63, 0xfe, 0x15, 0x58, 0x8d, 0x7d, 0x4e, 0x22, 0x69, 0xa6, 0xd4,
	0x9c, 0x26, 0xfb, 0x92, 0x5b, 0xd8, 0x97, 0xfc, 0x74, 0x5f, 0x5e, 0x40, 0xf9, 0x22, 0x8d, 0xa1,
	0x5f, 0x87, 0x62, 0xba, 0x1e, 0xf5, 0x2e, 0x2b, 0x7a, 0x05, 0xb3, 0x1f, 0x97, 0xae, 0xdf, 0xfd,
	0x7b, 0x05, 0xca, 0x47, 0x5a, 0x20, 0x27, 0x84, 0x0f, 0x69, 0x97, 0x9c, 0xba, 0xf8, 0x1c, 0x8a,
	0xe3, 0x15, 0x8e, 0x77, 0xe7, 0xe8, 0x68, 0x7a, 0xc1, 0xd7, 0x2a, 0x29, 0x28, 0x7d, 0x2f, 0xec,
	0xc3, 0xe4, 0x31, 0xb1, 0xee, 0xbf, 0xff, 0xeb, 0xc3, 0xcf, 0xb9, 0xbd, 0xd6, 0x5d, 0x67, 0xe8,
	0x76, 0x88, 0xf4, 0xbf, 0x70, 0xde, 0xa6, 0x35, 0x3f, 0x31, 0xc3, 0x26, 0x9c, 0x56, 0xf2, 0x74,
	0x09, 0xa7, 0xf5, 0x0e, 0xff, 0x84, 0x60, 0x7d, 0x6a, 0x97, 0xe0, 0x7b, 0x4b, 0xef, 0xe7, 0x5a,
	0x6b, 0x19, 0xa8, 0xd9, 0x80, 0x5b, 0xaa, 0xb2, 0x8a, 0x75, 0x33, 0x79, 0x3a, 0xcd, 0x34, 0xec,
	0x9f, 0x27, 0xe0, 0x7d, 0xd4, 0xc2, 0xef, 0x11, 0xac, 0x4d, 0x0e, 0x1a, 0x6e, 0xce, 0x9b, 0xa7,
	0x79, 0xbb, 0xa5, 0x76, 0x6f, 0x09, 0xa4, 0xa9, 0xa2, 0xae, 0xaa, 0xd8, 0xb0, 0xca, 0xd9, 0x2a,
	0x42, 0x2a, 0x64, 0x52, 0xc4, 0xef, 0x08, 0x1a, 0x97, 0x0d, 0x03, 0xde, 0xff, 0x48, 0xb2, 0x25,
	0x26, 0xb5, 0xf6, 0xf8, 0x7f, 0xdd, 0x35, 0xa5, 0x37, 0x55, 0xe9, 0x16, 0x6e, 0x24, 0xa5, 0xf7,
	0x17, 0x95, 0xc8, 0xa1, 0x90, 0x8a, 0x17, 0x5b, 0x1f, 0xef, 0xcd, 0xb8, 0xac, 0xdd, 0x85, 0x18,
	0x93, 0x7e, 0x5b, 0xa5, 0xdf, 0xc4, 0x1b, 0x49, 0xfa, 0xb7, 0x7a, 0xc4, 0x9e, 0xb4, 0x9c, 0xd6,
	0x3b, 0x25, 0xa6, 0x83, 0x5f, 0x10, 0x6c, 0x74, 0x59, 0x7f, 0x36, 0xd2, 0xc1, 0x27, 0x46, 0xee,
	0xc7, 0x89, 0x52, 0x8f, 0xd1, 0xcb, 0x87, 0x06, 0x12, 0xb0, 0xd0, 0x8f, 0x02, 0x9b, 0xf1, 0xc0,
	0x09, 0x48, 0xa4, 0x74, 0xec, 0x68, 0x97, 0x1f, 0x53, 0x91, 0xf9, 0xa3, 0xf5, 0xd8, 0xfc, 0xfc,
	0x17, 0xa1, 0x5f, 0x73, 0x9b, 0x5f, 0xe9, 0xdb, 0x4f, 0x43, 0x36, 0xe8, 0xd9, 0x26, 0xb4, 0x7d,
	0xea, 0xfe, 0x91, 0x7a, 0xce, 0x94, 0xe7, 0xcc, 0x78, 0xce, 0x4e, 0xdd, 0xce, 0xaa, 0x8a, 0xfd,
	0xe5, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x96, 0xc5, 0x3e, 0x3a, 0x0a, 0x00, 0x00,
}
