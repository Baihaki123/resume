package handler

import (
	"net/http"
	"resume/cv"
	"resume/helper"

	"github.com/gin-gonic/gin"
)

type resumeHandler struct {
	service cv.Service
}

func NewResumeHandler(service cv.Service) *resumeHandler {
	return &resumeHandler{service}
}

func (h *resumeHandler) ListResume(c *gin.Context) {
	resumes, err := h.service.GetResume()
	if err != nil {
		response := helper.APIResponse("Error Get Resume", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of Resume", http.StatusOK, "success", cv.FormatResumes(resumes))
	c.JSON(http.StatusOK, response)
}

func (h *resumeHandler) GetResume(c *gin.Context) {
	var input cv.GetResumeDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Error Get Resume", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	GetResume, err := h.service.GetResumeByID(input)
	if err != nil {
		response := helper.APIResponse("Error Get Resume", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get Resume success", http.StatusOK, "success", cv.FormatResume(GetResume))
	c.JSON(http.StatusOK, response)
}

func (h *resumeHandler) CreateResume(c *gin.Context) {
	var input cv.ResumeInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Resume gagal dibuat!", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newResume, err := h.service.CreateResume(input)
	if err != nil {
		response := helper.APIResponse("Resume gagal dibuat!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Resume berhasil dibuat!", http.StatusOK, "success", cv.FormatResume(newResume))
	c.JSON(http.StatusOK, response)
}

func (h *resumeHandler) UpdateResume(c *gin.Context) {
	var inputID cv.GetResumeDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Error Update Resume", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData cv.ResumeInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Gagal update resume!", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedResume, err := h.service.UpdateResume(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Gagal update resume!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Resume berhasil diupdate!", http.StatusOK, "success", cv.FormatResume(updatedResume))
	c.JSON(http.StatusOK, response)
}

func (h *resumeHandler) DeleteResume(c *gin.Context) {
	var inputID cv.GetResumeDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Error Delete Resume", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	deleteResume, err := h.service.DeleteResume(inputID)
	if err != nil {
		response := helper.APIResponse("Error Delete Resume", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Delete Resume success", http.StatusOK, "success", cv.FormatResume(deleteResume))
	c.JSON(http.StatusOK, response)
}
