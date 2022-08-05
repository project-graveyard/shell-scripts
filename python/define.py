#!/usr/bin/env python3

import os
import sys

try:
    os.system(f"dict {sys.argv[1]} | less")
except:
    print("\033[1;31;40m Error:\033[1;37;40m Missing Argument")
