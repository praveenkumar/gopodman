# gopodman
Podman varlink API consumer using go.

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

    $ ./out/gopodman listImages
    REPOSITORY                  TAG      IMAGEID                                                            CREATED       SIZE
    docker.io/library/busybox   1.28.4   8c811b4aec35f259572d0f79207bc0678df4c736eeec50bc9fec37ed936a472a   2 weeks ago   1.36MB
    docker.io/library/busybox   latest   8c811b4aec35f259572d0f79207bc0678df4c736eeec50bc9fec37ed936a472a   2 weeks ago   1.36MB

           
How to build
------------

    $ make build
