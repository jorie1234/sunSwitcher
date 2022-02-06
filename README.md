# SunSwitcher

Lampen ein und ausschalten.....

## BUILD
```
docker build -t jorie70/sunswitcher:0.0.8 -t jorie70/sunswitcher:latest .
```

## PUSH

```
docker push jorie70/sunswitcher:latest
docker push jorie70/sunswitcher:0.0.8
```

### VERSIONS
0.0.6 Change to Sonoff 8125 

0.0.7 Fix error handling when send mqtt failed

0.0.8 Fix keepalive timeout for mosquitto 2  