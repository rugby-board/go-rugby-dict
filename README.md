# go-rugby-dict

Rugby dict lib in Go

## Usage

### Installation

```shell
go get github.com/rugby-board/go-rugby-dict
```

### Example

```go
import (
    "fmt"

    dict "github.com/rugby-board/go-rugby-dict"
)

d := dict.NewDict("dict.yaml")
err := d.Load()
if err != nil {
    fmt.Println("load error")
}
chineseTranslation, err := d.Query("Frans Steyn")
fmt.Println(chineseTranslation) // 弗朗索瓦·斯特恩
```
