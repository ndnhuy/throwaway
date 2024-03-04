package com.example.demo;

import org.springframework.stereotype.Service;

@Service
public class HeavyTaskService {
  private void printEven() {
    System.out.println("even");
  }

  private void printOdd() {
    System.out.println("odd");
  }

  public void start() {
    for (int i = 0; i < 10000; i++) {
      if (i % 2 == 0) {
        printEven();
      } else {
        printOdd();
      }
    }
  }
}
