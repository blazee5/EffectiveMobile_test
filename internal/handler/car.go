package handler

import (
	"github.com/blazee5/EffectiveMobile_test/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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

	if err != nil {
		s.log.Infof("error while update car: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}

	c.JSON(http.StatusOK, cars)
}

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
