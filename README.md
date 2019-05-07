# Chat Service

> Simple multi-client chat server

Implementation of a tcp chat server and client. There are separate server and client `.go` files and in turn binaries. This package uses the go standard library [net package](https://golang.org/pkg/net/) to create a tcp listener for the server which accepts tcp dialing from clients.

There are some issues with the windows binaries. I suspect it's due to either the bufio buffer size, line terminators, or control characters in the string format for clients. The binaries will execute, but the messages received/printed by the clients will be malformed.

## Getting Started

For creating and running the chat server & chat client binaries, see [here](#creating-the-binaries).
For running the chat server and chat client(s) see [here](#developing).

### Basic Usage

Per the instructions below, build and start a chat server.

Per the instructions below, build and start any number of clients. When a client is started you will be asked to give a username which you must provide to continue with the application. At current the program is blocked and only able to add one user at a time. After setting the username(s) for you client(s) you can send meessages to the server and other clients by typing in the terminal and indicating a message to be sent with a return character (enter/return).

### Prerequisites

go 1.12+ downloaded & installed
[download go](https://golang.org/dl/)

---

### Installing & Running

Note that port 23 is the default port for the server. If the machine running the server has port 23 occupied, it will fail to start.

#### Creating the Binaries

To create the binaries run the following commands from the project root which will create a `chat_server` and `chat_client` in the root directory of your project.

```bash
# For Unix Systems
~/chat_service/> go build -o chat_server server/server.go && go build -o chat_client client/client.go
# For Windows System
~/chat_service/> go build -o chat_server.exe server/server.go && go build -o chat_client.exe client/client.go
```

#### Running the Server & Client Binaries

Summary

> Build the server and client binaries. Run the server binary and then any number of client binaries.

---

From the project root, assure that the built binaries have executable permissions, in unix based systems:

```bash
# validate they have execution permiss
~/chat_service/> ls -l chat_server chat_client
# If they do not, give them execution permission
~/chat_service/> chmod +x chat_server chat_client
```

To start the server, within the project root, execute the server binary:

```bash
# Unix
~/chat_service/> ./chat_server
# Windows
~/chat_service/> ./chat_server.exe
```

To start a client, within the project root, execute the client binary:

```bash
# Unix
~/chat_service/> ./chat_client
# Windows
~/chat_service/> ./chat_client.exe
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
