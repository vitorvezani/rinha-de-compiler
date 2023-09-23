## Introduction
A GoLang implementation that transpile AST rinha files into JS and runs with node

## Build
`docker build -t vitorvezani/rinha-de-compiler .`

## Run
`docker run -v {full_local_path}:/var/rinha/source.rinha.json vitorvezani/rinha-de-compiler`
