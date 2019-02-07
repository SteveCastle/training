## Compile c hello world

`gcc -o hello hello.c`

## Compile C++ hello world

`g++ -o hello-plus hello.c++`

## Compile rust hello world

`rustc main.rs`

## Compile go hello world (Or language of your choice)

`go build main.go`

## Decompile Binaries

`otool -tv hello > c-machine.txt && otool -tv hello-plus > cpp-machine.txt&& otool -tv cmd/main > go-machine.txt otool -tv hello-rust > rust-machine.txt`
