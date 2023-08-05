package adaptor

import (
	"context"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/system"

	"meowchat-system/biz/application/service"
	"meowchat-system/biz/infrastructure/config"
)

type SystemServerImpl struct {
	*config.Config
	SystemService service.SystemService
}

func (s *SystemServerImpl) RetrieveNotice(ctx context.Context, req *system.RetrieveNoticeReq) (resp *system.RetrieveNoticeResp, rr error) {
	return s.SystemService.RetrieveNotice(ctx, req)
}

func (s *SystemServerImpl) ListNotice(ctx context.Context, req *system.ListNoticeReq) (resp *system.ListNoticeResp, rr error) {
	return s.SystemService.ListNotice(ctx, req)
}

func (s *SystemServerImpl) CreateNotice(ctx context.Context, req *system.CreateNoticeReq) (resp *system.CreateNoticeResp, rr error) {
	return s.SystemService.CreateNotice(ctx, req)
}

func (s *SystemServerImpl) UpdateNotice(ctx context.Context, req *system.UpdateNoticeReq) (resp *system.UpdateNoticeResp, rr error) {
	return s.SystemService.UpdateNotice(ctx, req)
}

func (s *SystemServerImpl) DeleteNotice(ctx context.Context, req *system.DeleteNoticeReq) (resp *system.DeleteNoticeResp, rr error) {
	return s.SystemService.DeleteNotice(ctx, req)
}

func (s *SystemServerImpl) RetrieveNews(ctx context.Context, req *system.RetrieveNewsReq) (resp *system.RetrieveNewsResp, rr error) {
	return s.SystemService.RetrieveNews(ctx, req)
}

func (s *SystemServerImpl) ListNews(ctx context.Context, req *system.ListNewsReq) (resp *system.ListNewsResp, rr error) {
	return s.SystemService.ListNews(ctx, req)
}

func (s *SystemServerImpl) CreateNews(ctx context.Context, req *system.CreateNewsReq) (resp *system.CreateNewsResp, rr error) {
	return s.SystemService.CreateNews(ctx, req)
}

func (s *SystemServerImpl) UpdateNews(ctx context.Context, req *system.UpdateNewsReq) (resp *system.UpdateNewsResp, rr error) {
	return s.SystemService.UpdateNews(ctx, req)
}

func (s *SystemServerImpl) DeleteNews(ctx context.Context, req *system.DeleteNewsReq) (resp *system.DeleteNewsResp, rr error) {
	return s.SystemService.DeleteNews(ctx, req)
}

func (s *SystemServerImpl) RetrieveAdmin(ctx context.Context, req *system.RetrieveAdminReq) (resp *system.RetrieveAdminResp, rr error) {
	return s.SystemService.RetrieveAdmin(ctx, req)
}

func (s *SystemServerImpl) ListAdmin(ctx context.Context, req *system.ListAdminReq) (resp *system.ListAdminResp, rr error) {
	return s.SystemService.ListAdmin(ctx, req)
}

func (s *SystemServerImpl) CreateAdmin(ctx context.Context, req *system.CreateAdminReq) (resp *system.CreateAdminResp, rr error) {
	return s.SystemService.CreateAdmin(ctx, req)
}

func (s *SystemServerImpl) UpdateAdmin(ctx context.Context, req *system.UpdateAdminReq) (resp *system.UpdateAdminResp, rr error) {
	return s.SystemService.UpdateAdmin(ctx, req)
}

func (s *SystemServerImpl) DeleteAdmin(ctx context.Context, req *system.DeleteAdminReq) (resp *system.DeleteAdminResp, rr error) {
	return s.SystemService.DeleteAdmin(ctx, req)
}

func (s *SystemServerImpl) RetrieveUserRole(ctx context.Context, req *system.RetrieveUserRoleReq) (resp *system.RetrieveUserRoleResp, rr error) {
	return s.SystemService.RetrieveUserRole(ctx, req)
}

func (s *SystemServerImpl) ListUserIdByRole(ctx context.Context, req *system.ListUserIdByRoleReq) (resp *system.ListUserIdByRoleResp, rr error) {
	return s.SystemService.ListUserIdByRole(ctx, req)
}

func (s *SystemServerImpl) UpdateUserRole(ctx context.Context, req *system.UpdateUserRoleReq) (resp *system.UpdateUserRoleResp, rr error) {
	return s.SystemService.UpdateUserRole(ctx, req)
}

func (s *SystemServerImpl) ContainsRole(ctx context.Context, req *system.ContainsRoleReq) (resp *system.ContainsRoleResp, rr error) {
	return s.SystemService.ContainsRole(ctx, req)
}

func (s *SystemServerImpl) CreateApply(ctx context.Context, req *system.CreateApplyReq) (resp *system.CreateApplyResp, rr error) {
	return s.SystemService.CreateApply(ctx, req)
}

func (s *SystemServerImpl) HandleApply(ctx context.Context, req *system.HandleApplyReq) (resp *system.HandleApplyResp, rr error) {
	return s.SystemService.HandleApply(ctx, req)
}

func (s *SystemServerImpl) ListApply(ctx context.Context, req *system.ListApplyReq) (resp *system.ListApplyResp, rr error) {
	return s.SystemService.ListApply(ctx, req)
}

func (s *SystemServerImpl) RetrieveCommunity(ctx context.Context, req *system.RetrieveCommunityReq) (resp *system.RetrieveCommunityResp, rr error) {
	return s.SystemService.RetrieveCommunity(ctx, req)
}

func (s *SystemServerImpl) ListCommunity(ctx context.Context, req *system.ListCommunityReq) (resp *system.ListCommunityResp, rr error) {
	return s.SystemService.ListCommunity(ctx, req)
}

func (s *SystemServerImpl) CreateCommunity(ctx context.Context, req *system.CreateCommunityReq) (resp *system.CreateCommunityResp, rr error) {
	return s.SystemService.CreateCommunity(ctx, req)
}

func (s *SystemServerImpl) UpdateCommunity(ctx context.Context, req *system.UpdateCommunityReq) (resp *system.UpdateCommunityResp, rr error) {
	return s.SystemService.UpdateCommunity(ctx, req)
}

func (s *SystemServerImpl) DeleteCommunity(ctx context.Context, req *system.DeleteCommunityReq) (resp *system.DeleteCommunityResp, rr error) {
	return s.SystemService.DeleteCommunity(ctx, req)
}
