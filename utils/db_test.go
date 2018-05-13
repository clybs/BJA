package utils

import (
	"reflect"
	"testing"

	"github.com/globalsign/mgo"
)

func TestNewDB(t *testing.T) {
	tests := []struct {
		name string
		want *DB
	}{
		{
			"should match",
			&DB{
				"icecreams",
				"bja",
				"users",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDB(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_GetSession(t *testing.T) {
	tests := []struct {
		name string
		db   *DB
		want *mgo.Session
	}{
		{
			"should not match",
			&DB{
				"icecreams",
				"bja",
				"users",
			},
			&mgo.Session{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.GetSession(); reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.GetSession() = %v, don't want %v", got, tt.want)
			}
		})
	}
}
