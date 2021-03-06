# nm.com Go Test ( Test `Question` Details )

Welcome to the [nm.com](nm.com) Go test. The purpose of this assignment is to test your familiarity with Go, distributed systems concepts, performance benchmarking and TDD.

## Background

The source code that you are given is a very simple imitation of a key/value store:

* `Database` is the central store that takes a while (500ms) to store and retrieve data.
* `DistributedCache` represents a distributed cache that takes much less time to turn around (100ms to store or retrieve).

This scenario is a simplified example of a typical high performance server cluster with a database, a distributed cache and multiple worker nodes.

## Assumptions 

After startup:

* Data in `Database` never changes and can be cached forever.
* If `Database.Value()` returns `nil` for a key, the requested data item does not exist and will never exist.
* `DistributedCache` is initially empty.

## Task

Complete the 2 parts below & submit the solution. 
If the solution is incomplete, please state what hasn't been finished and outline how you are planning on solving it.

* Provided code can be modified at will.
* The whole solution must build with no errors.

### Part 1

Create an implementation for the `DataSource` interface to create a mechanism to retrieve data from `Database` with lowest possible latency.
For a frequently-requested item your `DataSource.Value()` implementation should have a better response time than the distributed cache store (ie < 100ms).

* The user of the `DataSource` interface must not have to deal with thread synchronisation.
* You must write unit tests for all new functionality, and all tests must pass.
* Ideally limit your use of libraries to the standard library only.

### Part 2

Complete `main()` to test your `DataSource` implementation; it must:

* Populate `Database` with the following data at startup:
<pre>
    | key          | value         |
    --------------------------------
    | key0         | value0        |
    | key1         | value1        |
    | key2         | value2        |
    | key3         | value3        |
    | key4         | value4        |
    | key5         | value5        |
    | key6         | value6        |
    | key7         | value7        |
    | key8         | value8        |
    | key9         | value9        |
</pre>
* Use 10 goroutines, each making 50 consecutive requests for a random key in the range (key0-key9). I.e. there should be a total of 500 requests.
* For each request, print the requested key name, returned value, time to complete that request; similar to the following example:
<pre>
    [1] Request 'key1', response 'value1', time: 50.05 ms
    [2] Request 'key2', response 'value2', time: 50.05 ms
</pre>
* Write a benchmark that measures the average run time of these goroutines performing the requests above.

## Submission instructions

* **DO NOT** fork this repository or create pull requests on it, because we don't want other candidates to see your solution.
* Provide your solution as a `.zip` or .`gz` archive file, either via email or some Dropbox-like service, to your nm contact.

# Solution setup & run instructions

* **TO RUN PROGRAM** `go run main.go` to run the program. 

* **TO RUN UNIT TESTS** `go test ./...` to run all unit test. Please make sure you have installed the third party test packages `testify/mock` and `testify/assert` before running the unit test.

* **IT IS VERY IMPORTANT** to make sure that the location of my project code folder `nmgt` should be `$GOPATH/src/github.com/gmpatel/nmgt` for project to build and run successfully!

* **HAVE NOT** used any third party packages for the code exercise itself, but **HAVE USED** `testify/mock` and `testify/assert` packages for `unit-testing` as they are nice and the assertion statements of `testify/assert` are sensible!.

* **HAVE `DEP ENSURE`** third-party packages used for the test, but **IN CASE THERE IS A TROUBLE RUNNING THE `UNIT TESTS`** please install `testify/mock` and `testify/assert` packages for `unit-testing`. To install the packages you can follow the instructions from the next step to install those dependencies to your machine manually.

* **PLEASE RUN** `install.sh` bash script from the root of the project to install dependent packages mentioned above. Command to run bash script is `bash install.sh`. 
