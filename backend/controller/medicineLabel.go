package controller

import (

	"fmt"
	"net/http"
	
	"github.com/asaskevich/govalidator"

	"github.com/sut65/team15/entity"

	"github.com/gin-gonic/gin"

	
)

// POST /ambulances
func CreateMedicineLabel(c *gin.Context) {
	var medicineLabel entity.MedicineLabel
	var order entity.Order
	var effect entity.Effect
	var pharmacist entity.User
	var suggestion entity.Suggestion
	
	//เช็คว่าตรงกันมั้ย
	if err := c.ShouldBindJSON(&medicineLabel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//  ค้นหา user ด้วย id
	if tx := entity.DB().Where("id = ?", medicineLabel.PharmacistID).First(&pharmacist); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "pharmacist not found"})
		return
	}
	// 8: ค้นหา ambulancetype ด้วย id
	if tx := entity.DB().Where("id = ?", medicineLabel.OrderID).First(&order); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicine not found"})
		return
	}
	// 9: ค้นหา brand ด้วย id
	if tx := entity.DB().Where("id = ?", medicineLabel.EffectID).First(&effect); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "side effect not found"})
		return
	}
	// 10: ค้นหา status ด้วย id
	if tx := entity.DB().Where("id = ?", medicineLabel.SuggestionID).First(&suggestion); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "suggestion not found"})
		return
	}
	// 11: สร้าง Ambulance
	wv := entity.MedicineLabel{
		Instruction:          medicineLabel.Instruction,
		Property:             medicineLabel.Property,
		Consumption:          medicineLabel.Consumption,
		Date:                 medicineLabel.Date,
		Order:                order,
		Suggestion:           suggestion,
		Effect:               effect,
		Pharmacist:           pharmacist,     // โยงความสัมพันธ์กับ Entity user
	}

	fmt.Println(wv.Date)

	if _, err := govalidator.ValidateStruct(wv); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 12: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wv})
}

// GET /ambulance/:id
func GetMedicineLabel(c *gin.Context) {
	var medicineLabel entity.MedicineLabel
	id := c.Param("id")
	if err := entity.DB().Preload("Order").Preload("Order.Medicine").Preload("Suggestion").Preload("Effect").Preload("Pharmacist").Raw("SELECT * FROM medicine_labels WHERE id = ?", id).Find(&medicineLabel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": medicineLabel})
}

// GET /ambulances
func ListMedicineLabel(c *gin.Context) {
	var medicineLabels []entity.MedicineLabel
	if err := entity.DB().Preload("Order").Preload("Order.Medicine").Preload("Suggestion").Preload("Effect").Preload("Pharmacist").Raw("SELECT * FROM medicine_labels").Find(&medicineLabels).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": medicineLabels})
}

// DELETE /ambulances/:id
func DeleteMedicineLabel(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM medicine_labels WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ambulance not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /ambulances
func UpdateMedicineLabel(c *gin.Context) {
	var medicineLabelold entity.MedicineLabel 
	var medicineLabel entity.MedicineLabel

	var order entity.Order
	var effect entity.Effect
	var pharmacist entity.User
	var suggestion entity.Suggestion

	if err := c.ShouldBindJSON(&medicineLabel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	if tx := entity.DB().Where("id = ?", medicineLabel.ID).First(&medicineLabelold); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("MedicineLabel id = %d not found", medicineLabel.ID)})
		c.Abort()
		return
	}

	if medicineLabel.Consumption == "" {
		medicineLabel.Consumption = medicineLabelold.Consumption
	}

	if medicineLabel.Instruction == "" {
		medicineLabel.Instruction = medicineLabelold.Instruction
	}

	if medicineLabel.Date.String() == "0001-01-01 00:00:00 +0000 UTC" {
		medicineLabel.Date = medicineLabelold.Date
	}

	if medicineLabel.PharmacistID != nil {
		if tx := entity.DB().Where("id = ?", medicineLabel.PharmacistID).First(&pharmacist); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "not found Order"})
			return
		}
		fmt.Print("NOT NULL")
		medicineLabel.Order = order
	} else {
		if tx := entity.DB().Where("id = ?", medicineLabel.PharmacistID).First(&pharmacist); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "not found Order"})
			return
		}
		fmt.Print("NULL")
		medicineLabel.Pharmacist = pharmacist
	}

	if medicineLabel.OrderID != nil {
		if tx := entity.DB().Where("id = ?", medicineLabel.OrderID).First(&order); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "not found Order"})
			return
		}
		fmt.Print("NOT NULL")
		medicineLabel.Order = order
	} else {
		if tx := entity.DB().Where("id = ?", medicineLabel.OrderID).First(&order); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "not found Order"})
			return
		}
		fmt.Print("NULL")
		medicineLabel.Order = order
	}
	
	if medicineLabel.SuggestionID != nil {
		if tx := entity.DB().Where("id = ?", medicineLabel.SuggestionID).First(&suggestion); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "not found suggedtion"})
			return
		}
		fmt.Print("NOT NULL")
		medicineLabel.Suggestion = suggestion
	} else {
		if tx := entity.DB().Where("id = ?", medicineLabel.SuggestionID).First(&suggestion); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "not found suggestion"})
			return
		}
		fmt.Print("NULL")
		medicineLabel.Suggestion = suggestion
	}

	if medicineLabel.EffectID != nil {
		if tx := entity.DB().Where("id = ?", medicineLabel.EffectID).First(&effect); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "not found suggedtion"})
			return
		}
		fmt.Print("NOT NULL")
		medicineLabel.Effect = effect
	} else {
		if tx := entity.DB().Where("id = ?", medicineLabel.EffectID).First(&effect); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "not found suggestion"})
			return
		}
		fmt.Print("NULL")
		medicineLabel.Effect = effect
	}
// Update abl in database
	if err := entity.DB().Save(&medicineLabel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "Update Success",
		"data":   medicineLabel,
	})

}