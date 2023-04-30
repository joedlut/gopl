package main

import (
	"io"
	"log"
	"net"
	"os"
)

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	//tcpconn, err := net.DialTCP("tcp","localhost:8000","localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	//不要忘了语句末尾的()   函数调用
	go func() {
		io.Copy(os.Stdout, conn)
		//if err != nil {
		//	log.Fatalf("error1111 %v", err)
		//} else {
		//	log.Printf("%d bytes received", n)
		//}

		//服务器断开连接后执行
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	// goroutine退出之后，主goroutine才退出，不然一直阻塞
	<-done
}
