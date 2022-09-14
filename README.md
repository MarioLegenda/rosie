# Rosie

Rosie will help you load test your server.

## Installation

For now, there is no streamlined way of installing this package. You can download the already
built binaries in the build directory or clone this repository and build it yourself. Currently,
this package does not work on Windows. If it gets traction, I will create an installation
process and setup for all platforms.

## How it works

Rosie creates "users" that will send requests from provided urls in random intervals. After all urls
are sent, this user is done and a new one is created, repeating the process. For example, if you 
want 50 users to visit the urls that you provided, 50 users will send request concurrently in random
intervals.

## Usage

```
./rosie --user=20 --urls=https://facebook,https://google.com --interval=3-20
```

This command will create 20 virtual users that will send concurrent requests in a random interval between 3 and 
20 seconds for 60 seconds. Every user first clicks the first URL, waits for the random number of seconds between
3 and 20, and then clicks the next One. After all URLs are "clicked", user "logs out" and starts again. 

```
./rosie --user=20 --urls=https://facebook,https://google.com --interval=3
```

This command will create 20 virtual users that will send concurrent requests in a random interval between 3 and
15 seconds for 60 seconds. When {min}-{max} notation is not specified, `min` defaults to 3 and `max` defaults to 15.

```
./rosie --user=20 --urls=https://facebook,https://google.com --interval=3 --throttle --duration=120
```

This command will create 20 virtual users that will send concurrent requests in a random interval between 3 and
15 seconds. ``--throttle`` will start "gently" creating 10 users per second as a preparation for real load testing. 
``--duration`` tells Rosie that this load testing will last for 120 seconds. 60 is the default. 

## Arguments

``--urls (required)``

A comma separated list of URLs that will be clicked i.e. requests will be sent to these URLs. At this 
moment, this package only supports GET requests.

``--users : (optional|int)`` 

Number of users to simulate. Defaults to 50.

``--interval : (optional|string)`` 

Interval with which this package will simulate a single user clicking on a URL. Notation is either ``--interval={min}``
or ``--interval={min}-{max}``. `min` cannot be less than 3.

``--duration : (optional|int)``

Amount of time load testing will last in seconds. Infinite load testing is not possible but you can specify any 
amount of seconds you want. Defaults to 60.

``--throttle``

Prepares server for load testing by creating 10 users per second. Useful if you don't want to brutally DOS your
server right away. 


