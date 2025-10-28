FROM alpine:latest AS buildenv
LABEL authors="rye"

RUN apk add go
RUN apk add npm
RUN npm install -g @go-task/cli
WORKDIR /build
COPY . .
RUN task

FROM alpine:latest

RUN apk add pandoc-cli
RUN apk add typst
RUN apk add curl

ARG DEBUG_LOGGING=true

ENV LISTEN_ON="0.0.0.0:3030"
ENV LATEX_COMMAND="/usr/bin/pdflatex"
ENV GOTEX_ENABLE="false"
ENV DEBUG=$DEBUG_LOGGING
ENV MATH_RENDERING_ENGINE=mathml
EXPOSE 3030

WORKDIR /app
COPY --from=buildenv /build/out/simple-pandoc-server /app/simple-pandoc-server

HEALTHCHECK --interval=30s --timeout=5s CMD curl -f http://localhost:3030/health || exit 1

ENTRYPOINT ["/app/simple-pandoc-server"]