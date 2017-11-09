First build docker image by running `docker build . -t canonical-test`

Run the docker image `docker run -it canonical-test sh`

Try running `dep ensure -v`

The command will keep running for ever. If you stop it via Ctrl-C and rerun `dep ensure -v` it works.