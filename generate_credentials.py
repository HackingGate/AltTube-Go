import os
import re
import base64
import secrets
import sys

def generate_secret():
    return secrets.token_hex(16)

def replace_in_file(file_name, pattern, replacement):
    with open(file_name, 'r+') as file:
        content = file.read()
        content_new = re.sub(pattern, replacement, content, flags=re.M)
        file.seek(0)
        file.write(content_new)
        file.truncate()

def main():
    if len(sys.argv) != 2:
        print("Usage: python3 generate_credentials.py <filename>")
        sys.exit(1)

    file_name = sys.argv[1]

    new_secret = generate_secret()
    replace_in_file(file_name, r'JWT_SECRET=.*', f'JWT_SECRET={new_secret}')

    new_password = generate_secret()
    replace_in_file(file_name, r'DB_PASSWORD=.*', f'DB_PASSWORD={new_password}')

if __name__ == "__main__":
    main()
