package serializer

/**
 * @author: biao
 * @date: 2025/9/21 19:46
 * @code: 彼方尚有荣光在
 * @description: 同意返回格式设置
 */

type DataList struct {
	Total uint        `json:"total"`
	Item  interface{} `json:"item"`
}

func BuildDataList(total uint, item interface{}) *DataList {
	return &DataList{
		Total: total,
		Item:  item,
	}
}
