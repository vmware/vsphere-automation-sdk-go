#! /bin/bash
set -e

modulePath=$(echo $CI_COMMIT_TAG | rev | cut -d "/" -f 2- | rev)
mkdir -p $CI_PROJECT_DIR/$DESTINATION_REPO/$modulePath

cd /usr/local/infrasdk_go_gradle
./gradlew prepareGithubBindings -PmodulePath=$modulePath
