package controllers

import (
	"reflect"
	"testing"

	"github.com/globalsign/mgo"
)

func TestNewIceCreamController(t *testing.T) {
	type args struct {
		session *mgo.Session
	}
	tests := []struct {
		name string
		args args
		want *IceCreamController
	}{
		{
			"should not match",
			args{
				&mgo.Session{},
			},
			&IceCreamController{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIceCreamController(tt.args.session); reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIceCreamController() = %v, don't want %v", got, tt.want)
			}
		})
	}
}
