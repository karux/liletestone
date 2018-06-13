echo -e "The lile skeleton and service model generator requires template files\n"
echo -e "TODO: modify these template files to follow our conventions\n"
mkdir -p ./src/github.com/lileio/lile/template
mkdir -p ./src/github.com/lileio/lile/protoc-gen-lile-server/templates
cp ./tooling/src/github.com/lileio/lile/template/* ./src/github.com/lileio/lile/template/
cp ./tooling/src/github.com/lileio/lile/protoc-gen-lile-server/templates/* ./src/github.com/lileio/lile/protoc-gen-lile-server/templates

