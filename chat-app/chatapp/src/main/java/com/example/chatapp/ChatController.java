package com.example.chatapp;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
public class ChatController {

    @Autowired
    private ChatService chatService;

    @GetMapping("/messages")
    public List<MessageDTO> getAllMessages() {
        return chatService.getAllMessages();
    }

    @PostMapping("/send")
    public void sendMessage(@RequestBody MessageDTO msg) {
        chatService.sendMessage(msg);
    }
}
