# channelswap

This is a small project that i used to explore the viability of having multiple services that share a target channel writer, and ability to HOT-SWAP change the channel assignments at runtime. This would allow a ServiceRegistry to alter the target of the other end of the channel w/o the service even being aware it changed.

The need we had in my organization was to take multiple assigned channels, and have one take a mutually exclusive log on the shared channel, to get full priority of messages. 

The service was an OTA firmware updater, and needed to insure no other packets were written into the wireless mesh simultaneously. 

By created a 'fake' channel, or what is referred to as a `devNullChannel` in the code, the service that needed exclusive rights to the mesh writer, called the ServiceRegistry to swap out shared service(s) with the `devNullChannel`. 

The main.go started (2) services (there is no limit based on the ServiceRegistry rule). The second service `updater` then initiates the new `devNullReader` to consume all the other services packets, while the updater is in progress. The `updater` on it's way out, swaps everything back to normal, and the other services are none-the-less aware. 

The only trick is to be sure your receivers that need to change addresses on the channels are setup as pointers in the receiver. This was my only gotcha of it not working correctly the first time.

## Warning: 
This works great for channel writers. If you need to hot swap a channel reader, odds are it might be in a select block. If so, be sure to interrrupt the blocking channel in select so it can start it's read again on the newly acquired channel address.

# Build instructions:
to run the demo, simple use golang build in  `go get` as such. 
```
go get github.com/dfense/channelswap/main
```
This should checkout the repo, use go.mod to resolve all dependencies (none so far) and build the binary. This will put the artifacts in your 
```
$HOME/go/src
        /bin
```
respectively
