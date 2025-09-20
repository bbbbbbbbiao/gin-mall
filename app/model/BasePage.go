package model

/**
 * @author: biao
 * @date: 2025/9/2 10:40
 * @code: 彼方尚有荣光在
 * @description: 分页信息
 */

type BasePage struct {
	PageNum  int `form:"pageNum"`
	PageSize int `form:"pageSize"`
}
