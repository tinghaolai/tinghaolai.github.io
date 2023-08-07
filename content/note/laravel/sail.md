---
title: "sail.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["laravel"]
---



---

## Permission issue

run `./vendor/bin/sail test` and got

> PHP Warning:  file_put_contents(/var/www/html/.phpunit.result.cache): Failed to open stream: Permission denied in /var/www/html/vendor/phpunit/phpunit/src/Runner/ResultCache/DefaultResultCache.php on line 143

* Enter container through exec -it or `./vendor/bin/sail root-shell`
* chown -R sail:sail {laravel_project_dir}



---

