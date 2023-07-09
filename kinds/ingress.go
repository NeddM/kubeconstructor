package kinds

import (
	"fmt"
	"log"
	"os"
)

type Data struct {
	name      string
	namespace string
	labels    []string
	host      string
	service   string
	port      string
}

func Ingress() {

	labels := 1
	var key, value string

	// Formulario
	form := Data{}

	fmt.Printf("Name (ingress): ")
	fmt.Scanln(&form.name)

	fmt.Printf("Namespace (default): ")
	fmt.Scanln(&form.namespace)

	fmt.Printf("How many labels (Default 1): ")
	fmt.Scanln(&labels)

	fmt.Println("You are going to enter data KEY VALUE, one by one")
	fmt.Println("End example -> name: grafana")
	for i := 0; i < labels; i++ {
		fmt.Printf("%dº Label key: ", i+1)
		fmt.Scanln(&key)
		fmt.Printf("%dº Label value: ", i+1)
		fmt.Scanln(&value)

		newLabel := fmt.Sprintf("    %s: %s", key, value)
		form.labels = append(form.labels, newLabel)
	}

	fmt.Printf("Host: ")
	fmt.Scanln(&form.host)

	fmt.Printf("Service: ")
	fmt.Scanln(&form.service)

	fmt.Printf("Port (80): ")
	fmt.Scanln(&form.port)

	// Form validation
	if form.name == "" {
		form.name = "ingress"
	}
	if form.namespace == "" {
		form.namespace = "default"
	}

	if form.port == "" {
		form.port = "80"
	}

	// Creating file
	createFile(form)
}

func createFile(form Data) {
	// Creating the file
	f, err := os.Create("ingress-" + form.name + ".yaml")
	if err != nil {
		log.Fatal(err)
	}

	data1 := fmt.Sprintf(`apiVersion: networking.k8s.io/v1\n
kind: Ingress
metadata:
  name: %s
  namespace: %s
  labels:
`, form.name, form.namespace)

	f.WriteString(data1)

	for _, label := range form.labels {
		f.WriteString(label + "\n")
	}

	data2 := fmt.Sprintf(`spec:
  rules:
  - host: %s
    http:
      paths:
      - pathType: Prefix
        path: "/"
        backend:
          service:
            name: %s
            port:
              number: %s`, form.host, form.service, form.port)
	f.WriteString(data2)

	defer f.Close()
}
