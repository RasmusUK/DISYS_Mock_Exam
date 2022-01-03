# DISYS_Mock_Exam

## How to run
### Run servers
- Open 1 to 90 terminal windows in DISYS_Mock_Exam/Server/ and run the following command in each to run the servers: 
```console
go run .
```
- You should open atleast two to try and crash one of them by closing the terminal window or clicking Crtl + C on Windows.
- Otherwise the system is not reliable and cannot tolerate a crash-failure of 1 node. 
- The system can tolerate N-1 crashes, where N is the number of nodes (servers) that you run. 
### Run clients
- Open as many terminal windows in DISYS_Mock_Exam/Client/ as you would like and type the following command in each to run the clients:
```console
go run .
```
## How to increment
- To increment, simply type an integer and press enter.
- You can only enter natural numbers.
- If subsequent calls, each must be greater than the prevoius one.
- The call will return the previous value
## Good to know
- Make sure to start all the servers first as the clients will ping the servers and find all the available servers.

