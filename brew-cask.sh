#!/bin/bash

z_help(){
    echo "1, brew-cask.sh list"
    echo "2, brew-cask.sh outdated"
    echo "3, brew-cask.sh upgrade"
}

z_action(){
    for c in $(brew cask list); do
        latest_ver=$(brew cask info $c | sed -n '1p' | sed 's/^'$c'\:[[:space:]]//g')
        local_ver=$(brew cask info $c | sed -n '3p' | sed 's/^\/usr\/local\/Caskroom\/'$c'\///g' | sed 's/[[:space:]](.*)//g')

        if [ "$1" = "list" ]; then
            if [ "$latest_ver" != "$local_ver" ]; then
                echo "++++upgrade++++"
            fi
            echo $c':'
            echo "latest is '$latest_ver', local is '$local_ver'\n"
        elif [ "$1" = "outdated" -a "$latest_ver" != "$local_ver" ]; then
            echo $c':'
            echo "latest is '$latest_ver', local is '$local_ver'\n"
        elif [ "$1" = "upgrade" -a "$latest_ver" != "$local_ver" ]; then
            echo $c':'
            echo "latest is '$latest_ver', local is '$local_ver'\n"
            brew cask reinstall $c
        fi
    done
}

if [ "$1" = "list" ]; then
    echo "========list========"
    z_action "$1"
elif [ "$1" = "outdated" ]; then
    echo "========outdated========"
    z_action "$1"
elif [ "$1" = "upgrade" ]; then
    echo "========upgrade========"
    z_action "$1"
elif [ "$1" = "help" ]; then
    echo "========help========"
    z_help
else
    echo "========error========"
    z_help
fi
