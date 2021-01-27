package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func handle(ctx context.Context) error {
	return nil
}

func main() {
	lambda.Start(handle)
}
