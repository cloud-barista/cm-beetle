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
    export LOGFILE_PATH=cm-beetle.log
    ```

- `config.yaml` has:
    ```yaml
    logfile:
        path: ./cm-beetle.log
    ```

#### How to use it

- Use a blank import in your package (e.g., main, logger, and so on)
- Get a value using Viper

Note - It's just my preference. `config.Init()` can be used.

```go
import (
    // other packages

    // Loads configurations from setup.env and config.yaml
    _ "github.com/cloud-barista/cm-beetle/pkg/config"
)

func main() {
    logFilePath := viper.GetString("logfile.path")
    // Application logic follows
}
```

#### Wrapping Up

This setup illustrates the package's ability to harmonize settings from both `setup.env` and `config.yaml`,
showcasing its versatility and ease of use.
