FROM golang:alpine as builder

ADD go.mod /build/
ADD main.go /build/
ADD sumupper/ /build/sumupper/
WORKDIR /build/
RUN go build main.go

FROM golang:alpine

COPY VERSION .
WORKDIR /home/
COPY --from=builder /build/main .

ENTRYPOINT [ "./main" ]