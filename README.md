# GoSiren

At the moment, there isn't really much to this repository.  It is
simply an attempt to represent the structure of @kevinswiber's
[Siren](https://github.com/kevinswiber/siren) hypermedia format natively in Golang.

The data structures involved have been marked up with struct tags so
that a `SirenEntity` can be marshaled up as JSON using the standard
`encoding/json` package.  Apart from that, there really isn't much
capability natively in this package.  I built this for another project
and recognized that a) there wasn't already an implementation of this
available and b) I might was well release this so others might benefit
from it.

That is all.
