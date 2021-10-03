# learn-golang

This repository is my place to practice and make notes when i'm learning go language

#sumarry so far

- the language is statically typed
- most syntax dependent on "packages"
- function name start with Capital letter is exported by default
- the syntaxt similiar to C or Java
- := is the way in golang to declare new variabel and initialize at the same time
- "variabel := []string{value1, value2, value}" is the way to make new array in golang
- you have to add pointer whenever you want to put variabel values in print statements
- init() function will always only called once whenever a packages is called/used, mostly used for make connection to database, or any action that only occour once
- you have to import every package you want to use in the code
- you have to initialize the module with `go mod init <prefix>/<descriptive-text>` in production the prefix could be your repository in github.com or where you store your package in the cloud

#questions

- does it have object-oriented paradigms? or OOP way to code?

# last step

- build
- install package

# build

- go to hello directory
- type
  ` go build` to compile the code into an executable
- now you can run the new hello executable to confirm that the code is working
  ` hello.exe` in windows
  `./hello` in Linux or Mac

# install the executable so you can run without specify the path, means that you set executable into env path

- still in hello directory
- type
  ` go list -f '{{.Target}}'` to show where is the executable is installed, in my case is

  `C:\Users\riko\go\bin\hello.exe`

- then type
  ` go env -w GOBIN=C:\path\to\your\bin` example ` go env -w GOBIN=C:\Users\riko\go\bin\`

- then type
  ` go install`

- now you can run the package just by calling its package name
  `hello`
  `map[Darrin:Hail, Darrin! Well met! Gladys:Great to see you, Gladys! Samantha:Hail, Samantha! Well met!]`
