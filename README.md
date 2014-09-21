# package httpcontext

`go get github.com/nbio/httpcontext`

Flexible per-request contexts for vanilla Go `http.Handler`s. Inspired by, and largely mirrors the interface of [gorilla/context](https://github.com/gorilla/context).

`httpcontext` stores the request context directly in the http.Request by mutating the `request.Body`’s `io.ReadCloser`, thus avoiding the use of a global mutex and per-request teardown.

© 2014 nb.io, LLC
