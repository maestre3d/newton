#!/bin/bash

function build_lambdas() {
  for dir in ../../cmd/library-api/lambda/*/
  do
    dir=${dir%*/}      # remove the trailing "/"
    echo "Building: $dir"
    build_lambda_function "$dir"
  done
  echo "Building done"
}

function build_lambda_function() {
  GOOS=linux GOARCH=amd64 go release -o "./bin/" "$1"
}

function compress_lambdas() {
  for file in ./bin/*
  do
    file=${file%*/}      # remove the trailing "/"
    file=$(echo "$file" | cut -d'/' -f3-)
    echo "Compressing: $file"
    compress_compiled_lambda "$file"
  done
  echo "Compression done"
}

function compress_compiled_lambda() {
  zip "./release/$1" "./bin/$1"
}

function upload_lambdas() {
  for file in ./release/*
  do
    file=${file%*/}      # remove the trailing "/"
    file=$(echo "$file" | cut -d'/' -f3-)
    echo "Uploading to S3: $file"
    upload_lambda_to_s3 "$file" "$1"
  done
  echo "S3 uploading done"
}

function upload_lambda_to_s3() {
  stage="dev"
  if [ "$2" != "" ]; then
    stage="$2"
  fi
  aws s3 cp "./release/$1" s3://newton-serverless/"$stage"/"$1"
}

build_lambdas
compress_lambdas
upload_lambdas "$1"