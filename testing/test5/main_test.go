package fixture

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

// フィクスチャーを読み込むための型の定義
type Fixture struct {
	TestCases []testCase `json:"test_cases"`
}
type testCase struct {
	L      int `json:"l"`
	R      int `json:"r"`
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
