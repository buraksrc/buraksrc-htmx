## Description
My personal website (currently under construction and not final) made with Go (only standard libraries), HTMX and Tailwind CSS. I plan to use no JavaScript whatsoever. I will checkout HyperScript if necessary but I don't think I will need it.

## CSS 
Tailwind CSS can be used without PostCSS but tailwindcss binary has to be installed and added to your PATH

```
tailwindcss -i public/css/tailwind.css --content "./*.html,./**/*.html" -o public/css/app.css --minify
```

can be used while I finish Makefile (if I have time)

## TODO
Add makefile\
Fix public routes\
Add stuff (about page, some demos made with HTMX and a blog page with markdown reader)\
Create Dockerfile for Docker\
Add LICENSE