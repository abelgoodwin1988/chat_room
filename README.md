# Chat Service

> Simple multi-client chat server

Implementation of a tcp chat server and client. There are separate server and client `.go` files and in turn binaries. This package uses the go standard library [net package](https://golang.org/pkg/net/) to create a tcp listener for the server which accepts tcp dialing from clients.

The clients are managed in a simple custom `client struct` slice. They require a little more work for accurately closing client connections from the server.

These applications make use for channels for blocking and managing the closing of connections as well as a means of communicating messages from one client, through the server then iterating over clients and writing the message to all clients.

## Getting Started

For creating and running the chat server binaries, see [here](#creating-the-binaries).
For running the chat server and chat client(s) see [here](#developing).

### Prerequisites

go 1.12+ downloaded & installed
[download go](https://golang.org/dl/)

---

### Installing & Running

Note that port 23 is the default port for the server. If the machine running the server has port 23 occupied, it will fail to start.

#### Creating the Binaries

To create the binaries run the following commands from the project root which will create a `chat_server` and `chat_client` in the root directory of your project.

```bash
~/chat_service/> go build -o chat_server server/server.go && go build -o chat_client client/client.go
```

#### Running the Server & Client

From the project root, assure that the built binaries have executable permissions, in unix based systems:

```bash
# validate they have execution permiss
~/chat_service/> ls -l chat_server chat_client
# If they do not, give them execution permission
~/chat_service/> chmod +x chat_server chat_client
```

To start the server, within the project root, execute the server binary:

```bash
~/chat_service/> ./chat_server
```

To start a client, within the project root, execute the client binary:

```bash
~/chat_service/> ./chat_client
```

### Developing

#### Running the Server and Clients

In a terminal within the root directory run a server:

```bash
~/chat_service/> go run server/server.go
```

In n+ terminals within the root directory, run n+ times:

```bash
~/chat_service/> go run client/client.go
```

<!-- ## Running the tests -->

## Built With

* [Go](https://golang.org/)

## Contributing

Please use the standard gitflow process, be sensical and kind in any commenting.

## Authors

* **Hamp Goodwin**- [GitHub](https://github.com/abelgoodwin1988)
