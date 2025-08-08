## Intro
`strip-ansi` is a simple pipe program to remove ANSI escape characters from its standard input.

## Usage
### As a pipe:
`cat [input-file] | strip-ansi > [output-file]`

### With command line arguments:
`strip-ansi -i [input-file] -o [output-file]`

### Get help 
`strip-ansi -h|--help`

## Installation
`go install github.com/alex-vosk/strip-ansi@latest`
