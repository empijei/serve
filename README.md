# serve
This is meant to be a utility to ease the process of sharing files between computers/phones/tablets that are in the same network

## Features
* Directory listing to download files
* Upload files to the shared directory
* Detects Avahi installation and gives the zeroconf (*.local) domain 
* Finds local network ip address
* Supports TLS

## Installation:
### Get the tool:
```sh
go get https://github.com/empijei/serve
```

### Setup TLS:
```sh
bash $GOPATH/src/github.com/empijei/serve/generate-certificate.sh
```

Follow the wizard.

WARNING: If you want to use zeroconf hostnames (*.local) remeber to put that hostname as name for your certificate.

## Usage: 
```
Usage of serve:
  -f string
    	The folder to serve (default PWD)
  -p string
    	The port to run on (default "8080")
  -s	Use TLS
  -u string
    	The webpath where the upload form is hosted (default "upload/")
  -w string
    	The root webpath (default "/")
  -h prints this help
```

Example to share the ~/share folder on a computer named icarus
```
╭─ (OK) rob@icarus:~
╰─$ serve
Local ip address:172.20.6.46
Serving local folder /home/rob/share/ on "http://icarus.local:8080/"
File Upload form is available at "http://icarus.local:8080/upload/"
Press Control+C to stop
```
