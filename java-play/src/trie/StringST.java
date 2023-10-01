package trie;

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

  void delete(String key) {
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
