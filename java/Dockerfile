FROM maven:3.6.3-jdk-8 

COPY ./ ./  
RUN mvn clean package 
RUN wget -O dd-java-agent.jar 'https://dtdg.co/latest-java-tracer'

CMD ["java", "-javaagent:dd-java-agent.jar", "-Ddd.agent.host=dd-agent", "-jar", "target/ddjavastack-0.0.1.jar"]