# BASE
###############
FROM golang:1.19.0-alpine3.15 AS base

WORKDIR /app
ADD . .
#RUN mv .netrc ~/.netrc

RUN apk update && apk add build-base git
RUN go mod download
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -buildvcs=false -o main .

# DEVELOPMENT
###############
FROM base AS dev

RUN go install github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon --build="go build -o main -buildvcs=false" --command="./main" 

# PRODUCTION
###############
FROM scratch AS prod

COPY --from=base /app/main ./main
CMD ["./main"]
