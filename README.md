URL Tabber
===

A small webapp to remotely open new Google Chrome tabs on macOS

# Installation

1. Grab a binary from the [release page](https://github.com/iwittkau/url-tabber)
2. [Download](https://github.com/iwittkau/url-tabber) the `index.html` template
3. Put the template in some folder
4. Start URL Tabber with:

```bash
$ ./urltabber -p SOMEPORT -t /path/to/folder/with/template -s APPLICATION_SECRET
```

The `-s` flag is optional. A hopfully random secret is generated on startup.

# Expose URL Tabber to the internet

Use [localtunnel](http://localtunnel.github.io) to expose the port from the start command.

```bash
lt --port SOMEPORT
``` 


# Full example

1st Terminal:

```bash
$ ./urltabber -p 8080 -t /template
Generating secret ...
Application secret: CJx562cU78joX04u31r9
Starting to listen on :8080
``` 

2nd Terminal:

```bash
$ lt --port 8080
your url is: https://jolly-octopus-67.localtunnel.me
```

