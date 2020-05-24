#!/bin/bash
for w in $(cat /usr/share/dict/words); do
       sha1="$(echo $w | sha1sum | cut -d ' ' -f 1)"
       if [ "$sha1" = "f164acde21f01cf3b2ba4f7dd650d9bf2a699b96" ]; then
               echo $w
               exit 0
       fi
done
exit 1

