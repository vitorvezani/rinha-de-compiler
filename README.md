## Introduction
A GoLang implementation that transpile AST rinha files into ES5 and runs using a [otto](https://github.com/robertkrimen/otto) (JavaScript parser and interpreter written natively in Go.)

## Build
`docker build -t vitorvezani/rinha-de-compiler .`

## Run
`docker run -v {full_local_path}:/var/rinha/source.rinha.json vitorvezani/rinha-de-compiler`
