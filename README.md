# Fybe CLI (Command-Line Interface)

`fybe` is the command-line interface (CLI) for managing resources from Fybe (http://fybe.com).

## Installation

1. Download the binary `fybe` for Windows, MacOS or Linux from the [releases page](https://github.com/fybecom/fybe/releases).

Hint: due to MacOS security controls, it is recommended to perform the following command:

```sh
curl -L '<url to release>' | tar xz
```

2. You might move the executable to any location on your disk. You may update your `PATH` environment variable for easier invocation.

## Getting Started

1. Setup `fybe` to use your credentials from [Cockpit](https://cockpit.fybe.com/account/security). This is only required once as the settings are stored on your disk. Whenever at least one of those settings are changed, please re-run the command again.

  ```sh
  fybe config set-credentials --oauth2-clientid=<ClientId from Cockpit> --oauth2-client-secret=<ClientSecret from Cockpit> --oauth2-user=<API User from Cockpit> --oauth2-password=<API Password from Cockpit>
  ```

1. You are ready to use `fybe`

```sh
fybe help
# List available images
fybe get images
### Start Compute Instance
fybe start instance 973743
```

## Enable Shell Completion

```sh
fybe completion
Bash:

        $ source <(fybe completion bash)

        # To load completions for each session, execute once:
        # Linux:
        $ fybe completion bash > /etc/bash_completion.d/fybe
        # macOS:
        $ fybe completion bash > /usr/local/etc/bash_completion.d/fybe

Zsh:

        # If shell completion is not already enabled in your environment,
        # you will need to enable it.  You can execute the following once:

        $ echo "autoload -U compinit; compinit" >> ~/.zshrc

        # To load completions for each session, execute once:
        $ fybe completion zsh > "${fpath[1]}/_fybe"

        # You will need to start a new shell for this setup to take effect.

fish:

        $ fybe completion fish | source

        # To load completions for each session, execute once:
        $ fybe completion fish > ~/.config/fish/completions/fybe.fish

PowerShell:

        PS> fybe completion powershell | Out-String | Invoke-Expression

        # To load completions for every new session, run:
        PS> fybe completion powershell > fybe.ps1
        # and source this file from your PowerShell profile.
```

## Build from source

1. Clone `fybe` project
2. Install `go` from here [Golang page](https://golang.org/doc/install)
3. Compile it via

```sh
go mod tidy && go mod download && go build -ldflags="-w -s"
```

4. Update apiclient (requires docker / podman)

```sh
rm -rf apiclient; docker run --rm -v "${PWD}:/local" --env JAVA_OPTS='-Dio.swagger.parser.util.RemoteUrl.trustAll=true -Dio.swagger.v3.parser.util.RemoteUrl.trustAll=true' openapitools/openapi-generator-cli:v5.2.1 generate --skip-validate-spec --input-spec 'https://api.fybe.com/api-v1.yaml' --generator-name go --output /local/apiclient
```

## Files used by `fybe`

`fybe` uses following files to work:

* `~/.fybe.yaml` is the default file for storing settings
* `~/.cache/fybe/` folder is used for caching purposes. Can be deleted safely.

## License




GNU GENERAL PUBLIC LICENSE, Version 3
