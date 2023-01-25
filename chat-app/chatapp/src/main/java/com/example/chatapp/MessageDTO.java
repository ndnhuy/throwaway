package com.example.chatapp;

public record MessageDTO(
        Long id,
        String fromUser,
        String toUser,
        String content,
        String createdAt) {
}
