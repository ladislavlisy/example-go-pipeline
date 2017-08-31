UPDATE GOPATH, PATH

```shell
$ nano .bash_profile
```

```nano
edit export GOPATH=$HOME/Projekty/GoLangDir

Ctrl+O (Write Out)
ENTER (filename)
Ctrl+X (Exit)
```

INSTALL CASKROOM

```shell
$ brew install caskroom/ cask/ brew-cask
```
INSTALL DELVE

```shell
$ brew install go-delve/delve/delve
```

INSTALL DELVE CERT

```shell
~/Library/Caches/Homebrew/
$ gunzip delve-1.0.0-rc.1.tar.gz
$ gunzip delve-1.0.0-rc.1.tar
$ cd delve-1.0.0-rc.1/scripts/
$ ./gencert.sh
openssl req -new -newkey rsa:2048 -x509 -days 3650 -nodes -config dlv-cert.cfg -extensions codesign_reqext -batch -out dlv-cer
sudo security add-trusted-cert -d -r trustRoot -k /Library/Keychains/System.keychain dlv-cert.cer
```

INSTALL GO

```shell
$ brew install go
```

START DOCKER
```shell
$ docker-machine start default
$ docker-machine env default
$ eval $(docker-machine env default)
$ env | grep DOCKER
```

RUN DOCKER APP IMAGE
```shell
docker run -p 8080:8080 cloudnativego/book-hello
```

TEST DOCKER APP
```shell
curl http://192.168.99.100:8080 
```

INSTALL WERCKER CLI
```shell
brew tap wercker/wercker
brew install wercker-cli
--or--
curl -L https://s3.amazonaws.com/downloads.wercker.com/cli/stable/darwin_amd64/wercker -o /usr/local/bin/wercker
```
