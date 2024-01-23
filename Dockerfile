FROM golang:alpine as builder

WORKDIR /builder

ADD go.mod /builder/
RUN go mod download

ADD . /builder
RUN go build -o /bin/ciagent -ldflags "-s -w" .

FROM ubuntu:latest

RUN apt update && apt upgrade -y && \
    apt install build-essential python3 python3-pip python3-venv -y && \
    python3 -m pip install --upgrade pip
COPY requirements.txt /requirements.txt
RUN python3 -m pip install -r /requirements.txt && mkdir /sandbox && mkdir /data
COPY --from=builder /bin/ciagent /bin/ciagent

WORKDIR /sandbox
ENV PYTHONUNBUFFERED=1

ENTRYPOINT ["/bin/ciagent"]
