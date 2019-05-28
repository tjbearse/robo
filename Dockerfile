FROM node as node-builder

RUN apt-get update && apt-get install -y \
inkscape && rm -rf /var/lib/apt/lists/*

RUN mkdir /home/node/robo
WORKDIR /home/node/robo
COPY client/package.json ./client/
RUN (cd client && npm install)
COPY Makefile .
COPY client ./client/
RUN make client


# ---
FROM golang as golang-builder

COPY . "$GOPATH/src/github.com/tjbearse/robo"
WORKDIR "$GOPATH/src/github.com/tjbearse/robo"
RUN go get && go build .

EXPOSE 8080
COPY --from=node-builder /home/node/robo/client/dist ./client/dist

ENTRYPOINT ["./robo"]