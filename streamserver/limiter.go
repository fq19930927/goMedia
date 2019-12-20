package main

import "log"

type ConnLimiter struct {
	//限制连接数
	concurrentConn int
	//连接bucket
	bucket chan int
}

//获取连接
func (cl *ConnLimiter) GenConn() bool {
	if (len(cl.bucket)) >= cl.concurrentConn {
		log.Printf("Reached the rate limitation")
		return false
	}
	cl.bucket <- 1
	return true
}

//释放连接
func (cl *ConnLimiter) ReleaseConn() {
	c := <-cl.bucket
	log.Println("Release conn is ", c)
}
