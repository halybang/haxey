# haxey
Stand-alone application for hexya - https://github.com/hexya-erp/hexya

# Howto
* Install hexya
```bash
go get -u https://github.com/hexya-erp/hexya
go get -u github.com/hexya-erp/hexya-base
#go get -u bitbucket.org/yourname/hexya-addons
```
* Clone haxey and edit config
```bash
# git clone https://github.com/halybang/haxey.git $GOPATH/src/github.com/halybang/haxey
# cd $GOPATH/src/github.com/halybang/haxey
# vi config/config.go => add user's addons here.
```
* Build hexya
```bash
hexya generate
```
* Build and run haxey
```bash
go build . && ./haxey updatedb && ./haxey
```
