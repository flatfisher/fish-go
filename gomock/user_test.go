package main

import (
	"testing"
	"github.com/golang/mock/gomock"
)

func TestUser_GetName(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_main.NewMockUser(ctrl)
}