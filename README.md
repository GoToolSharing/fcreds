# Installation

`go install github.com/QU35T-code/fzf-creds@latest`

# Usage

```bash
Interactive execution of bash commands

Usage:
  fzf-creds [command]

Available Commands:
  help        Help about any command
  link        Link a tool with fzf-creds
  list        Display a table of linked tools
  reset       Resetting fzf-creds
  smart       fzf-creds wrapper
  unlink      Unlink a tool of fzf-creds

Flags:
  -h, --help   help for fzf-creds

Use "fzf-creds [command] --help" for more information about a command.
```

# Example

```bash
fzf-creds list
fzf-creds link smbclient.py
source /opt/.exegol_aliases
fzf-creds list
smbclient.py %DOMAIN/%USERNAME:%PASSWORD@%TARGET

> smbclient.py %DOMAIN/%USERNAME:%PASSWORD@%TARGET
> smbclient.py NLIMCWHZ/qu35t:password@127.0.0.1

> Type help for list of commands
# 

fzf-creds unlink smbclient.py
fzf-creds list
```