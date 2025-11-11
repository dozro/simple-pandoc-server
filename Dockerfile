FROM alpine:20250108 AS buildenv

RUN apk --no-cache add go
RUN apk --no-cache add npm
RUN npm install -g @go-task/cli
WORKDIR /build
COPY . .
RUN task

FROM alpine:20250108

LABEL org.opencontainers.image.authors="Rye <docker@itsrye.dev>"
LABEL org.opencontainers.image.source="https://github.com/dozro/simple-pandoc-server"
LABEL org.opencontainers.image.title="Simple Pandoc Server"
LABEL org.opencontainers.image.documentation="https://github.com/dozro/simple-pandoc-server/wiki"
LABEL org.opencontainers.image.vendor="itsrye.dev"
LABEL org.opencontainers.image.licenses="Hippocratic-2.1"

RUN apk --no-cache add pandoc-cli
RUN apk --no-cache add typst
RUN apk --no-cache add curl

# add user to run server
RUN adduser -S -u 8231 sps

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

USER sps
ENTRYPOINT ["/app/simple-pandoc-server"]