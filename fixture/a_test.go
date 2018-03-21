package fixture

import (
	"testing"
	"io/ioutil"
	"log"
	"encoding/json"
	"os"
)

// フィクスチャーを読み込むための型の定義
type Fixture struct {
	TestCases []testCase `json:"test_cases"`
}

type testCase struct {
	L int `json:"l"`
	R int `json:"r"`
	Result int `json:"result"`
}

// TestMain を使ってテスト実行前にフィクスチャーを読み込んでテスト環境を用意する
func TestMain(m *testing.M) {
	// フィクスチャーを読み込む
	b, err := ioutil.ReadFile("testdata/fixture.json")
	if err != nil {
		log.Fatal(err)
	}
	f := new(Fixture)
	if err := json.Unmarshal(b, f); err != nil {
		log.Fatal(err)
	}
	// ここではログ表示しているだけだが利用しているデータストアへデータの登録をするなどできる
	log.Printf("fixture: %#v", f)

	// テストの実行
	os.Exit(m.Run())
}

func TestA(t *testing.T) {
	t.Log("fixture/a_test.go")
}

/* 実行結果

$ go test -v ./...

2018/03/21 21:22:16 fixture: &fixture.Fixture{TestCases:[]fixture.testCase{fixture.testCase{L:1, R:1, Result:2}, fixture.testCase{L:0, R:0, Result:0}, fixture.testCase{L:0, R:-1, Result:-1}, fixture.testCase{L:1, R:-3, Result:-2}}}
=== RUN   TestA
--- PASS: TestA (0.00s)
        a_test.go:41: fixture/a_test.go

 */