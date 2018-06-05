# gopodman
Get specific file or directory from the a github repo without cloning.

How to Install/Get
------------------

    $ go get github.com/praveenkumar/gopodman
    $ export PATH=$PATH:$GOPATH/bin
    $ gopodman -h
     NAME:
        gopodman - podman go client

    USAGE:
        main [global options] command [command options] [arguments...]

    VERSION:
        0.0.1

    COMMANDS:
        ping           display podman status
        podmanVersion  display podman version info
        help, h        Shows a list of commands or help for one command

    GLOBAL OPTIONS:
        --help, -h     show help
        --version, -v  print the version
       

How to Use
----------

First make sure you have podman varlink api listening to respective port and port is allow from firewall
    $ podman varlink tcp::12345 --timeout=0

Check if podman able to ping the remote API.

    $ export PODMAN_VARLINK_URI="tcp:127.0.0.1:12345"
    $ ./out/gopodman ping
    OK
    $ ./out/gopodman podmanVersion
    0.6.1
    $ ./out/gopodman --version
    gopodman version 0.0.1
           
How to build
------------
 
    $ make build
