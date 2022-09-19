FROM golang:latest as builder

WORKDIR /app
ENV GIN_MODE release

COPY ./back/go.mod  go.mod
COPY ./back/go.sum  go.sum

RUN go mod download
RUN go mod verify

COPY ./back /app

RUN go build -o haikaraapp


FROM chromedp/headless-shell:latest as runner
WORKDIR /app
ENV GIN_MODE release

COPY --from=builder /app/haikaraapp .
COPY --from=builder /app/.env.production .env
COPY --from=builder /app/google-application-credentials.json google-application-credentials.json

CMD ["./haikaraapp"]