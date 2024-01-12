# Stress test CLI

# 🚀 Starting the application!

- Execute the command `make run`. and then we can see the container set in `docker-compose.yaml`.

# 🧪 How can I test?

- You need to execute this command
  ```shell
  docker compose run --rm stress-test-cli -h
  ```
- Then you can see this output:
  ```
    -c int
          Time in seconds of each request (default 1)
    -m string
          HTTP method to use (default "GET")
    -r int
          Amount of requests to send (default 1)
    -url string
          URL from service to test
  
  ```

- After that, you can execute this command (you should install `jq`)

    ```shell
    docker compose run --rm stress-test-cli -url http://google.com/ | jq
    ```

- You can use your custom flags

    ```shell
    docker compose run --rm stress-test-cli -url http://google.com/ -c 10 -r 100 | jq
    ```

- After you execute the command, you can se a response, something like this:
  ```json
  {
    "requests": 100,
    "workers": 10,
    "status_map_count": {
      "200": 90,
      "429": 9
    },
    "errors": [
      "Get \"http://www.google.com/\": stopped after 10 redirects"
    ],
    "error_count": 1,
    "total_duration": "625.48675ms"
  }
  ```

    - `requests`: the total amount of requests performed. This number must be equal to the sum of all
      the `status_map_count` values and `error_count` value.
    - `workers`: the total amount of concurrency requests performed.
    - `status_map_count`: the result map that counts the amount of HTTP requests status code, in this example we have 9
      requests with status **429** and 90 requests with the status **200**
    - `errors`: The array of errors on the request
    - `error_count`: the amount of errors that we have during the stress tests.
    - `total_duration`: the duration in string format, for this case, we perform the test in **~625.5ms**.

- If you have some error in the CLI, we display it as a JSON string as well, please execute this:
  ```shell
  docker compose run --rm stress-test-cli -url http://google.com/ -c 1000 -r 1 | jq
  ```
  we show this error:
  ```json
  {
    "error": "we have more workers (go routines) than requests"
  }
  ```
