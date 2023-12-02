package com.ndnhuy.simplecal;

public class Tokenizer {
  static enum TokenType {
    OPERATOR,
    OPERAND,
    LEFT_BRACKET,
    RIGHT_BRACKET
  }

  static class Token {
    String value;
    TokenType type;

    Token(String value, TokenType type) {
      this.value = value;
      this.type = type;
    }
  }

}
