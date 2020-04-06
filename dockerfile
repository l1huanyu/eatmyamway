FROM golang
WORKDIR /
RUN git clone https://github.com/l1huanyu/eatmyamway.git
WORKDIR /eatmyamway
RUN git checkout -b br_to_develop_scheduler origin/br_to_develop_scheduler
COPY config.yaml .
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build cmd/server.go
ENTRYPOINT ["/eatmyamway/server", "-c", "/eatmyamway/config.yaml"]