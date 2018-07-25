# vloadgenerator



##To run with docker-compose 
```
cd docker-compose/
FREQUENCY=<Number> DURATION=<Number> docker-compose -f docker-compose-hsl-withoutsl.yaml up
FREQUENCY=<Number> DURATION=<Number> docker-compose -f docker-compose-hsl-withsl.yaml up
```





##To Build 

```docker build -t vloadgenerator:v1 .```


##To Run vloadgenerator alone

```
docker run --rm -it vloadgenerator:v1 help

docker run --rm -it vloadgenerator:v1 attack --help
```

