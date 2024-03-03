# sign up verification

from pymongo import MongoClient

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
 
# function checks if password
def check_password(password):
    if len(password) < 8:  # at least 8 characters
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