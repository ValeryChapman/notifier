#!/bin/bash

NO_COLOR='\033[0m'
GREEN='\033[0;32m'
PURPLE='\033[0;35m'

SERVER_ADDRESS='http://127.0.0.1'

testCreateNotification() {
  echo -e "[$GREEN START$NO_COLOR –$PURPLE testCreateNotification $NO_COLOR]" \
          "\n[$GREEN RESPONSE $NO_COLOR]\n"
  http POST "$SERVER_ADDRESS/api/notifications/" \
    to:="[\"$1\"]" \
    subject:="\"$2\"" \
    body:="\"$3\""
  echo -e "[$GREEN FINISHED $NO_COLOR]\n\n"
}

testGetNotification() {
  echo -e "[$GREEN START$NO_COLOR –$PURPLE testGetNotification $NO_COLOR]" \
          "\n[$GREEN RESPONSE $NO_COLOR]\n"
  http "$SERVER_ADDRESS/api/notifications/$1"
  echo -e "[$GREEN FINISHED $NO_COLOR]\n\n"
}

testGetNotifications() {
  echo -e "[$GREEN START$NO_COLOR –$PURPLE testGetNotifications $NO_COLOR]" \
          "\n[$GREEN RESPONSE $NO_COLOR]\n"
  http GET "$SERVER_ADDRESS/api/notifications/?limit=$1&offset=$2&to=$3"
  echo -e "[$GREEN FINISHED $NO_COLOR]"
}

# Test create a notification
testCreateNotification ily@xenc.ru "Test subject" "<h1>Test body</h1>"

# Test get a notification by Id
#testGetNotification "fbafedb311393d601eb269616e4d57895b5e67c5adfaeba03f74d77aeb70aa7b"

# Test get notifications by filters
testGetNotifications 10 0 ily@xenc.ru
