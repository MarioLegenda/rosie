# Click simulator

This package is created to test resiliency of web server under real user load.

## Usage

```
./click-simulator --user=20 --links=https://facebook,https://google.com --interval=3-20
```

This command will create 20 virtual users that will send parallel request in a random interval between 3 seconds and 
20 seconds. Every user first clicks the first link, waits for the random number of seconds between 3 and 20, and then clicks 
the next link. After all links are clicked, user "logs out", waits for 10 seconds, and starts again. 

```
./click-simulator --user=20 --links=https://facebook,https://google.com --interval=3
```

This command will create 20 virtual users that will send parallel request in a random interval between 3 seconds and
15 seconds. When {min}-{max} notation is not specified, `min` defaults to 3 and `max` defaults to 15.

## Reference

``--links : required|string``
A comma separated list of URLs that will be clicked i.e. request will be sent to these URLs. At this 
moment, this package only supports GET request.

``--users : optional|int`` Number of users to simulate

``--interval : optional|string`` Interval with which this package will simulate a single user clicking on a link.


