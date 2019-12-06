package example

import (
	"os"
	"testing"
)

//テスト実行前の処理を記述する。
func TestMain(m *testing.M) {
	// fmt.Println("before test --------------------------")
	//テストコードの実行
	code := m.Run()
	// fmt.Println("after test --------------------------")
	os.Exit(code)
}

func TestEcho(t *testing.T) {
	expect := "hoge"
	s := Echo(expect)
	if s != expect {
		t.Errorf("Unexpected Value %s. We Expected %s", s, expect)
	}
}
