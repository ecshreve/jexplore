FROM "mcr.microsoft.com/devcontainers/base:ubuntu-22.04"

RUN export DEBIAN_FRONTEND=noninteractive \
    && apt-get update && apt-get upgrade \
    && apt-get -y install --no-install-recommends \
    build-essential \
    cmake \
    git \
    libboost-all-dev \
    libeigen3-dev \
    libgflags-dev \
    libgoogle-glog-dev \
    libopencv-dev \
    libprotobuf-dev \
    protobuf-compiler \
    python3-dev \
    python3-pip \
    python3-setuptools \
    python3-wheel \
    curl \
    wget \
    software-properties-common \
    && apt-get autoremove -y && apt-get clean -y && rm -rf /var/lib/apt/lists/*
