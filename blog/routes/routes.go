package routes

import (
	"github.com/gin-contrib/gzip"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tingwei628/pgo/blog/models"
	"github.com/tingwei628/pgo/blog/utils"
	"gorm.io/gorm"
)

var (
	addHtmlPath   = "add.html"
	editHtmlPath  = "add.html"
	indexHtmlPath = "index.html"
	viewHtmlPath  = "view.html"
)
var db *gorm.DB

func SetupRouter(database *gorm.DB) *gin.Engine {
	db = database
	r := gin.Default()

	r.Use(gzip.Gzip(gzip.DefaultCompression))

	// Templates with helper functions
	templatePath := filepath.Join(utils.Basepath, "templates/*.html")
	r.SetFuncMap(template.FuncMap{
		"seq": func(start, end int) []int {
			s := make([]int, end-start+1)
			for i := range s {
				s[i] = start + i
			}
			return s
		},
		"add": func(a, b int) int { return a + b },
		"sub": func(a, b int) int { return a - b },
	})
	r.LoadHTMLGlob(templatePath)

	// Routes
	r.GET("/", Index)
	r.GET("/blog/add", AddPage)
	r.POST("/api/blog", CreateBlog)
	r.GET("/blog/edit/:id", EditPage)
	r.POST("/api/blog/put/:id", UpdateBlog)
	r.POST("/api/blog/delete/:id", DeleteBlog)
	r.GET("/blog/view/:id", ViewPage)
	r.GET("/api/blog/:id", GetBlog)
	r.GET("/api/blogs", GetBlogsByPage)

	return r
}

type PostWithDate struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func Index(c *gin.Context) {
	pageStr := c.Query("page")
	page, _ := strconv.Atoi(pageStr)
	if page < 1 {
		page = 1
	}
	const pageSize = 5

	var total int64
	db.Model(&models.Post{}).Count(&total)

	var posts []models.Post
	if err := db.Order("created_at DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&posts).Error; err != nil {
		log.Printf("query failed: %v", err)
		c.String(http.StatusInternalServerError, "query failed")
		return
	}

	postsFormatted := make([]PostWithDate, len(posts))
	for i, p := range posts {
		postsFormatted[i] = PostWithDate{
			ID:        p.ID,
			Title:     p.Title,
			CreatedAt: p.CreatedAt.Format("2006-01-02"),
			UpdatedAt: p.UpdatedAt.Format("2006-01-02"),
		}
	}

	totalPages := (int(total) + pageSize - 1) / pageSize
	c.HTML(http.StatusOK, indexHtmlPath, gin.H{
		"Posts":      postsFormatted,
		"Page":       page,
		"TotalPages": totalPages,
	})
}

func GetBlogsByPage(c *gin.Context) {
	pageStr := c.Query("page")
	page, _ := strconv.Atoi(pageStr)
	if page < 1 {
		page = 1
	}
	const pageSize = 5

	var total int64
	db.Model(&models.Post{}).Count(&total)

	var posts []models.Post
	db.Order("created_at DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&posts)

	postsFormatted := make([]PostWithDate, len(posts))
	for i, p := range posts {
		postsFormatted[i] = PostWithDate{
			ID:        p.ID,
			Title:     p.Title,
			CreatedAt: p.CreatedAt.Format("2006-01-02"),
			UpdatedAt: p.UpdatedAt.Format("2006-01-02"),
		}
	}

	totalPages := (int(total) + pageSize - 1) / pageSize
	c.JSON(http.StatusOK, gin.H{
		"posts":      postsFormatted,
		"page":       page,
		"totalPages": totalPages,
	})
}

func AddPage(c *gin.Context) {
	c.HTML(http.StatusOK, addHtmlPath, gin.H{
		"IsEdit": false,
	})
}

func CreateBlog(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "create failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":        post.ID,
		"title":     post.Title,
		"content":   post.Content,
		"createdAt": post.CreatedAt.Format("2006-01-02 15:04:05"),
		"updatedAt": post.UpdatedAt.Format("2006-01-02 15:04:05"),
	})
}

func EditPage(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var post models.Post
	db.First(&post, id)
	c.HTML(http.StatusOK, editHtmlPath, gin.H{
		"IsEdit": true,
		"Post":   post,
	})
}

func UpdateBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Model(&models.Post{}).Where("id = ?", id).Updates(post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}
	var updatedPost models.Post
	if err := db.First(&updatedPost, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "fetch updated post failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        updatedPost.ID,
		"title":     updatedPost.Title,
		"content":   updatedPost.Content,
		"createdAt": updatedPost.CreatedAt.Format("2006-01-02 15:04:05"),
		"updatedAt": updatedPost.UpdatedAt.Format("2006-01-02 15:04:05"),
	})
}

func DeleteBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	db.Unscoped().Delete(&models.Post{}, id)
	c.Redirect(http.StatusSeeOther, "/")
}

func ViewPage(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var post models.Post
	if err := db.First(&post, id).Error; err != nil {
		c.String(http.StatusNotFound, "blog does not exist")
		return
	}
	c.HTML(http.StatusOK, viewHtmlPath, gin.H{"Post": post})
}

func GetBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var post models.Post
	if err := db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "blog does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":        post.ID,
		"title":     post.Title,
		"content":   post.Content,
		"createdAt": post.CreatedAt.Format("2006-01-02"),
		"updatedAt": post.UpdatedAt.Format("2006-01-02"),
	})
}
