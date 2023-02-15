package com.example.chatapp;

import org.junit.jupiter.api.Test;

public class MessageDTOTest {
    @Test
    public void foo() {
        String s = "{\n" +
                "\t\"id\": \"1\",\n" +
                "\t\"fromUser\": \"test2\",\n" +
                "\t\"toUser\": \"test2\",\n" +
                "\t\"content\": \"content\",\n" +
                "\t\"createdAt\": \"2023-01-24\"\n" +
                "}";
        char c = 257;
        System.out.println(sizeofChar()*s.getBytes().length);
    }

    public int sizeofChar() {
        char i = 1, j = 0;
        while (i != 0) {
            i = (char) (i<<8); j++;
        }
        return j;
    }
}
