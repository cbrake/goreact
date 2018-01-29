Go/React Web app
================

Overview
--------

This is a web application that can be used to receive data over REST interface, store in a db, and then display it.

Technologies
------------

* Go Programming language
  * https://golang.org/
  * used on the server (backend)
  * greatly simplifies development and deployment
  * very popular in new companies and startups
  * cross platform -- can easily target Linux, MacOS, or Windows.
* Elm 
  * http://elm-lang.org/
  * used for web application frontend
  * very robust environment for developing reliable web applications
  * allows for interactive single page applications
  * language encourages best practices and reliable application design
  * is considered the best practice for frontend code with projects like React Redux following it
* Bootstrap 4
  * http://elm-bootstrap.info
  * CSS toolkit to make the web UI look nice
  * well integrated with elm
  * provides correct rendering on a large number of browsers
  * enables responsive design (mobile browsers)
* MongoDB
  * https://www.mongodb.com/
  * https://labix.org/mgo (Go Mongodb driver)
  * allows easy simple of documents such as device config files, etc.

Additionally, the single page web architecture also enables the creation of mobile applications using Cordova later if desired.

Development
-----------

To set up a build environment (only tested under Linux):

* install nodejs and the Go language
* set up your Go development directory and GOPATH (http://bec-systems.com/site/1275/setting-up-a-go-development-environment)
* mkdir -p $GOPATH/src/github.com/cbrake/
* git clone git@git.bec-systems.com:cbrake/goelm.git $GOPATH/src/github.com/cbrake/goelm
* cd $GOPATH/src/github.com/cbrake/goelm
* source envsetup.sh
* app\_setup (only run once)
* app\_build (build Linux binaries)
* app\_build\_windows (build windows binaries)
* ./goelm

Implementation details:

* all assets and files (including the frontend) are embedded in the golang server binary.  This makes deployment of new versions very simple for development or production -- all you need is a single binary.
* There is no separate run-time to install (like Python, Ruby, Nodejs, Java, or C# environments).  Everything is statically included in the binary.  Again, this makes deployment very simple.
* there are no dependencies or stack that needs to be installed on the server -- only a single binary.
* this app runs well even on low end embedded ARM systems as the memory requirements and speed of Go are similiar to C.

Editors/Tooling
---------------

There are a number of editors that support Golang and React development very well including Vim, Atom, Visual Studio Code, etc.  If you don't have a strong preference, Visual Studio Code is a good one to start with.

### Visual Studio

Visual Studio Code is an advanced editor that runs well under Linux, and has extensions for both Go and Elm development.  It also integrates well with Git and allows you to commit/push changes to a git repo directly from the editor.

To install:

* install Visual Studio Code under Linux (https://code.visualstudio.com/)
* install Go extension: Ctrl+P, then enter "ext install Go"

Visual Studio is configured to auto-format Go and Javascript code automatically when saved.  
This helps keep code clean looking and consistent.  

VS Code also has an integrated terminal that will allow you to run the app\_build from within your editor.




