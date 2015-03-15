Go Exercisms
=============

Uploads from my personal [exercisms](http://exercism.io)...

Feel free to leave comments via **issues**.

**Side note on Go:** 
Best to create a "main.go" file in "exercism/go" folder, 
then import exercise package via relative path...

##### File Structure
```shell

exercism/
	go/
		ex1/
			ex1.go
			ex1_test.go
    main.go

```
<br>

##### "main.go" Structure
```go
package main

import (
	ex "./ex1"
	"fmt"
)

func main() {
	a := ex.Func(10)
	fmt.Println(a)
}
```