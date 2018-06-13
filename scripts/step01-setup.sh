echo -e "create ./bin ./model ./src ./tooling \n"
mkdir bin
mkdir model
mkdir src
mkdir tooling

## INSTALL TOOL SOURCE INTO /tooling subdirectory
echo -e "install tooling into ./tooling subdirectory\n"
cd tooling
export GOPATH=${PWD}
export GOBIN=${GOPATH}/bin
echo -e "build grpc tool - wire protocol"
go get -u google.golang.org/grpc
echo -e "build lile tool - microservice framework"
go get -u github.com/lileio/lile/...
echo -e "build protoc-gen-go tool - protobuff compiler"
go get -u github.com/golang/protobuf/protoc-gen-go
echo -e "build grpcurl tool - curl for grpc"
go get -u github.com/fullstorydev/grpcurl
echo -e "grpc command line interface tool - cobra"
go get -u github.com/fiorix/protoc-gen-cobra
go get -u github.com/gogo/protobuf/protoc-gen-gofast
echo -e "regular cobra "
go get -u github.com/spf13/cobra/cobra
echo -e "vgo vendoring tool"
go get -u golang.org/x/vgo
echo -e "dep vendoring tool"
go get -u github.com/golang/dep/cmd/dep
echo -e "installed tool listing\n"
ls ${GOBIN}

## PRINT NEXT STEP TO USER
echo -e "\n\nCopy model.proto file into the /model subdirectory\n"
echo -e "Project source will be ./src/github.com/{repo-user}/{repo-name}"
