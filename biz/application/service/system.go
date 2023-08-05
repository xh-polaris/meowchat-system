package service

import (
	"context"
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/consts"
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/data/db"
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/mapper"
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/util"
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/util/log"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/system"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"github.com/google/wire"
)

var (
	RoleTypeValue = map[string]system.RoleType{
		"unknown":        system.RoleType_unknown,
		"user":           system.RoleType_user,
		"communityAdmin": system.RoleType_communityAdmin,
		"superAdmin":     system.RoleType_superAdmin,
		"developer":      system.RoleType_developer,
	}
)

type SystemService interface {
	RetrieveNotice(ctx context.Context, req *system.RetrieveNoticeReq) (resp *system.RetrieveNoticeResp, err error)
	ListNotice(ctx context.Context, req *system.ListNoticeReq) (resp *system.ListNoticeResp, err error)
	CreateNotice(ctx context.Context, req *system.CreateNoticeReq) (resp *system.CreateNoticeResp, err error)
	UpdateNotice(ctx context.Context, req *system.UpdateNoticeReq) (resp *system.UpdateNoticeResp, err error)
	DeleteNotice(ctx context.Context, req *system.DeleteNoticeReq) (resp *system.DeleteNoticeResp, err error)
	RetrieveNews(ctx context.Context, req *system.RetrieveNewsReq) (resp *system.RetrieveNewsResp, err error)
	ListNews(ctx context.Context, req *system.ListNewsReq) (resp *system.ListNewsResp, err error)
	CreateNews(ctx context.Context, req *system.CreateNewsReq) (resp *system.CreateNewsResp, err error)
	UpdateNews(ctx context.Context, req *system.UpdateNewsReq) (resp *system.UpdateNewsResp, err error)
	DeleteNews(ctx context.Context, req *system.DeleteNewsReq) (resp *system.DeleteNewsResp, err error)
	RetrieveAdmin(ctx context.Context, req *system.RetrieveAdminReq) (resp *system.RetrieveAdminResp, err error)
	ListAdmin(ctx context.Context, req *system.ListAdminReq) (resp *system.ListAdminResp, err error)
	CreateAdmin(ctx context.Context, req *system.CreateAdminReq) (resp *system.CreateAdminResp, err error)
	UpdateAdmin(ctx context.Context, req *system.UpdateAdminReq) (resp *system.UpdateAdminResp, err error)
	DeleteAdmin(ctx context.Context, req *system.DeleteAdminReq) (resp *system.DeleteAdminResp, err error)
	RetrieveUserRole(ctx context.Context, req *system.RetrieveUserRoleReq) (resp *system.RetrieveUserRoleResp, err error)
	ListUserIdByRole(ctx context.Context, req *system.ListUserIdByRoleReq) (resp *system.ListUserIdByRoleResp, err error)
	UpdateUserRole(ctx context.Context, req *system.UpdateUserRoleReq) (resp *system.UpdateUserRoleResp, err error)
	ContainsRole(ctx context.Context, req *system.ContainsRoleReq) (resp *system.ContainsRoleResp, err error)
	CreateApply(ctx context.Context, req *system.CreateApplyReq) (resp *system.CreateApplyResp, err error)
	HandleApply(ctx context.Context, req *system.HandleApplyReq) (resp *system.HandleApplyResp, err error)
	ListApply(ctx context.Context, req *system.ListApplyReq) (resp *system.ListApplyResp, err error)
	RetrieveCommunity(ctx context.Context, req *system.RetrieveCommunityReq) (resp *system.RetrieveCommunityResp, err error)
	ListCommunity(ctx context.Context, req *system.ListCommunityReq) (resp *system.ListCommunityResp, err error)
	CreateCommunity(ctx context.Context, req *system.CreateCommunityReq) (resp *system.CreateCommunityResp, err error)
	UpdateCommunity(ctx context.Context, req *system.UpdateCommunityReq) (resp *system.UpdateCommunityResp, err error)
	DeleteCommunity(ctx context.Context, req *system.DeleteCommunityReq) (resp *system.DeleteCommunityResp, err error)
}

type SystemServiceImpl struct {
	AdminModel     mapper.AdminModel
	ApplyModel     mapper.ApplyModel
	CommunityModel mapper.CommunityModel
	NewsModel      mapper.NewsModel
	NoticeModel    mapper.NoticeModel
	UserRoleModel  mapper.UserRoleModel
}

var SystemSet = wire.NewSet(
	wire.Struct(new(SystemServiceImpl), "*"),
	wire.Bind(new(SystemService), new(*SystemServiceImpl)),
)

func (s *SystemServiceImpl) CheckCommunityIdExist(ctx context.Context, id string) (primitive.ObjectID, error) {
	if id == "" {
		return primitive.NilObjectID, nil
	}
	r, err := s.CommunityModel.FindOne(ctx, id)
	if err != nil {
		return primitive.NilObjectID, consts.ErrCommunityIdNotFound
	}
	return r.ID, nil
}

func (s *SystemServiceImpl) CheckParentCommunityId(ctx context.Context, parentId string) (primitive.ObjectID, error) {
	if parentId == "" {
		return primitive.NilObjectID, nil
	}
	r, err := s.CommunityModel.FindOne(ctx, parentId)
	if err != nil {
		return primitive.NilObjectID, consts.ErrCommunityIdNotFound
	}
	if r.ParentId != primitive.NilObjectID {
		return primitive.NilObjectID, consts.ErrChildCommunityNotAllowed
	}
	return r.ID, nil
}

func (s *SystemServiceImpl) RetrieveNotice(ctx context.Context, req *system.RetrieveNoticeReq) (resp *system.RetrieveNoticeResp, err error) {
	notice, err := s.NoticeModel.FindOne(ctx, req.Id)
	if err != nil {
		return nil, consts.Switch(err)
	}

	return &system.RetrieveNoticeResp{
		Notice: util.ConvertNotice(notice),
	}, nil
}

func (s *SystemServiceImpl) ListNotice(ctx context.Context, req *system.ListNoticeReq) (resp *system.ListNoticeResp, err error) {
	notices, count, err := s.NoticeModel.ListNotice(ctx, req)
	if err != nil {
		return nil, err
	}

	var res = make([]*system.Notice, len(notices))
	for i, n := range notices {
		res[i] = util.ConvertNotice(n)
	}

	return &system.ListNoticeResp{
		Notices: res,
		Count:   count,
	}, nil
}

func (s *SystemServiceImpl) CreateNotice(ctx context.Context, req *system.CreateNoticeReq) (resp *system.CreateNoticeResp, err error) {
	id := primitive.NewObjectID()

	err = s.NoticeModel.Insert(ctx, &db.Notice{
		ID:          id,
		CommunityId: req.CommunityId,
		Text:        req.Text,
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &system.CreateNoticeResp{
		Id: id.Hex(),
	}, nil
}

func (s *SystemServiceImpl) UpdateNotice(ctx context.Context, req *system.UpdateNoticeReq) (resp *system.UpdateNoticeResp, err error) {
	err = s.NoticeModel.UpdateNotice(ctx, req)

	if err != nil {
		return nil, consts.Switch(err)
	}

	return &system.UpdateNoticeResp{}, nil
}

func (s *SystemServiceImpl) DeleteNotice(ctx context.Context, req *system.DeleteNoticeReq) (resp *system.DeleteNoticeResp, err error) {
	err = s.NoticeModel.Delete(ctx, req.Id)
	if err != nil {
		return nil, consts.Switch(err)
	}

	return &system.DeleteNoticeResp{}, nil
}

func (s *SystemServiceImpl) RetrieveNews(ctx context.Context, req *system.RetrieveNewsReq) (resp *system.RetrieveNewsResp, err error) {
	news, err := s.NewsModel.FindOne(ctx, req.Id)
	if err != nil {
		return nil, consts.Switch(err)
	}

	return &system.RetrieveNewsResp{
		News: util.ConvertNews(news),
	}, nil
}

func (s *SystemServiceImpl) ListNews(ctx context.Context, req *system.ListNewsReq) (resp *system.ListNewsResp, err error) {
	news, count, err := s.NewsModel.ListNews(ctx, req)
	if err != nil {
		return nil, err
	}

	var res = make([]*system.News, len(news))
	for i, n := range news {
		res[i] = util.ConvertNews(n)
	}

	return &system.ListNewsResp{
		News:  res,
		Count: count,
	}, nil
}

func (s *SystemServiceImpl) CreateNews(ctx context.Context, req *system.CreateNewsReq) (resp *system.CreateNewsResp, err error) {
	id := primitive.NewObjectID()

	err = s.NewsModel.Insert(ctx, &db.News{
		ID:          id,
		CommunityId: req.CommunityId,
		ImageUrl:    req.ImageUrl,
		LinkUrl:     req.LinkUrl,
		Type:        req.Type,
		IsPublic:    req.IsPublic,
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &system.CreateNewsResp{
		Id: id.Hex(),
	}, nil
}

func (s *SystemServiceImpl) UpdateNews(ctx context.Context, req *system.UpdateNewsReq) (resp *system.UpdateNewsResp, err error) {
	err = s.NewsModel.UpdateNews(ctx, req)

	if err != nil {
		return nil, consts.Switch(err)
	}

	return &system.UpdateNewsResp{}, nil
}

func (s *SystemServiceImpl) DeleteNews(ctx context.Context, req *system.DeleteNewsReq) (resp *system.DeleteNewsResp, err error) {
	err = s.NewsModel.Delete(ctx, req.Id)
	if err != nil {
		return nil, consts.Switch(err)
	}

	return &system.DeleteNewsResp{}, nil
}

func (s *SystemServiceImpl) RetrieveAdmin(ctx context.Context, req *system.RetrieveAdminReq) (resp *system.RetrieveAdminResp, err error) {
	admin, err := s.AdminModel.FindOne(ctx, req.Id)
	if err != nil {
		return nil, consts.Switch(err)
	}

	return &system.RetrieveAdminResp{
		Admin: util.ConvertAdmin(admin),
	}, nil
}

func (s *SystemServiceImpl) ListAdmin(ctx context.Context, req *system.ListAdminReq) (resp *system.ListAdminResp, err error) {
	admins, count, err := s.AdminModel.ListAdmin(ctx, req)
	if err != nil {
		return nil, err
	}

	var res = make([]*system.Admin, len(admins))
	for i, admin := range admins {
		res[i] = util.ConvertAdmin(admin)
	}

	return &system.ListAdminResp{
		Admins: res,
		Count:  count,
	}, nil
}

func (s *SystemServiceImpl) CreateAdmin(ctx context.Context, req *system.CreateAdminReq) (resp *system.CreateAdminResp, err error) {
	id := primitive.NewObjectID()

	err = s.AdminModel.Insert(ctx, &db.Admin{
		ID:          id,
		CommunityId: req.CommunityId,
		Name:        req.Name,
		Title:       req.Title,
		Phone:       req.Phone,
		Wechat:      req.Wechat,
		AvatarUrl:   req.AvatarUrl,
		UpdateAt:    time.Now(),
		CreateAt:    time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &system.CreateAdminResp{
		Id: id.Hex(),
	}, nil
}

func (s *SystemServiceImpl) UpdateAdmin(ctx context.Context, req *system.UpdateAdminReq) (resp *system.UpdateAdminResp, err error) {
	err = s.AdminModel.UpdateAdmin(ctx, req)

	if err != nil {
		return nil, consts.Switch(err)
	}

	return &system.UpdateAdminResp{}, nil
}

func (s *SystemServiceImpl) DeleteAdmin(ctx context.Context, req *system.DeleteAdminReq) (resp *system.DeleteAdminResp, err error) {
	err = s.AdminModel.Delete(ctx, req.Id)

	if err != nil {
		return nil, consts.Switch(err)
	}

	return &system.DeleteAdminResp{}, nil
}

func (s *SystemServiceImpl) RetrieveUserRole(ctx context.Context, req *system.RetrieveUserRoleReq) (resp *system.RetrieveUserRoleResp, err error) {

	userRole, err := s.UserRoleModel.FindOne(ctx, req.UserId)

	if err != nil {
		switch err {
		case mapper.ErrNotFound:
			return &system.RetrieveUserRoleResp{
				Roles: make([]*system.Role, 0),
			}, nil
		case mapper.ErrInvalidObjectId:
			return nil, consts.ErrInvalidObjectId
		default:
			return nil, err
		}
	}

	var res = make([]*system.Role, len(userRole.Roles))
	for i, role := range userRole.Roles {
		log.Info(RoleTypeValue[role.Type].String())
		res[i] = &system.Role{
			RoleType:    RoleTypeValue[role.Type],
			CommunityId: &role.CommunityId,
		}
	}

	return &system.RetrieveUserRoleResp{
		Roles: res,
	}, nil
}

func (s *SystemServiceImpl) ListUserIdByRole(ctx context.Context, req *system.ListUserIdByRoleReq) (resp *system.ListUserIdByRoleResp, err error) {
	Users, err := s.UserRoleModel.FindMany(ctx, req.Role.RoleType.String(), *req.Role.CommunityId)

	if err != nil {
		switch err {
		case mapper.ErrNotFound:
			return &system.ListUserIdByRoleResp{
				UserId: make([]string, 0),
			}, nil
		case mapper.ErrInvalidObjectId:
			return nil, consts.ErrInvalidObjectId
		default:
			return nil, err
		}
	}

	var res = make([]string, len(Users))
	for i, user := range Users {
		res[i] = user.ID.Hex()
	}

	return &system.ListUserIdByRoleResp{
		UserId: res,
	}, nil
}

func (s *SystemServiceImpl) UpdateUserRole(ctx context.Context, req *system.UpdateUserRoleReq) (resp *system.UpdateUserRoleResp, err error) {
	id, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		return nil, consts.ErrInvalidObjectId
	}

	roles := make([]db.Role, len(req.Roles))
	for i, role := range req.Roles {
		if role.RoleType.String() == db.RoleCommunityAdmin {
			id, _ := s.CheckCommunityIdExist(ctx, *role.CommunityId)
			if id == primitive.NilObjectID {
				return nil, consts.ErrCommunityIdNotFound
			}
		}
		roles[i] = db.Role{
			Type:        role.RoleType.String(),
			CommunityId: *role.CommunityId,
		}
	}

	_, err = s.UserRoleModel.Upsert(ctx, &db.UserRole{
		ID:    id,
		Roles: roles,
	})
	if err != nil {
		return nil, consts.Switch(err)
	}

	return &system.UpdateUserRoleResp{}, nil
}

func (s *SystemServiceImpl) subCommunityOf(ctx context.Context, cid1 string, cid2 string) bool {
	if cid1 == cid2 {
		return true
	}
	c1, _ := s.CommunityModel.FindOne(ctx, cid1)
	return c1 != nil && c1.ParentId.Hex() == cid2
}

func (s *SystemServiceImpl) ContainsRole(ctx context.Context, req *system.ContainsRoleReq) (resp *system.ContainsRoleResp, err error) {
	resp = &system.ContainsRoleResp{}

	if req.Role == nil {
		req.Role = &system.Role{}
	}

	userRole, _ := s.UserRoleModel.FindOne(ctx, req.UserId)
	if userRole == nil {
		return
	}

	for _, role := range userRole.Roles {
		switch role.Type {
		case db.RoleSuperAdmin:
			if req.AllowSuperAdmin || req.Role.RoleType.String() == db.RoleSuperAdmin {
				resp.Contains = true
				return
			}
		case db.RoleCommunityAdmin:
			if req.Role.RoleType.String() == db.RoleCommunityAdmin &&
				(*req.Role.CommunityId == "" || s.subCommunityOf(ctx, *req.Role.CommunityId, role.CommunityId)) {
				resp.Contains = true
				return
			}
		default:
			if req.Role.RoleType.String() == role.Type {
				resp.Contains = true
				return
			}
		}
	}

	return
}

func (s *SystemServiceImpl) CreateApply(ctx context.Context, req *system.CreateApplyReq) (resp *system.CreateApplyResp, err error) {
	if err := s.ApplyModel.Insert(ctx, &db.Apply{
		ApplicantId: req.ApplicantId,
		CommunityId: req.CommunityId,
	}); err != nil {
		return nil, err
	}

	return &system.CreateApplyResp{}, nil
}

func (s *SystemServiceImpl) HandleApply(ctx context.Context, req *system.HandleApplyReq) (resp *system.HandleApplyResp, err error) {
	if req.IsRejected == false {
		apply, err := s.ApplyModel.FindOne(ctx, req.ApplyId)
		if err != nil {
			return nil, err
		}
		userRole, err := s.UserRoleModel.FindOne(ctx, apply.ApplicantId)
		if err != nil {
			return nil, err
		}
		userRole.Roles = append(userRole.Roles, db.Role{
			Type:        db.RoleCommunityAdmin,
			CommunityId: apply.CommunityId,
		})
	}
	_, err = s.ApplyModel.Delete(ctx, req.ApplyId)
	if err != nil {
		return nil, err
	}
	return &system.HandleApplyResp{}, nil
}

func (s *SystemServiceImpl) ListApply(ctx context.Context, req *system.ListApplyReq) (resp *system.ListApplyResp, err error) {
	res, err := s.ApplyModel.FindAllApplyByCommunityId(ctx, req)
	if err != nil {
		return nil, err
	}
	apply := make([]*system.Apply, 0, len(res))
	for _, x := range res {
		apply = append(apply, &system.Apply{
			ApplyId:     x.ID.Hex(),
			ApplicantId: x.ApplicantId,
			CommunityId: x.CommunityId,
		})
	}
	return &system.ListApplyResp{Apply: apply}, nil
}

func (s *SystemServiceImpl) RetrieveCommunity(ctx context.Context, req *system.RetrieveCommunityReq) (resp *system.RetrieveCommunityResp, err error) {
	community, err := s.CommunityModel.FindOne(ctx, req.Id)

	if err != nil {
		return nil, consts.Switch(err)
	}

	return &system.RetrieveCommunityResp{
		Community: util.ConvertCommunity(community),
	}, nil
}

func (s *SystemServiceImpl) ListCommunity(ctx context.Context, req *system.ListCommunityReq) (resp *system.ListCommunityResp, err error) {
	communities, count, err := s.CommunityModel.ListCommunity(ctx, req)
	if err != nil {
		return nil, consts.Switch(err)
	}

	var res = make([]*system.Community, len(communities))
	for i, community := range communities {
		res[i] = util.ConvertCommunity(community)
	}

	return &system.ListCommunityResp{
		Communities: res,
		Count:       count,
	}, nil
}

func (s *SystemServiceImpl) CreateCommunity(ctx context.Context, req *system.CreateCommunityReq) (resp *system.CreateCommunityResp, err error) {
	id := primitive.NewObjectID()

	parentId, err := s.CheckParentCommunityId(ctx, req.ParentId)
	if err != nil {
		return nil, err
	}

	err = s.CommunityModel.Insert(ctx, &db.Community{
		ID:       id,
		Name:     req.Name,
		ParentId: parentId,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &system.CreateCommunityResp{
		Id: id.Hex(),
	}, nil
}

func (s *SystemServiceImpl) UpdateCommunity(ctx context.Context, req *system.UpdateCommunityReq) (resp *system.UpdateCommunityResp, err error) {
	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, consts.ErrInvalidObjectId
	}

	parentId, err := s.CheckParentCommunityId(ctx, req.ParentId)
	if err != nil {
		return nil, err
	}

	_, err = s.CommunityModel.Update(ctx, &db.Community{
		ID:       id,
		Name:     req.Name,
		ParentId: parentId,
	})

	if err != nil {
		return nil, consts.Switch(err)
	}

	return &system.UpdateCommunityResp{}, nil
}

func (s *SystemServiceImpl) DeleteCommunity(ctx context.Context, req *system.DeleteCommunityReq) (resp *system.DeleteCommunityResp, err error) {
	err = s.CommunityModel.DeleteCommunity(ctx, req.Id)

	if err != nil {
		return nil, consts.Switch(err)
	}

	return &system.DeleteCommunityResp{}, nil
}
