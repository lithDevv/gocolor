# Go color is a simple golang library for adding color to terminal output.
This library Supports the Following:
-> Normal colors
-> Bold colors 
-> Underlined colors
-> High Intensity Colors 
-> Bold High Intensity Colors

```commands -> go get github.com/lithDevv/gocolor/```

usage : 
```golang
package main

import (
  "fmt"
  "gocolor"
)

func main() {
  fmt.Println(gocolor.RedBold("Some Bold Red Text") + gocolor.Blank())
}
```
![image](https://user-images.githubusercontent.com/115331024/194685062-9c5728c9-6be3-47cc-a9f7-834b0f88c8c8.png)
