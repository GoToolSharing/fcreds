# Installation

`go install github.com/QU35T-code/fzf-creds@latest`

# Requirements

`golang 1.19`

# Configuration

The files needed by fzf-creds are found by default at the following path :

`$HOME/.fzf-creds/`

# Usage

```bash
Interactive execution of bash commands

Usage:
  fzf-creds [command]

Available Commands:
  completion  Manage autocompletion command
  help        Help about any command
  link        Link a tool with fzf-creds
  list        Display a table of linked tools
  options     Displays the list of supported variables
  reset       Resetting fzf-creds
  smart       fzf-creds wrapper
  unlink      Unlink a tool of fzf-creds

Flags:
  -h, --help   help for fzf-creds

Use "fzf-creds [command] --help" for more information about a command.
```

# Autocompletion

```bash
fzf-creds completion enable
fzf-creds completion disable
```

# Example

```bash
fzf-creds list
fzf-creds link smbclient.py rpcclient ffuf sqlmap
source /root/.fzf-creds/aliases
fzf-creds list
smbclient.py %DOMAIN/%USERNAME:%PASSWORD@%TARGET

> smbclient.py %DOMAIN/%USERNAME:%PASSWORD@%TARGET
smbclient.py NLIMCWHZ/qu35t:password@127.0.0.1

> Type help for list of commands
# 

fzf-creds reset
fzf-creds list
```