## Introduction
A GoLang implementation that transpile AST rinha files into ES5 and runs using a [otto](https://github.com/robertkrimen/otto) (JavaScript parser and interpreter written natively in Go.)

## Build and Run
`docker build -t vitorvezani/rinha-de-compiler .`
`docker run -v {full_local_path}:/var/rinha/source.rinha.json vitorvezani/rinha-de-compiler`

## On Windows
`docker run -v "C:\Users\Vitor Vezani\projects\rinha-de-compiler\files\fib.json":/var/rinha/source.rinha.json vitorvezani/rinha-de-compiler`