#! /bin/bash

# Store the script directory in a variable
script_dir=$(dirname $0)

# Go to parent directory from current directory (which is project root directory)
cd $script_dir/..

go build -o cmd/cvm-go/cvm-go