#!/bin/bash

curl                                                    \
    -d 'client_id=test-app'                             \
    -d 'client_secret=BuBYYDxHHRtyyoE9rFWr8G3JiBgTAHF1' \
    -d 'username=test-user'                             \
    -d 'password=12345'                                 \
    -d 'grant_type=password'                            \
    'http://localhost:8080/realms/master/protocol/openid-connect/token' | jq .access_token | tr -d '"'
    #'http://localhost:8080/auth/realms/master/protocol/openid-connect/token' | jq .access_token | tr -d '"'
     