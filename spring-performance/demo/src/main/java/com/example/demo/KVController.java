package com.example.demo;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.GetMapping;

@RestController
@RequestMapping(path = "/cache")
public class KVController {

  @Autowired
  private Cache<String, String> cache;

  @PostMapping
  public String put(@RequestParam("key") String key, @RequestParam("value") String value) {
    cache.put(key, value);
    return String.format("put %s %s\n", key, value);
  }

  @GetMapping
  public String get(@RequestParam("key") String key) {
    var v = cache.get(key);
    if (v.isPresent()) {
      return String.format("get %s %s\n", key, v.get());
    } else {
      return String.format("not found key %s\n", key);
    }
  }

}
