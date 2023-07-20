#syntax=docker/dockerfile:1.2
FROM golang:1.20-alpine as builder

WORKDIR /app 
COPY . /app
COPY .github/scripts/git-askpass-helper /

RUN --mount=type=ssh --mount=type=secret,id=github_token \
    apk --no-cache add git && . /git-askpass-helper && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" .

FROM alpine
EXPOSE 8080
COPY --from=builder /app/app /usr/local/bin/
CMD ["/usr/local/bin/app"]
