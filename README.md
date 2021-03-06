# ThroneStats

Nuclear Throne statistics and OBS overlay service hosted at [thronestats.com](http://thronestats.com)

Tracks statistics of players both on an individual and a global level.

Tracks:

 - Weapon pickups
 - Causes of death
 - Mutation choices
 - Crowns
 - Characters

Shows popups on overlay with information such as:

```
Picked up a *Shotgun*, which you pick up on *12.34%* of your runs. Globally
*Shotgun* is picked up on *7.89%* of runs.
```


## Building

You'll need Go installed in your environment, get it from [https://golang.org/dl/](https://golang.org/dl/). Make sure you can run `go` on your CLI before trying to continue.

Tested on Go 1.5.1.

Make sure you get all dependencies and then build it

```
go get github.com/lietu/thronestats/cmd/thronestats
cd $GOPATH/src/github.com/lietu/thronestats/cmd/thronestats
go build
```

## Dependencies for WWW

Firstly you'll need [Node.js]() and NPM installed. Then you need to install `bower` and `gulp`.

```bash
npm install -g gulp bower
npm install
```

To install 3rd party libraries:

```bash
bower install
```

To compile the SASS stylesheets etc.:

```bash
gulp
```


# Financial support

This project has been made possible thanks to [Cocreators](https://cocreators.ee) and [Lietu](https://lietu.net). You can help us continue our open source work by supporting us on [Buy me a coffee](https://www.buymeacoffee.com/cocreators).

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/cocreators)
