# Virtus - a simple command-line interface to serve HTML files written in Go 


virtus is a simple web server for serving static html files.<br />

Github: https://github.com/Szafa99/Virtus 

# Features 

Serving a static HTML file on a web server


# Instalation
Use the command below to install virtus

got get -u github.com/Szafa99/virtus/cmd


# Available Commands:
 ```diff
Available Commands
  help        Help about any command
  run         Serve a html file on the server
  version     Check virtus version


Flags:
  -h, --help   help for virtus
 ```
 run flags
 ```diff
  run
  Flags:
  -f, --file string   path to html file
  -h, --help          help for run
  ```
  
  # Examples 
  Example 1: Start a server and serve a HTML file on port 8080
 ```diff
  virtus run --file index.html
 ```
   Example 2: Check current version you're running
   ```diff
   virtus version
 ```

