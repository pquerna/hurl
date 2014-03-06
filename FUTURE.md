# This is README driven development.

## Repeatable Config File

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
hurl http -cluster :8000,:8001 -s secret -c 100 -n 10000 http://127.0.0.1/
```

TODO: Consider SSH based -cluster?