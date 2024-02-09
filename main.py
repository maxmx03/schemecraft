from src.neovim import Neovim

def main():
    import sys
    import yaml

    if len(sys.argv) < 2:
        print("Usage: python main.py <yaml_file>")
        sys.exit(1)

    yaml_path = sys.argv[1]
    with open(yaml_path, "r") as yaml_file:
        yaml_content = yaml.safe_load(yaml_file)

    neovim = Neovim(
        name=yaml_content["name"],
        author=yaml_content["author"],
        contact=yaml_content["contact"],
        repo=yaml_content["repo"],
        license=yaml_content["license"],
        palette=yaml_content["palette"],
        default_config=yaml_content["default_config"],
        highlights=yaml_content["highlights"],
    )

    neovim.create_colorscheme()

    print("Colorscheme was generated with success")


if __name__ == "__main__":
    main()
