package handler

import (
	"errors"
	"github.com/blazee5/EffectiveMobile_test/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"net/http"
	"strconv"
)

// CreateCars godoc
// @Summary Create new cars
// @Description Create new cars with the provided specifications
// @Tags cars
// @Accept  json
// @Produce  json
// @Param   cars body domain.CreateCarsRequest true "Create Cars Request"
// @Success 200 {object} []domain.Car
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cars [post]
func (s *Server) CreateCars(c *gin.Context) {
	var input domain.CreateCarsRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	cars, err := s.services.Car.CreateCars(c.Request.Context(), input)

	if err != nil {
		s.log.Infof("error while create car: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}

	c.JSON(http.StatusOK, cars)
}

// GetCars godoc
// @Summary Get cars
// @Description Get cars based on query parameters
// @Tags cars
// @Accept  json
// @Produce  json
// @Param   mark query string false "Car mark"
// @Param   model query string false "Car model"
// @Param   year query int false "Car year"
// @Param   regNum query string false "Car registration number"
// @Param   limit query int false "Limit for pagination"
// @Param   offset query int false "Offset for pagination"
// @Param   name query string false "Owner's name"
// @Param   surname query string false "Owner's surname"
// @Param   patronymic query string false "Owner's patronymic"
// @Success 200 {object} []domain.Car
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cars [get]
func (s *Server) GetCars(c *gin.Context) {
	var input domain.GetCarsRequest

	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	cars, err := s.services.Car.GetCars(c.Request.Context(), input)

	if err != nil {
		s.log.Infof("error while get cars: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}

	c.JSON(http.StatusOK, cars)
}

// UpdateCar godoc
// @Summary Update a car
// @Description Update car details by ID
// @Tags cars
// @Accept  json
// @Produce  json
// @Param   id path int true "Car ID"
// @Param   car body domain.UpdateCarRequest true "Update Car Request"
// @Success 200 {object} domain.Car
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cars/{id} [put]
func (s *Server) UpdateCar(c *gin.Context) {
	var input domain.UpdateCarRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "invalid id",
		})
		return
	}

	cars, err := s.services.Car.UpdateCar(c.Request.Context(), id, input)

	if errors.Is(err, pgx.ErrNoRows) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "car not found",
		})
		return
	}

	if err != nil {
		s.log.Infof("error while update car: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}

	c.JSON(http.StatusOK, cars)
}

// DeleteCar godoc
// @Summary Delete a car
// @Description Delete a car by ID
// @Tags cars
// @Accept  json
// @Produce  json
// @Param   id path int true "Car ID"
// @Success 200 {string} string "OK"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cars/{id} [delete]
func (s *Server) DeleteCar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "invalid id",
		})
		return
	}

	err = s.services.Car.DeleteCar(c.Request.Context(), id)

	if err != nil {
		s.log.Infof("error while delete car: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}

	c.JSON(http.StatusOK, "OK")
}
