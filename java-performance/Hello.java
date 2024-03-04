import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.Executors;
import java.util.concurrent.Future;

class Hello {
  public static void main(String[] args) throws InterruptedException, ExecutionException {
    System.out.println(sum(10000000));
  }

  static long sum(long n) throws InterruptedException, ExecutionException {
    var lock = new Object();
    var executor = Executors.newFixedThreadPool(5);
    List<Future<Long>> futures = new ArrayList<>();
    for (long i = 0; i < n; i+=2) {
        var k = i+1;
        if (i+1 >= n) {
            futures.add(executor.submit(() -> {
                synchronized (lock) {
                    return k;
                }
            }));
        } else {
            futures.add(executor.submit(() -> {
                synchronized (lock) {
                    return 2*k + 1;
                }
            }));
        }
    }
    var sum = 0L;
    for (var f : futures) {
        sum += f.get();
    }
    executor.shutdown();
    return sum;
  }
}