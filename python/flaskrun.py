import os

port = input("Enter port number: ")

host = int(input("\nChoose host type\n1.Open\n2.Close\n::"))



if port and host:
	if host == 1:
		command = f"flask run --host=0.0.0.0 --port={port}"
	elif host == 2:
		command = f"flask run --port={port}"
	elif host != 1 and host !=2:
		command = ""
		print("\nEnter a valid option")
else:
	command = "flask run"


os.system(f"export FLASK_APP=app && export FLASK_ENV=development && {command}")
