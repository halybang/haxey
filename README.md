# haxey
Stand-alone application for hexya - https://github.com/hexya-erp/hexya

# Howto
* Install hexya
```bash
go get -u https://github.com/hexya-erp/hexya
go get -u github.com/hexya-erp/hexya-base
go get -u bitbucket.org/yourname/hexya-addons
```
* Clone haxey and edit config
```bash
# git clone https://github.com/halybang/haxey.git $GOPATH/src/github.com/halybang/haxey
# cd $GOPATH/src/github.com/halybang/haxey
```
* Select addons to use (Module must be exist in both file hexya.yaml and config.go)
```bash
# vi config/config.go
# vi hexya.yaml
```
* Change configuration
```bash
# vi hexya.yaml
```
* Generate ORM code
```bash
$GOPATH/bin/hexya generate
```
* Build haxey, update symlink and update database
```bash
go build -i . && ./haxey link && ./haxey updatedb
```
* Run haxey
```bash
./haxey [server]
```
