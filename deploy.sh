echo build solPack ...
go build 
echo complete : now at "$(pwd)"
echo "  $(ls -al --color | grep solPack*)"
echo move solPack ... "=> $(echo $GOPATH  /bin)"
mv solPack* $GOPATH/bin
echo complete : now at $GOPATH/bin/
echo "  $(ls -al --color $GOPATH/bin | grep solPack*)"
