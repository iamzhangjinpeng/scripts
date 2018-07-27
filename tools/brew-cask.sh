#!/bin/bash

z_help(){
    echo "1, brew-cask.sh list"
    echo "2, brew-cask.sh outdated"
    echo "3, brew-cask.sh upgrade"
}

z_action(){
    for c in $(brew cask list); do
        latest_ver=$(brew cask info $c | sed -n '1p' | sed 's/^'$c'\:[[:space:]]//g' | sed 's/[[:space:]](.*)//g')
        local_ver=$(brew cask info $c | sed -n '3p' | sed 's/^\/usr\/local\/Caskroom\/'$c'\///g' | sed 's/[[:space:]](.*)//g')

        if [ "$1" = "list" ]; then
            if [ "$latest_ver" != "$local_ver" ]; then
                echo "\033[36m'$c':\033[0m"
                echo "\033[36mlatest is '$latest_ver', local is '$local_ver'\033[0m\n"
            else
                echo $c':'
                echo "latest is '$latest_ver', local is '$local_ver'\n"
            fi
        elif [ "$1" = "outdated" -a "$latest_ver" != "$local_ver" ]; then
            echo "\033[36m'$c':\033[0m"
            echo "\033[36mlatest is '$latest_ver', local is '$local_ver'\033[0m\n"
        elif [ "$1" = "upgrade" -a "$latest_ver" != "$local_ver" ]; then
            echo "\033[36m'$c':\033[0m"
            echo "\033[36mlatest is '$latest_ver', local is '$local_ver'\033[0m\n"
            brew cask reinstall $c
        fi
    done
}

if [ "$1" = "list" ]; then
    echo "========list========\n"
    z_action "$1"
elif [ "$1" = "outdated" ]; then
    echo "========outdated========\n"
    z_action "$1"
elif [ "$1" = "upgrade" ]; then
    echo "========upgrade========\n"
    z_action "$1"
elif [ "$1" = "help" ]; then
    echo "========help========\n"
    z_help
else
    echo "========error========\n"
    z_help
fi
