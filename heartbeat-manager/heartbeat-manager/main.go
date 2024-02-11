package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/prometheus/alertmanager/template"
)

// Alert represents the parsed alert data
type Alert struct {
	Data   template.Data
	Metric MetricData
}

// MetricData represents the data needed to create a CloudWatch metric
type MetricData struct {
	Namespace string
	Dimension map[string]string
}

// parseAlert parses the incoming webhook alert payload
func parseAlert(body string) (*Alert, bool, error) {
	var data template.Data
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		return nil, false, err
	}

	// Check if the alert is a "watchdog" alert with all required labels
	var isWatchdog bool
	for _, alert := range data.Alerts {
		if alert.Labels["alertname"] == "Watchdog" &&
			alert.Labels["region"] != "" &&
			alert.Labels["cloud"] != "" &&
			alert.Labels["env"] != "" &&
			alert.Labels["cell"] != "" &&
			alert.Labels["cluster_type"] != "" &&
			alert.Labels["kubernetes_cluster_name"] != "" {
			isWatchdog = true
			break
		}
	}

	return &Alert{Data: data}, isWatchdog, nil
}

// createMetric creates a CloudWatch metric based on the parsed alert data
func createMetric(ctx context.Context, alert *Alert) error {
	cwClient := cloudwatch.New(session.New())

	for _, alert := range alert.Data.Alerts {
		metricData := MetricData{
			Namespace: alert.Labels["namespace"], // Adjust accordingly based on your alert structure
			Dimension: map[string]string{
				"region":                  alert.Labels["region"],
				"cloud":                   alert.Labels["cloud"],
				"env":                     alert.Labels["env"],
				"cell":                    alert.Labels["cell"],
				"cluster_type":            alert.Labels["cluster_type"],
				"kubernetes_cluster_name": alert.Labels["kubernetes_cluster_name"],
			},
		}

		_, err := cwClient.PutMetricData(&cloudwatch.PutMetricDataInput{
			Namespace: aws.String(metricData.Namespace),
			MetricData: []*cloudwatch.MetricDatum{
				{
					MetricName: aws.String("WatchdogAlert"),
					Dimensions: []*cloudwatch.Dimension{},
					Timestamp:  aws.Time(time.Now()),
					Value:      aws.Float64(1), // Set the metric value
				},
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

// handleAlert handles incoming Alertmanager webhook alerts
func handleAlert(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Parse the alert body
	//alert, isWatchdog, err := parseAlert(request.Body)
	_, isWatchdog, err := parseAlert(request.Body)
	if err != nil {
		log.Printf("Failed to parse Alertmanager webhook payload: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Failed to parse Alertmanager webhook payload",
		}, nil
	}

	if !isWatchdog {
		log.Println("Received non-watchdog alert or missing required labels")
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Received non-watchdog alert or missing required labels",
		}, nil
	}

	//// Create CloudWatch metric
	//if err := createMetric(ctx, alert); err != nil {
	//    log.Printf("Failed to create CloudWatch metric: %v", err)
	//    return events.APIGatewayProxyResponse{
	//        StatusCode: http.StatusInternalServerError,
	//        Body:       "Failed to create CloudWatch metric",
	//    }, nil
	//}

	// Respond with success message
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "Heartbeat Custom Metric published successfully",
	}, nil
}

func main() {
	lambda.Start(handleAlert)
}
