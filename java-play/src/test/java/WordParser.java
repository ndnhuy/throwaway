import static org.junit.jupiter.api.Assertions.assertEquals;

import java.io.FileReader;
import java.io.IOException;
import java.io.Reader;
import java.io.StreamTokenizer;
import java.io.StringReader;

import org.junit.Test;

import trie.StringST;

public class WordParser {

  @Test
  public void streamTokenizer() throws IOException {
    String file = "src/test/resources/moby.txt";
    var st = new StringST<Integer>();
    try (FileReader reader = new FileReader(file)) {
      doTokenize(reader, st);
    }
    assertEquals(7, st.get("sleeps"));
    assertEquals(7, st.size());
  }

  @Test
  public void testTokenize() throws IOException {
    var st = new StringST<Integer>();
    doTokenize(new StringReader("hello world! The world, now have (9 billions people)"), st);
    assertEquals(8, st.size());
    assertEquals(1, st.get("hello"));
    assertEquals(2, st.get("world"));
    assertEquals(1, st.get("The"));
    assertEquals(null, st.get("the"));
    assertEquals(1, st.get("9"));
    assertEquals(1, st.get("now"));
    assertEquals(1, st.get("have"));
    assertEquals(1, st.get("billions"));
    assertEquals(1, st.get("people"));
  }

  @Test
  public void testTokenize2() throws IOException {
    var st = new StringST<Integer>();
    doTokenize(new StringReader("Ahab’s"), st);
    assertEquals(1, st.size());
  }

  @Test
  public void testTokenize3() throws IOException {
    var st = new StringST<Integer>();
    doTokenize(new StringReader("""
        “While you take in hand to school others, and to teach them by what
        name a whale-fish is to be called in our tongue, leaving out, through
        ignorance, the letter H, which almost alone maketh up the signification
        of the word, you deliver that which is not true.” —Hackluyt.
          """), st);
    assertEquals(2, st.get("you"));
    assertEquals(1, st.get("While"));
    assertEquals(1, st.get("whale-fish"));
    assertEquals(1, st.get("true"));
  }

  @Test
  public void testTokenizeWithQuote() throws IOException {
    var st = new StringST<Integer>();
    doTokenize(new StringReader("I say \"I am very sorry to say that!\""), st);
    assertEquals(7, st.size());
    assertEquals(1, st.get("sorry"));
    assertEquals(2, st.get("I"));
    assertEquals(2, st.get("say"));
  }

  @Test
  public void testTokenizeWithNumber() throws IOException {
    var st = new StringST<Integer>();
    doTokenize(new StringReader("1, 1.0, 1.00, 001 2.5, 2.5, 2.6; 3.12; 0.0  !!??"), st);
    assertEquals(5, st.size());
    assertEquals(4, st.get("1"));
    assertEquals(2, st.get("2.5"));
    assertEquals(2, st.get("2.500"));
    assertEquals(1, st.get("2.6"));
    assertEquals(1, st.get("02.60"));
    assertEquals(null, st.get("2.61"));
    assertEquals(1, st.get("3.12"));
    assertEquals(1, st.get("0"));
    assertEquals(1, st.get("0.0"));
    assertEquals(1, st.get("0.0000"));
  }

  private void doTokenize(Reader reader, StringST<Integer> st) throws IOException {
    StreamTokenizer tokenizer = new StreamTokenizer(reader);
    while (tokenizer.nextToken() != StreamTokenizer.TT_EOF) {
      var tokenType = tokenizer.ttype;
      switch (tokenType) {
        case StreamTokenizer.TT_WORD:
          st.putIfAbsent(tokenizer.sval, 1, v -> v + 1);
          break;
        case StreamTokenizer.TT_NUMBER:
          st.putIfAbsent(Double.toString(tokenizer.nval), 1, v -> v + 1);
          break;
        case '\"':
          String quotedStr = tokenizer.sval;
          if (quotedStr != null) {
            var quotedStrReader = new StringReader(quotedStr);
            doTokenize(quotedStrReader, st);
          }
          break;
      }
    }
  }
}
