package gslice

import (
	"reflect"
	"sort"
	"testing"
)

func TestInt64Intersection(t *testing.T) {
	type args struct {
		src1 []int64
		src2 []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		// TODO: Add test cases.
		{
			name: "case-1",
			args: args{
				src1: []int64{1, 2, 3, 4},
				src2: []int64{3, 4, 5, 6},
			},
			want: []int64{3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersection(tt.args.src1, tt.args.src2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64IsIntersection(t *testing.T) {
	type args struct {
		src    []int64
		target []int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "case-1",
			args: args{
				src:    []int64{3, 4},
				target: []int64{1, 2, 3, 5},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIntersection(tt.args.src, tt.args.target); got != tt.want {
				t.Errorf("IsIntersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64SliceToSet(t *testing.T) {
	type args struct {
		list [][]int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		// TODO: Add test cases.
		{
			name: "case-1",
			args: args{
				list: [][]int64{
					{1, 2, 3, 4},
					{3, 4, 5, 6},
					{4, 5, 6, 7},
				},
			},
			want: []int64{1, 2, 3, 4, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SliceToSet(tt.args.list...)
			sort.Slice(got, func(i, j int) bool { return got[i] > got[j] })
			sort.Slice(tt.want, func(i, j int) bool { return tt.want[i] > tt.want[j] })
			if len(got) != len(tt.want) {
				t.Errorf("SliceToSet.len = %v, want %v", len(got), len(tt.want))
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsIntersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
