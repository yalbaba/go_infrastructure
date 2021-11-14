package utils

import (
	"reflect"
	"testing"
)

func TestGetStreamObject(t *testing.T) {
	type args struct {
		streamName   string
		streamSource PlatType
	}
	tests := []struct {
		name    string
		args    args
		want    *LiveRes
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				streamName:   "tv_dongfang",
				streamSource: TencentStream,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetStreamObject(tt.args.streamName, tt.args.streamSource)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStreamObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStreamObject() got = %v, want %v", got, tt.want)
			}
		})
	}
}
