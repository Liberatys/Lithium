# Service 

A service is the code running on a maschine that is handling one functionality of your system.
An example could be a login or user service. This service is only concerned with login in users or storing their direct information.

## Configuration

A service needs the following elements to operate:
* Name (unique)
* Current-IP
* Port (For http server or for connection over rpc)
* Balancer-Information (if a balancer is used)
* Database-Information (if one is used)
* Server-type ("Account-Management")
* Short-Description ("What is the purpose of the service ? ")
* RPC / HTTP -- Information