import os

val = int(input("Enter length of secret key: "))
print(os.urandom(val))
