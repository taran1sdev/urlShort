# urlShort

urlShort is my first real project using Go

The application takes in a yaml/json file containing paths and urls. 

It hosts a local http server and when the user navigates to a path specified in the yaml/json file it redirects the user to the corresponding url. 

The server starts on port 80 but you can change this if you already have a server running. 

Build:
```bash
go install gopkg.in/yaml.v2@latest
go mod init urlshort
go build urlShort.go
```

Usage:
```bash
$ ./urlShort -pathsFile=yourfile.yaml||json
```

Then use your browser to navigate to http://localhost/yourpath - You should get redirected!


