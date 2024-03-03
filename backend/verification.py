# sign up verification

from pymongo import MongoClient
def get_database():
 
   # Provide the mongodb atlas url to connect python to mongodb using pymongo
   CONNECTION_STRING = "mongodb+srv://user:pass@cluster.mongodb.net/myFirstDatabase"
 
   # Create a connection using MongoClient. You can import MongoClient or use pymongo.MongoClient
   client = MongoClient(CONNECTION_STRING)
 
   # Create the database for our example (we will use the same database throughout the tutorial
   return client['user_shopping_list']
  
# This is added so that many files can reuse the function get_database()
if __name__ == "__main__":   
  
   # Get the database
   dbname = get_database()


# function checks if email ends in calpoly.edu
def check_calpoly(email):
    domain = '@calpoly.edu'
    parts = email.split('@')
    return len(parts) == 2 and parts[1].lower() == domain
 
# function checks if password is less than 
 def check_password(password):
    if len(password) < 8: # at least 8 characters
        print('Password needs to be at least 8 characters long.')
    elif 'ABCDEFGHIJKLMNOPQRSTUVWXYZ' not in password: # 1 capital letter
        print('Password needs to have a capital letter.')
    elif 'abcdefghijklmnopqrstuvwxyz' not in password: # 1 lowercase
        print('Password needs to have lowercase letter.')
    elif not any(char.isdigit for char in password)# 1 number
        print('Password needs to have a number.')
    # 1 unique character
    elif '1@#456&*()' not in password:
        print('Password needs to have a special character.')



# not account w existing email
# authenticate email exists in database
# checks if password is right 

# give client rand generated token


