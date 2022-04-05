#! /usr/bin/python3

# Generates an ssh key for github account

import os

email = input("Enter your email: ")
os.system(f"ssh-keygen -t ed25519 -C {email}")
os.system("eval '$(ssh-agent -s)'")
os.system("ssh-add ~/.ssh/id_ed25519")
print("\n\n")
print("Copy the ssh key below: \n")
os.system("cat ~/.ssh/id_ed25519.pub")
