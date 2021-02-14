# Config

[![release github](https://badge.fury.io/gh/usvc%2Fgo-config.svg)](https://github.com/usvc/go-config/releases)
[![build status](https://travis-ci.org/usvc/go-config.svg?branch=master)](https://travis-ci.org/usvc/go-config)
[![pipeline status](https://gitlab.com/usvc/modules/go/config/badges/master/pipeline.svg)](https://gitlab.com/usvc/modules/go/config/-/commits/master)
[![Test Coverage](https://api.codeclimate.com/v1/badges/aa75f20cdfd8f0d5785b/test_coverage)](https://codeclimate.com/github/usvc/go-config/test_coverage)
[![Maintainability](https://api.codeclimate.com/v1/badges/aa75f20cdfd8f0d5785b/maintainability)](https://codeclimate.com/github/usvc/go-config/maintainability)

A Go package to deal with configuration.

| | |
| --- | --- |
| Github | [https://github.com/usvc/go-config](https://github.com/usvc/go-config) |
| Gitlab | [https://gitlab.com/usvc/modules/go/config](https://gitlab.com/usvc/modules/go/config) |

- - -

- [Config](#config)
	- [Usage](#usage)
		- [Importing](#importing)
		- [Defining configuration](#defining-configuration)
		- [Consuming from environment](#consuming-from-environment)
		- [Applying to Cobra (`github.com/spf13/cobra` package)](#applying-to-cobra-githubcomspf13cobra-package)
		- [Retrieving values frrom the configuration](#retrieving-values-frrom-the-configuration)
		- [Note on `*`Slice types](#note-on-slice-types)
		- [Note on configuration names](#note-on-configuration-names)
	- [Example CLI Application](#example-cli-application)
	- [Development Runbook](#development-runbook)
		- [Getting Started](#getting-started)
		- [Continuous Integration (CI) Pipeline](#continuous-integration-ci-pipeline)
			- [On Github](#on-github)
				- [Releasing](#releasing)
			- [On Gitlab](#on-gitlab)
				- [Version Bumping](#version-bumping)
				- [DockerHub Publishing](#dockerhub-publishing)
	- [Licensing](#licensing)

## Usage

### Importing

```go
import "github.com/usvc/go-config"
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
}
```

### Consuming from environment

> Following example assumes the above `conf` variable was defined.

```go
conf.LoadFromEnvironment()
```

### Applying to Cobra (`github.com/spf13/cobra` package)

> Following example assumes the above `conf` variable was defined.

```go
cmd := cobra.Command { /* ... config ... */ }
/* ... everything else ... */
conf.ApplyToCobra(&cmd)
// OR
conf.ApplyToCobraPersistent(&cmd)
cmd.Execute()
```

When using `conf.ApplyToCobra()` or `conf.ApplyToCobraPersistent()`, **the flag takes priority** if the equivalent environment variable was also defined. This is the behaviour since 0.2.5.

### Retrieving values frrom the configuration

> Following example assumes the above `conf` variable was defined.

```go
fmt.Println("bool         : %v", conf.GetBool("bool"))
fmt.Println("float        : %v", conf.GetFloat("float"))
fmt.Println("int          : %v", conf.GetInt("int"))
fmt.Println("int slice    : %v", conf.GetIntSlice("int slice"))
fmt.Println("string       : %s", conf.GetString("string"))
fmt.Println("string slice : %v", conf.GetStringSlice("string slice"))
fmt.Println("uint         : %v", conf.GetUint("uint"))
```

If you've applied the above `config.Map` instance to Cobra, running the following:

```sh
go run ./cmd/config \
	--bool=true \
	--float 3.142 \
	--int 1 \
	--int-slice 1,-2,3,-4,5 \
	--string a \
	--string-slice a,b,c,d,e \
	--uint 1
```

Should result in the following output:

```
uint: 1
bool: true
float: 3.142
int: 1
int slice: [1 -2 3 -4 5] (length: 5)
string: a
string slice: [a b c d e] (length: 5)
```

If you ran the following:

```sh

```


### Note on `*`Slice types

For the slice types (`IntSlice`, `StringSlice`), the delimiter is a comma (`,`) despite the underlying package (`github.com/spf13/viper`) handling the splitting a little differently. A tiny hack (famous last words: sorry, me-in-2025) was made to parse these using commas instead of spaces.

### Note on configuration names

The `config.Map` is used to define a dictionary of your configuration values using a `map[string]config.Config` data structure. The `string` becomes the name of the configuration and this `string` is manipulated before being passed to the downstream library.

For environment variables, all non-alphanumeric characters will be converted to an underscore (`_`), and all alphabetic characters will be converted to `UPPERCASE_IN_SNAKE_CASE`.

For flags, all non-alphanumeric characters will be converted to a hypen (`-`), and all alphabetic characters will be converted to `lowercase-in-kebab-case`.

## Example CLI Application

The example CLI application can be found in the [`./cmd/config` directory](./cmd/config) and includes the configuration as found above. To test it out you can run the following:

```sh
# bool
go run ./cmd/config --bool
BOOL=1 go run ./cmd/config

# float
go run ./cmd/config --float 3.142
FLOAT=3.142 go run ./cmd/config

# int
go run ./cmd/config --int -12345
INT=-12345 go run ./cmd/config

# []int
go run ./cmd/config --int-slice -12345 --int-slice -67890
INT_SLICE="-12345 -67890" go run ./cmd/config

# string
go run ./cmd/config --string "hello world"
STRING="hello world" go run ./cmd/config

# []string
go run ./cmd/config --string-slice "hello" --string-slice "world"
STRING_SLICE="hello world" go run ./cmd/config

# uint
go run ./cmd/config --uint 12345
UINT=12345 go run ./cmd/config
```

## Development Runbook

### Getting Started

1. Clone this repository
2. Run `make deps` to pull in external dependencies
3. Write some awesome stuff
4. Run `make test` to ensure unit tests are passing
5. Push

### Continuous Integration (CI) Pipeline

#### On Github

Github is used to deploy binaries/libraries because of it's ease of access by other developers.

##### Releasing

Releasing of the binaries can be done via Travis CI.

1. On Github, navigate to the [tokens settings page](https://github.com/settings/tokens) (by clicking on your profile picture, selecting **Settings**, selecting **Developer settings** on the left navigation menu, then **Personal Access Tokens** again on the left navigation menu)
2. Click on **Generate new token**, give the token an appropriate name and check the checkbox on **`public_repo`** within the **repo** header
3. Copy the generated token
4. Navigate to [travis-ci.org](https://travis-ci.org) and access the cooresponding repository there. Click on the **More options** button on the top right of the repository page and select **Settings**
5. Scroll down to the section on **Environment Variables** and enter in a new **NAME** with `RELEASE_TOKEN` and the **VALUE** field cooresponding to the generated personal access token, and hit **Add**

#### On Gitlab

Gitlab is used to run tests and ensure that builds run correctly.

##### Version Bumping

1. Run `make .ssh`
2. Copy the contents of the file generated at `./.ssh/id_rsa.base64` into an environment variable named **`DEPLOY_KEY`** in **Settings > CI/CD > Variables**
3. Navigate to the **Deploy Keys** section of the **Settings > Repository > Deploy Keys** and paste in the contents of the file generated at `./.ssh/id_rsa.pub` with the **Write access allowed** checkbox enabled

- **`DEPLOY_KEY`**: generate this by running `make .ssh` and copying the contents of the file generated at `./.ssh/id_rsa.base64`

##### DockerHub Publishing

1. Login to [https://hub.docker.com](https://hub.docker.com), or if you're using your own private one, log into yours
2. Navigate to [your security settings at the `/settings/security` endpoint](https://hub.docker.com/settings/security)
3. Click on **Create Access Token**, type in a name for the new token, and click on **Create**
4. Copy the generated token that will be displayed on the screen
5. Enter the following varialbes into the CI/CD Variables page at **Settings > CI/CD > Variables** in your Gitlab repository:

- **`DOCKER_REGISTRY_URL`**: The hostname of the Docker registry (defaults to `docker.io` if not specified)
- **`DOCKER_REGISTRY_USERNAME`**: The username you used to login to the Docker registry
- **`DOCKER_REGISTRY_PASSWORD`**: The generated access token

## Licensing

Code in this package is licensed under the [MIT license (click to see full text))](./LICENSE)
