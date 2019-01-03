# -*- encoding: utf-8 -*-

import sys

convmap = {
    "<": "ôwo", # left,
    ">": "owô", # right,
    "+": "òwó", # add,
    "-": "ówò", # sub,
    ".": "OwO", # print,
    ",": "owo", # read,
    "[": "ÕwO", # startLoop,
    "]": "OwÕ", # endLoop,
}

if len(sys.argv) >= 2:
    chars = []
    with open(sys.argv[1]) as f:
        chars = [char for char in f.read() if char in ["<",">","+","-",".",",","[","]"]]
        print("".join(chars))
    with open(sys.argv[1][:-2] + ".owo", "w") as f:
        f.write(" ".join([convmap[x] for x in chars]))
        print(" ".join([convmap[x] for x in chars]))