# Change Log

All notable changes to this project will be documented in this file.

--- 

## [1.0.0](https://github.com/QU35T-code/fzf-creds/releases/tag/1.0.0) (17/06/2023)

#### Summary
This release adds the `core` fzf-creds commands as well as `crackmapexec integration`.

## Changes
- The `link` command has been added.
- The `list` command has been added.
- The `reset` command has been added.
- The `smart` command has been added.
- The `unlink` command has been added.
- Integration with `crackmapexec` databases for : `DOMAIN`, `USERNAME`, `PASSWORD`, `TARGET` variables has been added.

## [1.1.0](https://github.com/QU35T-code/fzf-creds/releases/tag/1.1.0) (18/06/2023)

#### Summary
New commands and the first bug fixes.

## Changes
- The `options` command has been added. This shows variables that can interact with fzf-creds.
- A `new column` has been added for the `list` command. It is used to display the `alias` associated with the tool.
- The `reset` command has been fixed and deletes all the modifications that may have been made by fzf-creds.
- Problem installing fzf-creds with `go install` has been resolved.
- The `completion` command has been added. It `enables` and `disables` autocompletion of fzf-creds commands.
- Configuration issues have been resolved.
- Added ability to `link` and `unlink` multiple commands `at once` on a single line.
- Removal of the `viper` module and `environment variables`.
- Added more control over arguments passed to commands.
- Separation of the alias file from the default one.

## [1.1.1](https://github.com/QU35T-code/fzf-creds/releases/tag/1.1.1) (18/06/2023)

#### Summary
Fix for not blocking commands due to flags.

## Changes
- Flags are `ignored` for the `smart` command, preventing a conflict between tools and fzf-creds flags