# Deprecated
Hexya already supported build stand-alone executable file.

* Install
```bash
go get -u github.com/hexya-erp/hexya
cd <projectDir>
hexya project init github.com/halybang/haxey
hexya generate .
go build .
./haxey updatedb -o
./haxey server -o
```
* Tips 
- Copy haxey.yaml to <projectDir> before run hexya generate.
- Change go.mod to use custom/patched hexya.

# haxey
Stand-alone application for hexya - https://github.com/hexya-erp/hexya

# Howto
* Install origin hexya
```bash
go get -u github.com/hexya-erp/hexya
go get -u github.com/hexya-erp/hexya-base
go get -u bitbucket.org/hexya-erp/hexya-addons
```
* Or install custom hexya with some change:
- support json data type
- change PostgreSQL's integer data type to bigint

```bash
mkdir -p $GOPATH/src/github.com/hexya-erp
cd $GOPATH/src/github.com/hexya-erp
git clone https://github.com/halybang/hexya.git hexya
go get -u github.com/hexya-erp/hexya-base
go get -u bitbucket.org/hexya-erp/hexya-addons
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
