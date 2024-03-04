package com.example.demo;

import java.util.Optional;

public interface Cache<K,V> {
  void put(K k, V v); 
  Optional<V> get(K k);
}
