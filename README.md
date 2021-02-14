# Serverus

Serverus is a experimental gRPC server that should help small teams
to up and running a very basic stream server. The idea behind this is
be able to play aroung gRPC easly and have a server up and running fast.

> No recommended to use in production environment since is an experiment.

## Install

```
go get github.com/asccigcc/serverus
```

```
import "github.com/asccigcc/serverus"
```

## Usage

Serverus provides methods that helps the developer to start easly a gRCP application.
But it is important to understand few constraints before to start running the service.

The developer should know about the following topics:

- Go lang basic.
- gRPC basics you can found more information [here](https://grpc.io/docs/languages/go/quickstart/).


### Preparing Protobuf

Before to integrate Serverus is very important to have defined your protobuf files since those will be used for
your gRPC server. You can find a really good guide in the [here](https://grpc.io/docs/languages/go/quickstart/). 

We are going to use the same quickstart protobuf as example for this documentation.
 
In the example from gRPC documentation you will have a protobuf with the following content:

```
// The greeting service definition.
service Greeter {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

// The request message containing the user's name.
message HelloRequest {
  string name = 1;
}

// The response message containing the greetings
message HelloReply {
  string message = 1;
}
```

After you copiled the proto file you should have two files:

- `helloworld.pb.go`
- `helloworld_grpc.pb.go`


Those files should be imported into the `main.go` from your project (we suggest to keep those files inside a directory like `hello/`).

### Preparing integration

In your main application; after to import the package and protobuf, you can create a server from serverus. The following
guide should help you to have an idea about how should works.

Import your protobuf files:

```
import proto "github.com/my_user/application/hello/"
```

Define `server struct` with the `UnimplementedGretterServer` this configure gRPC to what to do when an endpoint
does not exist.

```
type server struct {
	proto.UnimplementedGretterServer
}
```

Create `RegisterHandlerServer` this function will be used to inject into the server function
so Serverus will recognize the proto.

```
func RegisterHandlerServer(s grpc.ServiceRegistrar) {
	log.Println("Calling RegisterHandlerServer")
	opa.RegisterGretterServer(s, &server{})
}
```

 Create catcher method for `SayHello` request.

```
func (*server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {

	log.Println(req)

	return &proto.HelloReply{Message: "World"}, nil
}
```

In your main function start a Serverus service.

```
func main() {
	server := serverus.NewServerus(":3000")
	server.InitGRPC()
	server.RegisterServer(RegisterHandlerServer)
	server.StartServerus()
}
```