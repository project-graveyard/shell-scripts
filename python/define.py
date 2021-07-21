import os

val = input("Enter word: ")

os.system(f"dict {val} | less")
