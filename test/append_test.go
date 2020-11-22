package test

import "testing"

func BenchmarkApp(b *testing.B){
	for i:=0;i<b.N;i++{
		Append()
	}
}

//go test -run=NONE -bench=BenchmarkApp -cpuprofile=cpu.log append_test.go append.go
//go tool pprof -text -nodecount=10 ./test.test cpu.log
/*
	File: test.test
	Type: cpu
	Time: Sep 28, 2020 at 6:03pm (CST)
	Duration: 1.84s, Total samples = 1.54s (83.49%)
	Showing nodes accounting for 1.54s, 100% of 1.54s total
	flat  flat%   sum%        cum   cum%
	0.77s 50.00% 50.00%      1.44s 93.51%  command-line-arguments.App
	0.46s 29.87% 79.87%      0.61s 39.61%  strconv.FormatInt
	0.13s  8.44% 88.31%      0.13s  8.44%  strconv.small
	0.08s  5.19% 93.51%      0.08s  5.19%  runtime.nanotime1
	0.06s  3.90% 97.40%      0.67s 43.51%  strconv.Itoa
	0.02s  1.30% 98.70%      0.02s  1.30%  runtime.newstack
	0.01s  0.65% 99.35%      1.45s 94.16%  command-line-arguments.BenchmarkApp
	0.01s  0.65%   100%      0.01s  0.65%  runtime.usleep
	0     0%   100%      0.01s  0.65%  runtime.findrunnable
	0     0%   100%      0.01s  0.65%  runtime.mcall
	0     0%   100%      0.08s  5.19%  runtime.mstart
	0     0%   100%      0.08s  5.19%  runtime.mstart1
	0     0%   100%      0.08s  5.19%  runtime.nanotime (inline)
	0     0%   100%      0.01s  0.65%  runtime.park_m
	0     0%   100%      0.01s  0.65%  runtime.schedule
	0     0%   100%      0.08s  5.19%  runtime.sysmon
	0     0%   100%      1.45s 94.16%  testing.(*B).launch
	0     0%   100%      1.45s 94.16%  testing.(*B).runN
 */
