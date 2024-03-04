package com.example.demo;

import java.util.HashMap;
import java.util.Map;
import java.util.Optional;

import org.springframework.stereotype.Service;

@Service
public class CacheImpl implements Cache<String, String> {
  private Map<String, String> storage = new HashMap<>();

  @Override
  public synchronized void put(String k, String v) {
    storage.put(k, v);
  }

  @Override
  public synchronized Optional<String> get(String k) {
    if (storage.containsKey(k)) {
      return Optional.of(storage.get(k));
    } else {
      return Optional.empty();
    }
  }

}
