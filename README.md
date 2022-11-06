# SunSwitcher

Lampen ein und ausschalten.....

## BUILD
```
docker build -t jorie70/sunswitcher:0.0.11 -t jorie70/sunswitcher:latest .
```

## PUSH

```
docker push jorie70/sunswitcher:latest
docker push jorie70/sunswitcher:0.0.11
```

### VERSIONS
- 0.0.11 fix DST
- 0.0.10 Switch to paho-mqtt package
- 0.0.9 Fix mqtt error handling
- 0.0.8 Fix keepalive timeout for mosquitto 2
- 0.0.7 Fix error handling when send mqtt failed
- 0.0.6 Change to Sonoff 8125 


