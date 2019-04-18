FROM centos:7 as build-tools
LABEL maintainer "Devtools <devtools@redhat.com>"
LABEL author "Devtools <devtools@redhat.com>"
ENV LANG=en_US.utf8
ENV GOPATH /tmp/go
ARG GO_PACKAGE_PATH=github.com/redhat-developer/devconsole-api

ENV GIT_COMMITTER_NAME devtools
ENV GIT_COMMITTER_EMAIL devtools@redhat.com

RUN yum install epel-release -y \
    && yum install --enablerepo=centosplus install -y --quiet \
    findutils \
    git \
    golang \
    make \
    procps-ng \
    tar \
    wget \
    which \
    bc \
    && yum clean all

# install dep
RUN mkdir -p $GOPATH/bin && chmod a+rwx $GOPATH \
    && curl -L -s https://github.com/golang/dep/releases/download/v0.5.1/dep-linux-amd64 -o dep \
    && echo "7479cca72da0596bb3c23094d363ea32b7336daa5473fa785a2099be28ecd0e3  dep" > dep-linux-amd64.sha256 \
    && sha256sum -c dep-linux-amd64.sha256 \
    && rm dep-linux-amd64.sha256 \
    && chmod +x ./dep \
    && mv dep $GOPATH/bin/dep

ENV PATH=$PATH:$GOPATH/bin

RUN mkdir -p ${GOPATH}/src/${GO_PACKAGE_PATH}/

WORKDIR ${GOPATH}/src/${GO_PACKAGE_PATH}

ENTRYPOINT [ "/bin/bash" ]

#--------------------------------------------------------------------

FROM build-tools as builder
ARG VERBOSE=2
COPY . .
RUN make VERBOSE=${VERBOSE} build
