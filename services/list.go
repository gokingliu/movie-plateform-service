package services

import (
	"MovieService/config"
	"MovieService/logic"
	"MovieService/utils"
	"context"
	pb "git.woa.com/crotaliu/pb-hub"
)

type ListImpl struct{}

// GetList 获取视频列表
func (s *ListImpl) GetList(ctx context.Context, req *pb.GetListReq, rsp *pb.GetListRsp) error {
	// ConnDB 实例
	db := utils.ConnDB()
	// 判断 token，并获取用户名、用户角色
	tokenBol, _, role, tokenErr := logic.PreHandleTokenLogic(db, ctx)
	// struct 转 map
	mapData, mapErr := utils.StructToMap(req)
	if mapErr != nil {
		rsp.Code, rsp.Msg = config.ResFail.Code, utils.GetErrorMap(mapErr.Error()).Msg
		return nil
	}
	// 查询电影列表逻辑
	resultList, count, dbErr := logic.GetListLogic(db, mapData, role, req.PageNo, req.PageSize)
	if dbErr != nil {
		rsp.Code, rsp.Msg = config.InnerReadDbError.Code, config.InnerReadDbError.Msg
		return nil
	}
	// 客户端接口，判断 token 解析是否正常
	// 不正常时返回相应的错误码，同时返回列表信息，前端正常展示列表，但清空用户信息
	if !tokenBol {
		rsp.Code, rsp.Msg = config.ClientUserInfoError.Code, utils.GetErrorMap(tokenErr.Error()).Msg
	} else {
		rsp.Code, rsp.Msg = config.ResOk.Code, config.ResOk.Msg
	}
	// 解构类型重新赋值
	// TODO 不优雅，后续要优化
	var listRsp []*pb.GetListRsp_List
	for _, result := range resultList {
		listRsp = append(listRsp, &pb.GetListRsp_List{
			Mid:           result.Mid,
			MName:         result.MName,
			MPoster:       result.MPoster,
			MTypeName:     result.MTypeName,
			MDouBanScore:  result.MDouBanScore,
			MDirector:     result.MDirector,
			MStarring:     result.MStarring,
			MCountryName:  result.MCountryName,
			MLanguageName: result.MLanguageName,
			MDateYear:     result.MDateYear,
			MDate:         result.MDate,
			MViews:        result.MViews,
			MLikes:        result.MLikes,
			MCollects:     result.MCollects,
			CreateTime:    result.CreateTime,
			UpdateTime:    result.UpdateTime,
		})
	}
	rsp.Result = &pb.GetListRsp_Result{
		List:  listRsp,
		Count: count,
	}

	return nil
}

// GetLeaderboard 获取视频排行榜
func (s *ListImpl) GetLeaderboard(ctx context.Context, req *pb.GetLeaderboardReq, rsp *pb.GetLeaderboardRsp) error {
	// ConnDB 实例
	db := utils.ConnDB()
	// 判断 token，并获取用户名、用户角色
	tokenBol, _, _, tokenErr := logic.PreHandleTokenLogic(db, ctx)
	// 查询视频排行榜逻辑
	resultList, dbErr := logic.GetLeaderboardLogic(db, req.MType)
	if dbErr != nil {
		rsp.Code, rsp.Msg = config.InnerReadDbError.Code, config.InnerReadDbError.Msg
		return nil
	}
	// 客户端接口，判断 token 解析是否正常
	// 不正常时返回相应的错误码，同时返回列表信息，前端正常展示列表，但清空用户信息
	if !tokenBol {
		rsp.Code, rsp.Msg = config.ClientUserInfoError.Code, utils.GetErrorMap(tokenErr.Error()).Msg
	} else {
		rsp.Code, rsp.Msg = config.ResOk.Code, config.ResOk.Msg
	}
	// TODO 不优雅，后续要优化
	var listRsp []*pb.GetLeaderboardRsp_List
	for _, result := range resultList {
		listRsp = append(listRsp, &pb.GetLeaderboardRsp_List{
			Mid:    result.Mid,
			MName:  result.MName,
			MTotal: result.MTotal,
		})
	}

	rsp.Result = listRsp

	return nil
}
