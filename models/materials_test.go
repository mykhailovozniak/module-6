package models

import (
	"context"
	"testing"
)

type CursorMock struct {

}

func (CursorMock) Next(ctx context.Context) bool {
	return false
}

func (CursorMock) Err() (err error) {
	return nil
}

func (CursorMock) Decode(val interface{}) (err error) {
	return nil
}

func TestMapMaterials(t *testing.T) {
	mock := CursorMock{}
	want := MapMaterials(mock)

	if want != nil {
		t.Errorf("want should be nil")
	}
}
