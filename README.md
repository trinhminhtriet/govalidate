# ✅ govalidate

```text

                             _  _      _         _
  __ _   ___  __   __  __ _ | |(_)  __| |  __ _ | |_   ___
 / _` | / _ \ \ \ / / / _` || || | / _` | / _` || __| / _ \
| (_| || (_) | \ V / | (_| || || || (_| || (_| || |_ |  __/
 \__, | \___/   \_/   \__,_||_||_| \__,_| \__,_| \__| \___|
 |___/

```

## ✨ Features

Validates your Go installation and dependencies:

✅ Checks the Go installation and version.

🔍 Ensures the PATH is correctly configured.

🔧 Verifies CGO dependencies are installed.

🛠️ Checks plugin support for available editors.

## 🚀 Installation

To install **govalidate**, simply clone the repository and follow the instructions below:

```bash
git clone git@github.com:trinhminhtriet/govalidate.git
cd govalidate

make release
```

## 💡 Usage

Run **govalidate** with the following command to start cleaning your filesystem:

```sh
[✔] Go (go1.23.3)
[✗] Checking if $PATH contains "/Users/triettrinh/go/bin"
    Add "/Users/triettrinh/go/bin" to your $PATH.
    On Unix systems:
    export PATH=$PATH:/Users/jbd/go/bin
[✔] Checking gcc for CGO support
[✔] Vim Go plugin
[!] VSCode Go extension
    VSCode Go extension is not installed.
    See https://code.visualstudio.com/docs/languages/go to install.

```

## TODO

## 🤝 How to contribute

We welcome contributions!

- Fork this repository;
- Create a branch with your feature: `git checkout -b my-feature`;
- Commit your changes: `git commit -m "feat: my new feature"`;
- Push to your branch: `git push origin my-feature`.

Once your pull request has been merged, you can delete your branch.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
