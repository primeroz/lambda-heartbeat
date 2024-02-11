package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	alert "github.com/prometheus/alertmanager/template"
)

// validateLabels checks if the required labels are present in the alert
func validateLabels(alert alert.Alert) bool {
	requiredLabels := []string{"region", "cloud", "env", "cell", "cluster_type", "kubernetes_cluster_name"}
	for _, label := range requiredLabels {
		if _, ok := alert.Labels[label]; !ok {
			return false
		}
	}
	return true
}

// handleAlert handles incoming Alertmanager webhook alerts
func parseAlert(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var data alert.Data
	if err := json.Unmarshal([]byte(request.Body), &data); err != nil {
		log.Printf("Failed to parse Alertmanager webhook payload: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Failed to parse Alertmanager webhook payload",
		}, nil
	}

	// Check if the alert is a "watchdog" alert
	if data.GroupLabels["alertname"] != "Watchdog" {
		log.Println("Received non-watchdog alert. Ignoring.")
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       "Only watchdog alerts are processed by this webhook.",
		}, nil
	}

	// Validate labels
	for _, alert := range data.Alerts {
		if !validateLabels(alert) {
			log.Println("Missing required labels in the alert. Ignoring.")
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       "Alert is missing required labels.",
			}, nil
		}
	}

	// Process "watchdog" alerts
	log.Println("Received watchdog alert with status:", data.Status)

	// Respond with success message
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "Watchdog alert received successfully",
	}, nil
}

// handleAlert handles incoming Alertmanager webhook alerts
func handleAlert(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  response, err := parseAlert(ctx, request)
  return response, err
}

func main() {
	lambda.Start(handleAlert)
}
