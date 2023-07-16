FROM harbor.inner.galaxy.xxxxx.com/xxxxx/golang:1.20.5 as Builder
ARG COMMIT=""
ARG VERSION=""
ARG SERVER_DIR="."
ARG PKG_NAME="ezone.xxxxx.com/xxxxx/xxxxx/alarm"

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /workshop
COPY ./ ./
RUN GoVersion=$(go version) && \
    BuildTime=$(date) && \
    if [ ! -d vendor ]; then go mod vendor; fi && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor \
    -ldflags "-X '${PKG_NAME}/common/version/version.GoVersion=${GoVersion}' \
    -X '${PKG_NAME}/common/version/version.GitHash=${COMMIT}' \
    -X '${PKG_NAME}/common/version/version.BuildTime=${BuildTime}' \
    -X '${PKG_NAME}/common/version/version.Version=${VERSION}'" \
    -a -o ./_output/alarmv2 ${SERVER_DIR}

FROM alpine:3.17 as base
WORKDIR /app/
COPY --from=builder /usr/share/zoneinfo/ /usr/share/zoneinfo/
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /workshop/_output/alarmv2 ./alarmv2
ENTRYPOINT ["/app/alarmv2"]