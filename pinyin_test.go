package pinyin

import (
	"testing"
)

func TestPinyin(t *testing.T) {
	py1 := New("中华人民共和国")
	t.Log(py1.FullPinyin)
	t.Log(py1.Initials)
	py2 := New("China")
	t.Log(py2.FullPinyin)
	t.Log(py2.Initials)
	py3 := New("长城")
	t.Log(py3.FullPinyin)
	t.Log(py3.Initials)
	py4 := New("长大")
	t.Log(py4.FullPinyin)
	t.Log(py4.Initials)
}
