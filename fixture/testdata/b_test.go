package testdata

import "testing"

// テスト関数が定義されているが testdata ディレクトリ以下は無視されるので実行されないのを確認する
func TestB(t *testing.T) {
	t.Log("fixture/testdata/b_test.go")
}
