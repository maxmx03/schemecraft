dockerfile_template="""
FROM archlinux
WORKDIR /{{ colorscheme.name }}
COPY . .
RUN pacman -Syu --noconfirm
RUN pacman -S luarocks neovim gcc lua51 --noconfirm
RUN luarocks install luacheck && luarocks --lua-version=5.1 install vusted
CMD [ "vusted", "./tests" ]
"""

compose_template="""
services:
  roseline:
    build: .
    volumes:
      - .:/{{ colorscheme.name }}
"""
