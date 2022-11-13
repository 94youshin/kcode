/**
 * 使用Containerd API进行容器操作
 */
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/containerd/containerd"
	_ "github.com/containerd/containerd/cio"
	"github.com/containerd/containerd/namespaces"
)

const endpoint = "/run/containerd/containerd.sock"

func main() {
	fmt.Println("Hello Fleet!")
	client, err := containerd.New(endpoint)
	if err != nil {
		log.Print("Failed connect containerd!")
	}
	defer client.Close()

	ctx := namespaces.WithNamespace(context.Background(), "k8s.io")

	// Pull Image
	image, err := client.Pull(ctx, "docker.io/library/busybox:latest")
	if err != nil {
		log.Printf("Failed pull image: %s", err.Error())
	}

	bdata, _ := json.Marshal(image)
	log.Printf("Redis image info: %v", bdata)

	//

}
