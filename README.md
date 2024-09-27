## About

This is a coding challenge for a company I'm interviewing with. I decided to learn Go to do this, as it's been a language I've been wanting to learn for a while now. 

This code allows you to provide a YAML configuration file to a binary, which will then perform health checks on the configured health check items.

## Usage

- Grab your configuration file
- Download a binary that matches your OS/arch from [here](https://github.com/arkag/healthcheck/releases)
- Make it executable (MacOS or Linux)

```
./healthcheck -f some_filename.yml
```

You should see some output like so:

```
www.fetchrewards.com has 0% availability percentage
fetchrewards.com has 0% availability percentage
www.google.com has 66% availability percentage
duckduckgo.com has 50% availability percentage
```

## Other Options

If you'd like to build this yourself:

- [Install Go](https://go.dev/doc/install)
- `git clone https://github.com/arkag/healthcheck.git`
- `cd healthcheck`
- `go get`
- `go build`
- `go test`
- `./healthcheck -f some_filename.yml`

If you'd like to run this using Go directly:

- [Install Go](https://go.dev/doc/install)
- `git clone https://github.com/arkag/healthcheck.git`
- `cd healthcheck`
- `go run healthcheck.go -f some_filename.yml`