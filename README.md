# Unsafe Utils

The unsafe-utils library allows for use and access of lower level Dragonfly tooling (such as
direct packet writing and access to the network sessions) for processes which would otherwise be impossible without using reflection.

_Note: This library is called unsafe for a reason; it uses the Go `unsafe` library to access private fields of Go structs. This behavior is very hacky and parts of this library should only be used if they are not natively possible by any means via Dragonfly/Gophertunnel_

## Importing

You may import Unsafe Utils via the Go CLI

```
go get github.com/bedrock-gophers/living
```

## Examples

You may write a network packet directly to a player via the use of `unsafeutils.WritePacket`

```go
unsafeutils.WritePacket(p, packet.SetActorData{
    ...
})
```
