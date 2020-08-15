# Hikvision ISAPI Go SDK

[![PkgGoDev](https://pkg.go.dev/badge/github.com/loozhengyuan/hikvision-sdk/hikvision)](https://pkg.go.dev/github.com/loozhengyuan/hikvision-sdk/hikvision)
[![Go Report Card](https://goreportcard.com/badge/github.com/loozhengyuan/hikvision-sdk)](https://goreportcard.com/report/github.com/loozhengyuan/hikvision-sdk)
![ci](https://github.com/loozhengyuan/hikvision-sdk/workflows/ci/badge.svg)

> **Important**: This package is still under heavy development, only a partial implementation of the API is available.

Hikvision Go SDK for Intelligence Security API (ISAPI).

This implementation is heavily dependent on the [ISAPI Core Protocol](/resources/isapi.pdf). The latest available copy of this specification document was updated on Sep 2019.

# Installation

To install the SDK, use `go get` to fetch the latest version:

```shell
go get -u github.com/loozhengyuan/hikvision-sdk/hikvision
```

Once installed, you may import it directly into your Go application:

```shell
import "github.com/loozhengyuan/hikvision-sdk/hikvision"
```

# Usage

```go
package main

import (
	"fmt"

	"github.com/loozhengyuan/hikvision-sdk/hikvision"
)

func main() {
	// Create client object
	c, err := hikvision.NewClient(
		"YOUR_SERVER_HOST",
		"YOUR_SERVER_USERNAME",
		"YOUR_SERVER_PASSWORD",
	)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	// Send request
	d, err := c.GetDeviceInfo()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("%+v\n", d)
}
```

# Contributing

There's still lots of work to be done! Only a small subset of the APIs are implemented, so pull requests are welcome.

# License

[GPL-3.0](https://choosealicense.com/licenses/gpl-3.0/)
