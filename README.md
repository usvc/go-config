# Config

[![pipeline status](https://gitlab.com/usvc/modules/go/config/badges/master/pipeline.svg)](https://gitlab.com/usvc/modules/go/config/-/commits/master)


A Go package to deal with configurations easily.

## Usage

### Importing

```go
import "github.com/usvc/config"
```

### Creating a String configuration

```go
var conf = config.Map{
  // settable via environment variable STRING
  // or command flag --string
  // or alias -s
  "string": &config.String{
    Default: "default",
    Shorthand: "s",
    Usage: "specifies a string value",
  },
}
```

### Creating an Int configuration

```go
var conf = config.Map{
  // settable via environment variable INT
  // or command flag --int
  // or alias -s
  "int": &config.Int{
    Default: 1,
    Shorthand: "i",
    Usage: "specifies a string value",
  },
}
```

## Development Runbook

### Getting Started

1. Clone this repository
2. Run `make deps` to pull in external dependencies
3. Write some awesome stuff
4. Run `make test` to ensure unit tests are passing
5. Push

### Continuous Integration (CI) Pipeline

To set up the CI pipeline in Gitlab:

1. Run `make .ssh`
2. Copy the contents of the file generated at `./.ssh/id_rsa.base64` into an environment variable named **`DEPLOY_KEY`** in **Settings > CI/CD > Variables**
3. Navigate to the **Deploy Keys** section of the **Settings > Repository > Deploy Keys** and paste in the contents of the file generated at `./.ssh/id_rsa.pub` with the **Write access allowed** checkbox enabled

- **`DEPLOY_KEY`**: generate this by running `make .ssh` and copying the contents of the file generated at `./.ssh/id_rsa.base64`

## Licensing

Code in this package is licensed under the [MIT license (click to see full text))](./LICENSE)