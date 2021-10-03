# learn-golang

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
