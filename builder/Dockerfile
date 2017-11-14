FROM ubuntu

RUN apt-get update \
    && apt-get install -y curl software-properties-common python-software-properties apt-transport-https ca-certificates \
    && apt-get install -y python-pip python-dev build-essential \
    && rm -rf /var/lib/apt/lists/*

# set up source of bazel
RUN echo "deb [arch=amd64] http://storage.googleapis.com/bazel-apt stable jdk1.8" | tee /etc/apt/sources.list.d/bazel.list
RUN curl https://bazel.build/bazel-release.pub.gpg | apt-key add -
RUN curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
RUN apt-key fingerprint 0EBFCD88
RUN add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"


# install java and bazel and docker
RUN apt-get update \
    && apt-get install -y openjdk-8-jdk \
    && apt-get install -y bazel \
    && apt-get install -y docker-ce \
    && rm -rf /var/lib/apt/lists/*

VOLUME ["/source"]
WORKDIR "/source"
ENTRYPOINT [ "bazel" ]
CMD ["help"]
