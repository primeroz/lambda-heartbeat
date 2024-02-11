#!/bin/bash

curl -X POST -H "Content-Type: application/json" -d '{
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
}' $1
