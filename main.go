package main

import (
	"encoding/json"
	"flag"
	"log"

	k8s "github.com/mkm29/vulncron/pkg/kubernetes"
	"github.com/mkm29/vulncron/pkg/reports"
)

var (
	kubeconfig string
	cluster    string
)

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file")
	flag.StringVar(&cluster, "cluster", "", "cluster name")
	flag.Parse()
}

func main() {
	// Connect to Kubernetes API
	client, err := k8s.Connect(kubeconfig)
	if err != nil {
		log.Fatalf("failed to connect to Kubernetes API: %v", err)
	}

	// get VulnerabilityReportList
	vrl, err := reports.GetVulnerabilityReportList(client)
	if err != nil {
		log.Fatalf("failed to get VulnerabilityReportList: %v", err)
	}
	summary, summaries := reports.GetReportSummaries(vrl)
	_ = summaries

	// Convert summary to string
	js, err := json.Marshal(summary)
	if err != nil {
		log.Fatalf("failed to marshall to JSON: %v", err)
	}

	log.Printf("Sending email: %s", string(js))

	// err = reports.SendMail("descdevops@rtx.com", "Trivy Report Summary", string(js), []string{"descdevops@rtx.com"})
	// if err != nil {
	// 	log.Fatalf("failed to send email: %v", err)
	// }
}
