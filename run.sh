#!/bin/bash

set -e
set -u
set -o pipefail

go build -o ./cmd/build/bcc-canteen ./cmd/api
./cmd/build/bcc-canteen
