package com.example.demo;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;

@SpringBootApplication
@RestController
public class DemoApplication {

	@Autowired
	private HeavyTaskService service;

	public static void main(String[] args) {
		SpringApplication.run(DemoApplication.class, args);
	}

	@GetMapping("/start")
	public String startHeavyTask() {
		service.start();
		return "done";
	}

	@GetMapping("/hello")
	public String getMethodName(@RequestParam String name) {
		return "hello " + name;
	}

}
