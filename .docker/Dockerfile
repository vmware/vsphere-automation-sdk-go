FROM golang:1.13-alpine3.12

RUN apk update && apk add --no-cache bash git openjdk11

RUN mkdir -p /usr/local/infrasdk_go_gradle
ADD gradle /usr/local/infrasdk_go_gradle

ENV JAVA_HOME /usr/lib/jvm/java-11-openjdk
ENV PATH $JAVA_HOME/bin:$PATH
CMD [ "/bin/bash"]