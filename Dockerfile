FROM alpine:latest AS buildenv
LABEL authors="rye"

RUN apk add go
WORKDIR /build
COPY . .
WORKDIR /build/src
RUN go build

FROM alpine:latest

RUN apk add pandoc-cli
RUN apk add typst
RUN apk add curl

ENV LISTEN_ON="0.0.0.0:3030"
ENV LATEX_COMMAND="/usr/bin/pdflatex"
ENV GOTEX_ENABLE="false"
EXPOSE 3030

WORKDIR /app
COPY --from=buildenv /build/src/simple-pandoc-server /app/simple-pandoc-server

HEALTHCHECK --interval=30s --timeout=5s CMD curl -f http://localhost:3030/health || exit 1

ENTRYPOINT ["/app/simple-pandoc-server"]