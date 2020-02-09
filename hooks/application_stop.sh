#!/bin/bash

echo hogehogeout
hoge=`ps aux | grep goapp | grep -v grep | wc -l`
if [ $hoge != 0 ]; then
  echo hogehogein
  killall goapp -q
fi
