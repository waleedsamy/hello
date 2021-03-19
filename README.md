# let's go
 GO hello world based on https://golang.org/doc/code.html


Note: user in github url mean your username(I use mine).

### steps
```bash
# update your ~/.bashrc
export GOPATH=$HOME/code/go
export PATH=$PATH:$GOPATH/bin

# now you create a project hello
mkdir -p $GOPATH/src/github.com/user/letsgo

# start write your go code
cd $GOPATH/src/github.com/user/letsgo && touch hello.go
```

### build, install and run

```bash
go install
go run . -config.file let.yaml
curl http://localhost:8080/baz
```

### push to github
```bash
cd $GOPATH/src/github.com/user/letsgo
git init
git add --all && git commit -am "Hello GO"
git remote add origin git@github.com:user/letsgo.git
git push origin master
```
