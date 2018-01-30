# Go/React Web app

## Overview

This is a web application that can be used to receive data over REST interface, store in a db, and then display it.

## Technologies

* Go (backend)
  * https://golang.org/
  * used on the server (backend)
  * greatly simplifies development and deployment
  * very popular in new companies and startups
  * cross platform -- can easily target Linux, MacOS, or Windows.
  * relatively "safe" language
    * statically typed so compiler catches a lot of problems
    * garbage collected (memory leaks are very rare)
  * Go is efficient and well suited for embedded Linux targets
    * compiles to a single, statically linked binary that is relatively small
    * compiled and runs similar to C/C++
* React (frontend)
  * https://reactjs.org/
  * very popular frontend language (https://bestof.js.org/)
  * allows for interactive and real-time applications
  * allows migrate to react-native for mobile apps in the future
* Bootstrap 4 (css toolkit)
  * https://getbootstrap.com/
  * CSS toolkit to make the web UI look nice
  * works well with react (https://reactstrap.github.io/)
  * provides correct rendering on a large number of browsers
  * enables responsive design (mobile browsers)
* https://parceljs.org/ (frontend build)
  * very powerful and simple frontend build tool
  * development and production modes
  * in development, automatically updates page in browser if sources change.

Additional Note:

* all assets and files (including the frontend) are embedded in the golang server binary. This makes deployment of new versions very simple for development or production -- all you need is a single binary.
* There is no separate run-time to install (like Python, Ruby, Nodejs, Java, or C# environments). Everything is statically included in the binary. Again, this makes deployment very simple.
* there are no dependencies or stack that needs to be installed on the server -- only a single binary.

## Setup

To set up a build environment (only tested under Linux):

* install nodejs and the Go language
* set up your Go development directory and GOPATH (http://bec-systems.com/site/1275/setting-up-a-go-development-environment)
* mkdir -p $GOPATH/src/github.com/cbrake/
* git clone git@git.bec-systems.com:cbrake/goreact.git $GOPATH/src/github.com/cbrake/goreact
* cd $GOPATH/src/github.com/cbrake/goreact
* source envsetup.sh
* app_setup (only run once)

## Production Build

* app_build (build Linux production binaries)
* app_build_windows (build windows binaries)
* ./goreact

## Development

Install the following:
* http://entrproject.org/

The following commands can be run (in separate terminal windows) during development.
Any time one of the source files in the project are changed, then appropriate parts of
the project are rebuilt, and automatically reloaded (if frontend change).

* app_watch_backend
* app_watch_frontend

## Editors/Tooling

There are a number of editors that support Golang and React development very well including Vim, Atom, Visual Studio Code, etc. If you don't have a strong preference, Visual Studio Code is a good one to start with.

### Visual Studio

Visual Studio Code is an advanced editor that runs well under Linux, and has extensions for both Go and Elm development. It also integrates well with Git and allows you to commit/push changes to a git repo directly from the editor.

To install:

* install Visual Studio Code under Linux (https://code.visualstudio.com/)
* install Go extension: Ctrl+P, then enter "ext install Go"

Visual Studio is configured to auto-format Go and Javascript code automatically when saved.  
This helps keep code clean looking and consistent.

VS Code also has an integrated terminal that will allow you to run the app_build from within your editor.
