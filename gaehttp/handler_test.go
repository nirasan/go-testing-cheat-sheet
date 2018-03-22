package gaehttp

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

// HelloHandler 単体でテストする.
// ユーザーが存在しないので失敗するケース.
func TestHelloHandler_Failure(t *testing.T) {
	// App Engine テスト用の環境を立ち上げる
	instance, err := aetest.NewInstance(&aetest.Options{StronglyConsistentDatastore: true})
	if err != nil {
		t.Fatal(err)
	}
	defer instance.Close()

	// App Engine テスト用のリクエストを作成
	req, err := instance.NewRequest("GET", "http://example.com/?id=user1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// テスト用のレスポンス作成
	res := httptest.NewRecorder()

	// ハンドラーの実行
	HelloHandler(res, req)

	// ユーザーが存在しないのでエラーになる
	if res.Code != http.StatusInternalServerError {
		t.Errorf("invalid code: %d", res.Code)
	}
}

// HelloHandler 単体でテストする.
// 事前にユーザーを作成して成功するケース.
func TestHelloHandler_Success(t *testing.T) {
	// App Engine テスト用の環境を立ち上げる
	instance, err := aetest.NewInstance(&aetest.Options{StronglyConsistentDatastore: true})
	if err != nil {
		t.Fatal(err)
	}
	defer instance.Close()

	// App Engine テスト用のリクエストを作成
	req, err := instance.NewRequest("GET", "http://example.com/?id=user1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// テスト用のレスポンス作成
	res := httptest.NewRecorder()

	// 事前にユーザーを作成する
	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, "User", "user1", 0,nil)
	if _, err = datastore.Put(ctx, key, &User{Name: "world"}); err != nil {
		t.Error(err)
	}

	// ハンドラーの実行
	HelloHandler(res, req)

	// ユーザーが存在しないのでリクエストに成功する
	if res.Code != http.StatusOK {
		t.Errorf("invalid code: %d", res.Code)
	}

	// レスポンスのボディのテスト
	if res.Body.String() != "hello world" {
		t.Errorf("invalid response: %#v", res)
	}

	t.Logf("%#v", res)
}
