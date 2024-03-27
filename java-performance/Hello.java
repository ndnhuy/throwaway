import java.util.concurrent.ExecutionException;

class Hello {
  public static void main(String[] args) throws InterruptedException, ExecutionException {

  }

  static class Banker {
    public void transfer(Account fromAcc, Account toAcc, long amount) {
        fromAcc.decrease(amount);
        toAcc.increase(amount);
    }
  }
  static class Account {
    long id;
    long balance;

    void increase(long amount) {
        this.balance += amount;
    }

    void decrease(long amount) {
        if (this.balance < amount) {
            throw new RuntimeException(String.format("Account [id=%d, balance=%d] has not enough to withdraw %d", this.id, amount));
        }
        this.balance -= amount;
    }

  }
}