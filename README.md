# Templator
CLI application that reverse engineers go templates to make it easy to use templates without knowing the input variables beforehand.

## TL;DR
1. Point to a template
2. Answer some questions
3. Get that nice document you do not want to write everytime!

## Usage
```bash
go get github.com/vorstenbosch/templator
${GOPATH}/bin/templator -h
```

## Supported template types
- ActionNode (plain strings)
- RangeNode (lists)