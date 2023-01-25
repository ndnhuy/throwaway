package com.example.chatapp;

import org.springframework.data.repository.ListCrudRepository;

public interface ChatRepository extends ListCrudRepository<Message, Long> {
}
