package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// HelloHandler 単体でテストする
func TestHelloHandler(t *testing.T) {
	// テスト用のリクエスト作成
	req := httptest.NewRequest("GET", "http://example.com/?name=world", nil)
	// テスト用のレスポンス作成
	res := httptest.NewRecorder()
	// ハンドラーの実行
	HelloHandler(res, req)

	// レスポンスのステータスコードのテスト
	if res.Code != http.StatusOK {
		t.Errorf("invalid code: %d", res.Code)
	}

	// レスポンスのボディのテスト
	if res.Body.String() != "hello world" {
		t.Errorf("invalid response: %#v", res)
	}

	t.Logf("%#v", res)
}

// JsonHandler 単体でテストする.
// POST のボディで JSON を受け取ってレスポンスで JSON を返すハンドラー.
func TestJsonHandler(t *testing.T) {
	// テスト用の JSON ボディ作成
	b, err := json.Marshal(JsonRequest{Name: "world"})
	if err != nil {
		t.Fatal(err)
	}
	// テスト用のリクエスト作成
	req := httptest.NewRequest("POST", "http://example.com/", bytes.NewBuffer(b))
	// テスト用のレスポンス作成
	res := httptest.NewRecorder()
	// ハンドラーの実行
	JsonHandler(res, req)

	// レスポンスのステータスコードのテスト
	if res.Code != http.StatusOK {
		t.Errorf("invalid code: %d", res.Code)
	}

	// レスポンスの JSON ボディのテスト
	resp := JsonResponse{}
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		t.Errorf("errpr: %#v, res: %#v", err, res)
	}
	if resp.Message != "hello world" {
		t.Errorf("invalid response: %#v", resp)
	}

	t.Logf("%#v", resp)
}
