package trie;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.Test;

public class StringSTTest {

  @Test
  public void testPutKey() {
    StringST<Integer> st = new StringST<>();
    doTestPutKey(st, "sea", 6);
    doTestPutKey(st, "by", 4);
    doTestPutKey(st, "sells", 1);
    doTestPutKey(st, "shells", 3);
    doTestPutKey(st, "she", 0);
    doTestPutKey(st, "the", 5);
    doTestPutKey(st, "shore", 7);

    st.print();
  }

  private void doTestPutKey(StringST<Integer> st, String key, Integer value) {
    st.put(key, value);
    var want = st.get(key);
    assertEquals(value, want);
  }
}
