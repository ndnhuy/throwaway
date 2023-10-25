package trie;

import java.util.function.BiFunction;
import java.util.function.Function;

import org.junit.Assert;
import org.openjdk.jmh.runner.RunnerException;

public class StringST<Value> {
  private static int R = 256;

  private static class Node {
    private Node[] next = new Node[R];
    private Object val;

    void print(int level) {
      for (var c = 0; c < next.length; c++) {
        if (next[c] != null) {
          var n = next[c];
          var prefix = "";
          for (var i = 0; i < level; i++) {
            prefix += " ";
          }
          var valText = n.val == null ? "" : "->" + n.val;
          System.out.println(String.format("%s|%c%s", prefix, c, valText));
          n.print(level + 1);
        }
      }
    }
  }

  private Node root = null;

  public StringST() {
  }

  public String sanitizeKey(String key) {
    key = key.toLowerCase();
    key = key.replaceAll("’", "'");
    key = key.replaceAll("\\.", "");
    key = key.replaceAll("\"", "");
    for (char i = 0; i < key.length(); i++) {
      char c = key.charAt(i);
      if (c > 255) {
        key = key.replaceAll(String.valueOf(c), "");
        return sanitizeKey(key);
      }
    }
    return key;
  }

  public void put(String key, Value val) {
    Assert.assertNotNull("mus not be null", val);
    try {
      key = sanitizeKey(key);
      root = put(root, key, val, 0);
    } catch (Exception e) {
      throw new RuntimeException(String.format("error when put (key=%s,value=%s)", key, val), e);
    }
  }

  public void putIfAbsent(String key, Value defaultValue, Function<Value, Value> valueGetter) {
    var v = get(key);
    if (v == null) {
      put(key, defaultValue);
    } else {
      put(key, valueGetter.apply(v));
    }
  }

  private Node put(Node node, String key, Value val, int depth) {
    if (node == null) {
      node = new Node();
    }
    if (depth > key.length() - 1) {
      node.val = val;
      return node;
    }
    char c = key.charAt(depth);
    node.next[c] = put(node.next[c], key, val, depth + 1);
    return node;
  }

  public Value get(String key) {
    key = sanitizeKey(key);
    try {
      boolean isNumber = false;
      try {
        Double.parseDouble(key);
        isNumber = true;
      } catch (NumberFormatException e) {
        isNumber = false;
      }
      if (isNumber) {
        double d = Double.parseDouble(key);
        key = Double.toString(d);
      }
      if (root == null) {
        return null;
      }
      var iter = root;
      for (int i = 0; i < key.length(); i++) {
        char c = key.charAt(i);
        iter = iter.next[c];
        if (iter == null) {
          return null;
        }
      }
      return (Value) iter.val;
    } catch (Exception e) {
      throw new RuntimeException(String.format("error when put (key=%s)", key), e);
    }
  }

  public int numberOfNodes() {
    return numberOfNodes(root);
  }

  public int size() {
    return size(root);
  }

  private int size(Node node) {
    if (node == null)
      return 0;

    int cnt = 0;
    if (node.val != null) {
      cnt++;
    }
    for (int i = 0; i < node.next.length; i++) {
      cnt += size(node.next[i]);
    }
    return cnt;
  }

  private int numberOfNodes(Node node) {
    if (node == null) {
      return 0;
    }

    var sum = 0;
    for (int i = 0; i < node.next.length; i++) {
      Node t = node.next[i];
      if (t != null) {
        // if the link is not null, that means 1 node
        sum++;
      }
      // then sum up all size of subtries
      sum += numberOfNodes(t);
    }
    return sum;
  }

  public void delete(String key) {
    root = delete(root, key, 0);
  }

  private Node delete(Node node, String key, int depth) {
    if (node == null) {
      return node;
    }
    if (depth == key.length()) {
      node.val = null;
    } else {
      char c = key.charAt(depth);
      node.next[c] = delete(node.next[c], key, depth + 1);
    }

    if (node.val != null) {
      return node;
    } else {
      for (int i = 0; i < node.next.length; i++) {
        if (node.next[i] != null) {
          return node;
        }
      }
    }
    // remove the node if all links are null
    return null;
  }

  boolean contains(String key) {
    throw new RuntimeException();
  }

  boolean isEmpty() {
    return false;
  }

  String longestPrefixOf(String s) {
    throw new RuntimeException();
  }

  public void print() {
    root.print(0);
  }
}
