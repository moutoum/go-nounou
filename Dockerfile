# syntax=docker/dockerfile:1.2
# Alpine
FROM alpine

RUN apk --no-cache --no-progress add ca-certificates tzdata git \
    && rm -rf /var/cache/apk/*

ARG TARGETPLATFORM
COPY ./dist/$TARGETPLATFORM/payslip /
COPY ./templates/get-payslip.html /templates/get-payslip.html

ENTRYPOINT ["/payslip"]
EXPOSE 80
