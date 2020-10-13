docker build -f da0x/go1.15-py3-proto-alpine -t da0x/go1.15-py3-proto-alpine da0x/
docker build -f da0x/go1.15-proto-alpine -t da0x/go1.15-proto-alpine da0x/
docker push da0x/go1.15-py3-proto-alpine
docker push da0x/go1.15-proto-alpine
