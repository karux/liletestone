if [ $# -eq 0 ]
 then
    echo -e "usage: createSkeleton.sh {repository-name} {project-name}\n"
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
mkdir -p "./src/github.com/"$1"/"$2
echo -e "creating new project using lile.  project name  github.com/"$1"/"$2

lile new $1/$2
