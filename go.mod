module github.com/howge/tfs28

go 1.20

replace (
	github.com/tensorflow/tensorflow/tensorflow/go/core/example/example_protos_go_proto => ./vendor/github.com/tensorflow/tensorflow/tensorflow/go/core/example/example_protos_go_proto
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/allocation_description_go_proto => ./vendor/github.com/tensorflow/tensorflow/tensorflow/go/core/framework/allocation_description_go_proto
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/api_def_go_proto => ./vendor/github.com/tensorflow/tensorflow/tensorflow/go/core/framework/api_def_go_proto
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/attr_value_go_proto => ./vendor/github.com/tensorflow/tensorflow/tensorflow/go/core/framework/attr_value_go_proto
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/cost_graph_go_proto => ./vendor/github.com/tensorflow/tensorflow/tensorflow/go/core/framework/cost_graph_go_proto
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/device_attributes_go_proto => ./vendor/github.com/tensorflow/tensorflow/tensorflow/go/core/framework/device_attributes_go_proto
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/full_type_go_proto => ./vendor/github.com/tensorflow/tensorflow/tensorflow/go/core/framework/full_type_go_proto
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/function_go_proto => ./vendor/github.com/tensorflow/tensorflow/tensorflow/go/core/framework/function_go_proto
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/graph_go_proto => ./vendor/github.com/tensorflow/tensorflow/tensorflow/go/core/framework/graph_go_proto
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/model_go_proto => ./vendor/github.com/tensorflow/tensorflow/tensorflow/go/core/framework/model_go_proto
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/node_def_go_proto => ./vendor/github.com/tensorflow/tensorflow/tensorflow/go/core/framework/node_def_go_proto
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/op_def_go_proto => ./vendor/github.com/tensorflow/tensorflow/tensorflow/go/core/framework/op_def_go_proto
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/resource_handle_go_proto => ./vendor/github.com/tensorflow/tensorflow/tensorflow/go/core/framework/resource_handle_go_proto
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/step_stats_go_proto => ./vendor/github.com/tensorflow/tensorflow/tensorflow/go/core/framework/step_stats_go_proto
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_description_go_proto => ./vendor/github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_description_go_proto
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_go_proto => ./vendor/github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_go_proto
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_shape_go_proto => ./vendor/github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_shape_go_proto
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_slice_go_proto => ./vendor/github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_slice_go_proto
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/types_go_proto => ./vendor/github.com/tensorflow/tensorflow/tensorflow/go/core/framework/types_go_proto
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/variable_go_proto => ./vendor/github.com/tensorflow/tensorflow/tensorflow/go/core/framework/variable_go_proto
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/versions_go_proto => ./vendor/github.com/tensorflow/tensorflow/tensorflow/go/core/framework/versions_go_proto
	github.com/tensorflow/tensorflow/tensorflow/go/core/protobuf/for_core_protos_go_proto => ./vendor/github.com/tensorflow/tensorflow/tensorflow/go/core/protobuf/for_core_protos_go_proto
	github.com/tensorflow/tensorflow/tensorflow/go/stream_executor => ./vendor/github.com/tensorflow/tensorflow/tensorflow/go/stream_executor
)

require (
	github.com/tensorflow/tensorflow/tensorflow/go/core/example/example_protos_go_proto v0.0.0-00010101000000-000000000000
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_go_proto v0.0.0-00010101000000-000000000000
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_shape_go_proto v0.0.0-00010101000000-000000000000
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/types_go_proto v0.0.0-00010101000000-000000000000
	github.com/tensorflow/tensorflow/tensorflow/go/core/protobuf/for_core_protos_go_proto v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.62.0
	google.golang.org/protobuf v1.32.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/allocation_description_go_proto v0.0.0-00010101000000-000000000000 // indirect
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/attr_value_go_proto v0.0.0-00010101000000-000000000000 // indirect
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/cost_graph_go_proto v0.0.0-00010101000000-000000000000 // indirect
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/device_attributes_go_proto v0.0.0-00010101000000-000000000000 // indirect
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/full_type_go_proto v0.0.0-00010101000000-000000000000 // indirect
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/function_go_proto v0.0.0-00010101000000-000000000000 // indirect
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/graph_go_proto v0.0.0-00010101000000-000000000000 // indirect
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/node_def_go_proto v0.0.0-00010101000000-000000000000 // indirect
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/op_def_go_proto v0.0.0-00010101000000-000000000000 // indirect
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/resource_handle_go_proto v0.0.0-00010101000000-000000000000 // indirect
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/step_stats_go_proto v0.0.0-00010101000000-000000000000 // indirect
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_description_go_proto v0.0.0-00010101000000-000000000000 // indirect
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_slice_go_proto v0.0.0-00010101000000-000000000000 // indirect
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/variable_go_proto v0.0.0-00010101000000-000000000000 // indirect
	github.com/tensorflow/tensorflow/tensorflow/go/core/framework/versions_go_proto v0.0.0-00010101000000-000000000000 // indirect
	github.com/tensorflow/tensorflow/tensorflow/go/stream_executor v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240123012728-ef4313101c80 // indirect
)
