#!/usr/bin/env bash

GED_ARCHIVE_NAME="ged-classic-$TRAVIS_OS_NAME-$(janus version -format='TAG_OR_NIGHTLY')"
zip "$GED_ARCHIVE_NAME.zip" ged
tar -zcf "$GED_ARCHIVE_NAME.tar.gz" ged

mkdir deploy
mv *.zip *.tar.gz deploy/
ls -l deploy/

janus deploy -to="builds.etcdevteam.com/go.earthdollar/$(janus version -format='v%M.%m.x')/" -files="./deploy/*" -key="./gcloud-travis.json.enc"
