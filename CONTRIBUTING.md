
## Editing Themes

For now the best way to work on a theme is to run `static-docs` to produce the site:

```
$ make
```

Which produces the following files:

```
build
├── index.html
└── theme
    └── apex
        ├── css
        │   └── index.css
        ├── js
        │   ├── index.js
        │   └── smoothscroll.js
        └── views
            └── index.html
```

Open the site and edit anything you like.

```
$ open build/index.html
```

Copy and paste any changes back into `./docs/themes/apex`.

This process will be improved at some point.
