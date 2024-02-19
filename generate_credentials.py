import re
import secrets


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
    new_password = generate_secret()
    replace_in_file('.env', r'DB_PASSWORD=.*', f'DB_PASSWORD={new_password}')
    replace_in_file('.env.docker', r'DB_PASSWORD=.*', f'DB_PASSWORD={new_password}')

    new_secret = generate_secret()
    replace_in_file('.env', r'JWT_KEY=.*', f'JWT_KEY={new_secret}')
    replace_in_file('.env.docker', r'JWT_KEY=.*', f'JWT_KEY={new_secret}')


if __name__ == "__main__":
    main()
