if [ $# -eq 0 ]
 then
    echo -e "usage: genProtoset.sh {repository-name} {project-name}\n"
    echo -e "    proto file must exist at repo/project/project.proto\n"
    exit 2
fi

if [[ -z "$1" ]]
 then
    echo -e "Please provide a project repository name\n"
    exit 1
fi

if [[ -z "$2" ]]
 then
    echo -e "Please provide a project name\n"
    exit 1
fi

FILE=${GOPATH}/src/github.com/$1/$2/$2.proto 
echo "file:"${FILE}

if [[ ! -f $FILE ]]; then
    echo -e "Proto File not found at "${FILE}" filename \n"
    exit 1
fi
echo -e "generating protobuff SET file from ${FILE}\n"
echo -e " project name:  github.com/"$1"/"$2


cd ${GOPATH}/src/github.com/$1/$2
protoc --proto_path=./ \
    --descriptor_set_out=$2.protoset \
    --include_imports \
    ./$2.proto



