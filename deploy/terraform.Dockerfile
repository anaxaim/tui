FROM alpine:3.17.3

ARG TERRAFORM_VERSION
ENV TERRAFORM_VERSION=${TERRAFORM_VERSION}

RUN apk add --update --no-cache ca-certificates curl unzip \
    && curl -sSL https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip -o /tmp/terraform.zip \
    && unzip /tmp/terraform.zip -d /usr/local/bin \
    && rm -f /tmp/terraform.zip \
    && apk del curl unzip \
    && rm -rf /var/cache/apk/* \
    && mkdir /terraform \
    && chown nobody:nogroup /terraform

USER nobody
CMD ["sh", "-c", "sleep infinity"]
