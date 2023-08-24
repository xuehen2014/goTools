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
