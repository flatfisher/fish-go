package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1,2)
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
        {100, 200, 0},
    }

    for _, testCase := range testCases {
        result := a.Add(testCase.L, testCase.R)
        if result != testCase.Result {
            t.Errorf("invalid result. testCase:%#v, actual:%d", testCase, result)
        }
    }
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