DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

export GOPATH=$GOPATH:$DIR

echo $GOPATH

go clean
go install org/x4444/ongo
