# Fingo

![GitHub](https://img.shields.io/github/license/shuyaeer/fingo)  ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/shuyaeer/fingo)  
Fingo is a command line interface written in Go for travarse file system.
Fingo is developed to substitute shell's `find`.  
This module is designated to integrate file searching program.

# Installation

As a premise, You have already installed Go. Then,

```bash
go install github.com/shuyaeer/fingo@latest
```

binary file is installed beneath your $GOPATH/

```bash
 $ ls $GOPATH/bin
 fingo
```

# Quick Start

When you want to find csv files on whole your file system.

```bash
$ fingo / --extention csv
file : /home/ubuntu/csv_files/2022-02-13.csv
file : /home/ubuntu/csv_files/2022-03-13.csv
file : /home/ubuntu/2022-04-13.csv
                :
                :
```

When you want to find file by file name/

```bash
$ fingo / --name 2022-04-13
file : /home/ubuntu/2022-04-13.csv
file : /home/ubuntu/2022-04-13.txt
file : /home/ubuntu/documents/2022-04-13.pdf
                :
                :
```


# Options

```text
-h, --help        ・print help text and exit
-e, --extention   ・extracting file match the FileExtention
-n --name         ・search file based on designated file name
```
