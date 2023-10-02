package trie;

import static org.junit.Assert.assertNull;
import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.Test;

public class StringSTTest {

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
