# `d2n`

Go:

* _from_ a Docker image for a program that runs on a port
* _to_ the simplest possible Nomad config that runs that program

In more technical words:

Create a boilerplate Nomad job specification for a given container image.

## Installing

```
go build
sudo cp d2n /usr/local/bin/d2n
```

## Usage

```
$ cd myprogram
$ ls
Dockerfile main.go
$ docker build -t registry/myprogram:v1 .
[...]
$ docker push registry/myprogram:v1
$ cd ..
$ mkdir myprogram-config
$ git init .
$ d2n build myprogram registry/myprogram:v1 80
$ git add *.nomad
$ git commit -am "Initial config"
```

## What next?

Once you've pushed your Nomad job specification to a git repository, the
practice known as _GitOps_ will help you get configs in version control reliably
into a running application.

<!-- TODO: link to a public doc about _GitOps_. -->
