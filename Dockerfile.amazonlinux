FROM public.ecr.aws/amazonlinux/amazonlinux:latest

# Install build-base
RUN yum install -y git golang gcc-c++ wget tar make python3 cmake3
RUN ln -sf /usr/bin/cmake3 /usr/bin/cmake

# Install go-khaiii
RUN mkdir -p /root/go/src
WORKDIR /root/go/src
RUN git clone https://github.com/AhaOfficial/go-khaiii.git

WORKDIR /root/go/src/go-khaiii
RUN bash install_khaiii.sh

# Export variables
ENV LD_LIBRARY_PATH="$LD_LIBRARY_PATH:/usr/local/lib"

# Run go
RUN go get github.com/AhaOfficial/go-khaiii
RUN go mod init
RUN go mod tidy
RUN go test

WORKDIR /root
