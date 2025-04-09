---
title: Settling on TailwindCSS
date: 2025-04-08
lastmod:
summary:
description:
draft: true
categories: []
tags: []
---

While experimenting with various static site generators, I've also been learning about [Bulma](https://bulma.io/) (a traditional CSS framework) and [TailwindCSS](https://tailwindcss.com/) (a 'new style' utility CSS framework). They come at the idea of a CSS framework from two very different perspectives, and if you've seen the Internet discussions, you'll know that opinions are decidedly mixed.

<!--more-->

A 'problem' in many situations is that most of a page's content is generated from a markdown file. The site author won't have access to the HTML in order to style it in the normal way. Both Bulma and TailwindCSS have classes, `content` for Bulma and `prose` for TailwindCSS, that can be applied to a `<div>` enclosing the markdown output. Something like:

```HTML
<div class="prose">
  {{ markdown }}
</div>
```

Those classes can then style the HTML tags output by the markdown renderer.
