FROM registry.access.redhat.com/ubi10/go-toolset@sha256:efa721f5d675f0494425ea582823d4324f546c0e19f13ca996942a1c354c8497

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
