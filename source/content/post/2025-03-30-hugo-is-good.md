---
title: Hugo is Good
date: 2025-03-30
lastmod:
summary:
description:
draft: true
categories: []
tags: []
---

This site is running on Hugo, a popular Static Site Generator (SSG). I tried [Jekyll](https://jekyllrb.com/), [Hugo](https://gohugo.io/), [Astro](https://astro.build/), and [Eleventy/11ty](https://www.11ty.dev/) on the way here, but settled on Hugo.

<!--more-->

All of those SSGs were quite good and in wide use, but didn't quite fit, for me.

Jekyll
: One of the original SSG, but its getting a bit old, now. Ruby isn't a popular as it once was and many of the available Jekyll themes were released years ago and not updated since. Gems can be a bit of a pain to work with, too, because each version of Ruby on your machine needs its own set of gems.

Astro
: A very popular option, but a bit too much for this site. It's really a complete app framework that goes far beyond what I need. It's just too 'heavy'.

Eleventy/11ty
: Another super popular SSG that's fast and extendable via plugins. A little _too_ extendable for my tastes---it's very basic out-of-the-box and needs those plugins. But it's cool how many template languages it supports natively.

An often reported, stand out feature of Hugo is its speed in generating content. This is not a big site but it takes Hugo only 1.2 seconds to generate 82 pages, 16 paginated pages, and move 45 images. Changing a couple of pages take Hugo milliseconds to regenerate and serve. It feels almost instant and that makes development pleasant.

As opposed to Jekyll's gems and 11ty's node modules, Hugo is a single binary with no dependencies to manage unless you want to. I've added only [TailwindCSS](https://tailwindcss.com/, [Flowbite](https://flowbite.com/), and [PostCSS](https://postcss.org/)---all popular and well-maintained packages.

There's built-in support for multiple taxonomy page displays and the posts for each category within. This site uses 'categories' and 'tags', but many more are possible. 11ty can't do that without plugins and extra custom code.

But there's a bit more of a learning curve for Hugo than the others. Hugo is an all Go environment---it's written in Go and uses only the Go Template Language. If you've been working in more widely used template languages, you'll need a bit to get up to speed. And I don't find its [documentation](https://gohugo.io/documentation/) easy to navigate so I often use a web search to find what I need.
