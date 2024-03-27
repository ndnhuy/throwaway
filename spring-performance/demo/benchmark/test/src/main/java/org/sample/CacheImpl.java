package org.sample;

import java.util.Map;
import java.util.Optional;
import java.util.concurrent.ConcurrentHashMap;

public class CacheImpl {
  private Map<String, String> storage = new ConcurrentHashMap<>();

  public void put(String k, String v) {
    storage.put(k, v);
  }

  public Optional<String> get(String k) {
    if (storage.containsKey(k)) {
      return Optional.of(storage.get(k));
    } else {
      return Optional.empty();
    }
  }

}
