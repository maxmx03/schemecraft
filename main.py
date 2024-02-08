import yaml
from jinja2 import Template
import subprocess
import sys
import os
from templates.palette import palette_template
from templates.config import config_template
from templates.init import init_template


class Colorscheme:
    def __init__(self, name, palette, default_config):
        self.name = name
        self.default_config = default_config
        self.palette = palette


def create_folder(folder_path):
    try:
        os.makedirs(folder_path)
        print(f"Folder '{folder_path}' created successfully.")
    except FileExistsError:
        print(f"Folder '{folder_path}' already exists.")


def create_file(file_path, content=""):
    try:
        with open(file_path, "w") as file:
            file.write(content)
        print(f"File '{file_path}' created successfully.")
    except Exception as e:
        print(f"Error creating file '{file_path}': {e}")


def run_commands(commands):
    for command in commands:
        process = subprocess.Popen(command, shell=True)
        process.wait()


def main():
    if len(sys.argv) < 2:
        print("Usage: python main.py <yaml_file>")
        sys.exit(1)

    yaml_path = sys.argv[1]
    with open(yaml_path, "r") as yaml_file:
        yaml_content = yaml.safe_load(yaml_file)

    colorscheme = Colorscheme(
        yaml_content["name"], yaml_content["palette"], yaml_content["default_config"]
    )

    palette_code = Template(palette_template).render(colorscheme=colorscheme)
    config_code = Template(config_template).render(colorscheme=colorscheme)
    init_code = Template(init_template).render(colorscheme=colorscheme)

    create_folder("lua")
    create_folder("colors")
    create_folder(f"lua/{colorscheme.name}")
    create_folder(f"lua/{colorscheme.name}/highlights")
    create_folder(f"lua/{colorscheme.name}/highlights/modules")
    create_file(f"lua/{colorscheme.name}/init.lua", init_code)
    create_file(f"lua/{colorscheme.name}/config.lua", config_code)
    create_file(f"lua/{colorscheme.name}/palette.lua", palette_code)

    print("Plugin Neovim gerado com sucesso em 'colorscheme.lua'")


if __name__ == "__main__":
    main()
    commands = []
    commands.append("stylua . -f ./stylua.toml")
    run_commands(commands)
