# 时间相关

datetime 日期时间处理包

## 用法:

```go
import "github.com/yizhuying/luban/datetime"
```

## 目录

* [IsTimeInRangeStr](#IsTimeInRangeStr)
* [RandomTimeInRange](#RandomTimeInRange)

### IsTimeInRangeStr

时间是否在内

#### 函数定义:

```go
func IsTimeInRangeStr(timeStr, startStr, endStr, layout string) (bool, error)
```

#### 示例:

```go
package main

import (
	"fmt"

	"github.com/yizhuying/luban/datetime"
)

func main() {
	inRange, err := datetime.IsTimeInRangeStr("2023-01-15 10:00:00", "2023-01-01 00:00:00", "2023-01-31 23:59:59", "2006-01-02 15:04:05")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(inRange)

	// 输出: true

}

```

### RandomTimeInRange

在时间范围内随机返回一个时间

#### 函数定义:

```go
func RandomTimeInRange(start, end time.Time) (time.Time, error)
```

#### 示例:

```go
package main

import (
	"fmt"
	"time"

	"github.com/yizhuying/luban/datetime"
)

func main() {
	start := time.Date(2020, 6, 15, 13, 0, 0, 0, time.UTC)
	end := time.Date(2020, 6, 15, 13, 30, 0, 0, time.UTC)
	randomTime, err := datetime.RandomTimeInRange(start, end)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(randomTime)
}

// 输出: 2020-06-15 13:09:05.000000000 +0000 UTC

```