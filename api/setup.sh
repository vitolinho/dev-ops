#!/bin/bash

echo "Installing golangci-lint..."
cd api
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.63.4

echo "Done! You can now run 'golangci-lint run'"
