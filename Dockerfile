FROM registry.access.redhat.com/ubi9/go-toolset@sha256:97e30a01caeece72ee967013e7c7af777ea4ee93840681ddcfe38a87eb4c084a

LABEL name="my-name"
LABEL summary="my-summary"
LABEL description="my-description"
LABEL com.redhat.component="my-redhat-component"
LABEL io.openshift.tags="my-openshift-tags"
LABEL io.k8s.display-name="my-k8s-display-name"
LABEL io.k8s.description="my-k8s-description"

ARG REVISION
LABEL commit_sha=$REVISION

USER root

ARG USERNAME=nonroot
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN groupadd --gid $USER_GID $USERNAME && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME

WORKDIR /licenses

COPY LICENSE .

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o rest-api

EXPOSE 8080

USER $USERNAME

CMD ["./rest-api"]
