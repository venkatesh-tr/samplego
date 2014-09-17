package ping_test

import (
	"github.com/venkatesh-tr/samplego/ping"
	"testing"
)

func TestPing(t *testing.T) {
	str := ping.Ping("Hi There")
	if str != "Hi There" {
		t.Error("Test Failed : Returned string didn't match!")
	}
}
