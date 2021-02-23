#! /bin/bash
set -e

BASEDIR=$(dirname "$0")
cd $BASEDIR

param=$1
dryrun=false

if [ param = "--dryrun" ]
then
    dryrun=true
fi

# Create gradle wrapper
cd gradle
gradle wrapper -PmodulePath=github.com/vmware/vmware-infrastructure-sdk-go/modules
cd ..

DEFAULT_ARTIFACTORY="vmcswaggers-docker-local.artifactory.eng.vmware.com"
DEFAULT_ARTIFACTORY_ID="svc.sdk-deployer"

# Login to Docker Repo (Artifactory)
echo "Press enter if you want to use Default Value."

read -p "Docker Repository (Artifactory) [$DEFAULT_ARTIFACTORY]: " ARTIFACTORY
ARTIFACTORY=${ARTIFACTORY:-$DEFAULT_ARTIFACTORY}
echo $ARTIFACTORY

read -p "User [svc.sdk-deployer]: " ARTIFACTORY_ID
ARTIFACTORY_ID=${ARTIFACTORY_ID:-$DEFAULT_ARTIFACTORY_ID}
echo $ARTIFACTORY_ID

read -p "Password: " -s ARTIFACTORY_PASSWORD
echo $ARTIFACTORY_PASSWORD | docker login -u $ARTIFACTORY_ID --password-stdin $ARTIFACTORY

read -p "Docker Image Version: " DOCKER_VERSION
DOCKER_TAG=$ARTIFACTORY/infra_sdk_go:$DOCKER_VERSION

# Build latest docker image with the TAG
docker build --tag=$DOCKER_TAG .
# Push the latest build docker image to Docker Repo (Artifactory)
$dryrun || docker push $DOCKER_TAG
# Logout of Docker Repo (Artifactory)
docker logout $ARTIFACTORY

# Remove local gradle wrapper files
rm -rf gradle/.gradle gradle/gradle gradle/gradlew gradle/gradlew.bat
