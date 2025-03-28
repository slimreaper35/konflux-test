FROM registry.access.redhat.com/ubi9/go-toolset@sha256:10ae14c5f0a22d8eb4a5dad45c0330233cb9dd8cc01c370a1d3351f5243d5489

USER root

LABEL maintainer="Michal Šoltis <msoltis@redhat.com>"

WORKDIR /licenses

COPY LICENSE .

WORKDIR /app

COPY go.mod go.sum .

RUN go mod download

COPY . .

RUN go build

LABEL name="name" \
      summary="summary" \
      description="description" \
      com.redhat.component="component" \
      io.k8s.description="description" \
      io.k8s.display-name="display-name" \
      io.openshift.expose-services="8080:http" \
      io.openshift.tags="tags"

EXPOSE 8080

RUN chown -R 1001:0 /app

USER 1001

CMD ["./konflux-test"]
