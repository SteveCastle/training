# Session 1: Comparative Languages Introduction

This session is designed simply to familiarize yourself with
the installation and build process for the languages we will use in this course.
For fun run a decompile of each language binary to
get a feel for how much each compiler is doing. Later we will look at how all of
the compilers work and what makes them different.

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
