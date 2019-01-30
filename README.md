# handlekeeper
[![GoDoc](https://godoc.org/github.com/danielb42/handlekeeper?status.svg)](https://godoc.org/github.com/danielb42/handlekeeper) 
[![Go Report Card](https://goreportcard.com/badge/github.com/danielb42/handlekeeper)](https://goreportcard.com/report/github.com/danielb42/handlekeeper) 
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)  

`handlekeeper` is a wrapper for `os.OpenFile()`. It intends to help applications with keeping track of active textfiles by presenting "stable" file handles even when the opened files are moved or deleted. This is achieved by instantly creating a new, empty file in the location of the original file, and re-opening the corresponding filehandle internally.

## Usage / Example
Here, if `/var/log/myApp.log` is moved or deleted, the file handle does not need to be re-opened. 

```go
package main

import (
	"bufio"
	"time"

	"github.com/danielb42/handlekeeper"
)

func main() {
	hk, _ := handlekeeper.NewHandlekeeper("/var/log/myApp.log")
	defer hk.Close()

	for {
		scanner := bufio.NewScanner(hk.Handle)

		for scanner.Scan() {
			println(scanner.Text())
		}

		time.Sleep(time.Second)
	}
}
```
