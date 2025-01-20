FROM registry.access.redhat.com/ubi9/go-toolset@sha256:6ec9c3ce36c929ff98c1e82a8b7fb6c79df766d1ad8009844b59d97da9afed43

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
