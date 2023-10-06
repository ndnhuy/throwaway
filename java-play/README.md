** manually compile and run
javac -cp "lib/*" HelloWorld.java
java -jar lib/junit-platform-console-standalone-1.9.3.jar -cp . --select-class HelloWorld

** maven run
1/ mvn clean compile assembly:single
2/ java -jar target/javaplay-1.0-SNAPSHOT-jar-with-dependencies.jar