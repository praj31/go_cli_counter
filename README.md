# Go CLI Counter

Just a utility Golang based and Redis backed CLI counter application that can help one keep a track of numerable tasks.

## Requirements

1. To create a build file, you must have Golang installed on your system.
2. Have Redis server installed and running in order to use the CLI.
   
## Instructions on Installation

1. Create a build file according to your target machine.
```sh
go build -o executable_file ./cmd/main.go
```
2. Provide the path to executable file in your environment PATH variable.


## Functionality
```
$ counter --help

Utility counter storage tool backed by Redis.

Usage:
   counter {flags}
   counter <command> {flags}

Commands:
   add                           adds a new counter
   decr                          decrements a counter by 1
   get                           gets a counter's value
   help                          displays usage informationn
   incr                          increments a counter by 1
   list                          displays counter list
   remove                        removes a counter
   set                           sets a counter's value
   version                       displays version number

Flags:
   -h, --help                    displays usage information of the application or a command (default: false)
   -v, --version                 displays version number (default: false)
```

## Usage

### Add a counter

```sh
$ counter add Posts 
```

### Remove a counter

```sh
$ counter remove Posts
```

### List all the counters

```sh
$ counter list
```

### Get a counter

```sh
$ counter get Posts
```

### Set a counter

```sh
$ counter set Posts 2
```

### Increment a counter's value by 1

```sh
$ counter incr Posts
```

### Decrement a counter's value by 1

```sh
$ counter decr Posts
```