#!/bin/sh
docker-compose -f ${GOPATH}/src/github.com/revolutionchain/charon/docker/quick_start/docker-compose.regtest.yml up -d
sleep 3 #executing too fast causes some errors
docker cp ${GOPATH}/src/github.com/revolutionchain/charon/docker/fill_user_account.sh revo_regtest:.
docker exec revo_regtest /bin/sh -c ./fill_user_account.sh