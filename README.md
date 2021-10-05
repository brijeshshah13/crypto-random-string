# Crypto Random String

[![Go Reference](https://pkg.go.dev/badge/github.com/brijeshshah13/crypto-random-string.svg)](https://pkg.go.dev/github.com/brijeshshah13/crypto-random-string)

You can use this library to generate a cryptographically strong random string which can be useful for creating an
identifier, slug, salt, PIN code, fixture, etc.

## Installation

```shell
go get -u github.com/brijeshshah13/crypto-random-string
```

## Usage

### Using only `length`

```golang
package main

import (
	"fmt"
	cryptorandomstring "github.com/brijeshshah13/crypto-random-string"
)

func main() {
	generator := cryptorandomstring.New()
	if str, err := generator.WithLength(10).Generate(); err != nil {
		panic(err)
	} else {
		fmt.Println(str) // => "c152f80d02"
	}
}
```

### Using `kind: "base64"`

```golang
package main

import (
	"fmt"
	cryptorandomstring "github.com/brijeshshah13/crypto-random-string"
)

func main() {
	generator := cryptorandomstring.New()
	if str, err := generator.WithLength(10).WithKind("base64").Generate(); err != nil {
		panic(err)
	} else {
		fmt.Println(str) // => "e3WdumTFMK"
	}
}
```

### Using `kind: "url-safe"`

```golang
package main

import (
	"fmt"
	cryptorandomstring "github.com/brijeshshah13/crypto-random-string"
)

func main() {
	generator := cryptorandomstring.New()
	if str, err := generator.WithLength(10).WithKind("url-safe").Generate(); err != nil {
		panic(err)
	} else {
		fmt.Println(str) // => "A5WT-V~iG4"
	}
}
```

### Using `kind: "numeric"`

```golang
package main

import (
	"fmt"
	cryptorandomstring "github.com/brijeshshah13/crypto-random-string"
)

func main() {
	generator := cryptorandomstring.New()
	if str, err := generator.WithLength(10).WithKind("numeric").Generate(); err != nil {
		panic(err)
	} else {
		fmt.Println(str) // => "7917906408"
	}
}
```

### Using `kind: "distinguishable"`

```golang
package main

import (
	"fmt"
	cryptorandomstring "github.com/brijeshshah13/crypto-random-string"
)

func main() {
	generator := cryptorandomstring.New()
	if str, err := generator.WithLength(10).WithKind("distinguishable").Generate(); err != nil {
		panic(err)
	} else {
		fmt.Println(str) // => "0WT1KUR4ER"
	}
}
```

### Using `kind: "ascii-printable"`

```golang
package main

import (
	"fmt"
	cryptorandomstring "github.com/brijeshshah13/crypto-random-string"
)

func main() {
	generator := cryptorandomstring.New()
	if str, err := generator.WithLength(10).WithKind("ascii-printable").Generate(); err != nil {
		panic(err)
	} else {
		fmt.Println(str) // => "b|QK|LS"LN"
	}
}
```

### Using `kind: "alphanumeric"`

```golang
package main

import (
	"fmt"
	cryptorandomstring "github.com/brijeshshah13/crypto-random-string"
)

func main() {
	generator := cryptorandomstring.New()
	if str, err := generator.WithLength(10).WithKind("alphanumeric").Generate(); err != nil {
		panic(err)
	} else {
		fmt.Println(str) // => "WyK7545i98"
	}
}
```

### Using `characters`

```golang
package main

import (
	"fmt"
	cryptorandomstring "github.com/brijeshshah13/crypto-random-string"
)

func main() {
	generator := cryptorandomstring.New()
	if str, err := generator.WithLength(10).WithCharacters("abc").Generate(); err != nil {
		panic(err)
	} else {
		fmt.Println(str) // => "abbbccbaab"
	}
}
```

### Default Scraper (Ad hoc)

In simple cases, you can use the default scraper without creating an object instance

```golang
package main

import (
	"fmt"
	cryptorandomstring "github.com/brijeshshah13/crypto-random-string"
)

func main() {
	if str, err := cryptorandomstring.WithLength(10).WithCharacters("abc").Generate(); err != nil {
		panic(err)
	} else {
		fmt.Println(str)
	}
}
```

## API Options

### `length` with `WithLength`

*Required*\
Type: `uint64`

### `kind` with `WithKind`

Type: `string`\
Default: `'hex'`\
Values: `'hex' | 'base64' | 'url-safe' | 'numeric' | 'distinguishable' | 'ascii-printable' | 'alphanumeric'`

Use only characters from a predefined set of allowed characters.

Cannot be set at the same time as the `characters` option.

The `distinguishable` set contains only uppercase characters that are not easily confused: `CDEHKMPRTUWXY012458`. It can
be useful if you need to print out a short string that you'd like users to read and type back in with minimal errors.
For example, reading a code off of a screen that needs to be typed into a phone to connect two devices.

The `ascii-printable` set contains
all [printable ASCII characters](https://en.wikipedia.org/wiki/ASCII#ASCII_printable_characters): ``!"#$%&\'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~``
Useful for generating passwords where all possible ASCII characters should be used.

The `alphanumeric` set contains uppercase letters, lowercase letters, and
digits: `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789`. Useful for
generating [nonce](https://developer.mozilla.org/en-US/docs/Web/API/HTMLOrForeignElement/nonce) values.

### `characters` with `WithCharacters`

Type: `string`\
Minimum length: `1`\
Maximum length: `65536`

Use only characters from a custom set of allowed characters.

Cannot be set at the same time as the `type` option.