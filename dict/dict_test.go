package dict

import (
	"testing"
)

func TestDictLoad(t *testing.T) {
	d := NewDict(DefaultDictPath)
	err := d.Load()
	if err != nil {
		t.Error("Load dict failed")
	}
	if d.Size() <= 0 {
		t.Error("Dict is empty")
	}
}

func TestDefaultDictLoad(t *testing.T) {
	d := NewDefaultDict()
	err := d.Load()
	if err != nil {
		t.Error("Load dict failed")
	}
	if d.Size() <= 0 {
		t.Error("Dict is empty")
	}
}

func TestQueries(t *testing.T) {
	d := NewDict(DefaultDictPath)
	err := d.Load()
	if err != nil {
		t.Error("Load dict failed")
	}
	var trans string
	trans, err = d.Query("All Blacks")
	if trans != "全黑" {
		t.Error("Translate All Black failed")
	}
	trans, err = d.Query("Frans Steyn")
	if trans != "弗朗索瓦·斯特恩" {
		t.Error("Translate Frans Steyn failed")
	}
	trans, err = d.Query("François Steyn")
	if trans != "弗朗索瓦·斯特恩" {
		t.Error("Translate François Steyn failed")
	}
}
