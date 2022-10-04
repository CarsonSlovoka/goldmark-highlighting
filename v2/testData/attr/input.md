# Attribute

> {style="dracula" hls=[1, 3-5] base=3 linenos=table}

```go {style="dracula" hls=[1, 3-5] base=3 linenos=table}
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	// Output:
	// Hello World
}
```

> {style="dracula" hls=[3, 5-7] linenos=inline}

```go {style="dracula" hls=[3, 5-7] linenos=inline}
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	// Output:
	// Hello World
}
```

> {style="dracula" hls=[3, 5-7] linenos=true}

```go {style="dracula" hls=[3, 5-7] linenos=true}
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	// Output:
	// Hello World
}
```

> {style="dracula" hls=[3, 5-7]}

```go {style="dracula" hls=[3, 5-7]}
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	// Output:
	// Hello World
}
```

錯誤，lPrefix不行加上`"`，此外就算加上了也還是不能允許空白，所以就乾脆拿掉`"`
> go {lPrefix="demo Line"}

## lPrefix

LinkableLine Prefix 可以讓line number新增a屬性並附於其id

之所以要特別有lPrefix是因為如果每個區塊的id都用L開頭，那麼L5，會不知道到底這個id=L5是指哪一個區塊，

所以才需要有這個選項來說明

```go {linenos=true lPrefix=L}
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	// Output:
	// Hello World
}
```

在直接指定lPrefix，會直接啟動linenos，不需要特別說明(除非是table, inline等情況才要加註)

```go {lPrefix=demo1}
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	// Output:
	// Hello World
}
```

```go {lPrefix=demo2 base=5 hls=[3, 5-7]}
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	// Output:
	// Hello World
}
```

## Table Width(TW)

> {tw=2}
```go {tw=2}
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	// Output:
	// Hello World
}
```
