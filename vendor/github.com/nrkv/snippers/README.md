# Snippers
Collection of useful Golang snippets for standard library.

## Install

```bash
go get github.com/nrkv/snippers
```

## Examples

### Logger middleware
```go
// Default logger
http.ListenAndServe(":80", snippers.Logger(app.Router, os.Stdout, snippers.DefaultLoggerConfig))
/*
    2016/12/24 18:50:33 | GET    | 200 |        14.616µs | /hello
    2016/12/24 19:03:01 | POST   | 200 |      2.975525ms | /api/files
    2016/12/24 19:03:25 | GET    | 200 |       305.774µs | /debug/pprof/
*/
```

### Response writers
```go
func MyHandler(w http.ResponseWriter, r *http.Request) {
    // Respond with just Status
    snippers.StatusResponse(w, http.StatusOK)
    /*
        HTTP/1.1 200 OK
        Content-Type: text/plain; charset=utf-8
        Content-Length: 0
    */

    // Respond with String and Status
    snippers.StringResponse(w, http.StatusOK, "Done")
    /*
        HTTP/1.1 200 OK
        Content-Type: text/plain; charset=utf-8
        Content-Length: 4

        Done
    */

    // Respond with JSON and Status
    snippers.JSONResponse(w, http.StatusOK, struct{
        Result string
    }{
        Result: "Done",
    })
    /*
        HTTP/1.1 200 OK
        Content-Type: application/json
        Content-Length: 18

        {"Result":"Done"}
    */

    // Respond with XML and Status
    type response struct {
        Result string `xml:"result"`
    }
    snippers.XMLResponse(w, http.StatusOK, response{"Done"})
    /*
        HTTP/1.1 200 OK
        Content-Type: application/json
        Content-Length: 43

        <response><result>Done</result></response>
    */

}
```