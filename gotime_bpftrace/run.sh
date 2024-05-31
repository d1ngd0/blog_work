#!/bin/bash
go build .
parallel --ungroup ::: "/home/paul/Projects/blog_work/gotime_bpftrace/gotime_bpftrace" "sudo ./trace.bt"
