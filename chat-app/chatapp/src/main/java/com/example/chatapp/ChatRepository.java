package com.example.chatapp;

import org.springframework.data.repository.ListCrudRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface ChatRepository extends ListCrudRepository<Message, Long> {
}
