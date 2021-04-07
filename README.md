# Go_Blog
本專案使用 Go 語言之 Gin 框架，實現簡易的 Blog 文章、標籤之 Restful API。

透過 API ( JWT 驗證 ) 與 PostgreSQL 資料庫進行 CRUD 操作。

---

### 目錄結構
>go_blog/<br>
├── conf<br>
├── middleware<br>
├── models<br>
├── pkg<br>
├── routers<br>
└── runtime

* conf: 用於儲存配置文件<br>
* middleware: 應用中間件
* models: 應用資料庫模型
* pkg: 自訂工具、第三方套件
* routers: 路由邏輯處理
* runtime: 應用運行時資料

---

### 配置文件
在 go_blog/conf/app.ini 文件中，輸入資料庫的帳號與密碼。

---

### 資料庫 - 資料表
![blog_tag](/table_structure/blog_tag.png)

標籤表

![blog_article](/table_structure/blog_article.png)

文章表

![blog_auth](/table_structure/blog_auth.png)

驗證表

---

### 預覽 Restful API
運行本專案後，前往下列網址查看
http://127.0.0.1:8000/swagger/index.html

JWT 驗證之使用者帳號: test

JWT 驗證之使用者密碼: test123456