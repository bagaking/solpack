echo build solpack ...
go build 
echo complete : now at "$(pwd)"
echo "  $(ls -al --color | grep solpack*)"
echo move solpack ... "=> $(echo $GOPATH  /bin)"
mv solpack* $GOPATH/bin
echo complete : now at $GOPATH/bin/
echo "  $(ls -al --color $GOPATH/bin | grep solpack*)"
