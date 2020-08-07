# pinyin
golang中文汉字转拼音
```go
package main

import (
    "fmt"
    "github.com/zhangrui-git/pinyin"
    "strings"
)

func main()  {
    str := "中华人民共和国"
    py := pinyin.New(str)
    fmt.Println(py.FullPinyin)
    fmt.Println(py.Initials)
    fmt.Println(strings.Join(py.FullPinyin, " "))
    fmt.Println(strings.Join(py.Initials, "-"))
}
```