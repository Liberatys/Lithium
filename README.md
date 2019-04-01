# Lithium

### Repository for the little microservice helper

Welcome to my little library I'm currently developing in order to have an easier time
in developing microservice architectures.


Lithium wants to solve the problem of service-discovery and easy load balancing.


### How does it work?
<br>
<br>

#### Service
For your service, you use the service Structure that is provided by the library. 
Now you can just plug in your routes and methods into the service-router.

    UserDataService := Service.NewService("Storage-Service #1", "Storage", "The service is writing given usedata into the database", "8002", true)
    UserDataService.Initialize("127.0.0.1", "8000", "/serviceDiscovery")
    UserDataService.AddRoute("/writeToDatabase",MyHTTPHandler)
    
    
After you have plugged in all your methods/endpoints, you can run the local http-server and the network discovery.

    TestService.OpenForConnection()
    
    
Your new Service, will now connect to the specified Load-Balancer IP, and will register with the server, if the balancer is running.

<br>
<br>
<br>

#### Balancer/Service-Registry


Before we can run our services, we have to create the load-balancer.
This is an easy step if you want to use the given package.


You just create the new Balancer:
    
    mainServer := Balancer.CreateNewBalancer("8000")
    mainServer.SpinUpServer()

    
You set the port that the local http-server should run on, after that you spin up the serverr.
Every balancer implements two endpoints that serve for service-discover and apigateway

    balancer.AddRoute("/apigateway", APIGateWay)
    balancer.AddRoute("/serviceDiscovery", DiscoverService) 
        
    
The DisocverService is handling the registration of new services that connect to the server via the implemented Pitcher.
The APIGateWay is running the distribution of the traffic to the services that can handle the load.

The only thing you have to do now, is to send a url and the service-type that should be handled.

    {"destination":"/api/endpoint","destinationtype":"Storage"}
    
The balancer will, based on the last service-speed-test chose the service with the best stats and will redirect the API, to this service.

If a service is not responding, it gets flagged and is no longer receaving traffic.
      
      
### Additional Information about Lithium

<b>Current-Version: v0.1.0.2<br>
Latest-Update: 06.03.2019
</b>

#### Next Steps

* Implementation of a distributed Database-System
    * Native Service for Cassandra and MYSQL
        * Logging for database-migration. 
* MiddleWare for Logging
* Settings loaded from a settings file      
        
        
        
        
#### Version-System
So to keep track of things threw the version system, I created a system that would fit my needs.</br>
The v in the version stands for : Version;</br>
The first Part of the version is the release-versions.</br>
The second displays rewrites and big changes.</br>
The third is addressing bug-fixes.</br>
The fourth and last is adressing the new features and maintenance.</br>       
