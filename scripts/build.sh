#! /bin/bash
set -e

echo "Building for GIT TAG: $CI_COMMIT_TAG"
modulePath=$(echo $CI_COMMIT_TAG | rev | cut -d "/" -f 2- | rev)

cd $DESTINATION_REPO/$modulePath
echo "Module Path: "
pwd

echo "Setting Go environment variables."
export GO111MODULE=on
export GONOPROXY="github.com/vmware,gitlab.eng.vmware.com"
export GONOSUMDB="github.com/vmware,gitlab.eng.vmware.com"

echo "Clean go cache"
go clean -modcache -cache
echo "Remove old go.sum file"
rm -f go.sum
echo "Building Go Module"
go build ./...
echo "Running go mod tidy"
go mod tidy
