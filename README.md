# pongo2-components

Adds web components to the existing [pongo2](github.com/flosch/pongo2) template engine.  
Now you can make components similar to how you would in popular JS frameworks.

## How to install and initialize

Install via `go get -u github.com/leandergangso/pongo2-components`

Import and initialize the package as seen bellow.

```go
package main

import (
    "github.com/flosch/pongo2"
    "github.com/leandergangso/pongo2-components"
)

func init() {
    err := pongo2components.Init() 
    if err != nil {
        panic(err)
    }
}

func main() {
    // ...
}
```

**Note**: you should init `pongo2components` before any calls to `pongo2`.

## How to use (basic example)

A basic button component.

1. Registering and setup button component.

```go
// components/button.go
package components

import (
    "github.com/leandergangso/pongo2-components"
)

func init() {
    button := pongo2components.Component{
        Name:     "button",
        FilePath: "components/button.html",
        Props:    []pongo2components.Prop{"value", "type"}
    }
    pongo2components.Register(button)
}
```

2. Create html component with available props.

```html
<!-- components/button.html -->
<button class="{{type}}">{{value}}</button>
```

3. Use button component.

```html
<!-- views/login.html -->
{% component 'button' value="Login" type="primary" %}
```

See [/example](github.com/leandergangso/pongo2-components/tree/master/examples) directory for more examples.
