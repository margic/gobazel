FROM ubuntu

RUN apt-get update \
    && apt-get install -y curl software-properties-common python-software-properties apt-transport-https ca-certificates pass \
    && apt-get install -y python-pip python-dev build-essential \
    && rm -rf /var/lib/apt/lists/*

# set up source of bazel
RUN echo "deb [arch=amd64] http://storage.googleapis.com/bazel-apt stable jdk1.8" | tee /etc/apt/sources.list.d/bazel.list
RUN curl https://bazel.build/bazel-release.pub.gpg | apt-key add -

RUN add-apt-repository ppa:gophers/archive
# install java and bazel and go
RUN apt-get update \
    && apt-get install -y openjdk-8-jdk \
    && apt-get install -y bazel \
    && apt-get install -y git \
    && apt-get install -y golang-1.9-go \
    && rm -rf /var/lib/apt/lists/*

ENV PATH=$PATH:/usr/lib/go-1.9/bin
ENV GOPATH=/go

# install docker cred helper
RUN mkdir -p /go/src/github.com/docker \
    && cd $GOPATH/src/github.com/docker \
    && git clone https://github.com/docker/docker-credential-helpers.git \
    && cd docker-credential-helpers \
    && make pass \
    && mkdir -p $GOPATH/bin \
    && mv bin/docker-credential-pass $GOPATH/bin/docker-credential-pass

ENV PATH=$PATH:$GOPATH/bin
RUN mkdir -p ~/.docker
COPY config.json ~/.docker/config.json

# init pass store create key
COPY initpass.sh .
RUN ./initpass.sh

# Volume for mounting project source
VOLUME ["/source"]
WORKDIR "/source"
ENTRYPOINT [ "bazel" ]
CMD ["help"]
