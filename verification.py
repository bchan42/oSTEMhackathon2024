# function checks if email ends in calpoly.edu
def check_calpoly(email):
    domain = '@calpoly.edu'
    parts = email.split('@')
    return len(parts) == 2 and parts[1].lower() == domain

# not account w existing email
# authenticate email exists in database
# checks if password is right 

# give client rand generated token
 
# function checks if password is less than 
 def check_password(password):
    if len(password) < 8: # at least 8 characters
        print('Password needs to be at least 8 characters long.')
    elif 'ABCDEFGHIJKLMNOPQRSTUVWXYZ' not in password: # 1 capital letter
        print('Password needs to have a capital letter.')
    elif 'abcdefghijklmnopqrstuvwxyz' not in password: # 1 lowercase
        print('Password needs to have lowercase letter.')
    # 1 number
    # 1 unique character
    return