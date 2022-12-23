#!/bin/bash


echo -e "\e[33m***** Runnning tests...\e[0m"
for d in $(ls -d */ | sort -V)
do
        if [[ -f "$d/main_test.go" ]]; then
                cd "$d" && go test && cd ".."
        fi
done
echo -e "\e[33m***** Done!\e[0m"

