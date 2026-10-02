package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AliAstanov/my_blogs/models"
	"github.com/AliAstanov/my_blogs/pkg/mail"
	"github.com/gin-gonic/gin"
	"github.com/saidamir98/udevs_pkg/logger"
)


func (h *handlers) CheckUser(ctx *gin.Context) {
	var reqBody models.CheckViewer

	if err := ctx.BindJSON(&reqBody); err != nil {
		h.log.Error("error binding request body", logger.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	isExists, err := h.storage.GetCommonRepo().CheckIsExists(ctx, &models.Common{
		TableName:  "viewrs",
		ColumnName: "gmail",
		ExpValue:   reqBody.Gmail,
	})
	if err != nil {
		h.log.Error("error on checking viewer", logger.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if isExists {
		ctx.JSON(http.StatusOK, models.CheckisExistsReq{
			IsExists: isExists,
			Status:   "log-in",
		})
		return
	}

	otp := models.OtpData{
		Gmail: reqBody.Gmail,
		Otp:   mail.GenerateOtp(6),
	}

	otpDataB, err := json.Marshal(otp)
	if err != nil {
		h.log.Error("error marshaling OTP data", logger.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	err = h.cache.Set(ctx, reqBody.Gmail, string(otpDataB), 60)
	if err != nil {
		h.log.Error("error setting OTP data to cache", logger.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	err = mail.SendMail([]string{reqBody.Gmail}, otp.Otp)
	if err != nil {
		h.log.Error("error sending OTP mail", logger.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusCreated, models.CheckisExistsReq{
		IsExists: isExists,
		Status:   "register",
	})

}

func (h *handlers) ChekOtp(ctx *gin.Context) {
	var reqBody models.OtpData

	err := ctx.Bind(&reqBody)
	if err != nil {
		return
	}

	gmail := reqBody.Gmail

	data, err := h.cache.GetDel(ctx, gmail)
	if err != nil {
		h.log.Error("error on checking-otp:",logger.Error(err))
		return
	}
	if data == "" {
		ctx.JSON(201,"otp is expired")
	}

	var cacheData models.OtpData

	json.Unmarshal([]byte(data), &cacheData)

	ctx.JSON(201,models.CheckOtpResp{
		IsRight: cacheData.Otp == reqBody.Otp})
}

func(h *handlers) SignUp(ctx *gin.Context){
	var regReqBody models.ViewerRegReq

	err := ctx.ShouldBind(&regReqBody)
	if err != nil {
		h.log.Error("error on sign-up",logger.Error(err))
		return
	}
	otpSData, err :=h.cache.GetDel(ctx,regReqBody.Gmail)
	if err != nil {
		h.log.Error("error on get-del:",logger.Error(err))
		return
	}
	if otpSData == "" {
		ctx.JSON(201,"otp is expired")
	}

	var otpData models.OtpData

	err = json.Unmarshal([]byte(otpSData),&otpData)
	if err != nil {
		h.log.Error("error on unmarshal otp-data:",logger.Error(err))
		return
	}

	if otpData.Otp != regReqBody.Otp {
		ctx.JSON(405,"otp is incorrect")
		return
	}

}
