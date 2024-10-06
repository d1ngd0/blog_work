#!/bin/bash
set -e

go build -o gotime_bpftrace main.go
go build -o parse_time parse_time.go
sudo cp gotime_bpftrace /usr/local/bin/
parallel --ungroup ::: "gotime_bpftrace" "sudo ./trace.bt | ./parse_time"
