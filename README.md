# goenvy

short and sweet, loads the environment variables from a .env

### Installation
```sh
$ go get github.com/ashdawson/goenvy
```

### Usage
Add your application configuration to a .env file in your root directory
In your go app you can simply call goenvy.Load()

```sh
package main

import (
    "github.com/ashdawson/goenvy"
)

func init() {
    goenvy.Load()
}
```
