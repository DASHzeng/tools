# gopls

I have ported some features such as references, rename, workspace symbol, implementation of [bingo](https://github.com/saibing/bingo)  to gopls

## Install

gopls is a go module project, so you need install [Go 1.12 or above](https://golang.google.cn/dl/),
to  install the `gopls`, please run

```bash
git clone -b bingo https://github.com/saibing/tools.git
cd tools/gopls
go install
```

## Language Client

### [vscode-go](https://github.com/Microsoft/vscode-go)

```json
{
    "go.useLanguageServer": true,
    "go.alternateTools": {
        "go-langserver": "gopls"
    },
    "go.languageServerExperimentalFeatures": {
        "format": true,
        "autoComplete": true
    },
    "[go]": {
        "editor.snippetSuggestions": "none",
        "editor.formatOnSave": true,
        "editor.codeActionsOnSave": {
            "source.organizeImports": true
        },
    },
    "gopls": {
        "usePlaceholders": true,
        "enhancedHover": true
    }
}
```

### [coc.nvim](https://github.com/neoclide/coc.nvim)

```json
{
  "languageserver": {
    "golang": {
      "command": "gopls",
      "args": [],
      "rootPatterns": ["go.mod", ".vim/", ".git/", ".hg/"],
      "filetypes": ["go"]
    }
  }
}
```

## Google offical gopls wiki

[https://github.com/golang/go/wiki/gopls](https://github.com/golang/go/wiki/gopls)

