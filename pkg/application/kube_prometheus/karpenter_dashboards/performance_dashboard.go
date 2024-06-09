package karpenter_dashboards

// JSON: https://github.com/aws/karpenter-provider-aws/blob/main/website/content/en/preview/getting-started/getting-started-with-karpenter/karpenter-performance-dashboard.json
//
//	kubectl create configmap karpenter-performance \
//	  --from-file=./karpenter-performance-dashboard.json \
//	  --dry-run=client -o yaml > performance_dashboard.yaml
//
// Must double escape because textTemplate transforms and Helm uses the same Golang template
// TextTemplate turns: "{{ "{{" }} {{ "{{controller}}" | printf "%q" }} {{ "}}" }}" _to_ "{{ "{{controller}}" }}"
// and Helm turns it _to_ {{controller}}
const performanceDashboardTemplate = `---
apiVersion: v1
kind: ConfigMap
metadata:
  name: karpenter-performance
data:
  karpenter-performance-dashboard.json: |-
    {
        "annotations": {
            "list": [
                {
                    "builtIn": 1,
                    "datasource": "-- Grafana --",
                    "enable": true,
                    "hide": true,
                    "iconColor": "rgba(0, 211, 255, 1)",
                    "name": "Annotations & Alerts",
                    "target": {
                        "limit": 100,
                        "matchAny": false,
                        "tags": [],
                        "type": "dashboard"
                    },
                    "type": "dashboard"
                }
            ]
        },
        "editable": true,
        "fiscalYearStartMonth": 0,
        "graphTooltip": 0,
        "id": 7,
        "links": [],
        "liveNow": true,
        "panels": [
            {
                "datasource": {
                    "type": "prometheus",
                    "uid": "${datasource}"
                },
                "fieldConfig": {
                    "defaults": {
                        "color": {
                            "mode": "palette-classic"
                        },
                        "custom": {
                            "axisLabel": "",
                            "axisPlacement": "auto",
                            "barAlignment": 0,
                            "drawStyle": "line",
                            "fillOpacity": 0,
                            "gradientMode": "none",
                            "hideFrom": {
                                "legend": false,
                                "tooltip": false,
                                "viz": false
                            },
                            "lineInterpolation": "linear",
                            "lineWidth": 1,
                            "pointSize": 5,
                            "scaleDistribution": {
                                "type": "linear"
                            },
                            "showPoints": "never",
                            "spanNulls": false,
                            "stacking": {
                                "group": "A",
                                "mode": "none"
                            },
                            "thresholdsStyle": {
                                "mode": "off"
                            }
                        },
                        "mappings": [],
                        "thresholds": {
                            "mode": "absolute",
                            "steps": [
                                {
                                    "color": "green",
                                    "value": null
                                },
                                {
                                    "color": "red",
                                    "value": 80
                                }
                            ]
                        },
                        "unit": "s"
                    },
                    "overrides": []
                },
                "gridPos": {
                    "h": 9,
                    "w": 24,
                    "x": 0,
                    "y": 0
                },
                "id": 4,
                "options": {
                    "legend": {
                        "calcs": [],
                        "displayMode": "list",
                        "placement": "bottom"
                    },
                    "tooltip": {
                        "mode": "single",
                        "sort": "none"
                    }
                },
                "targets": [
                    {
                        "datasource": {
                            "type": "prometheus",
                            "uid": "${datasource}"
                        },
                        "editorMode": "code",
                        "expr": "karpenter_nodes_termination_time_seconds{quantile=\"0\"}",
                        "legendFormat": "Min",
                        "range": true,
                        "refId": "A"
                    },
                    {
                        "datasource": {
                            "type": "prometheus",
                            "uid": "${datasource}"
                        },
                        "editorMode": "code",
                        "expr": "karpenter_nodes_termination_time_seconds{quantile=\"0.5\"}",
                        "hide": false,
                        "legendFormat": "P50",
                        "range": true,
                        "refId": "B"
                    },
                    {
                        "datasource": {
                            "type": "prometheus",
                            "uid": "${datasource}"
                        },
                        "editorMode": "code",
                        "expr": "karpenter_nodes_termination_time_seconds{quantile=\"0.9\"}",
                        "hide": false,
                        "legendFormat": "P90",
                        "range": true,
                        "refId": "C"
                    },
                    {
                        "datasource": {
                            "type": "prometheus",
                            "uid": "${datasource}"
                        },
                        "editorMode": "code",
                        "expr": "karpenter_nodes_termination_time_seconds{quantile=\"0.99\"}",
                        "hide": false,
                        "legendFormat": "P99",
                        "range": true,
                        "refId": "D"
                    },
                    {
                        "datasource": {
                            "type": "prometheus",
                            "uid": "${datasource}"
                        },
                        "editorMode": "code",
                        "expr": "karpenter_nodes_termination_time_seconds{quantile=\"1\"}",
                        "hide": false,
                        "legendFormat": "Max",
                        "range": true,
                        "refId": "E"
                    }
                ],
                "title": "Node Termination Latency",
                "type": "timeseries"
            },
            {
                "datasource": {
                    "type": "prometheus",
                    "uid": "${datasource}"
                },
                "fieldConfig": {
                    "defaults": {
                        "color": {
                            "mode": "palette-classic"
                        },
                        "custom": {
                            "axisLabel": "",
                            "axisPlacement": "auto",
                            "barAlignment": 0,
                            "drawStyle": "line",
                            "fillOpacity": 10,
                            "gradientMode": "none",
                            "hideFrom": {
                                "legend": false,
                                "tooltip": false,
                                "viz": false
                            },
                            "lineInterpolation": "linear",
                            "lineWidth": 1,
                            "pointSize": 5,
                            "scaleDistribution": {
                                "type": "linear"
                            },
                            "showPoints": "never",
                            "spanNulls": false,
                            "stacking": {
                                "group": "A",
                                "mode": "none"
                            },
                            "thresholdsStyle": {
                                "mode": "off"
                            }
                        },
                        "mappings": [],
                        "min": 0,
                        "thresholds": {
                            "mode": "absolute",
                            "steps": [
                                {
                                    "color": "green",
                                    "value": null
                                },
                                {
                                    "color": "red",
                                    "value": 80
                                }
                            ]
                        },
                        "unit": "s"
                    },
                    "overrides": []
                },
                "gridPos": {
                    "h": 8,
                    "w": 24,
                    "x": 0,
                    "y": 9
                },
                "id": 2,
                "options": {
                    "legend": {
                        "calcs": [],
                        "displayMode": "list",
                        "placement": "bottom"
                    },
                    "tooltip": {
                        "mode": "single",
                        "sort": "none"
                    }
                },
                "targets": [
                    {
                        "datasource": {
                            "type": "prometheus",
                            "uid": "${datasource}"
                        },
                        "editorMode": "code",
                        "expr": "karpenter_pods_startup_time_seconds{quantile=\"0\"}",
                        "format": "time_series",
                        "legendFormat": "Min",
                        "range": true,
                        "refId": "Minimum"
                    },
                    {
                        "datasource": {
                            "type": "prometheus",
                            "uid": "${datasource}"
                        },
                        "editorMode": "code",
                        "expr": "karpenter_pods_startup_time_seconds{quantile=\"0.5\"}",
                        "hide": false,
                        "legendFormat": "P50",
                        "range": true,
                        "refId": "Median"
                    },
                    {
                        "datasource": {
                            "type": "prometheus",
                            "uid": "${datasource}"
                        },
                        "editorMode": "code",
                        "expr": "karpenter_pods_startup_time_seconds{quantile=\"0.9\"}",
                        "hide": false,
                        "legendFormat": "P90",
                        "range": true,
                        "refId": "P90"
                    },
                    {
                        "datasource": {
                            "type": "prometheus",
                            "uid": "${datasource}"
                        },
                        "editorMode": "code",
                        "expr": "karpenter_pods_startup_time_seconds{quantile=\"0.99\"}",
                        "hide": false,
                        "legendFormat": "P99",
                        "range": true,
                        "refId": "P99"
                    },
                    {
                        "datasource": {
                            "type": "prometheus",
                            "uid": "${datasource}"
                        },
                        "editorMode": "code",
                        "expr": "karpenter_pods_startup_time_seconds{quantile=\"1\"}",
                        "hide": false,
                        "legendFormat": "Max",
                        "range": true,
                        "refId": "Maximum"
                    }
                ],
                "title": "Pod Startup Latency",
                "type": "timeseries"
            },
            {
                "datasource": {
                    "type": "prometheus",
                    "uid": "${datasource}"
                },
                "fieldConfig": {
                    "defaults": {
                        "color": {
                            "mode": "palette-classic"
                        },
                        "custom": {
                            "axisLabel": "",
                            "axisPlacement": "auto",
                            "barAlignment": 0,
                            "drawStyle": "line",
                            "fillOpacity": 10,
                            "gradientMode": "none",
                            "hideFrom": {
                                "legend": false,
                                "tooltip": false,
                                "viz": false
                            },
                            "lineInterpolation": "linear",
                            "lineWidth": 1,
                            "pointSize": 5,
                            "scaleDistribution": {
                                "type": "linear"
                            },
                            "showPoints": "never",
                            "spanNulls": false,
                            "stacking": {
                                "group": "A",
                                "mode": "none"
                            },
                            "thresholdsStyle": {
                                "mode": "off"
                            }
                        },
                        "mappings": [],
                        "thresholds": {
                            "mode": "absolute",
                            "steps": [
                                {
                                    "color": "green",
                                    "value": null
                                },
                                {
                                    "color": "red",
                                    "value": 80
                                }
                            ]
                        },
                        "unit": "s"
                    },
                    "overrides": []
                },
                "gridPos": {
                    "h": 8,
                    "w": 13,
                    "x": 0,
                    "y": 17
                },
                "id": 6,
                "options": {
                    "legend": {
                        "calcs": [],
                        "displayMode": "list",
                        "placement": "bottom"
                    },
                    "tooltip": {
                        "mode": "single",
                        "sort": "none"
                    }
                },
                "targets": [
                    {
                        "datasource": {
                            "type": "prometheus",
                            "uid": "${datasource}"
                        },
                        "editorMode": "code",
                        "expr": "histogram_quantile(0, rate(controller_runtime_reconcile_time_seconds_bucket{controller=\"$controller\"}[10m]))",
                        "hide": false,
                        "legendFormat": "Min",
                        "range": true,
                        "refId": "Minimum"
                    },
                    {
                        "datasource": {
                            "type": "prometheus",
                            "uid": "${datasource}"
                        },
                        "editorMode": "code",
                        "expr": "histogram_quantile(0.5, rate(controller_runtime_reconcile_time_seconds_bucket{controller=\"$controller\"}[10m]))",
                        "legendFormat": "P50",
                        "range": true,
                        "refId": "Median"
                    },
                    {
                        "datasource": {
                            "type": "prometheus",
                            "uid": "${datasource}"
                        },
                        "editorMode": "code",
                        "expr": "histogram_quantile(0.9, rate(controller_runtime_reconcile_time_seconds_bucket{controller=\"$controller\"}[10m]))",
                        "hide": false,
                        "legendFormat": "P90",
                        "range": true,
                        "refId": "P90"
                    },
                    {
                        "datasource": {
                            "type": "prometheus",
                            "uid": "${datasource}"
                        },
                        "editorMode": "code",
                        "expr": "histogram_quantile(0.99, rate(controller_runtime_reconcile_time_seconds_bucket{controller=\"$controller\"}[10m]))",
                        "hide": false,
                        "legendFormat": "P99",
                        "range": true,
                        "refId": "P99"
                    },
                    {
                        "datasource": {
                            "type": "prometheus",
                            "uid": "${datasource}"
                        },
                        "editorMode": "code",
                        "expr": "histogram_quantile(1, rate(controller_runtime_reconcile_time_seconds_bucket{controller=\"$controller\"}[10m]))",
                        "hide": false,
                        "legendFormat": "Max",
                        "range": true,
                        "refId": "Maximum"
                    }
                ],
                "title": "Controller Reconciliation Latency [$controller]",
                "type": "timeseries"
            },
            {
                "datasource": {
                    "type": "prometheus",
                    "uid": "${datasource}"
                },
                "fieldConfig": {
                    "defaults": {
                        "color": {
                            "mode": "thresholds"
                        },
                        "mappings": [],
                        "thresholds": {
                            "mode": "absolute",
                            "steps": [
                                {
                                    "color": "green",
                                    "value": null
                                },
                                {
                                    "color": "red",
                                    "value": 80
                                }
                            ]
                        },
                        "unit": "reqps"
                    },
                    "overrides": []
                },
                "gridPos": {
                    "h": 8,
                    "w": 11,
                    "x": 13,
                    "y": 17
                },
                "id": 8,
                "options": {
                    "displayMode": "gradient",
                    "minVizHeight": 10,
                    "minVizWidth": 0,
                    "orientation": "horizontal",
                    "reduceOptions": {
                        "calcs": [
                            "lastNotNull"
                        ],
                        "fields": "",
                        "values": false
                    },
                    "showUnfilled": true
                },
                "pluginVersion": "9.0.5",
                "targets": [
                    {
                        "datasource": {
                            "type": "prometheus",
                            "uid": "${datasource}"
                        },
                        "editorMode": "code",
                        "expr": "sum(rate(controller_runtime_reconcile_total[10m])) by (controller)",
                        "legendFormat": "{{ "{{" }} {{ "{{controller}}" | printf "%q" }} {{ "}}" }}",
                        "range": true,
                        "refId": "A"
                    }
                ],
                "title": "Controller Reconciliation Rate",
                "type": "bargauge"
            }
        ],
        "refresh": "5s",
        "schemaVersion": 36,
        "style": "dark",
        "tags": [],
        "templating": {
            "list": [
                {
                    "current": {
                        "selected": false,
                        "text": "Prometheus",
                        "value": "Prometheus"
                    },
                    "hide": 0,
                    "includeAll": false,
                    "label": "Data Source",
                    "multi": false,
                    "name": "datasource",
                    "options": [],
                    "query": "prometheus",
                    "refresh": 1,
                    "regex": "",
                    "skipUrlSync": false,
                    "type": "datasource"
                },
                {
                    "current": {
                        "selected": false,
                        "text": "provisioning",
                        "value": "provisioning"
                    },
                    "datasource": {
                        "type": "prometheus",
                        "uid": "${datasource}"
                    },
                    "definition": "label_values(controller_runtime_reconcile_time_seconds_count, controller)",
                    "hide": 0,
                    "includeAll": false,
                    "multi": false,
                    "name": "controller",
                    "options": [],
                    "query": {
                        "query": "label_values(controller_runtime_reconcile_time_seconds_count, controller)",
                        "refId": "StandardVariableQuery"
                    },
                    "refresh": 2,
                    "regex": "",
                    "skipUrlSync": false,
                    "sort": 0,
                    "type": "query"
                }
            ]
        },
        "time": {
            "from": "now-6h",
            "to": "now"
        },
        "timepicker": {},
        "timezone": "",
        "title": "Karpenter Performance",
        "uid": "_bdgC2g4z",
        "version": 3,
        "weekStart": ""
    }
`
