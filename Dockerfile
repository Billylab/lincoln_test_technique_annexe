FROM golang:1.13 as builder
WORKDIR /app
COPY invoke.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server

FROM python:3.9.9
USER root
# Install required system packages and cleanup to reduce image size
RUN apt-get update -y && \
  apt-get install --no-install-recommends -y -q \
  libpq-dev python-dev  && \
  apt-get clean && \
  rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Set working directory
WORKDIR /lincoln_test_technique

# Copy files to the image
COPY --from=builder /app/server ./
COPY script.sh ./
COPY feature ./feature

# Install librairies
RUN pip install -r /lincoln_test_technique/feature/requirements.txt

ENTRYPOINT "./server"