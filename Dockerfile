FROM maven:3.6.3-openjdk-14-slim AS build

COPY settings.xml /usr/share/maven/conf/

COPY pom.xml pom.xml
COPY tds-api/pom.xml tds-api/pom.xml
COPY tds-base/pom.xml tds-base/pom.xml

RUN mvn dependency:go-offline package -B

COPY tds-api/src tds-api/src
COPY tds-base/src tds-base/src

RUN mvn install

FROM openjdk:14-ea-jdk-alpine
USER root

RUN mkdir service

COPY --from=build /tds-base/target/ /service/
COPY config.yaml /service/

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.5.0/wait /wait

RUN chmod +x /wait

ENV JAVA_TOOL_OPTIONS -agentlib:jdwp=transport=dt_socket,server=y,suspend=n,address=*:5005

EXPOSE 5005

CMD /wait && java --enable-preview -jar /service/tds-base-1.0-SNAPSHOT.jar -Xdebug