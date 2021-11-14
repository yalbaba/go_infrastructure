package utils

import (
	"reflect"
	"testing"
	"time"
)

func TestGetBetweenDates(t *testing.T) {
	type args struct {
		startDate string
		endDate   string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{startDate: "2020-11-1", endDate: "2020-11-1"},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBetweenDates(tt.args.startDate, tt.args.endDate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBetweenDates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeekByDate(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{date: "2020-07-06"},
			want: 1,
		},
		{
			name: "2",
			args: args{date: "2020-07-05"},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WeekByDate(tt.args.date); got != tt.want {
				t.Errorf("WeekByDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTextToTime(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "1",
			args:    args{text: "2020-07-24"},
			want:    time.Time{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TextToTime(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("TextToTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TextToTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTextToTimeOfLayout(t *testing.T) {
	type args struct {
		text   string
		layout string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "1",
			args:    args{text: "2020-07-24", layout: "2006-01-02"},
			want:    time.Time{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TextToTimeOfLayout(tt.args.text, tt.args.layout)
			if (err != nil) != tt.wantErr {
				t.Errorf("TextToTimeOfLayout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TextToTimeOfLayout() got = %v, want %v", got, tt.want)
			}
		})
	}
}
