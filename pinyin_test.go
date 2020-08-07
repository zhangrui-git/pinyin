package pinyin

import (
	"fmt"
	"testing"
)

func TestPinyin(t *testing.T)  {
	str := " "
	py := New(str)
	fmt.Println(py.FullPinyin)
	fmt.Println(py.Initials)
}