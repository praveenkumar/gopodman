# gopodman
Podman varlink API consumer using go.

Works with go v>1.9 .

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
        [...]
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
    
    $ ./out/gopodman ps
    CONTAINERID                                                        IMAGE                              COMMAND            CREATED                STATUS                      NAME
    d5cf67ec0aea59145f7bac7e0f87ab5406229533a613b6ec3c7047e32c5bae0b   docker.io/library/busybox:1.28.4   /bin/sleep 10000   Up About an hour ago   Up 55m55.270374207s ago     relaxed_leakey
    fc694d68cfe5bb8e8c51d9b7d32f50ea86d5d74322e86037782c83eca0d2ad6e   docker.io/library/busybox:1.28.4   /bin/sleep 10000   Up 2 hours ago         Up 1h40m14.559023759s ago   kind_wescoff


           
How to build
------------

    $ make
