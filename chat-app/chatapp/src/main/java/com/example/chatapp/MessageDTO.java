package com.example.chatapp;

import jakarta.validation.constraints.NotBlank;
import org.springframework.util.Assert;

public record MessageDTO(
        Long id,
        String fromUser,
        String toUser,
        String content,
        String createdAt) {

    public MessageDTO {
        Assert.hasText(fromUser, "fromUser must not be null");
        Assert.hasText(toUser, "toUser must not be null");
        Assert.hasText(content, "content must not be null");
        Assert.hasText(createdAt, "createdAt must not be null");
    }
}
