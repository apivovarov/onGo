DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

export GOPATH=$GOPATH:$DIR

echo $GOPATH

go clean
go install github.com/apivovarov/ongo/main github.com/apivovarov/ongo/stringutil
go test github.com/apivovarov/ongo/stringutil
