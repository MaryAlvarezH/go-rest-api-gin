package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Campaign struct {
	ID   int
	Name string
}

var campaigns = []Campaign{
	{1, "Campaign 1"},
	{2, "Campaign 2"},
	{3, "Campaign 3"},
}

func main() {
	r := gin.Default()
	r.GET("/api/v1/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api/v1")
	{
		api.GET("/campaigns", getCampaigns)
		api.GET("/campaigns/:campaignID", getCampaign)
		api.POST("/campaigns", addCampaign)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// el servidor por peticion de cada cliente va a generando un contexto
func getCampaigns(c *gin.Context) {
	c.JSON(http.StatusOK, campaigns)
}

func getCampaign(c *gin.Context) {
	campaignID := c.Param("campaignID")
	// c.Query("timezone")
	cID, err := strconv.ParseInt(campaignID, 10, 64)
	if err != nil || cID < 1 || cID > 3 {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, campaigns[cID-1])
}

func addCampaign(c *gin.Context) {
	var newCampaign Campaign
	if err := c.Bind(&newCampaign); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{})
		return
	}

	campaigns = append(campaigns, newCampaign)
	c.JSON(http.StatusCreated, gin.H{})
}

// Buenas practicas
// 1. Versionamiento URI /api/v1/ping si cambia un el nombre de un endpoint la version
// 2. Nombre de los recursos en plural
// 3. Se recibe y se contesta con JSON
// 4. Responder con status code http adecuados
// 5. No poner verbos en los endpoints
// 6. Agrupar recursos que esten asociados tanto a nivel código como en la UR /campaigns/:campaignID/adsets/:adsetid/ads/:adID
// 7. Integrar filtrado, ordenado y paginación
// 8. Usar un cache para evitar computos repetitivos // lo hace gin por defecto de lo contrario se tendría que configurar
// 9. Documentar la API
// 10. Consistencia en las peticiones y respuestas de nuestra API
