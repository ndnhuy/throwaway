FROM eclipse-temurin:17-jdk-alpine as build

WORKDIR /app

COPY .mvn/ .mvn
COPY mvnw pom.xml ./
# RUN ./mvnw dependency:go-offline
COPY src ./src
RUN  ./mvnw install -DskipTests

FROM eclipse-temurin:17-jre-alpine
COPY --from=build /app/target/*.jar app.jar
ENTRYPOINT [ "java", "-jar", "-Dspring.profiles.active=default,docker", "/app.jar" ]
# CMD ["./mvnw", "spring-boot:run", "-Dspring-boot.run.profiles=default,docker"]
# CMD ["./mvnw", "spring-boot:run", "-Dspring.profiles.active=default,docker"]