# SANTANDER_GIT_CALCULATOR

This is a simple calculator made in Go (https://go.dev/).

To use it you must have Go installed in your computer, in Linux you can easy install using snap:

```
sudo snap install go --classic
```

or

```bash
sudo  rm -rf /usr/local/go && tar -C /usr/local -xzf go1.23.1.linux-amd64.tar.gz
```

than you must add the Go path to your .bashrc or .zshrc file:

```bash
export PATH=$PATH:/usr/local/go/bin
```

verify if the Go is installed:

```bash
go version
```


For other languagues, check the Go website:


 - [Golang Install Docs Page](https://go.dev/doc/install)


## Install or Run

### Run direct from command line:

### For the web interface:

```bash
  go run main.go
```

### Open your browser and access:

```bash
  http://localhost:8080
```

results will be printed in the web browser.

or run in terminal:

```bash
  go run main.go 5 + 3
```
### Result:

```bash
  8
```
### The program takes 3 arguments from the command line:

     - The First number, ex: 10
     - The operator (+, -, "*", /).
     - The second number, ex: 2

It checks if the arguments are valid and perform the operations.

Note: For multiplication, you must input as: "*" or \*, or the command line will output an error.

## Compilation

#### Linux or Mac:

```bash
  go build name_of_binary.go
```

#### Windows:

```bash
  go build -o name_of_binary.exe name_of_binary.go
```

### Using the binary:

#### Linux or Mac:

```bash
  ./calculator 10 + 2
  result: 12
```

#### Windows:

```bash
  calculator.exe 10 + 2
  result: 12
```

### Cross-Compiling for other OS:

#### Linux:

```bash
  GOOS=linux GOARCH=amd64 go build -o calculator-linux calculator.go
```

#### Windows:

```bash
  GOOS=windows GOARCH=amd64 go build -o calculator.exe calculator.go
```

#### Mac:

```bash
  GOOS=darwin GOARCH=amd64 go build -o calculator-mac calculator.go
```

## That's all folks, i hope you liked. :)

## Author

- [Luiz Campedelli](https://www.github.com/LuizCampedelli)




- [![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://github.com/LuizCampedelli/SANTANDER_GIT_CALCULATOR/blob/main/LICENSE)



## Contribute


Contributions are always welcome!

See `CONTRIBUTING.md` to learn how to get started.
