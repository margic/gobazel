FROM ubuntu

RUN apt-get update \
    && apt-get install -y \
        curl \
        software-properties-common \
        python-software-properties \
        apt-transport-https \
        ca-certificates pass \
        python-pip \
        python-dev \
        build-essential \
        openjdk-8-jdk \
        git \
    && rm -rf /var/lib/apt/lists/*

ENV PATH=$PATH:/usr/lib/go-1.9/bin
ENV GOPATH=/go

# add go repo
RUN add-apt-repository ppa:gophers/archive

# set up source of bazel
RUN echo "deb [arch=amd64] http://storage.googleapis.com/bazel-apt stable jdk1.8" | tee /etc/apt/sources.list.d/bazel.list
RUN curl https://bazel.build/bazel-release.pub.gpg | apt-key add -

# dependencies in place add bazel and go
RUN apt-get update \
    && apt-get install -y bazel \
    golang-1.9-go \
    && rm -rf /var/lib/apt/list/*

# install kubectl for deploy to kubernetes
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl \
    -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl \
    && chmod +x ./kubectl \
    && mv ./kubectl /usr/local/bin/kubectl

ENV KUBECONFIG=/source/kube/config

RUN mkdir -p /opt/gobazel
# init pass store for the docker credential pass to work
COPY runbazel.sh /opt/gobazel/runbazel.sh

# Volume for mounting project source
VOLUME ["/source"]
WORKDIR "/source"
ENTRYPOINT [ "/opt/gobazel/runbazel.sh" ]
CMD ["help"]
