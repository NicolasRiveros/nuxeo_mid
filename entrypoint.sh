#!/usr/bin/env bash

set -e
set -u
set -o pipefail

if [ -n "${PARAMETER_STORE:-}" ]; then
  export USER_NUXEO="$(aws ssm get-parameter --name /${PARAMETER_STORE}/nuxeo_mid/nuxeo/user --output text --query Parameter.Value)"
  export PASSWORD_NUXEO="$(aws ssm get-parameter --with-decryption --name /${PARAMETER_STORE}/nuxeo_mid/nuxeo/password --output text --query Parameter.Value)"
fi

exec ./main "$@"