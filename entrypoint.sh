#!/usr/bin/env bash

set -e
set -u
set -o pipefail

if [ -n "${PARAMETER_STORE:-}" ]; then
  export USER_NUXEO="$(aws ssm get-parameter --name /${PARAMETER_STORE}/novedades_api/db/username --output text --query Parameter.Value)"
  export PASSWORD_NUXEO="$(aws ssm get-parameter --with-decryption --name /${PARAMETER_STORE}/novedades_api/db/password --output text --query Parameter.Value)"
fi

exec ./main "$@"