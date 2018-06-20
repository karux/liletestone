if [ $# -eq 0 ]
 then
    echo -e "usage: build.sh {repository-name} {project-name}\n"
    echo -e "    build the binary from  repo/project/\n"
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

cd ./src/github.com/$1/$2
ls -al
export GOOS=linux
export GOARCH=amd64
make build
#go install
