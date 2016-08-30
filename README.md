# now-go
 GO hello world based on https://golang.org/doc/code.html


### steps
```bash
export GOPATH=$HOME/code/github.com/now-go
mkdir -p $GOPATH
export PATH=$PATH:$GOPATH/bin
mkdir $GOPATH/src/github.com/waleedsamy
mkdir -p $GOPATH/src/github.com/waleedsamy/hello
```

### build, install and run

```bash
go get github.com/golang/example/hello
go install github.com/waleedsamy/hello
$GOPATH/bin/hello
> Hello, world.
```

### push to github
```bash
cd $GOPATH/src/github.com/waleedsamy/hello
git init
git add --all && git commit -am "Hello GO"
git remote add origin git@github.com:waleedsamy/now-go.git
git push origin master
```
