### Config Package

#### Overview

The `config` package manages configurations in Go applications,
ensuring compatibility between `config.yaml` and `setup.env`.
`setup.env` is used to setup environment variables.

Note - When both environment variables and config.yaml settings are present,
the package prioritizes environment variables, overriding equivalent settings in config.yaml.

#### Compatible configurations example

The below configurations are compatible in this project.

- `setup.env` contains:

  ```
  export LOGFILE_PATH=beetle.log
  ```

- `config.yaml` has:
  ```yaml
  logfile:
    path: ./beetle.log
  ```

#### How to use it

- Get a value using Viper

Note - It's just my preference. `config.Init()` can be used.

```go
// Package main is the starting point of CM-Beetle
package main

import (
    // other packages

    "github.com/cloud-barista/cm-beetle/pkg/config"
)

func init() {

	common.SystemReady = false

	// Initialize the configuration from "config.yaml" file or environment variables
	config.Init()

}

func main() {
    // Application logic follows
}
```

#### Wrapping Up

This setup illustrates the package's ability to harmonize settings from both `setup.env` and `config.yaml`,
showcasing its versatility and ease of use.
