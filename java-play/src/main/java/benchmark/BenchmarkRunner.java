package benchmark;

import java.util.concurrent.TimeUnit;

import org.openjdk.jmh.annotations.Benchmark;
import org.openjdk.jmh.annotations.BenchmarkMode;
import org.openjdk.jmh.annotations.Mode;
import org.openjdk.jmh.annotations.OutputTimeUnit;
import org.openjdk.jmh.runner.Runner;
import org.openjdk.jmh.runner.options.Options;
import org.openjdk.jmh.runner.options.OptionsBuilder;

public class BenchmarkRunner {
  public static void main(String[] args) throws Exception {
    Options opt = new OptionsBuilder().include(BenchmarkRunner.class.getSimpleName()).forks(1).build();
    new Runner(opt).run();
  }

  @Benchmark
  @BenchmarkMode(Mode.Throughput)
  @OutputTimeUnit(TimeUnit.SECONDS)
  public void measureThroughput() throws InterruptedException {
    TimeUnit.MILLISECONDS.sleep(100);
  }

  @Benchmark
  @BenchmarkMode(Mode.AverageTime)
  @OutputTimeUnit(TimeUnit.SECONDS)
  public void measureAvgTime() throws InterruptedException {
    TimeUnit.MILLISECONDS.sleep(100);
  }
}
