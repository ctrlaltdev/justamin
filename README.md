# Duration for Humans

Provided a time.Duration or a time.Time, will return a human representation of the duration.

## Installation

```sh
go get -u github.com/ctrlaltdev/justamin
```

## How To Use

Humanize a time.Duration:
```go
package main

import (
	"github.com/ctrlaltdev/justamin"
)

func main() {

	// ...
	
	d := time.Duration(-49 * 3600 * 1000000000)

	humanized := justamin.Humanize(d)
	// will return "2 days ago"

	// ...
}
```

Humanize a time.Time (will humanize relatively to time.Now()):
```go
package main

import (
	"github.com/ctrlaltdev/justamin"
)

func main() {

	// ...
	
	t := time.Now().Add(-time.Duration(3600 * 1000000000))

	humanized := justamin.Duration(t)
	// will return "an hour ago"

	// ...
}
```
