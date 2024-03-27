package com.example.demo;

import static org.assertj.core.api.Assertions.assertThat;

import java.util.concurrent.TimeUnit;

import org.junit.Test;
import org.openjdk.jmh.annotations.Benchmark;
import org.openjdk.jmh.annotations.BenchmarkMode;
import org.openjdk.jmh.annotations.Mode;
import org.openjdk.jmh.annotations.OutputTimeUnit;
import org.openjdk.jmh.runner.Runner;
import org.openjdk.jmh.runner.RunnerException;
import org.openjdk.jmh.runner.options.Options;
import org.openjdk.jmh.runner.options.OptionsBuilder;

public class CacheImplTest {

  public static void main(String[] args) throws RunnerException {
    Options opt = new OptionsBuilder().include(CacheImplTest.class.getSimpleName()).forks(1).build();
    new Runner(opt).run();
  }

  @Test
  public void testCache() {
    Cache<String, String> cache = new CacheImpl();
    cache.put("A", "1");
    cache.put("B", "2");
    assertGetKey(cache, "A", "1");
    assertGetKey(cache, "B", "2");
  }

  private void assertGetKey(Cache<String, String> cache, String k, String expectValue) {
    assertThat(cache.get(k)).isPresent();
    assertThat(cache.get(k).get()).isEqualTo(expectValue);
  }
  // @Benchmark
  // @BenchmarkMode(Mode.Throughput)
  // @OutputTimeUnit(TimeUnit.SECONDS)
  // public void measureThroughput() throws InterruptedException {
  // }

  @Benchmark
  @BenchmarkMode(Mode.AverageTime)
  @OutputTimeUnit(TimeUnit.SECONDS)
  public void measureAvgTime() throws InterruptedException {
    TimeUnit.MILLISECONDS.sleep(100);
  }
}
