### vue-seo-playwright
让 `vue` 爬虫支持 `seo`，这种方法其实就是使用 `playwright` 浏览器进行渲染，然后返回 `html` 代码。 

#### nginx 配置文件
 
```text
upstream spider_server {
	  server localhost:8081;
}

server {
    listen 80;
    server_name site_name;
    access_log /any;
    error_log /any warn;
	charset utf-8;
    location / {
		set $flag 0;
		if ($http_user_agent ~* "Baiduspider|bingbot|360Spider") {
			set $flag 1;
		}

		if ($flag = 1) {
			proxy_pass  http://spider_server;
		}

		root /work/apps/rain_dog;
		try_files $uri $uri/ /index.html;
    }
    error_page 500 502 503 504 /50x.html;
    location = /50x.html {
        root /usr/share/nginx/html;
    }

    location ~ /\.ht {
        deny all;
    }
}
```

### 关联项目

[浏览器自动化框架](https://github.com/dmf-code/automationFramework)
