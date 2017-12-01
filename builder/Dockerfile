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

# install kubectl for deploy to kubernetes
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl \
    -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl \
    && chmod +x ./kubectl \
    && mv ./kubectl /usr/local/bin/kubectl

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
RUN mkdir -p /root/.docker
COPY config.json /root/.docker/config.json

RUN mkdir -p /opt/gobazel
# init pass store for the docker credential pass to work
COPY initpass.sh /opt/gobazel/initpass.sh
COPY addpass.sh /opt/gobazel/addpass.sh
COPY runbazel.sh /opt/gobazel/runbazel.sh
RUN /opt/gobazel/initpass.sh


# Volume for mounting project source
VOLUME ["/root/.kube"]
VOLUME ["/source"]
WORKDIR "/source"
ENTRYPOINT [ "/opt/gobazel/runbazel.sh" ]
CMD ["help"]
