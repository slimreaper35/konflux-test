FROM registry.access.redhat.com/ubi9/go-toolset@sha256:71d97f49c67b22c5e26f5670dea9dd2e8c42de78479e51a59d5a8abc917c1822

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
