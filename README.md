# Type of Error on TCP client server

This simple client-server app demonstrate common errors on TCP connection

_Start Server_

```
make server
```

_Start Client_

```
make client
```
and put some text as message request to the server

## I/O timeout

commonly occured in server or client during read or write process. 

### Server (Write Timeout)

This error occured when the server try to write request to client connection.

Sample Error from server

```
error writing to connection: write tcp [::1]:9292->[::1]:58243: i/o timeout
```

How to make this error? 

This line code `conn.Write([]byte(response))` takes more than 100 micro second.
So then by setting deadline to 100μs, like `conn.SetWriteDeadline(time.Now().Add(100 * time.Microsecond))`,
the error message will appear from server side and message `EOF` will be on client side.

Alternatively, put `time.Sleep(110 * time.Microsecond)` and deadline like 200 microsecond.

The cause?
- Posibility the error happend because of bad connection.
- server writing time too long
- message response too long

Solution?



### Server (Read Timeout)

This error occured when the server try to read request from client connection.

Sampe Error:

```
error reading from connection: read tcp [::1]:9292->[::1]:57979: i/o timeout
```

How to demo the error?
- Option 1
    - set deadline for reading connection like `conn.SetReadDeadline(time.Now().Add(3 * time.Second))` on the server
    - then put time sleep in the client
    - or you can just start client and hold couple seconds to enter the messaeg request
- Option 2
    - we can check the duration conn.Read and set the read deadline less than the duration


The cause?
- the connection estabilish but there is no message or request (client write request too long)
- problem of server slower than the read deadline

Solution?
