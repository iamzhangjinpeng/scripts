#!/bin/bash
# login ssh
# https://github.com/iamzhangjinpeng/scripts

user="username"
password="password"

serverList=(
"192.168.1.7"
"192.168.1.8"
"192.168.1.9"
"192.168.1.0"
)

login_ssh(){
    command="ssh "$user"@"${serverList[$(($1-1))]}
    echo $command
    echo

    expect -c "
        spawn $command
        set timeout -1
        expect {
            \"*continue connecting (yes/no)? \" {
                send \"yes\r\"
                exp_continue
            }
            \"*password: \" {
                send \"$password\r\";
            }
        }
        interact
    "

    return
}

secure_shell(){
    printf "your select is: "
    read your_select

    case $your_select in
        1)  login_ssh $your_select
            return
        ;;
        2)  login_ssh $your_select
            return
        ;;
        3)  login_ssh $your_select
            return
        ;;
        4)  login_ssh $your_select
            return
        ;;
        q)  echo "Bye-Bye !!!"
            return
        ;;
        *)  echo "your select error !!!"
            echo "please select again !!!"
            secure_shell
            return
    esac
}

echo "please select:"
length=${#serverList[@]}
for((i=0;i<$length;i++))
do
    echo $(($i+1))"," ${serverList[$i]}
done
echo "q, quit"
echo

secure_shell
