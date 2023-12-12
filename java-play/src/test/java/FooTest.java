import java.math.BigDecimal;
import org.junit.Test;

public class FooTest {

  @Test
  public void foo() {
    Double a = 1.23;
    Double b = 1.12;
    System.out.println(
      BigDecimal.valueOf(a).subtract(BigDecimal.valueOf(b)).doubleValue()
    );
  }
}
