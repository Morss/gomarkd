# gomarkd

**gomarkd** is a simple and fast tool written in Go that converts Markdown files into clean, styled HTML pages — served on-the-fly. It's ideal for building simple static websites powered by Markdown, with minimal setup and full control over layout via Go templates.

## Features

- Converts Markdown to HTML dynamically at runtime
- Create static websites using plain `.md` files
- Fully customizable via Go `html/template`
- Clean URL routing (`/page-name` from `page-name.md`)
- Can be used as a lightweight blog framework or note sharing

## Getting Started

```bash
git clone https://github.com/yourusername/gomarkd.git
cd gomarkd

# Run the development server
go run .

# Add new pages in /content dir. Will be loaded on-the-fly
echo "# My new page\nLorem ipsum" > content/New\ Page.md
