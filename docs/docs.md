# cli-experiment

## Usage
> This cli template does nothing useful.

cli-experiment

## Description

```
This is a template CLI application, which can be used as a boilerplate for awesome CLI tools written in Go.
This template currently does nothing useful.
```
## Examples

```bash
cli-template example
```

## Flags
|Flag|Usage|
|----|-----|
|`--debug`|enable debug messages|
|`--disable-update-checks`|disables update checks|
|`--raw`|print unstyled raw output (set it if output is written to a file)|

## Commands
|Command|Usage|
|-------|-----|
|`cli-experiment completion`|Generate the autocompletion script for the specified shell|
|`cli-experiment example`|Does nothing useful.|
|`cli-experiment help`|Help about any command|
# ... completion
`cli-experiment completion`

## Usage
> Generate the autocompletion script for the specified shell

cli-experiment completion

## Description

```
Generate the autocompletion script for cli-experiment for the specified shell.
See each sub-command's help for details on how to use the generated script.

```

## Commands
|Command|Usage|
|-------|-----|
|`cli-experiment completion bash`|Generate the autocompletion script for bash|
|`cli-experiment completion fish`|Generate the autocompletion script for fish|
|`cli-experiment completion powershell`|Generate the autocompletion script for powershell|
|`cli-experiment completion zsh`|Generate the autocompletion script for zsh|
# ... completion bash
`cli-experiment completion bash`

## Usage
> Generate the autocompletion script for bash

cli-experiment completion bash

## Description

```
Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(cli-experiment completion bash)

To load completions for every new session, execute once:

#### Linux:

	cli-experiment completion bash > /etc/bash_completion.d/cli-experiment

#### macOS:

	cli-experiment completion bash > /usr/local/etc/bash_completion.d/cli-experiment

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion fish
`cli-experiment completion fish`

## Usage
> Generate the autocompletion script for fish

cli-experiment completion fish

## Description

```
Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	cli-experiment completion fish | source

To load completions for every new session, execute once:

	cli-experiment completion fish > ~/.config/fish/completions/cli-experiment.fish

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion powershell
`cli-experiment completion powershell`

## Usage
> Generate the autocompletion script for powershell

cli-experiment completion powershell

## Description

```
Generate the autocompletion script for powershell.

To load completions in your current shell session:

	cli-experiment completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion zsh
`cli-experiment completion zsh`

## Usage
> Generate the autocompletion script for zsh

cli-experiment completion zsh

## Description

```
Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions for every new session, execute once:

#### Linux:

	cli-experiment completion zsh > "${fpath[1]}/_cli-experiment"

#### macOS:

	cli-experiment completion zsh > /usr/local/share/zsh/site-functions/_cli-experiment

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... example
`cli-experiment example`

## Usage
> Does nothing useful.

cli-experiment example

## Flags
|Flag|Usage|
|----|-----|
|`-f, --format string`|specify a custom date format (default "02 Jan 06")|
# ... help
`cli-experiment help`

## Usage
> Help about any command

cli-experiment help [command]

## Description

```
Help provides help for any command in the application.
Simply type cli-experiment help [path to command] for full details.
```


---
> **Documentation automatically generated with [PTerm](https://github.com/pterm/cli-template) on 31 August 2022**
