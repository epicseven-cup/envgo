# getOrElseEnv

A simple Go package for retrieving environment variables with a default value if they are not set.

## Description

The `getOrElseEnv` package provides a convenient function, `GetValueOrDefault`, to retrieve a value from an environment variable. If the environment variable is not set, it returns a specified default value. This is useful for handling situations where environment variables might be missing and you want to provide a fallback.

## Usage

```go
package main

import (
	"fmt"
  "github.com/epicseven-cup/getOrElseEnv" // Assuming the package is in the same directory
)

func main() {
	// Example 1: Retrieving a variable with a default
	name := getOrElseEnv.GetValueOrDefault("MY_NAME", "Guest")
	fmt.Println("Hello,", name)
  // "Hello, Guest"

  
	// Example 2: Retrieving a variable that exists
	apiKey := getOrElseEnv.GetValueOrDefault("API_KEY", "Not Set")
	fmt.Println("API Key:", apiKey)
  // "API Key: <existing key>"
}
