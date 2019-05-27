FROM node as node-builder

RUN mkdir /home/node/robo
WORKDIR /home/node/robo
COPY client/package.json .
RUN make
COPY client ./
RUN npm run build-prod


# ---
FROM golang as golang-builder

COPY . "$GOPATH/src/github.com/tjbearse/robo"
WORKDIR "$GOPATH/src/github.com/tjbearse/robo"
RUN go get && go build .

EXPOSE 8080
COPY --from=node-builder /home/node/robo/dist ./client/dist

ENTRYPOINT ["./robo"]