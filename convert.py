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
    owos = []
    with open(sys.argv[1]) as f:
        owos = [owo for owo in f.read() if owo in ["<",">","+","-",".",",","[","]"]]
        print("".join(owos))
    with open(sys.argv[1][:-2] + ".owo", "w") as f:
        f.write(" ".join([convmap[owo] for owo in owos]))
        print(" ".join([convmap[owo] for owo in owos]))