package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	tensor_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_go_proto"
	tensor_shape_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_shape_go_proto"
	types_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/types_go_proto"

	pb "github.com/howge/tfs28/apis"

	"google.golang.org/grpc"
)

func main() {
	servingAddress := flag.String("serving-address", "localhost:10000", "The tensorflow serving address")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("Usage: " + os.Args[0] + " --serving-address localhost:10000 path/to/img.png")
		os.Exit(1)
	}

	imgPath, err := filepath.Abs(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}

	imageBytes, err := ioutil.ReadFile(imgPath)
	if err != nil {
		log.Fatalln(err)
	}

	request := &pb.PredictRequest{
		ModelSpec: &pb.ModelSpec{
			Name:          "inception",
			SignatureName: "predict_images",
			// VersionChoice: &google_protobuf.Int64Value{
			// 	Value: int64(1),
			// },
		},
		Inputs: map[string]*tensor_go_proto.TensorProto{
			"images": &tensor_go_proto.TensorProto{
				Dtype: types_go_proto.DataType_DT_STRING,
				TensorShape: &tensor_shape_go_proto.TensorShapeProto{
					Dim: []*tensor_shape_go_proto.TensorShapeProto_Dim{
						&tensor_shape_go_proto.TensorShapeProto_Dim{
							Size: int64(1),
						},
					},
				},
				StringVal: [][]byte{imageBytes},
			},
		},
	}

	conn, err := grpc.Dial(*servingAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot connect to the grpc server: %v\n", err)
	}
	defer conn.Close()

	client := pb.NewPredictionServiceClient(conn)

	resp, err := client.Predict(context.Background(), request)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp)
}
