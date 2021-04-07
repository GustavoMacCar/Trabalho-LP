/*
Networking module for the grid computing part
Amélia O. F. da S. - 190037971
*/
package main

import (
    "net"
    "fmt"
    "log"
    "encoding/binary"
    "sync/atomic"
    "time"
)

/*
Protocol specification:
C = Client S = Server

C starts connection
S sends one byte (state)
IF state == 0:
    S is occupied. Try another peer.
ELSE:
    C sends four bytes (uint32 size)
    C sends size*4 bytes ([size]uint32)
*/

var processing=uint32(0)

func serve(port string){
    ln,err := net.Listen("tcp",port)
    if err!=nil {
        fmt.Println("NET: Couldn't start listening on port ",port)
        log.Fatal(err)
    }
    fmt.Println("NET: Listening on port ",port)
    for {
        conn,_ := ln.Accept()
        //conn.SetDeadline(0)
        go manageConnection(conn)
    }
}

func manageConnection(conn net.Conn){
    /*
    If we're not processing, we swap into "processing" and continue.
    Otherwise, if the swap doesn't happen, write "0" to the
    stream and close
    */
    if !atomic.CompareAndSwapUint32(&processing,0,1){
        conn.Write([]byte{0})
        conn.Close()
        return
    }
    conn.Write([]byte{1})
    //Then we receive the size
    var s uint32
    binary.Read(conn,binary.BigEndian,&s)
    fmt.Println("NET: Received processing request - size = ",s," lines. Address = ",conn.RemoteAddr())
    //And the addresses
    b := make([]uint32,s)
    binary.Read(conn,binary.BigEndian,b)
    /*
    Process the tree data.
    */
    time.Sleep(time.Second)
    conn.Close()
    fmt.Println("NET: Processing request fulfilled.")
    atomic.StoreUint32(&processing,0)
}