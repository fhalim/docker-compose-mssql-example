#!/usr/bin/env python3
import sys
import os

who = "world"
greeting = "hello"

if "GREETING" in os.environ:
    greeting = os.environ["GREETING"]

if len(sys.argv) > 1:
    who = sys.argv[1]
print(f"{greeting}, {who}")