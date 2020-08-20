# testify-matchers
A set of argumentMatchers for testify's mock package

To be used in conjunction with testify: http://github.com/stretchr/testify

### Usage

```go
package main

import (
    matcher "github.com/henryleduc/testify-matchers"
    "github.com/stretchr/testify/mock"
    "time"
)

func main() {
    // TimeWithGracePeriod matcher
    mock.MatchedBy(matcher.TimeWithGracePeriod(10 * time.Second))
    
    // ContextWithValue matcher
    mock.MatchedBy(matcher.ContextWithValue("myBool", true))
}
```