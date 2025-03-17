---
layout: post
title: Zed is Good
summary: Programming editors and Integrated Development Environments (IDEs) are very personal choices because developers spend so much of their day working with them. VS Code is immensely popular as are the various JetBrains IDEs. VS Code is free, configurable, and powerful which is why so many develops prefer it. The very new Zed is gaining popularity fast.
image:
categories: [Web Development]
tags: [Zed]
date: 2025-02-05 17:55:00
---

Programming editors and Integrated Development Environments (IDEs) are very personal choices because developers spend so much of their day working with them. VS Code is immensely popular as are the various JetBrains IDEs. VS Code is free, configurable, and powerful which is why so many develops prefer it. The very new Zed is gaining popularity fast.

<!--more-->

I've used many IDEs over the years, especially as programming editors such as BBEdit were slowly replaced by IDEs. Where once BBEdit was my go-to, it wasn't long after Coda was release that I switched to it. When it ran out of steam, I went to VS Code. But with the release of Nova, I was back to a Mac-native IDE.

As good as Nova is, it doesn't fit as well in my experiments with Ruby on Rails and Laravel. It's syntax highlighting is not so good for `html.erb` files nor `blade.php` files. Both of those frameworks make great use of a terminal and I often have two, or even three, open at at time -- development server, framework console, and a command line. Each Nova terminal is a separate tab in the editor which crowds out the code tabs. Nova is just not keeping up with web development beyond HTML, CSS, and JavaScript. Enter Zed.

My primary complaint with Zed is that it is not Mac-native. Beyond that, I am very impressed with what is pre-release software (version 0.171.6, as I write) and I'm using it for almost all of my coding. It is nicely designed, easily configurable (with text files), offers a bunch of themes, and is really fast -- almost instant. For my needs, its terminal pane at the bottom of the working window is just perfect; a click and it's there, a click and it's gone. Better yet, I can have a number of full-width terminals running at the same time. **Much** better than either Nova or VS Code.

Zed's foundations are somewhat unique (and beyond my full understanding) being written in Rust, tied in with AI models that help with coding, and using [Tree-sitter models for parsing syntax](https://zed.dev/blog/syntax-aware-editing). I'm very impressed with the quality and speed of the syntax highlighting. The colors are consistent across tokens and change live as you type. Zed does not, however, auto-complete closing tags such as `</div>` or `{%`, which is a little surprising. It is very good at auto-indention, so if you type `</div>`, Zed will line it up properly with its matching `</div>`, for example.

Right now, Zed is stuck between being a full IDE and a programmer's editor. Its file and directory listing in the left pane works really well and allows renaming, deleting, moving, creating, and more. But its support for Git is in the early stages limited to coloring the names in the file tree for new, changed, or ignored files. There's no attempt at FTP or SFTP -- I'm guessing those are too old fashioned now. The concept of projects, or switching between them, is limited to _Open Recent ..._. But, as handy-to-have as these things are, I don't miss them enough to make me want to change back to Nova or VS Code. Except for writing this blog - Zed doesn't have a full-featured spell checker and I'm a terrible speller -- so I use Nova for that.

Updates are frequent, happening at least every Wednesday. If you're more the bleeding edge type, preview and nightly builds are available. This is an IDE to keep your eye on.
