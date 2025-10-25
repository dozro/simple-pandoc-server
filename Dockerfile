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

ENV LISTEN_ON="0.0.0.0:3030"
ENV LATEX_COMMAND="/usr/bin/pdflatex"
ENV GOTEX_ENABLE="false"
EXPOSE 3030

WORKDIR /app
COPY --from=buildenv /build/src/simple-pandoc-server /app/simple-pandoc-server
# COPY --from=pandoc-latex /opt/texlive/texdir/bin /opt/texlive/texdir/bin

ENTRYPOINT ["/app/simple-pandoc-server"]