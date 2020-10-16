# Build the project
FROM ubuntu AS builder

LABEL maintainer="Nicolas Goldack <admin@ntec.io>"

ARG BUILD_TAG

ENV DEBIAN_FRONTEND noninteractive

# Configuring apt-get
RUN echo 'Acquire::Retries "10";' > /etc/apt/apt.conf.d/80retry && \
    echo 'APT::Install-Recommends "false";' > /etc/apt/apt.conf.d/80recommends && \
    echo 'APT::Install-Suggests "false";' > /etc/apt/apt.conf.d/80suggests && \
    echo 'APT::Get::Assume-Yes "true";' > /etc/apt/apt.conf.d/80forceyes && \
    echo 'APT::Get::Fix-Missing "true";' > /etc/apt/apt.conf.d/80fixmissin

# Installing dev/build dependencies
RUN apt-get update && apt-get upgrade && \
    apt-get install \
    build-essential \
    curl \
    chrpath \
    pkg-config \
    software-properties-common \
    rsync \
    unzip \
    apt-utils \
    seccomp \
    make \
    git \
    tzdata \
    wget


# Set default https://en.wikipedia.org/wiki/Umask
RUN umask 0000

# Install Go
ENV GOLANG_VERSION 1.15.3
RUN set -eux; \
	\
	url="https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz"; \
	wget -O go.tgz "$url"; \
	echo "010a88df924a81ec21b293b5da8f9b11c176d27c0ee3962dc1738d2352d3c02d *go.tgz" | sha256sum -c -; \
	tar -C /usr/local -xzf go.tgz; \
	rm go.tgz; \
	export PATH="/usr/local/go/bin:$PATH"; \
	go version

# Configure Go environment
ENV GOPATH /go
ENV GOBIN $GOPATH/bin
ENV PATH $GOBIN:/usr/local/go/bin:/root/.local/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
ENV TMPDIR /tmp
ENV GO111MODULE on
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

WORKDIR /nostradamus

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download -x all

# Copy project files
COPY . .

# Test project
RUN go test -v .

# Build the project
RUN GOOS=linux go build -a -v -o app .


# Runtime docker
FROM alpine:latest AS runner
RUN apk --no-cache add ca-certificates

# Add docker-compose-wait tool
ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

WORKDIR /root/

COPY --from=builder /nostradamus/app .
CMD ["./app"]  