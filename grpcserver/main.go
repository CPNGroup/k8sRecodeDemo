package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "pro/pb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedZkServiceServer
} //服务对象

// SayHello 实现服务的接口 在proto中定义的所有服务都是接口
func (s *server) Get(ctx context.Context, in *emptypb.Empty) (*pb.Message, error) {
	kubeconfig := getKubeConfig()

	config, err := clientConfig(kubeconfig)
	if err != nil {
		klog.Fatalf("Failed to create client config: %v", err)
	}

	config.Host = "https://168.11.7.126:6443" // 替换为master的IP

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatalf("Failed to create clientset: %v", err)
	}

	result := listNamespaces(clientset)
	return &pb.Message{Ns: result}, nil
}

// 列出所有命名空间
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer() //起一个服务

	pb.RegisterZkServiceServer(s, &server{})
	// 注册反射服务 这个服务是CLI使用的 跟服务本身没有关系
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func getKubeConfig() string {
	kubeconfig := "./admin.conf" // 注意，在容器中采用相对路径
	return kubeconfig
}

func clientConfig(kubeconfig string) (*rest.Config, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func listNamespaces(clientset *kubernetes.Clientset) []string {
	result := []string{}
	retries := 5
	for retries > 0 {
		namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			klog.Errorf("Failed to list namespaces: %v", err)
			retries--
			time.Sleep(5 * time.Second) // 等待一段时间后重试
			continue
		}

		fmt.Printf("Namespaces in the cluster:\n")
		for _, ns := range namespaces.Items {
			result = append(result, ns.Name)
		}
		break
	}
	return result
}
