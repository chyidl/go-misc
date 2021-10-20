/*
etcd v3 API
*/
package main

import (
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"172.30.1.14:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {

		return
	}
}
