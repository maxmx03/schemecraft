package template

func GithubWorkFlows() (string, string) {
	var release = `on:
  push:
    branches:
      - main

permissions:
  contents: write
  pull-requests: write

name: release-please

jobs:
  release-please:
    runs-on: ubuntu-latest
    steps:
      - uses: google-github-actions/release-please-action@v3
        with:
          release-type: simple
          package-name: release-please-action`

	var test = `name: Test

on: [push]

jobs:
  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: Stylua
        uses: JohnnyMorganz/stylua-action@v3
        with:
          token: {{ print "${{ secrets.GITHUB_TOKEN }}" }}
          version: latest
          args: --check .


  docs:
    runs-on: ubuntu-latest
    name: pandoc to vimdoc
    steps:
      - uses: actions/checkout@v3
      - name: panvimdoc
        uses: kdheepak/panvimdoc@main
        with:
          vimdoc: {{.Name}}
          treesitter: true
          version: "{{.Version}}"
      - uses: stefanzweifel/git-auto-commit-action@v4
        with:
          commit_message: 'chore(doc): auto generate docs'
          commit_user_name: "github-actions[bot]"
          commit_user_email: "github-actions[bot]@users.noreply.github.com"
          commit_author: "github-actions[bot] <github-actions[bot]@users.noreply.github.com>"

  test:
    name: Run Test
    runs-on: {{ print "${{ matrix.os }}" }}
    strategy:
      matrix:
        os: [ubuntu-latest, macos-latest]
    steps:
      - uses: actions/checkout@v3
      - uses: rhysd/action-setup-vim@v1
        with:
          neovim: true

      - name: luajit
        uses: leafo/gh-actions-lua@v10
        with:
          luaVersion: "luajit-2.1.0-beta3"

      - name: luarocks
        uses: leafo/gh-actions-luarocks@v4

      - name: run test
        shell: bash
        run: |
          luarocks install luacheck
          luarocks install vusted
          vusted ./tests`

	return release, test
}
