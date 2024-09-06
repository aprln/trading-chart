# Trading Chart Service

### Note on the solution

- The solution follows the popular clean code structure 
where the handler package is the equivalent of the controller 
layer. All external structures such as requests and events 
are converted to internal models to work with in the service 
and repository layers. 
This structure promotes modularity, testability, and extensibility. 
- Due to time constraint, authentication, error handling, logging, and
  comprehensive testing are not fully implemented or implemented at a very basic level.

### How to run the service


- Start GRPC server:

    `make run-server-docker`


- Start websocket listener:

    `make run-ws-listener-docker`

### Room for improvement

- Tests need to be added
- Error handling could be improved
- Could use a query builder library for DB queries 

