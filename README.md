# Config

[![pipeline status](https://gitlab.com/usvc/modules/go/config/badges/master/pipeline.svg)](https://gitlab.com/usvc/modules/go/config/-/commits/master)


A Go package to deal with configurations easily.

## Usage

### Importing

```go
import "github.com/usvc/config"
```

### Defining configuration

Following example can also be found at [`./cmd/config/config.go`](./cmd/config/config.go)

```go
var conf = config.Map{
	// with env : BOOL=true
	// with flag: --bool | -b
	"bool": &config.Bool{
		Default:   false,
		Shorthand: "b",
		Usage:     "specifies a boolean value",
	},
	// with env : FLOAT=-123
	// with flag: --float -123 | -f -123
	"float": &config.Float{
		Default:   1.6180339887498948482045868343,
		Shorthand: "f",
		Usage:     "specifies a floating point value",
	},
	// with env : INT=-123
	// with flag: --int -123 | -i -123
	"int": &config.Int{
		Default:   -1,
		Shorthand: "i",
		Usage:     "specifies a signed integer value",
	},
	// with env : INTS="-123 -456"
	// with flag: --int-slice -123,-456 | -I -123,-456
	"int slice": &config.IntSlice{
		Default:   []int{-2, -3},
		Shorthand: "I",
		Usage:     "specifies a slice of signed integers value",
	},
	// with env : STRING=value
	// with flag: --string value | -s value
	"string": &config.String{
		Default:   "default",
		Shorthand: "s",
		Usage:     "specifies a string value",
	},
	// with env : STRING_SLICE="value1 value2"
	// with flag: --string-slice value1,value2 | -S value1,value2
	"string slice": &config.StringSlice{
		Default:   []string{"hello", "world"},
		Shorthand: "S",
		Usage:     "specifies a slice of strings value",
	},
	// with env : UINT=123
	// with flag: --uint 123 | -u 123
	"uint": &config.Uint{
		Default:   1,
		Shorthand: "u",
		Usage:     "specifies an unsigned integer value",
	},
	// with env : UINT_SLICE="123 456"
	// with flag: --uint-slice 123,456 | -U 123,456
	"uint slice": &config.UintSlice{
		Default:   []uint{2, 3},
		Shorthand: "U",
		Usage:     "specifies a slice of unsigned integers value",
	},
}
```

### Consuming from environment

> Following example assumes the above `conf` variable was defined.

```go
conf.GetFromEnvironment()
```

### Applying to Cobra (`github.com/spf13/cobra` package)

> Following example assumes the above `conf` variable was defined.

```go
cmd := cobra.Command { /* ... config ... */ }
conf.ApplyToCobra(cmd)
```

### Deciding environment/flag precedence

> Following example assumes the above `conf` variable was defined.

To give priority to environment variables first, call `GetFromEnvironment()` outside of the `cobra.Command`'s runtime functions (eg. `PreRun`/`Run`)

```go
func main() {
  cmd := cobra.Command{ /* ... config ... */ }
  conf.ApplyToCobra(cmd)
  conf.GetFromEnvironment()
  cmd.Execute()
}
```

To give priority to values from flags first, call `GetFromEnvironment()` inside one of `cobra.Command`'s runtime functions:

```go
func main() {
  cmd := cobra.Command{
    // ... other config ...
    PreRun: func(c *cobra.Command, args []string) {
      conf.GetFromEnvironment()
    },
    // ... other config ...
  }
  conf.ApplyToCobra(cmd)
  cmd.Execute()
}
```

### Retrieving values frrom the configuration

> Following example assumes the above `conf` variable was defined.

```go
fmt.Println("bool     : %v", conf.GetValue())
fmt.Println("float    : %v", conf.GetValue())
fmt.Println("int      : %v", conf.GetValue())
fmt.Println("[]int    : %v", conf.GetValue())
fmt.Println("string   : %s", conf.GetValue())
fmt.Println("[]string : %v", conf.GetValue())
fmt.Println("uint     : %v", conf.GetValue())
fmt.Println("[]uint   : %v", conf.GetValue())
```

### Note on `*`Slice types

For the slice types (`IntSlice`, `StringSlice`, `UintSlice`), when running in flag mode, the delimiter is a comma (`,`), but in environment variable mode, the delimiter is a space (` `). This is because of how the underlying package (`github.com/spf13/viper`) does the string splitting.

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