// ==========================================================================
// LV自动生成model扩展代码列表、增、删，改、查、导出，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-07-30 21:59:38 +0800 CST
// 生成人：lv
// ==========================================================================
package dao

import (
	"github.com/lostvip-com/lv_framework/lv_db"
	"github.com/lostvip-com/lv_framework/lv_db/lv_batis"
	"github.com/lostvip-com/lv_framework/lv_db/lv_dao"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"main/internal/iot_dev/model"
	"main/internal/iot_dev/vo"
)

// 新增页面请求参数
type IotPrdPropertyDao struct{}

// 根据条件分页查询数据
func (d IotPrdPropertyDao) ListMapByPage(req *vo.PageIotPrdPropertyReq) (*[]map[string]string, int64, error) {
	ibatis := lv_batis.NewInstance("iot_dev/iot_prd_property_mapper.sql") //under the mapper directory
	// 约定用方法名ListByPage对应sql文件中的同名tagName
	limitSQL, err := ibatis.GetLimitSql("ListIotPrdProperty", req)
	//查询数据
	rows, err := lv_dao.ListMapByNamedSql(limitSQL, req, true)
	lv_err.HasErrAndPanic(err)
	count, err := lv_dao.CountByNamedSql(ibatis.GetCountSql(), req)
	lv_err.HasErrAndPanic(err)
	return rows, count, nil
}

// 根据条件分页查询数据
func (d IotPrdPropertyDao) ListByPage(req *vo.PageIotPrdPropertyReq) (*[]model.IotPrdProperty, int64, error) {
	ibatis := lv_batis.NewInstance("iot_dev/iot_prd_property_mapper.sql") //under the mapper directory
	// 对应sql文件中的同名tagName
	limitSQL, err := ibatis.GetLimitSql("ListIotPrdProperty", req)
	//查询数据
	rows, err := lv_dao.ListByNamedSql[model.IotPrdProperty](limitSQL, req)
	lv_err.HasErrAndPanic(err)
	count, err := lv_dao.CountByNamedSql(ibatis.GetCountSql(), req)
	lv_err.HasErrAndPanic(err)
	return rows, count, nil
}

// ListAll 导出excel使用
func (d IotPrdPropertyDao) ListAll(req *vo.PageIotPrdPropertyReq, isCamel bool) (*[]map[string]string, error) {
	ibatis := lv_batis.NewInstance("iot_dev/iot_prd_property_mapper.sql")
	// 约定用方法名ListByPage对应sql文件中的同名tagName
	sql, err := ibatis.GetSql("ListIotPrdProperty", req)
	lv_err.HasErrAndPanic(err)

	arr, err := lv_dao.ListMapByNamedSql(sql, req, isCamel)
	return arr, err
}

// FindByWhere 根据条件查询
func (d IotPrdPropertyDao) Find(where, order string) (*[]model.IotPrdProperty, error) {
	var list []model.IotPrdProperty
	err := lv_db.GetMasterGorm().Table("iot_prd_property").Where(where).Order(order).Find(&list).Error
	return &list, err
}

func (d IotPrdPropertyDao) DeleteByIds(ida []int64) int64 {
	db := lv_db.GetMasterGorm().Table("iot_prd_property").Where("id in ? ", ida).Update("del_flag", 1)
	if db.Error != nil {
		panic(db.Error)
	}
	return db.RowsAffected
}
