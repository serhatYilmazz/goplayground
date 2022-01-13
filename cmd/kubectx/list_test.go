package main

import (
	"reflect"
	"testing"
)

func Test_parseKubeConfig(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    kubeconfig
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseKubeConfig(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseKubeConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseKubeConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
