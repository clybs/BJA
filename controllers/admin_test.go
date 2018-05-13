package controllers

import (
	"reflect"
	"testing"

	"github.com/globalsign/mgo"
	"github.com/gocraft/web"
)

func TestNewAdminController(t *testing.T) {
	type args struct {
		session *mgo.Session
	}
	tests := []struct {
		name string
		args args
		want *AdminController
	}{
		{
			"should not match",
			args{
				&mgo.Session{},
			},
			&AdminController{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAdminController(tt.args.session); reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAdminController() = %v, don't want %v", got, tt.want)
			}
		})
	}
}

func TestAdminController_Login(t *testing.T) {
	type args struct {
		rw  web.ResponseWriter
		req *web.Request
	}
	tests := []struct {
		name string
		ac   *AdminController
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ac.Login(tt.args.rw, tt.args.req)
		})
	}
}
