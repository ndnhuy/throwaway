import static org.junit.Assert.assertNull;
import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.Random;

import org.junit.Test;

import trie.StringST;

public class StringSTTest {

  @Test
  public void testFoo() {
    // byte[] arr = new byte[7];
    // new Random().nextBytes(arr);
    // String s = new String(arr);
    // System.out.println(s);

    for (int i = 0; i < 100; i++) {
      int rnd = new Random().nextInt();
      // Byte.SIZE * (Integer.SIZE / Byte.SIZE - 1)
      // rnd = rnd >> 24;
      rnd = -11;
      // 1000000 0000000 00000000 00001011
      char c = (char) rnd;
      System.out.printf("%d\n", (int) c);
      // char c = (char) (byte) rnd;
      // System.out.printf("%d=%c\n", (int) c, c);
    }
  }

  @Test
  public void testPrintBits() {
    int x = -1;
    boolean[] bits = new boolean[Integer.SIZE];
    for (int i = Integer.SIZE-1; i > 0; i--) {
      bits[i] = (x & 1) == 1;
      x = x >> 1;
    }
    for (int i = 0; i < bits.length; i++) {
      System.out.print(bits[i] ? 1 : 0);
    }
  }

  @Test
  public void testPutIfAbsent() {
    StringST<Integer> st = new StringST<>();
    st.putIfAbsent("sea", 1, v -> v + 1);
    assertEquals(1, st.get("sea"));
    st.putIfAbsent("sea", 1, v -> v + 1);
    assertEquals(2, st.get("sea"));
    st.putIfAbsent("sea", 1, v -> v + 1);
    st.putIfAbsent("sea", 1, v -> v + 1);
    assertEquals(4, st.get("sea"));

    st.putIfAbsent("sells", 1, v -> v + 1);
    st.putIfAbsent("she", 1, v -> v + 1);
    st.putIfAbsent("she", 1, v -> v + 1);
    st.putIfAbsent("shells", 1, v -> v + 1);
    st.putIfAbsent("shells", 1, v -> v + 1);
    assertEquals(1, st.get("sells"));
    assertEquals(2, st.get("she"));
    assertEquals(2, st.get("shells"));
  }

  @Test
  public void testKeyPutDeleteSize() {
    StringST<Integer> st = new StringST<>();
    doTestPutKey(st, "sea", 6);
    assertEquals(3, st.numberOfNodes());
    assertEquals(1, st.size());
    doTestPutKey(st, "by", 4);
    assertEquals(5, st.numberOfNodes());
    assertEquals(2, st.size());
    doTestPutKey(st, "sells", 1);
    doTestPutKey(st, "shells", 3);
    assertEquals(4, st.size());
    assertEquals(13, st.numberOfNodes());
    doTestPutKey(st, "she", 0);
    doTestPutKey(st, "the", 5);
    doTestPutKey(st, "shore", 7);
    assertEquals(19, st.numberOfNodes());
    assertEquals(7, st.size());

    // st.print();

    // test delete key
    doTestDeleteKey(st, "sea");
    assertEquals(18, st.numberOfNodes());
    assertEquals(6, st.size());
    doTestDeleteKey(st, "by");
    assertEquals(16, st.numberOfNodes());
    assertEquals(5, st.size());
  }

  @Test
  public void testDeleteAllKeys() {
    StringST<Integer> st = new StringST<>();
    doTestPutKey(st, "sea", 6);
    doTestPutKey(st, "by", 4);
    doTestPutKey(st, "sells", 1);
    doTestPutKey(st, "shells", 3);
    doTestPutKey(st, "she", 0);
    doTestPutKey(st, "the", 5);
    doTestPutKey(st, "shore", 7);

    // test delete key
    doTestDeleteKey(st, "sea");
    doTestDeleteKey(st, "by");
    doTestDeleteKey(st, "sells");
    doTestDeleteKey(st, "shells");
    doTestDeleteKey(st, "she");
    doTestDeleteKey(st, "the");
    doTestDeleteKey(st, "shore");
  }

  @Test
  public void testGetKey() {
    StringST<Integer> st = new StringST<>();
    st.put("Ahab’s", 1);
    assertEquals(1, st.size());
    assertEquals(1, st.get("Ahab’s"));
  }

  @Test
  public void testGetKey2() {
    StringST<Integer> st = new StringST<>();
    st.put("חן", 1);
    assertEquals(1, st.size());
    assertEquals(1, st.get(""));
  }

  private void doTestDeleteKey(StringST<Integer> st, String key) {
    st.delete(key);
    var got = st.get(key);
    assertNull("key is already deleted", got);
  }

  private void doTestPutKey(StringST<Integer> st, String key, Integer value) {
    st.put(key, value);
    var got = st.get(key);
    assertEquals(value, got);
  }
}