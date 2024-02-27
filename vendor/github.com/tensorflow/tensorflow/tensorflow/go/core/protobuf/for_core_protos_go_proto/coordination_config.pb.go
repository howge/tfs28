// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.3
// source: tensorflow/core/protobuf/coordination_config.proto

package for_core_protos_go_proto

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

// Coordination service configuration parameters.
// The system picks appropriate values for fields that are not set.
type CoordinationServiceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type of coordination service implementation to enable.
	// For example, setting the service type as "standalone" starts a service
	// instance on the leader task to provide the coordination services such as
	// heartbeats and consistent key-value store.
	ServiceType string `protobuf:"bytes,1,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	// Address where the coordination service instance is hosted.
	ServiceLeader string `protobuf:"bytes,2,opt,name=service_leader,json=serviceLeader,proto3" json:"service_leader,omitempty"`
	// Whether to enable the health check mechanism.
	EnableHealthCheck bool `protobuf:"varint,3,opt,name=enable_health_check,json=enableHealthCheck,proto3" json:"enable_health_check,omitempty"`
	// Maximum wait time for all members in the cluster to be registered.
	ClusterRegisterTimeoutInMs int64 `protobuf:"varint,4,opt,name=cluster_register_timeout_in_ms,json=clusterRegisterTimeoutInMs,proto3" json:"cluster_register_timeout_in_ms,omitempty"`
	// Heartbeat timeout, if a worker does not record heartbeat in this time
	// window, it will be considered disconnected.
	HeartbeatTimeoutInMs int64 `protobuf:"varint,5,opt,name=heartbeat_timeout_in_ms,json=heartbeatTimeoutInMs,proto3" json:"heartbeat_timeout_in_ms,omitempty"`
	// The list of jobs that partipate in the coordination service. If empty, all
	// jobs will be included in the coordination service by default.
	CoordinatedJobs []string `protobuf:"bytes,6,rep,name=coordinated_jobs,json=coordinatedJobs,proto3" json:"coordinated_jobs,omitempty"`
}

func (x *CoordinationServiceConfig) Reset() {
	*x = CoordinationServiceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_coordination_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoordinationServiceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoordinationServiceConfig) ProtoMessage() {}

func (x *CoordinationServiceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_coordination_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoordinationServiceConfig.ProtoReflect.Descriptor instead.
func (*CoordinationServiceConfig) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_coordination_config_proto_rawDescGZIP(), []int{0}
}

func (x *CoordinationServiceConfig) GetServiceType() string {
	if x != nil {
		return x.ServiceType
	}
	return ""
}

func (x *CoordinationServiceConfig) GetServiceLeader() string {
	if x != nil {
		return x.ServiceLeader
	}
	return ""
}

func (x *CoordinationServiceConfig) GetEnableHealthCheck() bool {
	if x != nil {
		return x.EnableHealthCheck
	}
	return false
}

func (x *CoordinationServiceConfig) GetClusterRegisterTimeoutInMs() int64 {
	if x != nil {
		return x.ClusterRegisterTimeoutInMs
	}
	return 0
}

func (x *CoordinationServiceConfig) GetHeartbeatTimeoutInMs() int64 {
	if x != nil {
		return x.HeartbeatTimeoutInMs
	}
	return 0
}

func (x *CoordinationServiceConfig) GetCoordinatedJobs() []string {
	if x != nil {
		return x.CoordinatedJobs
	}
	return nil
}

var File_tensorflow_core_protobuf_coordination_config_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_coordination_config_proto_rawDesc = []byte{
	0x0a, 0x32, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6f, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x22, 0xbb, 0x02, 0x0a, 0x19, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x21,
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x13, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x42, 0x0a, 0x1e, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x5f, 0x69, 0x6e, 0x5f, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x1a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x49, 0x6e, 0x4d, 0x73, 0x12, 0x35, 0x0a, 0x17,
	0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x5f, 0x69, 0x6e, 0x5f, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x68,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x49,
	0x6e, 0x4d, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x63,
	0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x4a, 0x6f, 0x62, 0x73, 0x42, 0x57,
	0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66,
	0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5f, 0x67,
	0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_protobuf_coordination_config_proto_rawDescOnce sync.Once
	file_tensorflow_core_protobuf_coordination_config_proto_rawDescData = file_tensorflow_core_protobuf_coordination_config_proto_rawDesc
)

func file_tensorflow_core_protobuf_coordination_config_proto_rawDescGZIP() []byte {
	file_tensorflow_core_protobuf_coordination_config_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_protobuf_coordination_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_protobuf_coordination_config_proto_rawDescData)
	})
	return file_tensorflow_core_protobuf_coordination_config_proto_rawDescData
}

var file_tensorflow_core_protobuf_coordination_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_core_protobuf_coordination_config_proto_goTypes = []interface{}{
	(*CoordinationServiceConfig)(nil), // 0: tensorflow.CoordinationServiceConfig
}
var file_tensorflow_core_protobuf_coordination_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_coordination_config_proto_init() }
func file_tensorflow_core_protobuf_coordination_config_proto_init() {
	if File_tensorflow_core_protobuf_coordination_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_protobuf_coordination_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoordinationServiceConfig); i {
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
			RawDescriptor: file_tensorflow_core_protobuf_coordination_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_protobuf_coordination_config_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_coordination_config_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_protobuf_coordination_config_proto_msgTypes,
	}.Build()
	File_tensorflow_core_protobuf_coordination_config_proto = out.File
	file_tensorflow_core_protobuf_coordination_config_proto_rawDesc = nil
	file_tensorflow_core_protobuf_coordination_config_proto_goTypes = nil
	file_tensorflow_core_protobuf_coordination_config_proto_depIdxs = nil
}
