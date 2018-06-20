echo -e "config:  -c 5 -n 200 -t 5 -z 120s -cpus 1"
./tooling/bin/grpcannon -proto ./model/liletestone.proto -call liletestone.Liletestone.Get \
 -c 5 -n 200 -t 5 -z 120s -cpus 1 \
 -D ./req.json localhost:8000

 # -c 15

# -c  Number of requests to run concurrently. Total number of requests cannot
#	  be smaller than the concurrency level. Default is 50.
#  -n  Number of requests to run. Default is 200.
#  -q  Rate limit, in queries per second (QPS). Default is no rate limit.
#  -t  Timeout for each request in seconds. Default is 20, use 0 for infinite.
#  -z  Duration of application to send requests. When duration is reached,
###-x  Maximum duration of application to send requests with n setting respected.
  #    If duration is reached before n requests are completed, application stops and exits.
  #    Examples: -x 10s -x 3m.
#
#  -d  The call data as stringified JSON.
#  -D  Path for call data JSON file. For example, /home/user/file.json or ./file.json.
#  -m  Request metadata as stringified JSON.
#  -M  Path for call metadata JSON file. For example, /home/user/metadata.json or ./metadata.json.
#
#  -o  Output path. If none provided stdout is used.
