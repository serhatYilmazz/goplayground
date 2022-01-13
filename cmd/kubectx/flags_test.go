package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_parseArgs(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want Op
	}{
		{
			name: "nil args",
			args: nil,
			want: ListOp{},
		},
		{
			name: "empty args",
			args: []string{},
			want: ListOp{},
		}, {
			name: "help shorthand",
			args: []string{"-h"},
			want: HelpOp{},
		},
		{
			name: "help long form",
			args: []string{"--help"},
			want: HelpOp{},
		},
		{
			name: "switch by name",
			args: []string{"foo"},
			want: SwitchOp{Target: "foo"},
		},
		{
			name: "switch context by swapping prev/cur ('-')",
			args: []string{"-"},
			want: SwitchOp{Target: "-"},
		},
		{
			name: "unsupported (unknown option)",
			args: []string{"-x"},
			want: UnknownOp{Args: []string{"-x"}},
		},
		{
			name: "unsupported (unknown option) with multiple parameter",
			args: []string{"-x", "-d"},
			want: UnknownOp{Args: []string{"-x", "-d"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseArgs(tt.args)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("parseArgs(%#v), diff: %s", tt.args, diff)
			}
		})
	}

	// TODO add more unsupported cases

	// TODO consider these cases
	// - kubectx foo --help
	// - kubectx -h --help
	// - kubectx -d foo --h
}
