package mlog

import "testing"

func TestNew(t *testing.T) {
	l := NewLogger()
	l.DeBug.Println("111")
	l.Info.Println("111")
	l.Error.Println("111")
}
