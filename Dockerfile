#syntax=docker/dockerfile:1.2
FROM golang:1.20-alpine as builder

WORKDIR /app 
COPY . /app
COPY .github/scripts/git-askpass-helper /

RUN --mount=type=ssh --mount=type=secret,id=github_token \
    apk --no-cache add git alpine-sdk gcc && . /git-askpass-helper && \
    CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" .

FROM alpine
COPY --from=builder /app/app /usr/local/bin/
CMD ["/usr/local/bin/app"]
