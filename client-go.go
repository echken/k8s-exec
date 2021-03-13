package main

import (
  "flag"
  "fmt"
  "os"
  "path/filepath"
  "time"
  "context"

  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  "k8s.io/client-go/kubernetes"
  "k8s.io/client-go/tools/clientcmd"
)

func main() {
  var kubeconfig *string
  if home := homeDir(); home != "" {
    kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "/home")
  } else {
    kubeconfig = flag.String("kubeconfig", "", "/home")
  }
  flag.Parse()

  config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
  if err != nil {
    panic(err.Error())
  }

  clientSet, err := kubernetes.NewForConfig(config)
  if err != nil {
    panic(err.Error())
  }

  for {
    pods, err := clientSet.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
    if err != nil {
      panic(err.Error())
    }
    fmt.Printf("there are %d pods in cluster\n", len(pods.Items))
    time.Sleep(10 * time.Second)
  }
}

func homeDir() string {
  if h := os.Getenv("HOME"); h != "" {
    return h
  }
  return os.Getenv("USERPROFILE")
}
