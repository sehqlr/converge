---
title: "file.content"
slug: "file-content"
date: "2016-10-04T13:01:49-05:00"
menu:
  main:
    parent: resources
---


Content renders content to disk


## Example

```hcl
param "message" {
  default = "Hello, World in {{param `filename`}}"
}

param "filename" {
  default = "test.txt"
}

file.content "render" {
  destination = "{{param `filename`}}"
  content     = "{{param `message`}}"
}

```


## Parameters

- `content` (string)

  Content is the file content. This will be rendered as a template.

- `destination` (string)

  Destination is the location on disk where the content will be rendered.


