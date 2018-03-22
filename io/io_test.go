package io

import (
	"bytes"
	"io"
	"testing"
	"testing/iotest"
)

// io.Reader と io.Writer による入出力のある関数のテストをする.
func TestReaderWriter(t *testing.T) {

	// テスト用の io.Reader 実装を作成する.
	// io.Reader のモックは bytes.Buffer を使ってテストするのが手軽.
	text := "string from input buffer"
	input := bytes.NewBufferString(text)

	// テスト用の io.Writer 実装を作成する.
	// io.Writer のモックも bytes.Buffer を利用できる.
	output := new(bytes.Buffer)

	// io.Copy(io.Writer, io.Reader) を実行してテキストをコピーする.
	if _, err := io.Copy(output, input); err != nil {
		t.Error(err)
	}

	// テキストがコピーできていることを確認する.
	if output.String() != text {
		t.Errorf("failed to copy. output:%v", output)
	}
}

// io.Reader に iotest.NewReadLogger をラップさせると Read されると同時に内容を標準エラー出力に表示する.
// io.Writer に対する iotest.NewWriteLogger も同じように動作する.
func TestReaderWriterWithDebugPrint(t *testing.T) {
	input := bytes.NewBufferString("hello world")
	output := new(bytes.Buffer)
	_, err := io.Copy(
		iotest.NewWriteLogger("output: ", output),
		iotest.NewReadLogger("input: ", input))
	if err != nil {
		t.Error(err)
	}
	if output.String() != "hello world" {
		t.Errorf("faild to copy. output:%v", output)
	}
}

/* 実行結果

=== RUN   TestReaderWriterWithDebugPrint
2018/03/22 12:59:08 input:  68656c6c6f20776f726c64
2018/03/22 12:59:08 output:  68656c6c6f20776f726c64
2018/03/22 12:59:08 input:  : EOF
--- PASS: TestReaderWriterWithDebugPrint (0.00s)
 */