package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadFile 上传文件
func UploadFile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "上传成功",
		"data": gin.H{
			"file_id":        "file_123",
			"book_id":        1,
			"user_id":        1,
			"filename":       "发票照片.jpg",
			"original_name":  "IMG_1234.jpg",
			"mime_type":      "image/jpeg",
			"size":           "1.5MB",
			"size_bytes":     1572864,
			"path":           "/uploads/files/2023/01/05/file_123.jpg",
			"url":            "https://example.com/uploads/files/2023/01/05/file_123.jpg",
			"thumbnail_url":  "https://example.com/uploads/files/2023/01/05/thumb_file_123.jpg",
			"entity_type":    "transaction",
			"entity_id":      101,
			"description":    "发票照片",
			"created_at":     "2023-01-05T12:30:00Z",
			"updated_at":     "2023-01-05T12:30:00Z",
		},
	})
}

// DownloadFile 下载文件
func DownloadFile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "下载成功",
		"data": gin.H{
			"file_id":   "file_123",
			"filename":  "发票照片.jpg",
			"mime_type": "image/jpeg",
			"size":      "1.5MB",
		},
	})
}

// DeleteFile 删除文件
func DeleteFile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "文件删除成功",
		"data": gin.H{
			"file_id":     "file_123",
			"deleted_at":  "2023-01-05T12:35:00Z",
		},
	})
}

// PreviewFile 预览文件
func PreviewFile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "预览成功",
		"data": gin.H{
			"file_id":   "file_123",
			"filename":  "发票照片.jpg",
			"mime_type": "image/jpeg",
			"preview_url": "https://example.com/preview/file_123.jpg",
		},
	})
}

// GetFileList 获取文件列表
func GetFileList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"total":     5,
			"page":      1,
			"page_size": 20,
			"list": []gin.H{
				{
					"file_id":        "file_123",
					"filename":       "发票照片.jpg",
					"original_name":  "IMG_1234.jpg",
					"mime_type":      "image/jpeg",
					"size":           "1.5MB",
					"url":            "https://example.com/uploads/files/2023/01/05/file_123.jpg",
					"thumbnail_url":  "https://example.com/uploads/files/2023/01/05/thumb_file_123.jpg",
					"entity_type":    "transaction",
					"entity_id":      101,
					"description":    "发票照片",
					"created_at":     "2023-01-05T12:30:00Z",
				},
			},
		},
	})
}

// AttachTransactionFile 关联交易附件
func AttachTransactionFile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "附件关联成功",
		"data": gin.H{
			"transaction_id": 101,
			"attached_count": 2,
			"file_ids":       []string{"file_123", "file_456"},
		},
	})
}
