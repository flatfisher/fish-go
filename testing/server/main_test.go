package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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
