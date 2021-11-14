package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGeoDistances(t *testing.T) {
	type args struct {
		lon1 float64
		lat1 float64
		lon2 float64
		lat2 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				lon1: 5.494532,
				lat1: 61.962793,
				lon2: 104.10194,
				lat2: 30.65984,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GeoDistances(tt.args.lon1, tt.args.lat1, tt.args.lon2, tt.args.lat2); got != tt.want {
				t.Errorf("GeoDistances() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGaodeReGeo(t *testing.T) {
	result, err := new(geoUtil).GaodeReGeo(116, 29)
	require.NoError(t, err)
	t.Log("成功", result)
}

func TestGaodeIPPosition(t *testing.T) {
	result, err := new(geoUtil).GaodeIPPosition("114.247.50.22")
	require.NoError(t, err)
	t.Log("成功", result)
}
