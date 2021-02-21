package facade

import (
	"testing"
)

var except = "A module running\nB module running"

func Test_FacadeApi(t *testing.T) {
	api := NewApi()
	ret := api.Test()
	if ret != except {
		t.Fatalf("except %s, return %s", except, ret)
	}
}
