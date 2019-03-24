package main

import (
	"database/sql"
	"reflect"
	"testing"
)

func Test_connectToDatabase(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			connectToDatabase()
		})
	}
}

func Test_connectToDB(t *testing.T) {
	tests := []struct {
		name string
		want *sql.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connectToDB(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connectToDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dbConnect(t *testing.T) {
	type args struct {
		connStr string
	}
	tests := []struct {
		name string
		args args
		want *sql.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dbConnect(tt.args.connStr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dbConnect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pingDB(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pingDB()
		})
	}
}
