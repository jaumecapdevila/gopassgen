# gopassgen
My first go project that offers a cool cli binary to generate random passwords

## Installation
To install this project just run the command `git clone https://github.com/jaumecapdevila/gopassgen.git` inside your **$GOPATH** folder. Once the project is downloaded, just `cd` into the root folder and run the commands `go build` and `go install`. After all these steps you should be able to use the command `gopassgen` from you cli.

**Note:** Just be sure that you have defined the environment variables $GOAPTH and $GOBIN properly.

## Usage
In order to use this binary just type the command `gopassgen` from your command line. This program can accept two parameters:

1. The first one is the length of the password that you can specify as follow: `gopassgen -l 20`. The default value is 12.
2. The second one is a flag that will tell the program to use or avoid special characters once building the password. You can specify the flag's value as follow: `gopassgen -s=false`. By default this value is set to true.

So a full usage example could be: `gopassgen -l 5 -s=false`