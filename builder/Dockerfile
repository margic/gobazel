FROM ubuntu

RUN apt-get update \
    && apt-get install -y curl \
    && rm -rf /var/lib/apt/lists/*

# set up source of bazel
RUN echo "deb [arch=amd64] http://storage.googleapis.com/bazel-apt stable jdk1.8" | tee /etc/apt/sources.list.d/bazel.list
RUN curl https://bazel.build/bazel-release.pub.gpg | apt-key add -

# install java and bazel
RUN apt-get update \
    && apt-get install -y openjdk-8-jdk \
    && apt-get install -y bazel \
    && rm -rf /var/lib/apt/lists/*




