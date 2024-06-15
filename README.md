# QuéPS
QuéPS - A Lightweight **Q**ueries **P**er **S**econd Measuring Service

## Install
```sh
go install github.com/sha65536/queps@master
```

## Usage
You can configure and run the QuePS service using environment variables or command-line flags.

### Using env vars
```sh
export QPS_HOST="0.0.0.0"
export QPS_PORT="8080"
export QPS_VERBOSE="false"
export QPS_INTERVAL="10"
export QPS_LABEL_NAMES="label1,label2"
export QPS_LABEL_VALUES="value1,value2"
export QPS_METRIC_PATH="/metrics"

queps
```

### Using flags
```sh
queps --host 0.0.0.0 \
    --port 8080 \
    --verbose false \
    --interval 10 \
    --label-names "label1,label2" \ 
    --label-values "value1,value2" \
    --metric-path "/metrics"
```

### Help section
```sh
queps --help
NAME:
   queps - A lightweight service that measures QPS and prints it to console and prometheus metrics

USAGE:
   queps [global options] command [command options] 

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --host value          Host address to bind the server (default: "0.0.0.0") [$QPS_HOST]
   --port value          Port to bind the server (default: "8080") [$QPS_PORT]
   --verbose             Enable verbose logging (default: false) [$QPS_VERBOSE]
   --interval value      Interval in seconds to print QPS to stdout (default: 10) [$QPS_INTERVAL]
   --label-names value   Comma-separated list of label names for the Prometheus metric [$QPS_LABEL_NAMES]
   --label-values value  Comma-separated list of label values for the Prometheus metric [$QPS_LABEL_VALUES]
   --metric-path value   Path to expose Prometheus metrics (default: "/metrics") [$QPS_METRIC_PATH]
   --help, -h            show help
```