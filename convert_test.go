package goTools

import "testing"

//go test -v -run TestToString .
func TestToString(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "uint64-test", args: struct{ value interface{} }{value: 11111111}, want: "11111111"},
		{name: "float64-test", args: struct{ value interface{} }{value: 12.232323}, want: "12.232323"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToString(tt.args.value); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// go test -run=^$  -benchmem -bench ^BenchmarkGoPoolWithSpinLock$ .
/*
 * benchtime 表示时间或运行次数，比如 -benchtime=10s 表示基准测试运行 10 秒，-benchtime=100x 表示基准测试运行 100 次
 * benchmem 统计内存分配情况
 * cpuprofile CPU 性能剖析 -cpuprofile=cpu.out
 * memprofile=$FILE 内存 性能剖析 -memprofile=mem.out
 * blockprofile=$FILE 阻塞 性能剖析 blockprofile=block.out
 *
 */
func Benchmark_ToString(b *testing.B) {
	for i := 0; i < b.N; i++ { // b.N 表示测试用例运行的次数
		ToString(123.2238)
	}
}
