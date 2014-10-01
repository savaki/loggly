loggly
======

command line client for loggly written in go

## Usage

```
NAME:
   loggly - send logs from the specified port to loggly

USAGE:
   loggly [global options] command [command options] [arguments...]

VERSION:
   0.1

AUTHOR:
  Matt Ho

COMMANDS:
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --port '6030'			the default port loggly will listen to [$PORT]
   --token 				your loggly token [$TOKEN]
   --tag '--tag option --tag option'	the tags to apply to the stream [$TAGS]
   --help, -h				show help
   --version, -v			print the version
```

## Example

Starts up loggly to listen to the default port, 6030, using the token specified

```
loggly --token "your-loggly-token"
```

## Example 2 - Token from the Environment

Suppose you don't want to pass your token on the command.  No problem, pass it via the environment variable, TOKEN.

```
export TOKEN="your-loggly-token"
loggly
```