# serve
This is meant to be a utility to ease the process of sharing files between computers/phones/tablets that are in the same network

## Features
* Directory listing to download files
* Upload files to the shared directory
* Detects Avahi installation and gives the zeroconf (*.local) domain 
* Finds local network ip address

## Installation:
`go get https://github.com/empijei/serve`

## Usage: 
```
Usage of serve:
  -f string
    	The folder to serve (default PWD)
  -p string
    	The port to run on (default "8080")
  -s	Use TLS
  -u string
    	The webpath where the uplaod form is hosted (default "xyzzy/")
  -w string
    	The root webpath (default "/")
```

Example to share the ~/share folder on a computer named icarus
```
╭─ (OK) rob@icarus:~
╰─$ serve
Local ip address:172.20.6.46
Serving local folder /home/rob/share/ on "http://icarus.local:8080/"
File Upload form is available at "http://icarus.local:8080/xyzzy/"
Press Control+C to stop
```

## Repository info:
### tree:
```
serve
├── .gitignore
├── lib
│   ├── download.go
│   ├── upload.go
│   └── zeroconf.go
├── LICENSE
├── main.go
├── README.md
└── TODO
```
### cloc:
```
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               4             17              1            182
-------------------------------------------------------------------------------
SUM:                             4             17              1            182
-------------------------------------------------------------------------------
```
