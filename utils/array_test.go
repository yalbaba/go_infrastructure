package utils

import "testing"

func TestInterSection(t *testing.T) {
	type args struct {
		n1 []int
		n2 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{n1: []int{}, n2: []int{}},
			want: false,
		},
		{
			name: "2",
			args: args{n1: []int{1, 2, 34, 45}, n2: []int{}},
			want: false,
		},
		{
			name: "3",
			args: args{n1: []int{1, 2, 34, 45}, n2: []int{34}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterSection(tt.args.n1, tt.args.n2); got != tt.want {
				t.Errorf("InterSection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringArray_Contains(t *testing.T) {
	type args struct {
		item string
	}
	tests := []struct {
		name string
		s    StringArray
		args args
		want bool
	}{
		{
			name: "1",
			s:    StringArray([]string{"tom", "marry", "jack", "kit", "wowo"}),
			args: args{item: "jack"},
			want: true,
		},
		{
			name: "2",
			s:    StringArray([]string{"tom", "marry", "jack", "kit", "wowo"}),
			args: args{item: "jack2"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Contains(tt.args.item); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrPage(t *testing.T) {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	t.Log(a[0:0])
	type args struct {
		pageNumber int
		pageSize   int
		totalCount int
	}
	tests := []struct {
		name           string
		args           args
		wantStartIndex int
		wantEndIndex   int
	}{
		{
			name: "1",
			args: args{
				pageNumber: 1,
				pageSize:   10,
				totalCount: 9,
			},
			wantStartIndex: 0,
			wantEndIndex:   9,
		},
		{
			name: "2",
			args: args{
				pageNumber: 2,
				pageSize:   10,
				totalCount: 9,
			},
			wantStartIndex: 0,
			wantEndIndex:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStartIndex, gotEndIndex := ArrPage(tt.args.pageNumber, tt.args.pageSize, tt.args.totalCount)
			if gotStartIndex != tt.wantStartIndex {
				t.Errorf("ArrPage() gotStartIndex = %v, want %v", gotStartIndex, tt.wantStartIndex)
			}
			if gotEndIndex != tt.wantEndIndex {
				t.Errorf("ArrPage() gotEndIndex = %v, want %v", gotEndIndex, tt.wantEndIndex)
			}
		})
	}
}
