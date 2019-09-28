# envjson

## Requirement
Go (>= 1.11)

## Installation

```shell
go get github.com/hlts2/envjson
```

## Usage

Add your application configuration to your `.envjson` file in the root of your project:

```json
  
{
    "debug": "true",
    "db": {
        "user":   "user_1",
        "pass":   "pass_1",
        "dbname": "dbname_1"    
    },
    "services": {
        "user-provided": [
            {
                "name": "name",
                "instance_name": "instance_name"
            }
        ]
    }
}
```

Then in your Go app you can do something like

```go
package main

import (
    "github.com/hlts2/envjson"
    "log"
    "os"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    debug := os.Getenv("debug")
    db := os.Getenv("db")
}

```
