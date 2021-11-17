#!/bin/bash
# Configure the go application
go mod download
go mod tidy

# Set up kubernetes resources for the application
kubectl apply -f ./templates --recursive