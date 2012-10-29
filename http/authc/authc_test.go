package authc

import (
	"net/http"
	"testing"
)

func TestXYZ(t *testing.T) {
	s := DefaultSigner()
	r := http.Request{}
	s.Sign(&r, "foobar")
}
