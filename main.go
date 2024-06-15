package main

import (
	"log"
	"os"
	"strings"

	"github.com/sha65536/queps/queps"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "queps",
		Usage: "A lightweight service that measures QPS and prints it to console and prometheus metrics",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "host",
				Value:   "0.0.0.0",
				Usage:   "Host address to bind the server",
				EnvVars: []string{"QPS_HOST"},
			},
			&cli.StringFlag{
				Name:    "port",
				Value:   "8080",
				Usage:   "Port to bind the server",
				EnvVars: []string{"QPS_PORT"},
			},
			&cli.BoolFlag{
				Name:    "verbose",
				Value:   false,
				Usage:   "Enable verbose logging",
				EnvVars: []string{"QPS_VERBOSE"},
			},
			&cli.IntFlag{
				Name:    "interval",
				Value:   10,
				Usage:   "Interval in seconds to print QPS to stdout",
				EnvVars: []string{"QPS_INTERVAL"},
			},
			&cli.StringFlag{
				Name:    "label-names",
				Value:   "placeholder",
				Usage:   "Comma-separated list of label names for the Prometheus metric",
				EnvVars: []string{"QPS_LABEL_NAMES"},
			},
			&cli.StringFlag{
				Name:    "label-values",
				Value:   "placeholder",
				Usage:   "Comma-separated list of label values for the Prometheus metric",
				EnvVars: []string{"QPS_LABEL_VALUES"},
			},
			&cli.StringFlag{
				Name:    "metric-path",
				Value:   "/metrics",
				Usage:   "Path to expose Prometheus metrics",
				EnvVars: []string{"QPS_METRIC_PATH"},
			},
		},
		Action: func(c *cli.Context) error {
			labelNames := strings.Split(c.String("label-names"), ",")
			labelValues := strings.Split(c.String("label-values"), ",")

			// If in kubernetes, HOSTNAME will be the pod name
			if os.Getenv("HOSTNAME") != "" {
				labelNames = append(labelNames, "pod")
				labelValues = append(labelValues, os.Getenv("HOSTNAME"))
			}

			qpsMeter := &queps.QPSMeter{
				Host:        c.String("host"),
				Port:        c.String("port"),
				Verbose:     c.Bool("verbose"),
				Interval:    c.Int("interval"),
				LabelNames:  labelNames,
				LabelValues: labelValues,
				MetricPath:  c.String("metric-path"),
			}

			if err := qpsMeter.Start(); err != nil {
				log.Fatalf("Failed to start QPSMeter: %v", err)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
