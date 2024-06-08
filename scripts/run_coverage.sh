#! /bin/bash

# Store the script directory in a variable
script_dir=$(dirname $0)

output=coverage.out
report=coverage.html

# Go to parent directory from current directory (which is project root directory)
cd $script_dir/..

go test -v -coverprofile $output ./...
go tool cover -html $output -o $report
open $report