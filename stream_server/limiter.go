package main
//token bucket 实现

import (
	"log"
)

type ConnLimiter struct{
	concurrentConn int
	bucket   chan int
}

//构造函数
func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket: make(chan int, cc),
	}
}

//判断token是否拿到
func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Printf("Reached the rate limitation.")
		return false
	}
	cl.bucket <- 1
	log.Printf("Successfully got connection")
	return true
}

func (cl *ConnLimiter) ReleaseConn() {
	c := <-cl.bucket
	log.Printf("New connection coming: %d", c)
}