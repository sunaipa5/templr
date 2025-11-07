# templr

templr is a simple CLI tool that creates new project folders from predefined templates.
It copies a template directory into a new project directory so you can start coding quickly.

---

## Install

### From source

```bash
go install github.com/sunaipa5/templr@latest
```

Make sure your Go bin path is in your `PATH`.

### Download (Releases)

You can download prebuilt binaries from:

[Go releases](https://github.com/sunaipa5/templr/releases)

---

## Template Location

templr looks for templates in:

**Windows:**

```
C:\Users\<user>\Documents\templr\
```

**Linux / macOS:**

```
~/Documents/templr/
```

Each template must be a folder inside `templr/`.

Example:

```
templates/
    go-api/
    vue-app/
    flutter-app/
```

---

## Commands

### List available templates

```bash
templr list
```

### Create a new project

```bash
templr init <templateName> <projectName>
```

Example:

```bash
templr init go-api myproject
```

This will copy:

```
Documents/templr/templates/go-api/
```

into:

```
./myproject/
```
