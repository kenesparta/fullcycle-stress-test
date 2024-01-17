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
  Usage of ./stress-test-cli:
  -m string
      HTTP method to use (default "GET")
  -r int
      Amount of requests to send (default 10)
  -url string
      URL from service to test
  -w int
      Amount of concurrent requests (workers) (default 1)
  ```

- After that, you can execute this command (you should install `jq`)

    ```shell
    docker compose run --rm stress-test-cli -url http://google.com/ | jq
    ```

- You can use your custom flags

  ```shell
    docker compose run --rm stress-test-cli -url http://google.com/ -w 20 -r 1000 | jq | jq
  ```

- After you execute the command, you can se a response, something like this:
  ```json
  {
   "requests": 1000,
   "workers": 20,
   "status_map_count": {
     "200": 701,
     "301": 3,
     "302": 79,
     "429": 217
   },
   "errors": [],
   "error_count": 0,
   "total_duration": "33.274527583s"
  }
  ```

    - `requests`: the total amount of requests performed. This number must be equal to the sum of all
      the `status_map_count` values and `error_count` value.
    - `workers`: the total amount of concurrency requests performed.
    - `status_map_count`: the result map that counts the amount of HTTP requests status code, in this example we have
      701 requests with status **200**, 3 requests with the status **301**, 79 with **302** and 217 with **409**.
    - `errors`: The array of errors on the request. These errors are related with each request.
    - `error_count`: the amount of errors that we have during the stress tests.
    - `total_duration`: the duration in string format, for this case, we perform the test in **~33s**.

# 🚨 Errors

We have identified these possible errors that validated the input params **BEFORE** we proceed with the execution of the
stress tests.

| Condition        | Example        | Json Response                                                   |
|------------------|----------------|-----------------------------------------------------------------|
| `r <= 0`         | `-w 12 -r 0`   | `{"error": "request should be positive number"}`                |
| `w <= 0`         | `-w 0 -r 13`   | `{"error": "concurrency value should be positive number"}`      |
| `w > r`          | `-w 500 -r 5`  | `{"error": "we have more workers (go routines) than requests"}` |
| `valid(url)`     | `-url hhpp:/g` | `{"error": "should be a URL valid value"}`                      |
| `validMethod(m)` | `-m PMTD`      | `{"error": "should be a valid HTTP method"}`                    |

## Type of errors

### Error before executing the request tests

If we execute this:

```shell
docker compose run --rm stress-test-cli -url http://google.com/ -w 500 -r 5 | jq
```

we show this error:

```json
{
  "error": "we have more workers (go routines) than requests"
}
```

Since `w > r`, as indicated in the previously shown table, we are in a blocked condition, and therefore, we will not
proceed with the execution of the stress test.

### Errors in each request

If we execute this:

```shell
docker compose run --rm stress-test-cli -url http://google.com/ -w 5 -r 500 | jq
```

we have

```json
{
  "requests": 500,
  "workers": 5,
  "status_map_count": {
    "200": 498
  },
  "errors": [
    "Get \"http://www.google.com/\": stopped after 10 redirects",
    "Get \"http://www.google.com/\": stopped after 10 redirects"
  ],
  "error_count": 2,
  "total_duration": "4.470558169s"
}

```

the `errors` in the JSON are showing because we have these errors in each request.