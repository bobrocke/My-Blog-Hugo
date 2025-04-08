---
title: "New Draft"
description: ""
summary: ""
date: "2025-03-17T08:36:13-04:00"
thumbnail: ""
draft: true
categories: []
tags: []
---

Theme dependent:
Jekyll -- `description`
Hugo -- `summary` or
Astro -- `excerpt`

<!--more-->

Jekyll -- `date` format -- `2024-06-05 13:20:41 -0000`
Hugo -- `date` format -- `2024-06-05 13:20:41`
Astro -- `publishDate` format -- `2023-04-03 21:33:13`
Hugo -- `lostmod` format -- format -- `2024-06-05 13:20:41`

Path to images
Jekyll -- `/assets/image_path` with images in `/assets/images`
Hugo -- `/images/image_path` with images in `/static/images`
Astro -- `/image_path` with images in `/public`

Jekyll -- Gems, rbenv
Eleventy -- basically node modules, lots of Liquid warnings/errors in Zed, plugin are needed for the functionality I want and I worry bout them staying maintained.
Hugo -- no Zed support for Go template language

daisyUI has its own classes that combine TailwindCSS classes. Copy/paste HTML.

[Flowbite](https://flowbite.com/) has components using TailwindCSS classes and its own JavaScript library. Copy/paste HTML. CDN for JavaScript
