package main

import (
	"testing"
)

var err = "Error: i不可小於0"

func TestSum(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr string
	}{
		{"總和運算成功", args{100}, 5050, ""},
		{"遇到負數", args{-100}, -1, "Error: i不可小於0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Sum(tt.args.i)
			if (err != nil) && (err.Error() != tt.wantErr) {
				t.Errorf("Sum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr string
	}{

		{"相加運算成功", args{100, 100}, 200, ""},
		{"遇到負數", args{100, -1}, -1, "Error: j不可小於0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Add(tt.args.i, tt.args.j)
			if (err != nil) && (err.Error() != tt.wantErr) {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
