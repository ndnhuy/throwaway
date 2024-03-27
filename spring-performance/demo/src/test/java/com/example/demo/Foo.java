package com.example.demo;

import java.nio.charset.Charset;
import java.util.concurrent.ThreadLocalRandom;

import org.apache.commons.lang3.RandomStringUtils;
import org.junit.Test;

public class Foo {
  @Test
  public void foo() {
    System.out.println(RandomStringUtils.randomAlphabetic(3));
  }

}
