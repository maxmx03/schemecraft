from src.file_manager import Filemanager
from src.templates.neovim import init, config, palette, colors, docker
from src.templates.neovim.highlights import (
    init as highlights_init,
    editor,
    syntax,
    modules,
)
from src.templates.neovim.tests import colorscheme_spec
from jinja2 import Environment


def get_content(env, source, colorscheme):
    tmp = env.from_string(source)
    return tmp.render(colorscheme=colorscheme)


class Neovim:
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
        file_manager = Filemanager(root="build/neovim")
        folders = ["colors", f"lua/{self.name}/highlights", "tests"]

        for folder in folders:
            file_manager.create_folder(folder)

        file_manager.create_file(
            f"colors/{self.name}.lua",
            content=get_content(env=env, source=colors.template, colorscheme=self),
        )
        file_manager.create_file(
            f"lua/{self.name}/init.lua",
            content=get_content(env=env, source=init.template, colorscheme=self),
        )
        file_manager.create_file(
            f"lua/{self.name}/config.lua",
            content=get_content(env=env, source=config.template, colorscheme=self),
        )
        file_manager.create_file(
            f"lua/{self.name}/palette.lua",
            content=get_content(env=env, source=palette.template, colorscheme=self),
        )
        file_manager.create_file(
            f"lua/{self.name}/highlights/init.lua",
            content=get_content(
                env=env, source=highlights_init.template, colorscheme=self
            ),
        )
        file_manager.create_file(
            f"lua/{self.name}/highlights/editor.lua",
            content=get_content(env=env, source=editor.template, colorscheme=self),
        )
        file_manager.create_file(
            f"lua/{self.name}/highlights/syntax.lua",
            content=get_content(env=env, source=syntax.template, colorscheme=self),
        )
        file_manager.create_file(
            f"lua/{self.name}/highlights/modules.lua",
            content=get_content(env=env, source=modules.template, colorscheme=self),
        )
        file_manager.create_file(
            f"Dockerfile",
            content=get_content(
                env=env, source=docker.dockerfile_template, colorscheme=self
            ),
            format=False
        )
        file_manager.create_file(
            f"compose.yml",
            content=get_content(
                env=env, source=docker.compose_template, colorscheme=self
            ),
            format=False
        )
        file_manager.create_file(
            f"tests/colorscheme_spec.lua",
            content=get_content(
                env=env, source=colorscheme_spec.template, colorscheme=self
            ),
        )
