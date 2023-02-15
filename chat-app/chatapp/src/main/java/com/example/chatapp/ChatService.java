package com.example.chatapp;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.util.Assert;

import java.util.List;
import java.util.stream.Collectors;

@Service
public class ChatService {
    @Autowired
    private ChatRepository chatRepository;

    public List<MessageDTO> getAllMessages() {
        List<Message> messages = chatRepository.findAll();
        return messages.stream()
                .map(this::toDto)
                .collect(Collectors.toList());
    }

    public void sendMessage(MessageDTO msg) {
        Assert.notNull(msg, "message must not be null");
        Assert.hasText(msg.fromUser(), "fromUser must not be null");
        Assert.hasText(msg.toUser(), "toUser must not be null");
        Assert.hasText(msg.content(), "content must not be null");
        Assert.hasText(msg.createdAt(), "createdAt must not be null");

        chatRepository.save(toEntity(msg));
    }

    private Message toEntity(MessageDTO dto) {
        return new Message(dto.fromUser(), dto.toUser(), dto.content(), dto.createdAt());
    }

    private MessageDTO toDto(Message entity) {
        return new MessageDTO(entity.getId(), entity.getFromUser(), entity.getToUser(), entity.getContent(), entity.getCreatedAt());
    }
}
