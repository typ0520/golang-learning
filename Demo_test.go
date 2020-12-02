package main

import (
	"testing"
)

//go test

func TestAdd(t *testing.T) {
	excepted := 3
	got := 1 + 2
	if got != excepted {
		t.Errorf("excepted:%v, got:%v", excepted, got) // 测试失败输出错误提示
	}
}
