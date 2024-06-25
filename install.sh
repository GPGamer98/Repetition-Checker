#!/bin/bash

YELLOW='\033[0;33m'
GREEN='\033[0;32m'
R='\033[0m'

printf "Choose your language and enter the respective number:\nEN = 1\nIT = 2\n"
read lang

echo -e "${YELLOW}Installing the app...${R}"

mkdir -p "/tmp/rep-checker/"

curl -sL "https://github.com/GPGamer98/Repetition-Checker/releases/latest/download/CogitoErgoVet-linux-amd64.zip" -o "/tmp/rep-checker/latest.zip"

unzip -qq "/tmp/rep-checker/latest.zip" -d "/tmp/rep-checker/"

rm "/tmp/rep-checker/latest.zip"

mv /tmp/rep-checker/CogitoErgoVet-linux-amd64 /tmp/rep-checker/CogitoErgoVet

if [[ $lang == "1" ]]; then
    mv /tmp/rep-checker/CogitoErgoVet /home/$(whoami)/Desktop/
    chmod +x "/home/$(whoami)/Desktop/CogitoErgoVet"
elif [[ $lang == "2" ]]; then
    mv /tmp/rep-checker/CogitoErgoVet /home/$(whoami)/Scrivania/
    chmod +x "/home/$(whoami)/Scrivania/CogitoErgoVet"
else
    echo "ERROR! Please select a valid language."
    echo "Press any key to close..."
    read
    exit 1
fi

echo -e "${GREEN}Installation completed. The app should be on the desktop${R}"
echo "Press any key to close..."
read
