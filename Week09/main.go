package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Conn struct {
	IP       string
	Port     uint32
	TCPConn  *net.TCPConn
	MsgChan  chan []byte
	ExitChan chan bool
}

func NewConn(ip string, port uint32) *Conn {
	return &Conn{
		IP:      ip,
		Port:    port,
		MsgChan: make(chan []byte),
	}
}

func (c *Conn) Start() {
	log.Printf("%s:%d start...\n", c.IP, c.Port)
	go func() {
		addr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d", c.IP, c.Port))
		if err != nil {
			log.Println("resolve tcp addr err ", err)
			return
		}
		listener, err := net.ListenTCP("tcp4", addr)
		if err != nil {
			log.Println("listen tcp err ", err)
			return
		}

		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				log.Println("accept tcp err ", err)
				continue
			}
			c.TCPConn = conn
			// 读
			go c.StartRead()
			// 写
			go c.StartWrite()
		}
	}()
	select {}
}

func (c *Conn) StartRead() {
	log.Println("read is wait")
	defer c.Stop()
	defer log.Println("read goroutine exit")

	for {
		reader := bufio.NewReader(c.TCPConn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		go c.HandleMsg([]byte(recvStr))
	}
}

func (c *Conn) HandleMsg(data []byte) {
	res := fmt.Sprintf("res:%s", string(data))

	c.MsgChan <- []byte(res)
}

func (c *Conn) StartWrite() {
	log.Println("write is wait")
	defer log.Println("write goroutine exit")
	for {
		select {
		case data := <-c.MsgChan:
			if _, err := c.TCPConn.Write(data); err != nil {
				log.Println("conn write error ", err)
				return
			}
		case <-c.ExitChan:
			return
		}

	}
}

func (c *Conn) Stop() {
	c.ExitChan <- true

	c.TCPConn.Close()
	close(c.ExitChan)
	close(c.MsgChan)
}

func main() {
	c := NewConn("0.0.0.0", 8777)
	c.Start()
}
