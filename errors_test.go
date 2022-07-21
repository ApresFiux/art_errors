package art_errors

import (
	"reflect"
	"testing"
)

func TestFailLogHandler(t *testing.T) {
	type args struct {
		err       error
		failFirst bool
		errorList []string
		result    string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FailLogHandler(tt.args.err, tt.args.failFirst, tt.args.errorList, tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FailLogHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFailLogHandler1(t *testing.T) {
	type args struct {
		err       error
		failFirst bool
		errorList []string
		result    string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FailLogHandler(tt.args.err, tt.args.failFirst, tt.args.errorList, tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FailLogHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
