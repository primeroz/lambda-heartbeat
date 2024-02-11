package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

// TestHandleAlert tests the HandleAlert function
func TestHandleAlert(t *testing.T) {
	// Define a sample "watchdog" alert payload
	payload := `{
	  "version": "4",
	  "groupKey": "example_group",
	  "status": "firing",
	  "receiver": "example_receiver",
	  "groupLabels": {
	    "alertname": "Watchdog"
	  },
	  "commonLabels": {
	    "alertname": "Watchdog",
	    "severity": "none",
	    "region": "us-west-1",
	    "cloud": "aws",
	    "env": "production",
	    "cell": "cell-1",
	    "cluster_type": "kubernetes",
	    "kubernetes_cluster_name": "my-cluster"
	  },
	  "commonAnnotations": {
	    "summary": "Alertmanager instance has stopped receiving alerts."
	  },
	  "externalURL": "http://alertmanager.example.com",
	  "alerts": [
	    {
	      "labels": {
	        "alertname": "Watchdog",
	        "severity": "none",
	        "region": "us-west-1",
	        "cloud": "aws",
	        "env": "production",
	        "cell": "cell-1",
	        "cluster_type": "kubernetes",
	        "kubernetes_cluster_name": "my-cluster"
	      },
	      "annotations": {
	        "description": "Error rate is above 90% for the last 5 minutes."
	      }
	    }
	  ]
	}`

	// Define the API Gateway event
	event := events.APIGatewayProxyRequest{
		Body: payload,
	}

	// Call the Lambda function handler
	response, err := handleAlert(context.Background(), event)
	if err != nil {
		t.Errorf("HandleAlert failed: %v", err)
		return
	}

	// Verify the response
	if response.StatusCode != 200 {
		t.Errorf("Unexpected status code: %d", response.StatusCode)
	}
	expectedBody := "Watchdog alert received successfully"
	if response.Body != expectedBody {
		t.Errorf("Unexpected body: got %s, want %s", response.Body, expectedBody)
	}
}

// BenchmarkHandleAlert benchmarks the HandleAlert function
func BenchmarkHandleAlert(b *testing.B) {
	payload := `{
	  "version": "4",
	  "groupKey": "example_group",
	  "status": "firing",
	  "receiver": "example_receiver",
	  "groupLabels": {
	    "alertname": "Watchdog"
	  },
	  "commonLabels": {
	    "alertname": "Watchdog",
	    "severity": "none",
	    "region": "us-west-1",
	    "cloud": "aws",
	    "env": "production",
	    "cell": "cell-1",
	    "cluster_type": "kubernetes",
	    "kubernetes_cluster_name": "my-cluster"
	  },
	  "commonAnnotations": {
	    "summary": "Alertmanager instance has stopped receiving alerts."
	  },
	  "externalURL": "http://alertmanager.example.com",
	  "alerts": [
	    {
	      "labels": {
	        "alertname": "Watchdog",
	        "severity": "none",
	        "region": "us-west-1",
	        "cloud": "aws",
	        "env": "production",
	        "cell": "cell-1",
	        "cluster_type": "kubernetes",
	        "kubernetes_cluster_name": "my-cluster"
	      },
	      "annotations": {
	        "description": "Error rate is above 90% for the last 5 minutes."
	      }
	    }
	  ]
	}`

	event := events.APIGatewayProxyRequest{
		Body: payload,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := handleAlert(context.Background(), event)
		if err != nil {
			b.Fatalf("HandleAlert failed: %v", err)
		}
	}
}
