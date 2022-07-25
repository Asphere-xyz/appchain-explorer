#!/usr/bin/env bash
if ! [[ -e /usr/ready ]]
then
  echo "Replacing envs in the JS folder"
  for file in $(find /explorer/web/build/static/js/ -regex '.*\.js'); do
    echo "Working on file $file..."
    envsubst '${REACT_APP_DEFAULT_NETWORK} ${REACT_APP_API_ENDPOINT}' < $file > $file.tpm
    mv $file.tpm $file
  done
  touch /usr/ready
else
  echo "Files are already prepared"
fi
exec /explorer/node