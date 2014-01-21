# STATUS: WIP, nothing works, this is README driven development.

# hurl, a tool to hurl traffic at URLs

`hurl` provides a suite of network benchmarking and result analytic tools.

You can start using `hurl` with commands similar to Apache Bench `ab`, but as your needs grow, it exposes a framework for much more complicated and intense test plans.

## Installation

    go install github.com/pquerna/hurl/hurl

## Running

`hurl http` supports most of the basic arguments from Apache Bench.  A simple example:

    hurl http -c 100 -n 10000 http://127.0.0.1/

Would run 100 concurrent clients until 10,000 requests have been completed against `http://127.0.0.1/`.

`hurl` also has support for customized test plans.  Test plans are defined in a JSON file, for example
the previous example could be written as:

```json
    {
      name: '100 client test',
      tasks: [
        {'url': 'http://127.0.0.1/',
         'count': 10000,
         'concurrency': 100}
      ]
    }
```

And then ran as `hurl http --config 100clients.json`.   While this example is more verbose, the JSON format
enables building complicated and large test plans hitting multiple URLs, or ramping up traffic.

## Scaling Out

`hurl` also provides a daemon mode, which can be ran across many hosts to generate distributed load from many machines.

```sh
hurl daemon -p 8000 -s secret &
hurl daemon -p 8001 -s secret &
hurl http -cluster :8000,:8001 -c 100 -n 10000 http://127.0.0.1/
```


## Not just HTTP!

`hurl` also can benchmark other services, like MongoDB:

```sh

hurl mongo -c 100 -n 100000 -rwmix 80 mongodb://10.0.0.1/mydatabase
```

The scaling out abilities also work with the non-HTTP smasher too.

Available "smashers":

* `http`
* `etcd`
* `mongo`

