# System-monitor daemon

## Environment

```shell script
# parser tick interval in seconds, default is 1s
# each tick daemon will run all enabled parsers
TICK_INTERVAL 

# enabled parsers separated by ":", 
# by default all enabled "load_avg:cpu"`
PARSERS 

# host name or IP address daemon listens to
# by default it listens any interface, e.g. '0.0.0.0'
HOST

# GRPC port of server, default 9000
GRPC_PORT

# HTTP port for JS client to get access
# default 9001
HTTP_PORT

# logging level, default is 'info'
LOG_LEVEL

# production flag, by default is 'false'
PRODUCTION
```