package types

import (
	"testing"
)

func TestNewProductFromRequest(t *testing.T) {
	wrongProduct := ReqProduct{
		Name:        "test",
		Description: "",
	}
	_, err := NewProductFromRequest(wrongProduct)
	if err == nil {
		t.Error("expected error, got nothing")
	}
	correctProduct := ReqProduct{
		Name:        "test",
		Description: "This is a test product",
	}
	_, err = NewProductFromRequest(correctProduct)
	if err != nil {
		t.Errorf("expected not error, got %c", err)
	}
}
