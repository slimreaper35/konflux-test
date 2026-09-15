FROM registry.access.redhat.com/ubi10/go-toolset@sha256:290ba654458e9a269b1509d10e6ebbd3c2b2456570e73e73201adb3ee54fb244

ARG COLOR
ARG DOG
LABEL color="${COLOR}" dog="${DOG}"

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

ENTRYPOINT ["./konflux-test"]
