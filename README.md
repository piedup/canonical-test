First build docker image by running `docker build . -t canonical-test`

Run the docker image `docker run -it canonical-test sh`

Try running `dep ensure -v`

The command will keep running for ever. If you stop it via Ctrl-C and rerun `dep ensure -v` it works.

---

Forking [gibson042/canonicaljson-go](https://github.com/gibson042/canonicaljson-go) and deleting .gitmodules "solves" the issue. So it might have something to do with a recursion?
