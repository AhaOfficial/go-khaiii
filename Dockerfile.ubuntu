FROM ubuntu:20.04
ARG DEBIAN_FRONTEND=noninteractive
ENV TZ=Asia/Seoul

# Install build-base
RUN mkdir -p /downloads
RUN apt update && apt-get install -y tzdata build-essential git cmake qt5-default golang-go language-pack-ko libssl-dev python3

# Install go-khaiii
RUN mkdir -p /root/go/src
WORKDIR /root/go/src
RUN git clone https://github.com/AhaOfficial/go-khaiii.git

WORKDIR /root/go/src/go-khaiii
RUN make
RUN make install

# Export variables
ENV LD_LIBRARY_PATH="$LD_LIBRARY_PATH:/usr/local/lib"
ENV CGO_LDFLAGS="-L/usr/local/lib -lkhaiiic"

# Run go
RUN go get github.com/AhaOfficial/go-khaiii
RUN echo "en_US.UTF-8 UTF-8" > /etc/locale.gen && locale-gen
RUN go mod init
RUN go mod tidy
RUN go test

WORKDIR /root
