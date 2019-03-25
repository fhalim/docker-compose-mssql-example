#!/usr/bin/env python3
import sys

who = "world"
if len(sys.argv) > 1:
    who = sys.argv[1]
print(f"hello {who}")