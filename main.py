import yaml
from jinja2 import Template
import sys
import os
import subprocess
from templates.palette import palette_template
from templates.config import config_template
from templates.init import init_template
from templates.colors import colors_template
from templates.highlights.editor import editor_template
from templates.highlights.syntax import syntax_template
from templates.highlights.modules import modules_template


class Colorscheme:
    def __init__(
        self,
        name,
        author,
        repo,
        contact,
        license,
        palette,
        default_config,
        highlights,
    ):
        self.name = name
        self.author = author
        self.repo = repo
        self.contact = contact
        self.license = license
        self.default_config = default_config
        self.palette = palette
        self.highlights = highlights


def format_file(path):
    command = f"stylua {path} -f ./stylua.toml"
    subprocess.run(command, shell=True)


def create_folder(folder_path):
    try:
        os.makedirs("build/" + folder_path)
        print(f"Folder '{folder_path}' created successfully.")
    except FileExistsError:
        print(f"Folder '{folder_path}' already exists.")


def create_file(file_path, content=""):
    try:
        with open("build/" + file_path, "w") as file:
            file.write(content)
        format_file("build/" + file_path)
        print(f"File '{file_path}' created successfully.")
    except Exception as e:
        print(f"Error creating file '{file_path}': {e}")


def main():
    if len(sys.argv) < 2:
        print("Usage: python main.py <yaml_file>")
        sys.exit(1)

    yaml_path = sys.argv[1]
    with open(yaml_path, "r") as yaml_file:
        yaml_content = yaml.safe_load(yaml_file)

    colorscheme = Colorscheme(
        name=yaml_content["name"],
        author=yaml_content["author"],
        contact=yaml_content["contact"],
        repo=yaml_content["repo"],
        license=yaml_content["license"],
        palette=yaml_content["palette"],
        default_config=yaml_content["default_config"],
        highlights=yaml_content["highlights"],
    )

    colors_code = Template(
        colors_template, trim_blocks=True, lstrip_blocks=True
    ).render(colorscheme=colorscheme)

    modules_code = Template(
        modules_template, trim_blocks=True, lstrip_blocks=True
    ).render(colorscheme=colorscheme)

    syntax_code = Template(
        syntax_template, trim_blocks=True, lstrip_blocks=True
    ).render(colorscheme=colorscheme)

    editor_code = Template(
        editor_template, trim_blocks=True, lstrip_blocks=True
    ).render(colorscheme=colorscheme)

    palette_code = Template(
        palette_template, trim_blocks=True, lstrip_blocks=True
    ).render(colorscheme=colorscheme)

    config_code = Template(
        config_template,
        trim_blocks=True,
        lstrip_blocks=True,
    ).render(colorscheme=colorscheme)

    init_code = Template(init_template, trim_blocks=True, lstrip_blocks=True).render(
        colorscheme=colorscheme
    )

    create_folder("colors")
    create_folder(f"lua/{colorscheme.name}/highlights")
    create_file(f"colors/{colorscheme.name}.lua", colors_code)
    create_file(f"lua/{colorscheme.name}/init.lua", init_code)
    create_file(f"lua/{colorscheme.name}/config.lua", config_code)
    create_file(f"lua/{colorscheme.name}/palette.lua", palette_code)
    create_file(f"lua/{colorscheme.name}/highlights/editor.lua", editor_code)
    create_file(f"lua/{colorscheme.name}/highlights/editor.lua", syntax_code)
    create_file(f"lua/{colorscheme.name}/highlights/modules.lua", modules_code)

    print("Colorscheme was generated with success")


if __name__ == "__main__":
    main()
