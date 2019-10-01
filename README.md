# Sanctuary

## Usage

Sanctuary offers a simple solution to deploy pre-configured services and balancers that communicate over http (later rpc will also be available).

Easy setup of a new service and adding new routes to the service. Also handle the database connection when given the login information and ip and port.
Currently the service as well as the balancer have a small interface that is already available as a http api. I'm currently working on a ui for the balancer that can then be used to manage the services
and send notifications and commands over the wire.

You may also use the service without a balancer. This will allow you to create a http api in a very sh\*\*ort time and have it deployed in a fast manner.

## How to setup

### Service

    service := service.NewService("#001", "login", "A login service that handles login for users", "3440")
    service.ActivateHTTPServer()
    service.AddHTTPRoute("/route", Function)
    service.StartHTTPServer()

    Creating a service with a name, type, description and port
    Activation a http-server and add a new route to it.
    Start the http server on the given port and wait requests.

    --

### Balancer

    balancer := balancer.NewBalancer("My Balancer", "3400")
    balancer.Setup()
    balancer.Start()

    This code will create a new balancer on port 3400.
    The balancer has some preinstalled routes that can be used to register and query servies:
        * Register | /register - Registers a new service to the balancer
        * GetService | /service - Get the optimal servie by its type
        * GetServices | /servies - Get all servies by a type

### Database

    service.SetDatabaseInformation("127.0.0.1", "3306", "mysql", "username", "password", "users")
    service.PrepareQuery("SELECT * From ?")

    Adding support for a database and setup the connection with the given parameters
    The service can also handle the preperation of queries and executes them. (currently in development)

### Contribute

Because this is currently the site of a rewrite... there is not really anyway you could contribute. As soon as the version is up to date again and everything has been rewritten, I will add a section that will explain how to contribute.

Thanks for your interest

### Shoulders

- The golang programming language
- The golang standart library

### Testing

Currently the api for balancer and service are tested with a postman test-suit. I will later switch to a golang based testing framework to make testing easier and faster.
