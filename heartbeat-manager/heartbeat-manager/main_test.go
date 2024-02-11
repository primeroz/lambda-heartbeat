package main

import (
	"testing"
)

func TestParseAlert(t *testing.T) {
	// Test case: Valid "watchdog" alert with all required labels
	validWatchdogAlert := `{
        "version": "4",
        "groupKey": "{}:{alertname=\"Watchdog\", region=\"us-west-1\", cloud=\"aws\", env=\"production\", cell=\"cell-1\", cluster_type=\"kubernetes\", kubernetes_cluster_name=\"my-cluster\"}",
        "commonAnnotations": {},
        "commonLabels": {},
        "alerts": [
            {
                "status": "firing",
                "labels": {
                    "alertname": "Watchdog",
                    "region": "us-west-1",
                    "cloud": "aws",
                    "env": "production",
                    "cell": "cell-1",
                    "cluster_type": "kubernetes",
                    "kubernetes_cluster_name": "my-cluster"
                }
            }
        ]
    }`

	_, isWatchdog, err := parseAlert(validWatchdogAlert)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Verify that parsing was successful
	if !isWatchdog {
		t.Errorf("Expected alert to be a watchdog alert, got false")
	}

	// Test case: Non-watchdog alert
	nonWatchdogAlert := `{
        "version": "4",
        "groupKey": "{}:{alertname=\"NonWatchdog\"}",
        "commonAnnotations": {},
        "commonLabels": {},
        "alerts": [
            {
                "status": "firing",
                "labels": {
                    "alertname": "NonWatchdog"
                }
            }
        ]
    }`

	_, isWatchdog, err = parseAlert(nonWatchdogAlert)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Verify that parsing was successful
	if isWatchdog {
		t.Errorf("Expected alert not to be a watchdog alert, got true")
	}

	// Test case: Watchdog alert missing required labels
	missingLabelsAlert := `{
        "version": "4",
        "groupKey": "{}:{alertname=\"Watchdog\"}",
        "commonAnnotations": {},
        "commonLabels": {},
        "alerts": [
            {
                "status": "firing",
                "labels": {
                    "alertname": "Watchdog"
                }
            }
        ]
    }`

	_, isWatchdog, err = parseAlert(missingLabelsAlert)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Verify that parsing was successful
	if isWatchdog {
		t.Errorf("Expected alert not to be a watchdog alert, got true")
	}
}
