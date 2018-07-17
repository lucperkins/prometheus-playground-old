FROM golang:1.10 AS builder

ARG goPackage

WORKDIR $GOPATH/src/${goPackage}

ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /web .

FROM scratch
COPY --from=builder /web ./
ENTRYPOINT ["./web"]