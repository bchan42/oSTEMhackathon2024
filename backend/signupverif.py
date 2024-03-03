# # SIGN UP VERIF

from pymongo import MongoClient

import json
from functools import cached_property
from http.cookies import SimpleCookie
from http.server import HTTPServer, BaseHTTPRequestHandler
from urllib.parse import parse_qsl, urlparse


# function checks if email ends in calpoly.edu
# dict -> bool
def check_calpoly(user_info):
    email = user_info['email']
    domain = 'calpoly.edu'
    parts = email.split('@')
    return len(parts) == 2 and parts[1].lower() == domain
 
# function checks if password is less than 
# dict -> bool
def check_password(user_info):
    password = user_info['password']
    if len(password) < 8: # at least 8 characters
        # print('Password needs to be at least 8 characters long.')
        return False
    elif not any(char.isupper() for char in password):  # 1 capital letter
        # print('Password needs to have a capital letter.')
        return False
    elif not any(char.islower() for char in password):  # 1 lowercase
        # print('Password needs to have a lowercase letter.')
        return False
    elif not any(char.isdigit() for char in password):  # 1 number
        # print('Password needs to have a number.')
        return False
    elif not any(char in '!@#$%^&*()' for char in password):  # 1 special character
        # print('Password needs to have a special character.')
        return False
    else:
        # print('Password meets the criteria.')
        return True


class WebRequestHandler(BaseHTTPRequestHandler):

    def do_signup(self, data):
        status = 200
        # Sign up
        #if check_calpoly(data) is False or check_password(data) is False:
        #    return 402
        if users.find_one({'email': data['email']}) is not None:
            return 402
        
        users.insert_one(data)
        return status        

    def do_signin(self, data):
        status = 200
        # Sign in
        if users.find_one({'email': data['email'], 'password': data['password']}) is None:
            return 402
        return status

    # send data from frontend to backend
    def do_POST(self):
        status = 200 

        # send frontend data to server
        request = self.rfile.read(int(self.headers["Content-Length"]))
        data = json.loads(request)

        if self.path == "/user/signup":
            status = self.do_signup(data)
        elif self.path == "/user/signin":
            status = self.do_signin(data)
        else:
            status = 404

        self.send_response(status)
        self.send_header("Content-Type", "application/json")
        self.send_header("Access-Control-Allow-Origin", "*")
        self.send_header("Access-Control-Allow-Headers", "*")
        self.end_headers()

def get_database():
   # Provide the mongodb atlas url to connect python to mongodb using pymongo
   CONNECTION_STRING = "mongodb://user:pass@mongo.lone-faerie.xyz:27017"
   # Create a connection using MongoClient. You can import MongoClient or use pymongo.MongoClient
   client = MongoClient(CONNECTION_STRING)
   # Create the database for our example (we will use the same database throughout the tutorial
   return client['db']
  
# This is added so that many files can reuse the function get_database()
if __name__ == "__main__":   
   # Get the database
    dbname = get_database()
    global users
    users = dbname['users']

    server = HTTPServer(("", 8081), WebRequestHandler)
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


# create some global dictionary
user1 = {
    'name': 'Bern',
    'email': 'bern@gmail.com',
    'password': 'Abc@34we'
}

# SIGN IN VERIF

