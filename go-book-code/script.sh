#!/bin/env bash
#---
# Excerpted from "Distributed Services with Go",
# published by The Pragmatic Bookshelf.
# Copyrights apply to this code. It may not be used to create training material,
# courses, books, articles, and the like. Contact us if you are in doubt.
# We make no guarantees that this code is fit for any purpose.
# Visit http://www.pragmaticprogrammer.com/titles/tjgo for more book information.
#---

filepath="internal/server/server.go"
cur=""
for dir in $(cat .directories); do
  if [ -f "$dir/$filepath" ]; then
    if [ -z "$cur" ]; then
      cur="$dir/$filepath"
    else
      echo "---"
      echo "--- $cur vs. $dir/$filepath"
      echo "---"
      diff "$cur" "$dir/$filepath"
      cur="$dir/$filepath"
    fi
  fi
done
  
