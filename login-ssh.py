#!/usr/bin/python2
# -*- coding: utf-8 -*-
# login ssh
# https://github.com/iamzhangjinpeng/scripts

import pexpect

user = "username"
password = "password"
serverList = [
    "192.168.1.7",
    "192.168.1.8",
    "192.168.1.9",
    "192.168.1.0",
]

def secure_shell():

    serverListInfo = "please select:\n" + "".join([(str(index + 1) + '. ' + val + '\n') for index, val in enumerate(serverList)])
    serverListInfo += "q. quit\n"
    print serverListInfo

    your_select = raw_input("your select is: ").strip()
    if your_select == 'q':
        print "Bye-Bye !!!"
        return None

    while your_select not in ['1', '2', '3', '4']:
        print "your select error !!!"
        print "please select again !!!"
        your_select = raw_input("your select is: ").strip()
        if your_select == 'q':
            print "Bye-Bye !!!"
            return None

    cmd = "ssh " + user + "@" + serverList[int(your_select) - 1]
    print cmd

    try:
        child = pexpect.spawn(cmd)
        info = ["continue connecting (yes/no)?", "password:"]
        index = child.expect(info)
        if index == 0:
            child.sendline("yes")
            child.expect("password:")
        child.sendline(password)
        child.interact()
    except Exception as e:
        print "pexpect error: %s" % str(e)


if __name__ == "__main__":
    secure_shell()
