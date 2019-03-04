package main

import (
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("add failed. expect:%d, actual:%d", 3, result)
	}
	t.Logf("result is %d", result)
}

func TestAdder_Add(t *testing.T) {
	a := &Adder{}
	testCases := []struct {
		L      int
		R      int
		Result int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{0, -1, -1},
		{100, 200, 300},
	}

	for _, testCase := range testCases {
		result := a.Add(testCase.L, testCase.R)
		if result != testCase.Result {
			t.Errorf("invalid result. testCase:%#v, actual:%d", testCase, result)
		}
	}
}

// サブテストの実行.
// サブテストを使うことで容易にテスト前後の処理を定義できる.
// テスト関数内で t.Run を使うことでサブテストが実行できる.
func TestAdder_AddMulti(t *testing.T) {
	// テスト開始処理
	t.Log("setup")

	// サブテスト. t.Run() が順次実行される.
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

	// テスト終了処理
	t.Log("tear-down")
}

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

// t.Error はテスト失敗としてログを出すが以後の処理も実行される.
// t.Errorf は出力のフォーマット指定ができるだけで挙動は同じ.
// func TestAdd2(t *testing.T) {
//     t.Error("error")
//     t.Log("log")
// }

// t.Fatal はテスト失敗としてログを出し以後の処理は実行されない.
// t.Fatalf は出力のフォーマット指定ができるだけで挙動は同じ.
// func TestAdd3(t *testing.T) {
//     t.Fatal("fatal")
//     t.Log("log")
// }

// t.Fail はテスト失敗とするが以後の処理も実行される.
// func TestAdd4(t *testing.T) {
//     t.Fail()
//     t.Log("log")
// }

// t.FailNow はテスト失敗とし以後の処理は実行されない.
// func TestAdd5(t *testing.T) {
//     t.FailNow()
//     t.Log("log")
// }
