package main

import (
	"bytes"
	"errors"
	"github.com/v-sadovsky/ok_test_task/cmd/testdata"
	"github.com/v-sadovsky/ok_test_task/pkg/models"
	"io"
	"log"
	"testing"

	"github.com/v-sadovsky/ok_test_task/pkg/models/mock"
)

func newTestApplication(t *testing.T) *application {
	return &application{
		errorLog: log.New(io.Discard, "", 0),
		infoLog:  log.New(io.Discard, "", 0),
		stores:   &mock.StoresModel{},
	}
}

func TestHandleData(t *testing.T) {
	// Create a new instance of application struct which uses the mocked dependencies
	app := newTestApplication(t)

	// Set up table-driven tests to check the responses sent by application for different times
	tests := []struct {
		name        string
		workersPool int
		currentTime string
		wantData    [][]byte
		wantCode    error
	}{
		{"No records", 2, mock.NonWorkingTime, [][]byte{nil}, models.ErrNoRecord},
		{"Shop1 records only", 2, mock.Shop1WorkingTimeOnly, [][]byte{testdata.Shop1}, nil},
		{"Shop2 records only", 2, mock.Shop2WorkingTimeOnly, [][]byte{testdata.Shop2}, nil},
		{"Both shops records", 2, mock.BothShopsWorkingTime, [][]byte{testdata.BothShopsV1, testdata.BothShopsV2}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, code := app.requestData(tt.workersPool, tt.currentTime)

			if !errors.Is(code, tt.wantCode) {
				t.Errorf("want %s; got %s", tt.wantCode.Error(), code.Error())
			}

			if len(tt.wantData) == 1 {
				if !bytes.Equal(data, tt.wantData[0]) {
					t.Errorf("want %s got %s", string(tt.wantData[0]), string(data))
				}
			} else {
				if !bytes.Equal(data, tt.wantData[0]) && !bytes.Equal(data, tt.wantData[1]) {
					t.Errorf("want %s got %s", string(tt.wantData[0]), string(data))
				}
			}
		})
	}
}
