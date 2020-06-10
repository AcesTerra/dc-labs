User guide
==============================

This document teaches how to use REST API to send jobs to workers.


### REST API
This is going to be the second phase of design and implementation.
On this phase you are adding 3 new components that will start making more sense as a distributed system
- Controller
- Scheduler
- Worker

Your project will be divided on packages with very descriptive names where each system's component will be implemented.
Below you can see the details of each package and requirements for this partial:

- `api/`
  - From now, all request must be token-based authenticated
  - **Endpoint:** `/status` - Overall system status and logged user details. Also, print workers details (name, status and usage percentage)
  - **Endpoint:** `/status/<worker>` - Per worker details:  (name, tags, status and usage percentage)
  - **Endpoint:** `workloads/test` - This endpoint will trigger an initial end-to-end test from the `api` to a running `worker`

- `controller/`
  - Basic overall system and per node data store (it can be in-memory or a key-value datastore)
  - Request pre-validation before sending to `scheduler`
  - Controller will create a message-passing server for its interaction with workers

- `scheduler/`
  - Basic workloads scheduling that will be based on node's tags and # of running workloads
  - Scheduler is calling workers through RPC

- `worker/`
  - Standalone component with initial `test` RPC function.
  - Worker's command line will be as follows:
    - `./worker --controller <host>:<port> --node-name <node_name> --tags <tag1>,<tag2>...`


**Documentation**
- A detailed arquitecture document will be required for this initial phase in the [architecture.md](architecture.md) file. Diagrams and charts can be included$
- A detailed user guide must be written in the [user-guide.md](user-guide.md) file. This document explains how to install, configure and use your system.


Test Cases (from console)
-------------------------
- [Project's First Phase Test Cases](../second-partial/#test-cases-from-console)

- **Node Status**
```
$ curl -H "Authorization: Bearer <ACCESS_TOKEN>" http://localhost:8080/status/<worker>
{
        "Worker": "Worker-name",
        "Tags": "tag1,tag2,tag3",
        "Status": "Running",
        "Usage": "50%"
}
```

- **Execute Workload**
```
$ curl -H "Authorization: Bearer <ACCESS_TOKEN>" http://localhost:8080/workloads/test
{
        "Workload": "test",
        "Job ID": "1",
        "Status": "Scheduling",
        "Result: "Done in Worker: <worker_name>"
}
```

- **Execute Filter Workload**
```
$ curl -F 'data=@path/to/local/image.png' -F 'workload-id=my-filters' -F 'filter=grayscale' -H "Authorization: Bearer <ACCESS_TOKEN>" http://localhost:8080/workloads/filter
{
        "Workload": "test",
        "Job ID": "1",
        "Status": "Scheduling",
        "Result: "Done in Worker: <worker_name>"
}
```
