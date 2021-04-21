FROM golang:alpine

WORKDIR /build

COPY ./DataAnalysisgo/go.mod .

RUN go mod download

COPY ./DataAnalysis .

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .

EXPOSE 9004

CMD ["/dist/main"]