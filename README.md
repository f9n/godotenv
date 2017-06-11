# godotenv
Reading .env file in golang

# Download

```shell
  $ go get github.com/pleycpl/godotenv    # First download repo in go directory
```

# Usage

<h6>Example Env file (.env)</h6>

```
Key=Value
Mongo=20830
Token=asd123asdasd234
```

<h6>Usage in code</h6>

```go
  package main
  
  import (
    "fmt"
    "github.com/pleycpl/godotenv"
  )
  
  func main() {
    env := godotenv.Godotenv(".env")        // You should write your .env file's path.
    fmt.Println(env)                        // env variable is map.
    fmt.Println(env["Key"])
    fmt.Println(env["Mongo"])
    fmt.Println(env["Token"])
  }
```
