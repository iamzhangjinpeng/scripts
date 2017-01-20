#!/bin/bash
# brew cask upgrade
# https://github.com/iamzhangjinpeng/scripts

for c in $(brew cask list); do
    latest_ver=$(brew cask info $c | sed -n '1p' | sed 's/^'$c'\:[[:space:]]//g')
    local_ver=$(brew cask info $c | sed -n '3p' | sed 's/^\/usr\/local\/Caskroom\/'$c'\///g' | sed 's/[[:space:]](.*)//g')

    if [ "$latest_ver" != "$local_ver" ]; then
        echo $c':'
        echo "latest is '$latest_ver', local is '$local_ver'\n"
        brew cask reinstall $c
    fi
done

