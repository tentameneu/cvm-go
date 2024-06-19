#! /bin/bash

output=coverage.out
report=coverage.html

go test -v -coverprofile $output ./...
go tool cover -html $output -o $report
open $report