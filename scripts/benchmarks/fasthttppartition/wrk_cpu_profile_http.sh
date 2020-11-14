#!/usr/bin/env bash

# 1) Build the application.
echo "Building the cache"
cd ../../../src/cmd/fasthttpexample
go build -o simplecache .
cd - >/dev/null
mv ../../../src/cmd/fasthttpexample/simplecache .

# 2) Start the application.
echo "Starting the cache"
./simplecache &
APP_PID=$!

sleep 2

# 3) Insert some test data.
curl -X PUT -d'{"key":"value"}' http://localhost:8080/something

if [ ! -d results ] ; then
  mkdir results
fi

# 4) Run the tests and save results to file.
echo "== Running GET benchmark =="
cat >wrk_put.lua <<-END
request = function()
  param_value = math.random(1, 100000)
  path = "/something" .. param_value
  return wrk.format("GET", path)
end
END
curl localhost:8081/debug/pprof/profile --output get_cpu_profile.prof &
wrk -t12 -c400 -d30s http://localhost:8080/something | tee results/get_results.txt

echo "== Running PUT benchmark =="
cat >wrk_put.lua <<-END
request = function()
  param_value = math.random(1, 100000)
  path = "/something" .. param_value
  return wrk.format("POST", path, wrk.headers, "{\"key\":\"value\"}")
end
END
curl localhost:8081/debug/pprof/profile --output put_cpu_profile.prof &
wrk -t12 -c400 -d30s -s wrk_put.lua http://localhost:8080/something | tee results/put_results.txt
rm wrk_put.lua

echo "== Running DELETE benchmark =="
cat >wrk_del.lua <<-END
request = function()
  param_value = math.random(1, 100000)
  path = "/something" .. param_value
  return wrk.format("DELETE", path, wrk.headers, "{\"key\":\"value\"}")
end
END
curl localhost:8081/debug/pprof/profile --output del_cpu_profile.prof &
wrk -t12 -c400 -d30s -s wrk_del.lua http://localhost:8080/something | tee results/del_results.txt
rm wrk_del.lua

# 5) Kill the application.
echo "Killing the cache application"
kill $APP_PID
