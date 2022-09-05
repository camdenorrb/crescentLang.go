#!/bin/sh
set -- sh "$(which dind)" "$@"
exec "$@"