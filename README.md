# 🛠️ yask (YAML Skeleton)

**yask** is a handy tool that generates a project structure from a `YAML` file. Define your folders, files, and their content in a simple `YAML` format — and get your entire project skeleton created automatically.

## Perfect for:

- 🔧 Scaffolding new projects
- 🧪 Prototyping and experimenting
- 📚 Teaching and tutorials
- ♻️ Automating boilerplate setups

---

## ✨ Features

- 📁 Generate files and directories from YAML
- 🔄 Supports deeply nested structures
- 📝 Inline file content embedding
- 🧠 Special keywords for dynamic behavior
- ⚙️ Execute shell commands after generation
- 🔌 Include external files and reusable content blocks

---

## 📄 Example

Here's an example `config.yaml`:

```yaml
skel:
  backend:
    go.mod: |
      module my-project

      go 1.24.0
    cmd:
      main.go: |
        package main

        import "fmt"

        func main() {
            fmt.Println("Hello, YASK!")
        }
    configs:
      files:
        - dev.yaml
        - prod.yaml
    internal:
      dirs:
        - app
        - database
        - service
    pkg:
      tools:
        tools.go: package tools
    Makefile: '$./Makefile'
    README.md: '#/content/readme'
    exec:
      - go mod tidy
      - go get github.com/av-ugolkov/yask
content:
  readme: |
    # yask (YAML Skeleton)

    This tool can generate structure
```

---

## 🧾 Output Structure:

```css
backend/
├── cmd/
│ └── main.go
├── configs
│ ├── dev.yaml
│ └── prod.yaml
├── internal
│ ├── app
│ ├── database
│ └── service
├── go.mod
├── pkg/
│ └── tools/
│   └── tools.go
├── Makefile ← included from local file
└── README.md ← generated from content block
```

---

## 🚀 Get Started

1. Install `yask`:
   ```bash
   go install github.com/av-ugolkov/yask@latest
   ```
2. Create a `config.yaml` file.
3. Run:
   ```bash
   yask -c=config.yaml
   ```
   or example with placeholders
   ```bash
   yask -c ./config.yaml -p folder=TestFolder -p file1=about -p file2=dynamic.txt -p makefile="ls -a"
   ```

---

## 🆘 Help

### 📦 Available Commands:

- `help` — Help about any command
- `version` — Print the version number of yask

### 🏁 Flags:

- `-c`, `--config string` — Path to config YAML (**required**)
- `-h`, `--help` — Show help for yask
- `-p`, `--placeholder key=value` — Modifiable parameters that are substituted into the configuration file.

---

## 🔑 YAML Keywords

| Key       | Description                                                                       |
| --------- | --------------------------------------------------------------------------------- |
| `skel`    | Defines the directory and file structure                                          |
| `content` | Reusable content blocks (referenced by `#/path`)                                  |
| `exec`    | List of shell commands to run after generation                                    |
| `dirs`    | List of directories to create                                                     |
| `files`   | List of files to create.                                                          |
| `@`       | Read external files and write them to the created file (`@./relative/path`)       |
| `#`       | Read the contents of the path and write it to the created file (`#/content/path`) |
| `$`       | Sets a placeholder that can be changed (`${placeholder}`)                         |

---

## 💡 Tips

- Use `#/content/...` to reuse documentation blocks.
- Use `@./path` to include local file content (like Makefiles or licenses).
- Use `exec` to automatically set up dependencies or tools.
- Use `folder_1/folder_2/.../folder_N` to create nested folder hierarchy in one line.

---

## 🧩 Use Cases

- Initialize consistent project templates
- Share scaffolding configs across teams
- Teach beginners project structure visually
- Automate environment setup with minimal effort

---

## 📃 License

MIT — feel free to use and contribute!

---

## 🙌 Contribute

Pull requests and issues are welcome! Let's build better skeletons together 🦴

---

Made with ❤️ by [@av-ugolkov](https://github.com/av-ugolkov)
