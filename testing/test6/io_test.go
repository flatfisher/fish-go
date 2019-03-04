package io

import (
	"bytes"
	"io"
	"testing"
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
