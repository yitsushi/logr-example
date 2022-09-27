```
go run .

my realm				start process				
my realm				message					step: download
INFO[0000] start process                                 logger="my realm"
INFO[0000] message                                       logger="my realm" step=download
```

Or with a less fancy output with logrus:

```
go run . 2> output && cat output
my realm				start process				
my realm				message					step: download  
time="2022-09-27T12:39:50+02:00" level=info msg="start process" logger="my realm"
time="2022-09-27T12:39:50+02:00" level=info msg=message logger="my realm" step=download
```
