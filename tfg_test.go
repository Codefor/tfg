package tfg

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T){
	now := time.Now()

	//2006-01-02 15:04:05
	str1 := Format(now,"%Y-%m-%d %H:%M:%S")
	str2 := now.String()[:19]

	if str1 != str2 {
		 t.Error("Format Error")
	}
}

func TestParse(t *testing.T){
	value := "2006-01-02 15:04:05"
	layout:= "%Y-%m-%d %H:%M:%S"

	t1,_ := Parse(layout,value)
	t2,_ := time.Parse(value,value)

	if !t1.Equal(t2) {
		t.Error("Parse error")
	}
}
