echo build solidityPack ...
go build 
echo complete : now at "$(pwd)"
echo "  $(ls -al --color | grep solidityPack*)"
echo move solidityPack ... "=> $(echo $GOPATH  /bin)"
mv solidityPack* $GOPATH/bin
echo complete : now at $GOPATH/bin/
echo "  $(ls -al --color $GOPATH/bin | grep solidityPack*)"
