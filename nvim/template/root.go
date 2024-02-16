package template

func NixShell() string {
	return `{pkgs ? import <nixpkgs> {}}:
pkgs.mkShell {
  nativeBuildInputs = with pkgs.buildPackages; [luajitPackages.vusted];
}`
}

func DockerFile() string {
	return `FROM archlinux
WORKDIR /{{lower .Name}}
COPY . .
RUN pacman -Syu --noconfirm
RUN pacman -S luarocks neovim gcc lua51 --noconfirm
RUN luarocks install luacheck && luarocks --lua-version=5.1 install vusted
CMD [ "vusted", "./tests" ]`
}

func DockerCompose() string {
	return `services:
  {{lower .Name}}:
    build: .
    volumes:
      - .:/{{lower .Name}}`
}
