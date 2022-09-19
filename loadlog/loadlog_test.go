package main

import (
	"testing"

	"github.com/golang/mock/gomock"

	mock "github.com/maishiro/golang/loadlog/mocks"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fileMock := mock.NewMockFileLoadSave(ctrl)
	strInput := "[2020/04/26 19:44:41:123] func1 Start\r\n[2020/04/27 14:36:30:456] func1 End\r\n"
	fileMock.EXPECT().LoadFile("example.log").Return(strInput, nil)
	fileMock.EXPECT().SaveFileCSV(gomock.Any(), "output.csv")

	service := userService{stream: fileMock}
	service.Run("example.log", "output.csv")
}
