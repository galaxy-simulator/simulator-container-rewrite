package main

import (
	"reflect"
	"testing"

	"git.darknebu.la/GalaxySimulator/structs"
)

func Test_processStars(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processStars()
		})
	}
}

func Test_printStarsProcessed(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printStarsProcessed()
		})
	}
}

func Test_getStarID(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStarID(); got != tt.want {
				t.Errorf("getStarID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getStar(t *testing.T) {
	type args struct {
		starID int64
	}
	tests := []struct {
		name string
		args args
		want structs.Star2D
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStar(tt.args.starID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getStar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTimestep(t *testing.T) {
	type args struct {
		starID int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTimestep(tt.args.starID); got != tt.want {
				t.Errorf("getTimestep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcNewPos(t *testing.T) {
	type args struct {
		star     structs.Star2D
		force    structs.Vec2
		timestep float64
	}
	tests := []struct {
		name string
		args args
		want structs.Star2D
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcNewPos(tt.args.star, tt.args.force, tt.args.timestep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcNewPos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertStar(t *testing.T) {
	type args struct {
		star     structs.Star2D
		timestep int64
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertStar(tt.args.star, tt.args.timestep)
		})
	}
}
