package auth

import (
	"testing" // テストで使える関数・構造体が用意されているパッケージをimport
)

func TestGetGoogleToken(t *testing.T) {
	ok, err := GetGoogleToken("path-to-config-file")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if !ok {
		t.Fatal("failed test")
	}
}
