package basic

import (
	"testing"
	"time"
)

// テスト用の関数の定義.
// テスト用の関数名は TestXxx という形式で Add 関数のテストなら TestAdd とする.
func TestAdd(t *testing.T) {
	var result int

	// テストケースの検証.
	// テストしたい関数の実行結果を if などで判定して想定した値であるかを検証する.
	result = Add(1, 2)
	if result != 3 {
		// テスト失敗時には t.Error などでエラーを表示する.
		t.Errorf("add failed. expect:%d, actual:%d", 3, result)
	}

	// テスト中のロギング.
	// t.Log, t.Logf でログを出すと `go test -v` と実行したときのみ表示される.
	t.Logf("result is %d", result)
}

/* 実行結果

=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
        adder_test.go:20: result is 3
*/

// エラー関数の比較.
// t.Error はテスト失敗としてログを出すが以後の処理も実行される.
func TestAdd2(t *testing.T) {
	t.Error("error")
	t.Log("log")
}

// t.Fatal はテスト失敗としてログを出し以後の処理は実行されない.
func TestAdd3(t *testing.T) {
	t.Fatal("fatal")
	t.Log("log")
}

// t.Fail はテスト失敗とするが以後の処理も実行される.
func TestAdd4(t *testing.T) {
	t.Fail()
	t.Log("log")
}

// t.FailNow はテスト失敗とし以後の処理は実行されない.
func TestAdd5(t *testing.T) {
	t.FailNow()
	t.Log("log")
}

/* 実行結果

=== RUN   TestAdd2
--- FAIL: TestAdd2 (0.00s)
        adder_test.go:24: error
        adder_test.go:25: log
=== RUN   TestAdd3
--- FAIL: TestAdd3 (0.00s)
        adder_test.go:29: fatal
=== RUN   TestAdd4
--- FAIL: TestAdd4 (0.00s)
        adder_test.go:35: log
=== RUN   TestAdd5
--- FAIL: TestAdd5 (0.00s)
*/

// レシーバを持つ関数のテスト用の関数名は TestReceiver_Func という形式.
// Adder 型のレシーバを持つ Add 関数のテストなので TestAdder_Add
func TestAdder_Add(t *testing.T) {
	a := &Adder{}

	// 一つの処理に対して複数のテストケースをまとめてテストしたい場合はテストケース用の型の配列を定義して実行するパターンがある.
	testCases := []struct {
		L      int
		R      int
		Result int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{0, -1, -1},
		{100, 200, 0}, // 失敗するケース. 失敗時の出力を確認する.
	}

	for _, testCase := range testCases {
		result := a.Add(testCase.L, testCase.R)
		if result != testCase.Result {
			t.Errorf("invalid result. testCase:%#v, actual:%d", testCase, result)
		}
	}
}

/* 実行結果

# go test -run REGEX で REGEX にマッチするテストだけ実行できる
$ go test -v -run Adder

=== RUN   TestAdder_Add
--- FAIL: TestAdder_Add (0.00s)
        adder_test.go:92: invalid result. testCase:struct { L int; R int; Result int }{L:100, R:200, Result:0}, actual:200
*/

// サブテストの実行.
// テスト関数内で t.Run を使うことでサブテストが実行できる.
// サブテストを使うことで容易にテスト前後の処理を定義できる.
func TestAdder_AddMulti(t *testing.T) {
	t.Log("setup")

	t.Run("Len=1", func(t *testing.T) {
		t.Log("Len=1")
		if new(Adder).AddMulti(1) != 1 {
			t.Fail()
		}
	})

	t.Run("Len=2", func(t *testing.T) {
		t.Log("Len=2")
		if new(Adder).AddMulti(1, 2) != 3 {
			t.Fail()
		}
	})

	t.Run("Len=3", func(t *testing.T) {
		t.Log("Len=3")
		if new(Adder).AddMulti(1, 2, 3) != 6 {
			t.Fail()
		}
	})

	t.Log("tear-down")
}

/* 実行結果

$ go test -v -run AddMulti

=== RUN   TestAdder_AddMulti
=== RUN   TestAdder_AddMulti/Len=1
=== RUN   TestAdder_AddMulti/Len=2
=== RUN   TestAdder_AddMulti/Len=3
--- PASS: TestAdder_AddMulti (0.00s)
        adder_test.go:112: setup
    --- PASS: TestAdder_AddMulti/Len=1 (0.00s)
        adder_test.go:115: Len=1
    --- PASS: TestAdder_AddMulti/Len=2 (0.00s)
        adder_test.go:122: Len=2
    --- PASS: TestAdder_AddMulti/Len=3 (0.00s)
        adder_test.go:129: Len=3
        adder_test.go:135: tear-down

# -run で "/" 以後にサブテスト名を指定すると任意のサブテストだけ実行できる
$ go test -v -run AddMulti/Len=3

=== RUN   TestAdder_AddMulti
=== RUN   TestAdder_AddMulti/Len=3
--- PASS: TestAdder_AddMulti (0.00s)
        adder_test.go:112: setup
    --- PASS: TestAdder_AddMulti/Len=3 (0.00s)
        adder_test.go:129: Len=3
        adder_test.go:135: tear-down

 */

// サブテストを並列実行.
// サブテスト内で t.Parallel を実行することでサブテストが並列実行される.
func TestAdder_AddMulti2(t *testing.T) {
	// テストの開始終了と各サブテストの終了で時間を表示して実行順を確認する
	t.Logf("setup: %s", time.Now())

	// 並列実行されるサブテストを t.Run でラップすることで全てのサブテストの終了を待つ.
	// こうすることで全てのサブテストの終了を待ってテスト終了処理を実行することができる.
	t.Run("group", func(t *testing.T) {
		t.Run("Len=1", func(t *testing.T) {
			// サブテストを並列実行する
			t.Parallel()
			// 並列実行されていることを確認するため sleep で終了タイミングをずらす
			time.Sleep(time.Second * 2)
			if new(Adder).AddMulti(1) != 1 {
				t.Fail()
			}
			t.Logf("Len=1: %s", time.Now())
		})

		t.Run("Len=2", func(t *testing.T) {
			t.Parallel()
			time.Sleep(time.Second * 3)
			if new(Adder).AddMulti(1, 2) != 3 {
				t.Fail()
			}
			t.Logf("Len=2: %s", time.Now())
		})

		t.Run("Len=3", func(t *testing.T) {
			t.Parallel()
			time.Sleep(time.Second * 1)
			if new(Adder).AddMulti(1, 2, 3) != 6 {
				t.Fail()
			}
			t.Logf("Len=3: %s", time.Now())
		})
	})

	t.Logf("tear-down: %s", time.Now())
}

/* 実行結果

$ go test -v -run AddMulti2

=== RUN   TestAdder_AddMulti2
=== RUN   TestAdder_AddMulti2/group
=== RUN   TestAdder_AddMulti2/group/Len=1
=== RUN   TestAdder_AddMulti2/group/Len=2
=== RUN   TestAdder_AddMulti2/group/Len=3
--- PASS: TestAdder_AddMulti2 (3.00s)
        adder_test.go:174: setup: 2018-03-21 09:56:04.979988051 +0900 JST
    --- PASS: TestAdder_AddMulti2/group (0.00s)
        --- PASS: TestAdder_AddMulti2/group/Len=3 (1.01s)
                adder_test.go:205: Len=3: 2018-03-21 09:56:05.9852616 +0900 JST
        --- PASS: TestAdder_AddMulti2/group/Len=1 (2.00s)
                adder_test.go:187: Len=1: 2018-03-21 09:56:06.981221633 +0900 JST
        --- PASS: TestAdder_AddMulti2/group/Len=2 (3.00s)
                adder_test.go:196: Len=2: 2018-03-21 09:56:07.983102269 +0900 JST
        adder_test.go:209: tear-down: 2018-03-21 09:56:07.983340171 +0900 JST
 */

// t.Helper を実行するとヘルパースクリプト内で実行した t.Log や t.Error の出力元がヘルパースクリプト内ではなく呼び出し元のものになる.
// ヘルパースクリプトが同一テスト内で複数回呼ばれていると t.Error などがどの回のヘルパースクリプトから出力されているものかわかりづらいので判断するのに便利.
func TestAdder_AddMulti3(t *testing.T) {
	t.Run("group", func(t *testing.T) {
		helperFunc(t, false)
		helperFunc(t, true) // helperFunc 内で t.Helper を実行するので Log や Error が発生した際にこの行から出ていることになりデバッグが楽になる.
	})
}
func helperFunc(t *testing.T, useHelper bool) {
	if useHelper {
		t.Helper()
	}
	t.Logf("use helper: %v", useHelper)
}

/*
=== RUN   TestAdder_AddMulti3
=== RUN   TestAdder_AddMulti3/group
--- PASS: TestAdder_AddMulti3 (0.00s)
    --- PASS: TestAdder_AddMulti3/group (0.00s)
        adder_test.go:245: use helper: false
        adder_test.go:238: use helper: true
 */