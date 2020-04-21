# smc-sock

smc-sock is a simple command line tool for testing
[SMC](https://www.rfc-editor.org/info/rfc7609) connections. It opens an
`AF_SMC` socket and runs either as server or as a client. As client, it
connects to a server, sends a text message over the SMC connection, reads the
reply from the server, and terminates. As server, it waits for a client
connection, reads the message from the client and returns it to the client.

## Usage

You can run `smc-sock` with the following command line arguments:

```
  -a string
        server/client address (default "127.0.0.1")
  -c    run client
  -p int
        server/client port (default 50000)
  -s    run server
  -t string
        text to transfer (default "Hello, world!")
```
