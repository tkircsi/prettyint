# Pretty Int

This package extends the int basic type with printing the thousands separators. Currently the separator is fixed ','.

## Usage

```
package main

import (
	"fmt"

	"github.com/tkircsi/prettyint/prettyint"
)

func main() {
	var i prettyint.Int = -100000

	fmt.Println(i)
	fmt.Printf("%%d: %d\n", i)
	fmt.Printf("%%v: %v\n", i)
	fmt.Printf("%%s: %s\n", i)

	i = 99456
	fmt.Println(i)
	fmt.Printf("%%d: %d\n", i)
	fmt.Printf("%%v: %v\n", i)
	fmt.Printf("%%s: %s\n", i)

}
```

## Output

```
❯ go run main.go
-100,000
%d: -100,000
%v: -100,000
%s: -100,000
99,456
%d: 99,456
%v: 99,456
%s: 99,456
```
