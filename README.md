# Configit

#### Yet another tool for Go config management. This is a simple package that supports loading environment variables into a custom struct. 

![Build Status](https://github.com/michaelwomack/configit/workflows/build/badge.svg) [![License: MIT](https://img.shields.io/badge/License-MIT-default.svg)](https://opensource.org/licenses/MIT)



#### Install 
```shell
go get github.com/michaelwomack/configit
```
#### Usage
```shell
export API_URL=https://localhost:8080
export MAX_CONCURRENCY=10
export DB_PORT=5432
export DB_HOST=127.0.0.1
```

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/michaelwomack/configit"
)

func main() {
	type config struct {
		ApiUrl         string `env:"API_URL"`
		MaxConcurrency int64  `env:"MAX_CONCURRENCY"`
		Db             struct {
			Port string `env:"DB_PORT"`
			Host string `env:"DB_HOST"`
		}
	}

	_ = os.Setenv("API_URL", "http://localhost:8080")
	_ = os.Setenv("MAX_CONCURRENCY", "10")
	_ = os.Setenv("DB_PORT", "3306")
	_ = os.Setenv("DB_HOST", "127.0.0.1")

	conf := &config{}
	if err := configit.Load(conf); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", conf)
}
```