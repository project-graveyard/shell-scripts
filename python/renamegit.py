#!/usr/bin/env python3

import os

os.system("git remote -v")
print("\n")

cls = input("Do you want to clear old origin(Y/n): ").upper()
repo = input("Enter name of repo: ")
branch = input("Enter branch name: ")

def clear(cls):
    if cls == "" or cls == 'Y':
        os.system("git remote remove origin")

def rename(repo, branch):
    os.system(f"git remote add origin git@github.com:DaveSaah/{repo}.git && git push --set-upstream origin {branch}")


clear(cls)
rename(repo, branch)
print("\n")
os.system("git remote -v")
