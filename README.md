# merge-requests

## Introduce
In PHP single-threaded applications, it is used to accelerate access to multiple remote calls.

## Usage
```
./bin/merge.sh -k start

./bin/merge.sh -k stop

./bin/merge.sh -k restart
```

## Example
```
http://{your-server-ip}:8080/merge?urls=http://www.a.com/a.json&urls=http://www.b.com/b.json
```

## Performance
Depending on the response time of the remote call.
