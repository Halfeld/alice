package googleApis

import (
	"context"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/alice/helpers"
	pb "google.golang.org/genproto/googleapis/cloud/vision/v1"
)

// DetectFaces load faces from images
func DetectFaces(filePath string) ([]*pb.FaceAnnotation, error) {
	ctx := context.Background()
	client, err := vision.NewImageAnnotatorClient(ctx)
	helpers.ErrorHandler(err, "Failed to create client")

	file, err := os.Open(filePath)
	helpers.ErrorHandler(err, "Failed to read file")
	defer file.Close()

	image, err := vision.NewImageFromReader(file)
	helpers.ErrorHandler(err, "Failed load image")

	faces, err := client.DetectFaces(ctx, image, nil, 10)
	helpers.ErrorHandler(err, "Failed to detect faces on image")

	return faces, err
}
