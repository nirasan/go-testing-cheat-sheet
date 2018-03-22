package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// http.Handler からテスト用のサーバーを起動してサーバー単位でのテストを行う.
// ハンドラー単体のテストと比べてルーティング設定やミドルウェアの挙動などもテストできる.
func TestMyServer_ServeHttp(t *testing.T) {
	// テスト用のサーバーを起動
	ts := httptest.NewServer(&MyServer{})
	defer ts.Close()

	// HelloHandler のテスト
	res, err := http.Get(ts.URL + "/hello?name=world")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("invalid response: %v", res)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	res.Body.Close()
	if string(body) != "hello world" {
		t.Errorf("invalid body: %s", body)
	}

	// JsonHandler のテスト
	b, err := json.Marshal(JsonRequest{Name: "world"})
	if err != nil {
		t.Fatal(err)
	}
	res, err = http.Post(ts.URL+"/json", "application/json", bytes.NewBuffer(b))
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("invalid response: %v", res)
	}
	resp := JsonResponse{}
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		t.Error(err)
	}
	res.Body.Close()
	if resp.Message != "hello world" {
		t.Errorf("invalid message: %#v", resp)
	}
}
