FROM registry.access.redhat.com/ubi9/go-toolset@sha256:ead35188c5748efe2b9420352aba56b02b43d8fcd7e879cc96c6b9ac2548e454

USER root

WORKDIR /licenses

COPY LICENSE .

WORKDIR /app

COPY go.mod go.sum .

RUN go mod download

COPY . .

RUN go build

LABEL name="my-name" \
      summary="my-summary" \
      description="my-description" \
      com.redhat.component="my-redhat-component" \
      io.openshift.tags="my-openshift-tags" \
      io.k8s.display-name="my-k8s-display-name" \
      io.k8s.description="my-k8s-description"

EXPOSE 8080

USER 1001

ENTRYPOINT ["./konflux-test"]
