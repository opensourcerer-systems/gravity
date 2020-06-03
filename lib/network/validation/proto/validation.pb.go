// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: validation.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	agentpb "github.com/gravitational/satellite/agent/proto/agentpb"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// CheckPortsRequest describes a ports network test request
type CheckPortsRequest struct {
	// Listen specifies the listen endpoints
	Listen []*Addr `protobuf:"bytes,1,rep,name=listen,proto3" json:"listen,omitempty"`
	// Ping specifies the ping endpoints
	Ping []*Addr `protobuf:"bytes,2,rep,name=ping,proto3" json:"ping,omitempty"`
	// Duration specifies the maximum duration for the request
	Duration             *types.Duration `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CheckPortsRequest) Reset()         { *m = CheckPortsRequest{} }
func (m *CheckPortsRequest) String() string { return proto.CompactTextString(m) }
func (*CheckPortsRequest) ProtoMessage()    {}
func (*CheckPortsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfc2ab0b60b7792f, []int{0}
}
func (m *CheckPortsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckPortsRequest.Unmarshal(m, b)
}
func (m *CheckPortsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckPortsRequest.Marshal(b, m, deterministic)
}
func (m *CheckPortsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckPortsRequest.Merge(m, src)
}
func (m *CheckPortsRequest) XXX_Size() int {
	return xxx_messageInfo_CheckPortsRequest.Size(m)
}
func (m *CheckPortsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckPortsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckPortsRequest proto.InternalMessageInfo

func (m *CheckPortsRequest) GetListen() []*Addr {
	if m != nil {
		return m.Listen
	}
	return nil
}

func (m *CheckPortsRequest) GetPing() []*Addr {
	if m != nil {
		return m.Ping
	}
	return nil
}

func (m *CheckPortsRequest) GetDuration() *types.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

// CheckPortsResponse describes the results of a ports network test
type CheckPortsResponse struct {
	// Listen describes the listen test results
	Listen []*ServerResult `protobuf:"bytes,1,rep,name=listen,proto3" json:"listen,omitempty"`
	// Ping describes the ping test results
	Ping                 []*ServerResult `protobuf:"bytes,2,rep,name=ping,proto3" json:"ping,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CheckPortsResponse) Reset()         { *m = CheckPortsResponse{} }
func (m *CheckPortsResponse) String() string { return proto.CompactTextString(m) }
func (*CheckPortsResponse) ProtoMessage()    {}
func (*CheckPortsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfc2ab0b60b7792f, []int{1}
}
func (m *CheckPortsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckPortsResponse.Unmarshal(m, b)
}
func (m *CheckPortsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckPortsResponse.Marshal(b, m, deterministic)
}
func (m *CheckPortsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckPortsResponse.Merge(m, src)
}
func (m *CheckPortsResponse) XXX_Size() int {
	return xxx_messageInfo_CheckPortsResponse.Size(m)
}
func (m *CheckPortsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckPortsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckPortsResponse proto.InternalMessageInfo

func (m *CheckPortsResponse) GetListen() []*ServerResult {
	if m != nil {
		return m.Listen
	}
	return nil
}

func (m *CheckPortsResponse) GetPing() []*ServerResult {
	if m != nil {
		return m.Ping
	}
	return nil
}

// CheckBandwidthRequest describes a bandwidth check network test
type CheckBandwidthRequest struct {
	// Listen specifies the listen endpoint
	Listen *Addr `protobuf:"bytes,1,opt,name=listen,proto3" json:"listen,omitempty"`
	// Ping specifies the ping endpoints
	Ping []*Addr `protobuf:"bytes,2,rep,name=ping,proto3" json:"ping,omitempty"`
	// Duration specifies the maximum duration for the request
	Duration             *types.Duration `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CheckBandwidthRequest) Reset()         { *m = CheckBandwidthRequest{} }
func (m *CheckBandwidthRequest) String() string { return proto.CompactTextString(m) }
func (*CheckBandwidthRequest) ProtoMessage()    {}
func (*CheckBandwidthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfc2ab0b60b7792f, []int{2}
}
func (m *CheckBandwidthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckBandwidthRequest.Unmarshal(m, b)
}
func (m *CheckBandwidthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckBandwidthRequest.Marshal(b, m, deterministic)
}
func (m *CheckBandwidthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckBandwidthRequest.Merge(m, src)
}
func (m *CheckBandwidthRequest) XXX_Size() int {
	return xxx_messageInfo_CheckBandwidthRequest.Size(m)
}
func (m *CheckBandwidthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckBandwidthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckBandwidthRequest proto.InternalMessageInfo

func (m *CheckBandwidthRequest) GetListen() *Addr {
	if m != nil {
		return m.Listen
	}
	return nil
}

func (m *CheckBandwidthRequest) GetPing() []*Addr {
	if m != nil {
		return m.Ping
	}
	return nil
}

func (m *CheckBandwidthRequest) GetDuration() *types.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

// CheckBandwidthResponse describes the results of a bandwidth check
type CheckBandwidthResponse struct {
	// Bandwidth is the result of a bandwidth test
	Bandwidth            uint64   `protobuf:"varint,1,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckBandwidthResponse) Reset()         { *m = CheckBandwidthResponse{} }
func (m *CheckBandwidthResponse) String() string { return proto.CompactTextString(m) }
func (*CheckBandwidthResponse) ProtoMessage()    {}
func (*CheckBandwidthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfc2ab0b60b7792f, []int{3}
}
func (m *CheckBandwidthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckBandwidthResponse.Unmarshal(m, b)
}
func (m *CheckBandwidthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckBandwidthResponse.Marshal(b, m, deterministic)
}
func (m *CheckBandwidthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckBandwidthResponse.Merge(m, src)
}
func (m *CheckBandwidthResponse) XXX_Size() int {
	return xxx_messageInfo_CheckBandwidthResponse.Size(m)
}
func (m *CheckBandwidthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckBandwidthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckBandwidthResponse proto.InternalMessageInfo

func (m *CheckBandwidthResponse) GetBandwidth() uint64 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

// ServerResult defines the operation result for a server
type ServerResult struct {
	// Code specifies the result, with 0 for success
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// Error specifies an error message
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	// Server specifies which server the result is from
	Server               *Addr    `protobuf:"bytes,3,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerResult) Reset()         { *m = ServerResult{} }
func (m *ServerResult) String() string { return proto.CompactTextString(m) }
func (*ServerResult) ProtoMessage()    {}
func (*ServerResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfc2ab0b60b7792f, []int{4}
}
func (m *ServerResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerResult.Unmarshal(m, b)
}
func (m *ServerResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerResult.Marshal(b, m, deterministic)
}
func (m *ServerResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerResult.Merge(m, src)
}
func (m *ServerResult) XXX_Size() int {
	return xxx_messageInfo_ServerResult.Size(m)
}
func (m *ServerResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerResult.DiscardUnknown(m)
}

var xxx_messageInfo_ServerResult proto.InternalMessageInfo

func (m *ServerResult) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ServerResult) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ServerResult) GetServer() *Addr {
	if m != nil {
		return m.Server
	}
	return nil
}

// Addr defines an endpoint address
type Addr struct {
	// Network specifies the type of network (tcp, udp)
	Network string `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	// Addr specifies the address as IP or IP:port
	Addr                 string   `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Addr) Reset()         { *m = Addr{} }
func (m *Addr) String() string { return proto.CompactTextString(m) }
func (*Addr) ProtoMessage()    {}
func (*Addr) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfc2ab0b60b7792f, []int{5}
}
func (m *Addr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Addr.Unmarshal(m, b)
}
func (m *Addr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Addr.Marshal(b, m, deterministic)
}
func (m *Addr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Addr.Merge(m, src)
}
func (m *Addr) XXX_Size() int {
	return xxx_messageInfo_Addr.Size(m)
}
func (m *Addr) XXX_DiscardUnknown() {
	xxx_messageInfo_Addr.DiscardUnknown(m)
}

var xxx_messageInfo_Addr proto.InternalMessageInfo

func (m *Addr) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *Addr) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

// ValidateRequest defines a request to run validation
type ValidateRequest struct {
	// Manifest specifies the application manifest with
	// requirements
	Manifest []byte `protobuf:"bytes,1,opt,name=manifest,proto3" json:"manifest,omitempty"`
	// Profile specifies the node profile to validate against
	Profile string `protobuf:"bytes,2,opt,name=profile,proto3" json:"profile,omitempty"`
	// FullRequirements forces validation of all requirements
	// from the manifest.
	// This is used to validate requirements during installation.
	FullRequirements bool `protobuf:"varint,3,opt,name=full_requirements,json=fullRequirements,proto3" json:"full_requirements,omitempty"`
	// ValidateOptions is additional validation options
	Options *ValidateOptions `protobuf:"bytes,4,opt,name=options,proto3" json:"options,omitempty"`
	// Docker specifies the Docker configuration to validate
	Docker               *Docker  `protobuf:"bytes,5,opt,name=docker,proto3" json:"docker,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateRequest) Reset()         { *m = ValidateRequest{} }
func (m *ValidateRequest) String() string { return proto.CompactTextString(m) }
func (*ValidateRequest) ProtoMessage()    {}
func (*ValidateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfc2ab0b60b7792f, []int{6}
}
func (m *ValidateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateRequest.Unmarshal(m, b)
}
func (m *ValidateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateRequest.Marshal(b, m, deterministic)
}
func (m *ValidateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateRequest.Merge(m, src)
}
func (m *ValidateRequest) XXX_Size() int {
	return xxx_messageInfo_ValidateRequest.Size(m)
}
func (m *ValidateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateRequest proto.InternalMessageInfo

func (m *ValidateRequest) GetManifest() []byte {
	if m != nil {
		return m.Manifest
	}
	return nil
}

func (m *ValidateRequest) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *ValidateRequest) GetFullRequirements() bool {
	if m != nil {
		return m.FullRequirements
	}
	return false
}

func (m *ValidateRequest) GetOptions() *ValidateOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *ValidateRequest) GetDocker() *Docker {
	if m != nil {
		return m.Docker
	}
	return nil
}

// ValidateResponse describes a validation response
type ValidateResponse struct {
	// Failed lists the failed probes
	Failed               []*agentpb.Probe `protobuf:"bytes,1,rep,name=failed,proto3" json:"failed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ValidateResponse) Reset()         { *m = ValidateResponse{} }
func (m *ValidateResponse) String() string { return proto.CompactTextString(m) }
func (*ValidateResponse) ProtoMessage()    {}
func (*ValidateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfc2ab0b60b7792f, []int{7}
}
func (m *ValidateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateResponse.Unmarshal(m, b)
}
func (m *ValidateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateResponse.Marshal(b, m, deterministic)
}
func (m *ValidateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateResponse.Merge(m, src)
}
func (m *ValidateResponse) XXX_Size() int {
	return xxx_messageInfo_ValidateResponse.Size(m)
}
func (m *ValidateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateResponse proto.InternalMessageInfo

func (m *ValidateResponse) GetFailed() []*agentpb.Probe {
	if m != nil {
		return m.Failed
	}
	return nil
}

// ValidateOptions is additional validation options
type ValidateOptions struct {
	// VxlanPort is the custom overlay network port
	VxlanPort int32 `protobuf:"varint,1,opt,name=vxlan_port,json=vxlanPort,proto3" json:"vxlan_port,omitempty"`
	// DnsAddrs specifies the list of listen IP addresses for coredns
	DnsAddrs []string `protobuf:"bytes,2,rep,name=dns_addrs,json=dnsAddrs,proto3" json:"dns_addrs,omitempty"`
	// DnsPort specifies the DNS port for coredns
	DnsPort              int32    `protobuf:"varint,3,opt,name=dns_port,json=dnsPort,proto3" json:"dns_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateOptions) Reset()         { *m = ValidateOptions{} }
func (m *ValidateOptions) String() string { return proto.CompactTextString(m) }
func (*ValidateOptions) ProtoMessage()    {}
func (*ValidateOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfc2ab0b60b7792f, []int{8}
}
func (m *ValidateOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateOptions.Unmarshal(m, b)
}
func (m *ValidateOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateOptions.Marshal(b, m, deterministic)
}
func (m *ValidateOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateOptions.Merge(m, src)
}
func (m *ValidateOptions) XXX_Size() int {
	return xxx_messageInfo_ValidateOptions.Size(m)
}
func (m *ValidateOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateOptions proto.InternalMessageInfo

func (m *ValidateOptions) GetVxlanPort() int32 {
	if m != nil {
		return m.VxlanPort
	}
	return 0
}

func (m *ValidateOptions) GetDnsAddrs() []string {
	if m != nil {
		return m.DnsAddrs
	}
	return nil
}

func (m *ValidateOptions) GetDnsPort() int32 {
	if m != nil {
		return m.DnsPort
	}
	return 0
}

// Docker groups Docker-relevant attributes to validate
type Docker struct {
	// StorageDriver specifies the Docker storage driver
	StorageDriver string `protobuf:"bytes,1,opt,name=storage_driver,json=storageDriver,proto3" json:"storage_driver,omitempty"`
	// Device is the device used for devicemapper
	Device               string   `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Docker) Reset()         { *m = Docker{} }
func (m *Docker) String() string { return proto.CompactTextString(m) }
func (*Docker) ProtoMessage()    {}
func (*Docker) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfc2ab0b60b7792f, []int{9}
}
func (m *Docker) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Docker.Unmarshal(m, b)
}
func (m *Docker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Docker.Marshal(b, m, deterministic)
}
func (m *Docker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Docker.Merge(m, src)
}
func (m *Docker) XXX_Size() int {
	return xxx_messageInfo_Docker.Size(m)
}
func (m *Docker) XXX_DiscardUnknown() {
	xxx_messageInfo_Docker.DiscardUnknown(m)
}

var xxx_messageInfo_Docker proto.InternalMessageInfo

func (m *Docker) GetStorageDriver() string {
	if m != nil {
		return m.StorageDriver
	}
	return ""
}

func (m *Docker) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func init() {
	proto.RegisterType((*CheckPortsRequest)(nil), "proto.CheckPortsRequest")
	proto.RegisterType((*CheckPortsResponse)(nil), "proto.CheckPortsResponse")
	proto.RegisterType((*CheckBandwidthRequest)(nil), "proto.CheckBandwidthRequest")
	proto.RegisterType((*CheckBandwidthResponse)(nil), "proto.CheckBandwidthResponse")
	proto.RegisterType((*ServerResult)(nil), "proto.ServerResult")
	proto.RegisterType((*Addr)(nil), "proto.Addr")
	proto.RegisterType((*ValidateRequest)(nil), "proto.ValidateRequest")
	proto.RegisterType((*ValidateResponse)(nil), "proto.ValidateResponse")
	proto.RegisterType((*ValidateOptions)(nil), "proto.ValidateOptions")
	proto.RegisterType((*Docker)(nil), "proto.Docker")
}

func init() { proto.RegisterFile("validation.proto", fileDescriptor_bfc2ab0b60b7792f) }

var fileDescriptor_bfc2ab0b60b7792f = []byte{
	// 608 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xcd, 0x6e, 0xd4, 0x3c,
	0x14, 0x55, 0xda, 0x99, 0x74, 0x72, 0xfb, 0xf3, 0xb5, 0xee, 0x47, 0x49, 0x87, 0x16, 0xaa, 0xa0,
	0x42, 0xa5, 0x4a, 0x29, 0x2a, 0x3f, 0x0b, 0x58, 0x15, 0x2a, 0xb1, 0x42, 0x54, 0x46, 0xea, 0x0e,
	0x8d, 0x32, 0xe3, 0x3b, 0xd3, 0x30, 0xa9, 0x9d, 0xda, 0xce, 0x94, 0xc7, 0x60, 0xc1, 0x6b, 0xf1,
	0x04, 0xbc, 0x0c, 0xb2, 0x63, 0xcf, 0x3f, 0x5b, 0x56, 0xf1, 0x3d, 0xf7, 0xdc, 0xeb, 0xe3, 0x73,
	0xed, 0xc0, 0xf6, 0x28, 0x2b, 0x72, 0x96, 0xe9, 0x5c, 0xf0, 0xb4, 0x94, 0x42, 0x0b, 0xd2, 0xb4,
	0x9f, 0xf6, 0xe3, 0x81, 0x10, 0x83, 0x02, 0xcf, 0x6c, 0xd4, 0xad, 0xfa, 0x67, 0xac, 0x92, 0x53,
	0xb4, 0xf6, 0x6e, 0x36, 0x40, 0xae, 0xcb, 0xee, 0x99, 0xfd, 0xd6, 0x60, 0xf2, 0x23, 0x80, 0x9d,
	0x0f, 0x37, 0xd8, 0x1b, 0x5e, 0x09, 0xa9, 0x15, 0xc5, 0xbb, 0x0a, 0x95, 0x26, 0x4f, 0x21, 0x2c,
	0x72, 0xa5, 0x91, 0xc7, 0xc1, 0xd1, 0xea, 0xc9, 0xfa, 0xf9, 0x7a, 0xcd, 0x4e, 0x2f, 0x18, 0x93,
	0xd4, 0xa5, 0xc8, 0x13, 0x68, 0x94, 0x39, 0x1f, 0xc4, 0x2b, 0x8b, 0x14, 0x9b, 0x20, 0xaf, 0xa1,
	0xe5, 0x25, 0xc4, 0xab, 0x47, 0xc1, 0xc9, 0xfa, 0xf9, 0x7e, 0x5a, 0x6b, 0x4c, 0xbd, 0xc6, 0xf4,
	0xd2, 0x11, 0xe8, 0x98, 0x9a, 0x7c, 0x03, 0x32, 0xad, 0x48, 0x95, 0x82, 0x2b, 0x24, 0xa7, 0x73,
	0x92, 0x76, 0xdd, 0x7e, 0x5f, 0x50, 0x8e, 0x50, 0x52, 0x54, 0x55, 0xa1, 0xc7, 0xd2, 0x9e, 0xcf,
	0x48, 0x5b, 0x4a, 0xb5, 0x84, 0xe4, 0x67, 0x00, 0x0f, 0xec, 0x66, 0xef, 0x33, 0xce, 0xee, 0x73,
	0xa6, 0x6f, 0x96, 0x59, 0x10, 0xfc, 0x6b, 0x0b, 0xde, 0xc0, 0xde, 0xbc, 0x2a, 0x67, 0xc3, 0x01,
	0x44, 0x5d, 0x0f, 0x5a, 0x65, 0x0d, 0x3a, 0x01, 0x92, 0xaf, 0xb0, 0x31, 0x7d, 0x48, 0x42, 0xa0,
	0xd1, 0x13, 0x0c, 0x2d, 0xb1, 0x49, 0xed, 0x9a, 0xfc, 0x0f, 0x4d, 0x94, 0x52, 0xc8, 0x78, 0xe5,
	0x28, 0x38, 0x89, 0x68, 0x1d, 0x98, 0xe3, 0x2a, 0x5b, 0xe9, 0x64, 0xce, 0x1e, 0xb7, 0x4e, 0x25,
	0xaf, 0xa0, 0x61, 0x62, 0x12, 0xc3, 0x1a, 0x47, 0x7d, 0x2f, 0xe4, 0xd0, 0x76, 0x8e, 0xa8, 0x0f,
	0xcd, 0x86, 0x19, 0x63, 0xbe, 0xb7, 0x5d, 0x27, 0xbf, 0x02, 0xf8, 0xef, 0xba, 0xbe, 0xb3, 0xe8,
	0xdd, 0x6d, 0x43, 0xeb, 0x36, 0xe3, 0x79, 0x1f, 0x95, 0xb6, 0x2d, 0x36, 0xe8, 0x38, 0x36, 0xdd,
	0x4b, 0x29, 0xfa, 0x79, 0x81, 0xae, 0x8d, 0x0f, 0xc9, 0x29, 0xec, 0xf4, 0xab, 0xa2, 0xe8, 0x48,
	0xbc, 0xab, 0x72, 0x89, 0xb7, 0xc8, 0xb5, 0xb2, 0x7a, 0x5b, 0x74, 0xdb, 0x24, 0xe8, 0x14, 0x4e,
	0x5e, 0xc0, 0x9a, 0x28, 0x8d, 0x9b, 0x2a, 0x6e, 0xd8, 0x23, 0xed, 0xb9, 0x23, 0x79, 0x2d, 0x9f,
	0xeb, 0x2c, 0xf5, 0x34, 0x72, 0x0c, 0x21, 0x13, 0xbd, 0x21, 0xca, 0xb8, 0x69, 0x0b, 0x36, 0x5d,
	0xc1, 0xa5, 0x05, 0xa9, 0x4b, 0x26, 0x6f, 0x61, 0x7b, 0x72, 0x1c, 0x37, 0x96, 0x67, 0x10, 0xf6,
	0xb3, 0xbc, 0x40, 0xe6, 0x6e, 0xe7, 0x56, 0xea, 0x1e, 0x5b, 0x7a, 0x25, 0x45, 0x17, 0xa9, 0xcb,
	0x26, 0x37, 0x13, 0x2b, 0xdc, 0xf6, 0xe4, 0x10, 0x60, 0xf4, 0xbd, 0xc8, 0x78, 0xa7, 0x14, 0x52,
	0xbb, 0x49, 0x45, 0x16, 0x31, 0x0f, 0x80, 0x3c, 0x82, 0x88, 0x71, 0xd5, 0x31, 0x4e, 0x2a, 0x7b,
	0xcf, 0x22, 0xda, 0x62, 0x5c, 0x99, 0x39, 0x28, 0xb2, 0x0f, 0x66, 0x5d, 0x57, 0xae, 0xda, 0xca,
	0x35, 0xc6, 0x95, 0xa9, 0x4b, 0x3e, 0x42, 0x58, 0xeb, 0x26, 0xc7, 0xb0, 0xa5, 0xb4, 0x90, 0xd9,
	0x00, 0x3b, 0x4c, 0xe6, 0x66, 0xc4, 0xf5, 0xd0, 0x36, 0x1d, 0x7a, 0x69, 0x41, 0xb2, 0x07, 0x21,
	0xc3, 0x51, 0xde, 0xf3, 0xae, 0xbb, 0xe8, 0xfc, 0x77, 0x00, 0x70, 0x3d, 0xfe, 0xe5, 0x90, 0x0b,
	0x80, 0xc9, 0xeb, 0x24, 0xb1, 0xb3, 0x68, 0xe1, 0x17, 0xd2, 0xde, 0x5f, 0x92, 0x71, 0x66, 0x7d,
	0x82, 0xad, 0xd9, 0xdb, 0x4d, 0x0e, 0xa6, 0xc9, 0xf3, 0x4f, 0xb1, 0x7d, 0xf8, 0x97, 0xac, 0x6b,
	0xf7, 0x0e, 0x5a, 0xde, 0x53, 0x32, 0x3f, 0x63, 0xdf, 0xe2, 0xe1, 0x02, 0x5e, 0x17, 0x77, 0x43,
	0x8b, 0xbf, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x18, 0xa2, 0xb5, 0x56, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ValidationClient is the client API for Validation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ValidationClient interface {
	// CheckPorts executes a ports network test
	CheckPorts(ctx context.Context, in *CheckPortsRequest, opts ...grpc.CallOption) (*CheckPortsResponse, error)
	// CheckBandwidth executes a bandwidth network test
	CheckBandwidth(ctx context.Context, in *CheckBandwidthRequest, opts ...grpc.CallOption) (*CheckBandwidthResponse, error)
	// Validate validatest this node against the requirements
	// from a manifest.
	Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error)
}

type validationClient struct {
	cc *grpc.ClientConn
}

func NewValidationClient(cc *grpc.ClientConn) ValidationClient {
	return &validationClient{cc}
}

func (c *validationClient) CheckPorts(ctx context.Context, in *CheckPortsRequest, opts ...grpc.CallOption) (*CheckPortsResponse, error) {
	out := new(CheckPortsResponse)
	err := c.cc.Invoke(ctx, "/proto.Validation/CheckPorts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *validationClient) CheckBandwidth(ctx context.Context, in *CheckBandwidthRequest, opts ...grpc.CallOption) (*CheckBandwidthResponse, error) {
	out := new(CheckBandwidthResponse)
	err := c.cc.Invoke(ctx, "/proto.Validation/CheckBandwidth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *validationClient) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error) {
	out := new(ValidateResponse)
	err := c.cc.Invoke(ctx, "/proto.Validation/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ValidationServer is the server API for Validation service.
type ValidationServer interface {
	// CheckPorts executes a ports network test
	CheckPorts(context.Context, *CheckPortsRequest) (*CheckPortsResponse, error)
	// CheckBandwidth executes a bandwidth network test
	CheckBandwidth(context.Context, *CheckBandwidthRequest) (*CheckBandwidthResponse, error)
	// Validate validatest this node against the requirements
	// from a manifest.
	Validate(context.Context, *ValidateRequest) (*ValidateResponse, error)
}

// UnimplementedValidationServer can be embedded to have forward compatible implementations.
type UnimplementedValidationServer struct {
}

func (*UnimplementedValidationServer) CheckPorts(ctx context.Context, req *CheckPortsRequest) (*CheckPortsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPorts not implemented")
}
func (*UnimplementedValidationServer) CheckBandwidth(ctx context.Context, req *CheckBandwidthRequest) (*CheckBandwidthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckBandwidth not implemented")
}
func (*UnimplementedValidationServer) Validate(ctx context.Context, req *ValidateRequest) (*ValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}

func RegisterValidationServer(s *grpc.Server, srv ValidationServer) {
	s.RegisterService(&_Validation_serviceDesc, srv)
}

func _Validation_CheckPorts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPortsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidationServer).CheckPorts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Validation/CheckPorts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidationServer).CheckPorts(ctx, req.(*CheckPortsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Validation_CheckBandwidth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckBandwidthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidationServer).CheckBandwidth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Validation/CheckBandwidth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidationServer).CheckBandwidth(ctx, req.(*CheckBandwidthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Validation_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidationServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Validation/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidationServer).Validate(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Validation_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Validation",
	HandlerType: (*ValidationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckPorts",
			Handler:    _Validation_CheckPorts_Handler,
		},
		{
			MethodName: "CheckBandwidth",
			Handler:    _Validation_CheckBandwidth_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _Validation_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "validation.proto",
}
