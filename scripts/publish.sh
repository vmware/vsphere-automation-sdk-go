#! /bin/bash
set -e

modulePath=$(echo $CI_COMMIT_TAG | rev | cut -d "/" -f 2- | rev)

git clone -b master --single-branch https://oauth2:${ACCESS_TOKEN}@${DESTINATION_REPO}.git github-repo

mkdir -p github-repo/$modulePath
cd github-repo/$modulePath

# Delete old files of the module.
moduleRepoPath="${DESTINATION_REPO}/$modulePath/"
go list all | grep "$moduleRepoPath" | sed "s#$moduleRepoPath##" | sort -r -t /
rm -rf $(go list all | grep "$moduleRepoPath" | sed "s#$moduleRepoPath##" | sort -r -t /)
find . -type f -maxdepth 1 -delete

cp -r $DESTINATION_REPO/$modulePath/* github-repo/$modulePath/

cd github-repo
git add $modulePath
git commit -m "$CI_COMMIT_MESSAGE"
git tag -a -m "$CI_COMMIT_TAG" $CI_COMMIT_TAG
git push --follow-tags https://oauth2:${ACCESS_TOKEN}@${DESTINATION_REPO}.git HEAD:master $CI_COMMIT_TAG
