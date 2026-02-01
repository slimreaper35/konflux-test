FROM registry.access.redhat.com/ubi10/go-toolset@sha256:46c65f0a0b9dd5802490fb7908e8a99c037bb7633fcc66fd84a05982ac07cc6f

USER root

LABEL maintainer="Michal Šoltis <msoltis@redhat.com>"

WORKDIR /licenses

COPY LICENSE .

WORKDIR /app

COPY go.mod go.sum ./

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
