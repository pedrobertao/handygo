# handygo

**Lightweight helpers for Go — small, composable functions that make everyday coding faster and easier.**

Handygo is not a framework. It’s a growing collection of focused utilities for strings, dates, slices, and more — the little tools you wish the Go standard library had.
Zero dependencies, idiomatic APIs, and examples for every function.

> ⚡ Build faster. ✨ Write cleaner code. 🧩 Import only what you need.

---

## ✨ Why handygo?

- **Lightweight** — no heavy dependencies, just pure Go.
- **Composable** — import only the package you need (`strx`, `timex`, …).
- **Familiar** — functions follow predictable, consistent naming.
- **Tested** — edge cases covered (UTF-8, leap years, DST).
- **Documented** — examples for every helper, visible on pkg.go.dev.

---

## 📦 Installation

```bash
go get github.com/yourname/handygo/strx
go get github.com/yourname/handygo/timex
```

---

## 🚀 Quick Examples

### Strings (`strx`)

```go
strx.ToSnake("HelloWorld")        // "hello_world"
strx.Slug("Olá, Mundo!")          // "ola-mundo"
strx.TruncateWords("lorem ipsum dolor sit amet", 3)
// "lorem ipsum dolor…"
```

### Dates (`timex`)

```go
now := time.Now()
timex.StartOfMonth(now)           // 2025-08-01 00:00:00
timex.AddDays(now, 7)             // same time next week
timex.IsWeekend(now)              // true/false
```

---

## 📂 Packages

- **`strx`** — string utilities: case conversion, slugify, truncate, pad, reverse.
- **`timex`** — date/time helpers: add days/months, start/end of periods, diffs.
- _(coming soon)_ **`slice`**, **`envx`**, **`errx`**, **`httpx`**

---

## 🛠 Roadmap

- [x] Strings (`strx`) basics
- [x] Dates (`timex`) basics
- [ ] Slice helpers (`slice`)
- [ ] Env/config helpers (`envx`)
- [ ] HTTP helpers (`httpx`)
- [ ] Error helpers (`errx`)

---

## 🤝 Contributing

Contributions welcome! Handygo is designed to be community-driven.

- Open an issue for new helper ideas.
- Add examples/tests with every PR.
- Keep it idiomatic, simple, and dependency-free.

---

## 📜 License

MIT — do whatever you like, just keep it handy.

---

👉 This gives you a **professional OSS landing page** right away. Would you like me to also prepare a **minimal `ci.yml` GitHub Actions workflow** snippet for automated testing/linting so the repo feels production-ready from day one?
