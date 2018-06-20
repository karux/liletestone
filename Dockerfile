FROM scratch
#FROM alpine:latest

WORKDIR /app
COPY ./bin/liletestone-linux ./bin/liletestone-linux
COPY ./scripts/step06-run.sh ./scripts/step06-run.sh

ENV SERVICE_HOST_OVERRIDE=localhost:8000
ENV GOOS=linux
ENV GODEBUG=gcpacertrace=1,gctrace=1
# scheddetail=1,schedtrace=1000
ENV GOGC=30

EXPOSE 8000

CMD ./scripts/step06-run.sh
