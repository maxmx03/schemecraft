import os
import subprocess


def format_file(file_path):
    command = f"stylua {file_path} -f ./stylua.toml"
    subprocess.run(command, shell=True)


class Filemanager:
    def __init__(self, root):
        self.root = root

    def create_folder(self, folder_path):
        try:
            folder_path = f"{self.root}/{folder_path}"
            os.makedirs(folder_path)
            print(f"Folder '{folder_path}' created successfully.")
        except FileExistsError:
            print(f"Folder '{folder_path}' already exists.")

    def create_file(self, file_path, content="", format=True):
        try:
            file_path = f"{self.root}/{file_path}"
            with open(file_path, "w") as file:
                file.write(content)
            if format:
                format_file(file_path)
            print(f"File '{file_path}' created successfully.")
        except Exception as e:
            print(f"Error creating file '{file_path}': {e}")
