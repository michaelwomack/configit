# Configit

##### Yet another tool for Go config management. This is a simple package that supports loading environment variables into a custom struct. 

#### Example

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