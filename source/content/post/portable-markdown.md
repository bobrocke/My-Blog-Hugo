---
title: Writing Portable Markdown
date: 2025-04-09
lastmod:
summary:
description:
draft: true
categories: [Web Development]
tags: [Markdown]
---

After moving my blog from WordPress to Jekyll to Hugo, I came to better appreciate the need to write markdown that is independent of the system used to create a blog. That may mean forgoing a special feature in one system, but if (when) that markdown file is going to run though another blogging engine, the effort will be worth it.

<!--more-->

Shortcodes
`<!--more-->`
Linking to other blog posts

- 11ty post-name/index.html
- Hugo post-name

Some custom code will be needed to 'convert' from one directory structure to another.
