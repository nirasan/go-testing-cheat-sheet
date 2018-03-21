package testmain

import (
	"testing"
	"os"
	"log"
)

func TestMain(m *testing.M) {
	log.Print("setup")
	code := m.Run()
	log.Print("tear-down")
	os.Exit(code)
}

/* 実行結果

$ go test -v

2018/03/21 16:41:48 setup
=== RUN   TestA
--- PASS: TestA (0.00s)
        a_test.go:6: TestA
=== RUN   TestB
--- PASS: TestB (0.00s)
        b_test.go:6: TestB
PASS
2018/03/21 16:41:48 tear-down
 */
