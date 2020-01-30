FROM openjdk:11
MAINTAINER Gaurav Abbi <gaurav.abbi@sap.com>

WORKDIR /var/app

ADD my-java-extension-*.jar /var/app/my-java-extension.jar

ENV JAVA_OPTS=""

EXPOSE 8080

CMD java $JAVA_OPTS -jar /var/app/my-java-extension.jar