# Go lang HTTP + CURD

https://github.com/sausheong/gwp

Hello Word!

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
}

```

## 服务静态文件

```go
	files := http.FileServer(http.Dir("./public"))//  
	mux.Handle("/static/", http.StripPrefix("/static/", files))

```

