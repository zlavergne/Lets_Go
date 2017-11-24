Language Tools
---

#### Popular Go Compilers

There are two primary compilers available, and both are written by the official
Go developers:
  - `gc`: The original and default compiler toolchain
  - `gccgo`: A Go frontend for gcc

Compiler and associated tools are invoked by calling `go <command>`

Some examples follow:
  - `go run` and `go build`
  - `go test`
    - Can output test coverage with `-coverprofile`
    - Can test for race conditions with `-race`
  - `go get`: Package management

#### Editor Tools

There is not an official editor or IDE for Go. A list of common editors
and the tools associated with them can be found on 
[github](https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins).

Possible IDE's include:
  - Eclipse using [GoClipse](https://goclipse.github.io/)
  - [GoLand](https://www.jetbrains.com/go/), Jetbrains' Go IDE
  - Visual Studio

Possible text editors:
  - Vim can use [vim-go](https://github.com/fatih/vim-go) for autocompletion and refactoring
  - Atom can use [go-plus](https://github.com/joefitzgerald/go-plus)
  - Visual Studio Code has [vscode-go](https://github.com/Microsoft/vscode-go) available

#### Debugging 

  - `gdb`
    - `gc` requires some manual configuration
    - `gccgo` naturally has native support.

  - `delve`
    - Mostly compatible with commands you would normally use inside gdb
