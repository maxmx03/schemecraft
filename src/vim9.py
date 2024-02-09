from src.file_manager import Filemanager
from src.templates.vim9 import colors
from src.templates.vim9.lib import (
    config,
    editor,
    terminal,
    syntax,
    modules,
)
from jinja2 import Environment


def get_content(env, source, colorscheme):
    tmp = env.from_string(source)
    return tmp.render(colorscheme=colorscheme)


class Vim9:
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

    def create_colorscheme(self):
        env = Environment(trim_blocks=True, lstrip_blocks=True)
        file_manager = Filemanager(root="build/vim9")
        folders = ["colors", "lib"]

        for folder in folders:
            file_manager.create_folder(folder)

        file_manager.create_file(
            f"colors/{self.name}.vim",
            content=get_content(env=env, source=colors.template, colorscheme=self),
            format=False,
        )
        file_manager.create_file(
            "lib/config.vim",
            content=get_content(env=env, source=config.template, colorscheme=self),
            format=False,
        )
        file_manager.create_file(
            "lib/syntax.vim",
            content=get_content(env=env, source=syntax.template, colorscheme=self),
            format=False,
        )
        file_manager.create_file(
            "lib/editor.vim",
            content=get_content(env=env, source=editor.template, colorscheme=self),
            format=False,
        )
        file_manager.create_file(
            "lib/terminal.vim",
            content=get_content(env=env, source=terminal.template, colorscheme=self),
            format=False,
        )
        file_manager.create_file(
            "lib/modules.vim",
            content=get_content(env=env, source=modules.template, colorscheme=self),
            format=False,
        )
