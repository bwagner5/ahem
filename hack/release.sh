#!/bin/bash
set -euo pipefail


IMG_DIGEST=$(KO_DOCKER_REPO=${RELEASE_REPO} ko publish -B -t ${VERSION} ${RELEASE_PLATFORM} ${BUILD_DIR}/../cmd/ahem)
yq e -i ".image = \"${IMG_DIGEST}\"" charts/ahem/values.yaml
yq e -i ".appVersion = \"${VERSION#v}\"" charts/ahem/Chart.yaml
yq e -i ".version = \"${VERSION#v}\"" charts/ahem/Chart.yaml

HELM_CHART_FILE_NAME="ahem-${VERSION#v}.tgz"
cd ${BUILD_DIR}/../charts
helm package ahem --version ${VERSION}
helm push "${HELM_CHART_FILE_NAME}" "oci://${RELEASE_REPO}"
rm "${HELM_CHART_FILE_NAME}"