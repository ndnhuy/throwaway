package trie;

import org.junit.Assert;

public class StringST<Value> {
  private static int R = 256;

  private static class Node {
    private Node[] next = new Node[R];
    private Object val;

    void print() {
      for (var c = 0; c < next.length; c++) {
        if (next[c] != null) {
          var n = next[c];
          System.out.println(String.format("[%c, %d]", c, n.val));
          n.print();
        }
      }
    }
  }

  private Node root = null;

  public StringST() {
  }

  void put(String key, Value val) {
    Assert.assertNotNull("mus not be null", val);
    root = put(root, key, val, 0);
  }

  private Node put(Node node, String key, Value val, int depth) {
    if (node == null) {
      node = new Node();
    }
    if (depth > key.length() - 1) {
      node.val = val;
      return node;
    }
    var c = key.charAt(depth);
    node.next[c] = put(node.next[c], key, val, depth + 1);
    return node;
  }

  Value get(String key) {
    if (root == null) {
      return null;
    }
    Node iter = root;
    for (var i = 0; i < key.length(); i++) {
      var c = key.charAt(i);
      iter = iter.next[c];
      if (iter == null) {
        return null;
      }
    }
    return (Value) iter.val;
  }

  int size() {
    return size(root);
  }

  private int size(Node node) {
    if (node == null) {
      return 0;
    }

    var sum = 0;
    for (var i = 0; i < node.next.length; i++) {
      var t = node.next[i];
      if (t != null) {
        // if the link is not null, that means 1 node
        sum++;
      }
      // then sum up all size of subtries
      sum += size(t);
    }
    return sum;
  }

  void delete(String key) {
    root = delete(root, key, 0);
  }

  private Node delete(Node node, String key, int depth) {
    if (node == null) {
      return node;
    }
    if (depth == key.length()) {
      node.val = null;
    } else {
      var c = key.charAt(depth);
      node.next[c] = delete(node.next[c], key, depth + 1);
    }

    if (node.val != null) {
      return node;
    } else {
      for (var i = 0; i < node.next.length; i++) {
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

  void print() {
    root.print();
  }
}
