# go simple webapp template

This template is intended to show how one can easily setup
a tls encrypted webserver in go including routing, html templating and assets.

Finally all resources are compiled into the binary, to be ready to deploy.

As you can see from the go.mod file there are some third party dependencies
bound to this project.

Here are some reasons why I prefer to use them. 

## Routing (gorilla mux)
As router I chosse the gorilla.mux, which is pretty famous.
**https://github.com/gorilla/mux**

It is easy to use, allows short, expressive configuration and has
good regular expression support for route parameters.

## Assets
Assets are embedded into the built binary at compile time using go standard library package "embed".

## Logging (logrus)
The reasons for logrus are pretty straight forward.
I like it ...it has colors ... sometimes.... and ... it's used by docker!! ;-)
No, seriously. I've seen it can be used in production for logging to syslog,logstash and more.

## Getting started
* call ```make build```
* You will find the server binary in **.build-assets**

### local dev
For local development I suggest to use the **.env** file
which sets the assets to be load from local filesystem instead
of looking in the generated go files. 

