# sign up verification

from pymongo import MongoClient

# webapp.py

import json
from functools import cached_property
from http.cookies import SimpleCookie
from http.server import HTTPServer, BaseHTTPRequestHandler
from urllib.parse import parse_qsl, urlparse

class WebRequestHandler(BaseHTTPRequestHandler):
    # ...

    path = "/user"

    # # 
    # def do_GET(self):
    #     status = 200
    #     #insert logic (if any errors change value to 402)
        
    #     request = self.rfile.read(int(self.headers["Content-Length"]))
    #     data = json.loads(request)

    #     self.send_response(status)
    #     self.send_header("Content-Type", "application/json")
    #     self.send_header("Access-Control-Allow-Origin", "*")
    #     self.send_header("Access-Control-Allow-Headers", "*")
    #     self.end_headers()
    #     self.wfile.write(self.get_response().encode("utf-8"))

    # send data from frontend to backend
    def do_POST(self):
        status = 200 
        # insert logic (if any errors in logic, change value of status to 402)

        # send frontend data to server
        request = self.rfile.read(int(self.headers["Content-Length"]))
        data = json.loads(request)


        self.send_response(status)
        self.send_header("Content-Type", "application/json")
        self.send_header("Access-Control-Allow-Origin", "*")
        self.send_header("Access-Control-Allow-Headers", "*")
        self.end_headers()
        



def get_database():
   # Provide the mongodb atlas url to connect python to mongodb using pymongo
   CONNECTION_STRING = "mongodb+srv://user:pass@cluster.mongodb.net/myFirstDatabase"
   # Create a connection using MongoClient. You can import MongoClient or use pymongo.MongoClient
   client = MongoClient(CONNECTION_STRING)
   # Create the database for our example (we will use the same database throughout the tutorial
   return client['user_information']
  
# This is added so that many files can reuse the function get_database()
if __name__ == "__main__":   
   # Get the database
    dbname = get_database()

    server = HTTPServer(("0.0.0.0", 8000), WebRequestHandler)
    server.serve_forever()


# add this stuff to end
# item_1 = {
#   "_id" : "U1IT00001",
#   "item_name" : "Blender",
#   "max_discount" : "10%",
#   "batch_number" : "RR450020FRG",
#   "price" : 340,
#   "category" : "kitchen appliance"
# }

# item_2 = {
#   "_id" : "U1IT00002",
#   "item_name" : "Egg",
#   "category" : "food",
#   "quantity" : 12,
#   "price" : 36,
#   "item_description" : "brown country eggs"
# }
# collection_name.insert_many([item_1,item_2])


# function checks if email ends in calpoly.edu
def check_calpoly(email):
    domain = '@calpoly.edu'
    parts = email.split('@')
    return len(parts) == 2 and parts[1].lower() == domain
 
# function checks if password is less than 
def check_password(password):
    if len(password) < 8: # at least 8 characters
        print('Password needs to be at least 8 characters long.')
    elif not any(char.isupper() for char in password):  # 1 capital letter
        print('Password needs to have a capital letter.')
    elif not any(char.islower() for char in password):  # 1 lowercase
        print('Password needs to have a lowercase letter.')
    elif not any(char.isdigit() for char in password):  # 1 number
        print('Password needs to have a number.')
    elif not any(char in '!@#$%^&*()' for char in password):  # 1 special character
        print('Password needs to have a special character.')
    else:
        print('Password meets the criteria.')