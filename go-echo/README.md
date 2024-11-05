# go-echo


### Build image
```
$ export TAG=1.0.6
$ docker build -t ghcr.io/rajiv-k/go-echo:$TAG -f Dockerfile.echo .
```

### Run

```
$ docker run -itd --name go-echo -p 9000:9000 -e PORT=9000 -e HOST_ID=$(hostname) ghcr.io/rajiv-k/go-echo:1.0.6
```
